package controllers

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/MetaBloxIO/metablox-foundation-services/did"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metabloxStaking/comm/regutil"
	"github.com/metabloxStaking/dao"

	serviceModels "github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/MetaBloxIO/metablox-foundation-services/presentations"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/contract"
	"github.com/metabloxStaking/errval"
	"github.com/metabloxStaking/models"
)

const (
	NONCE_TIME_OUT = 120 * 1000
)

var NoncePool map[string]models.VpNonce = map[string]models.VpNonce{}

func DeleteTimeoutSession() {
	timer := time.NewTimer(time.Second * NONCE_TIME_OUT * 3)

	go func(timer *time.Timer) {
		select {
		case <-timer.C:
			t := time.Now()
			for session, c := range NoncePool {
				if t.UnixMilli()-c.Time.UnixMilli() > NONCE_TIME_OUT {
					delete(NoncePool, session)
				}
			}
		}
	}(timer)
}

func ApplyNonce(session string) (uint64, error) {
	_, ok := NoncePool[session]
	if ok {
		return 0, errval.ErrInvalidSession
	}

	NoncePool[session] = models.VpNonce{Session: session, Nonce: rand.Uint64(), Time: time.Now()}

	return NoncePool[session].Nonce, nil
}

func GetNonce(c *gin.Context) (uint64, error) {
	session := c.GetHeader("session")
	if len(session) == 0 {
		return 0, errval.ErrInvalidSession
	}
	nonce, ok := NoncePool[session]
	if !ok {
		return 0, errval.ErrInvalidSession
	}

	t := time.Now()
	if t.UnixMilli()-nonce.Time.UnixMilli() > NONCE_TIME_OUT {
		return 0, errval.ErrNonceTimeout
	}

	delete(NoncePool, session)

	return nonce.Nonce, nil
}

func ActivateExchange(c *gin.Context) error {
	session := c.GetHeader("session")
	if len(session) == 0 {
		return errval.ErrInvalidSession
	}

	var vp models.MiningRoleInput

	err := c.BindJSON(&vp)
	if err != nil {
		return err
	}

	err = ValidateDID(vp.SeedPresentation.Holder)
	if err != nil {
		return err
	}

	//TODO  Check VC credentialSubjects

	_, err = presentations.VerifyVP(&vp.SeedPresentation) //going to fail at the moment as we don't have all the info to do this verification
	if err != nil {                                       //skip this error check to avoid failures until we can properly verify seed presentations
		return err
	}

	role := models.MiningRole{
		DID:           vp.SeedPresentation.Holder,
		WalletAddress: vp.WalletAddress,
		Type:          vp.SeedPresentation.VerifiableCredential[0].Type[1],
	}

	err = dao.InsertMiningRole(&role)
	return err
}

func NewExchangeSeed(c *gin.Context) ([]*models.SeedExchangeOutput, error) {
	var input models.NewSeedExchangeInput
	var outputArray []*models.SeedExchangeOutput
	err := c.BindJSON(&input)

	if err != nil {
		return nil, err
	}

	valid := regutil.IsETHAddress(input.WalletAddress)
	if !valid {
		return nil, errval.ErrETHAddress
	}

	for _, seed := range input.Seeds {

		if seed.Confirm.Did != seed.Result.Target ||
			seed.Confirm.Target != seed.Result.Did {
			return nil, errval.ErrDIDMismatch
		}

		err = ValidateDID(seed.Confirm.Did)
		if err != nil {
			return nil, err
		}

		exists, err := dao.CheckIfDIDIsMiner(seed.Confirm.Did)
		if !exists || err != nil {
			return nil, errval.ErrInvalidMiner
		}

		err = ValidateDID(seed.Result.Did)
		if err != nil {
			return nil, err
		}

		exists, err = dao.CheckIfDIDIsValidator(seed.Result.Did)
		if !exists || err != nil {
			return nil, errval.ErrInvalidValidator
		}
		serviceModels.GenerateTestDIDDocument()

		ret, err := verifyNetworkReq(&seed.Confirm)
		if err != nil || !ret {
			return nil, errval.ErrSignatureVerifyError
		}

		ret, err = verifyNetworkResult(&seed.Result)
		if err != nil || !ret {
			return nil, errval.ErrSignatureVerifyError
		}

		role, err := dao.GetMiningRole(seed.Result.Did)
		if err != nil {
			return nil, err
		}
		if role == nil {
			return nil, errval.ErrMinerRoleNotFound
		}

		valid = regutil.IsETHAddress(input.WalletAddress)
		if !valid {
			return nil, errval.ErrETHAddress
		}

		sendSeedToken(seed.Confirm.Target, role.WalletAddress)
		output, err := sendSeedToken(seed.Confirm.Did, input.WalletAddress)
		if err != nil {
			return nil, err
		}
		outputArray = append(outputArray, output)
	}
	return outputArray, nil
}

func sendSeedToken(DID, addressString string) (*models.SeedExchangeOutput, error) {
	targetAddress := common.HexToAddress(addressString)
	//todo: may have to change calculation method
	txHash, err := contract.TransferTokens(targetAddress, int(placeholderExchangeRate)) //todo: need a proper method of converting exchangeValue into an int
	if err != nil {
		return nil, err
	}

	exchange := models.NewSeedExchange("", DID, placeholderExchangeRate, placeholderExchangeRate)

	//todo: uncomment when we have a valid value for the seed VcID. This function will fail if the VcID is an empty string
	/*err = dao.UploadSeedExchange(exchange)
	if err != nil {
		fmt.Println("check2")
		return nil, err
	}*/

	txTime := strconv.FormatFloat(float64(time.Now().UnixNano())/float64(time.Second), 'f', 3, 64)
	output := models.NewSeedExchangeOutput(exchange.Amount, txHash, txTime, exchange.ExchangeRate)

	return output, nil
}

func verifyNetworkReq(req *models.NetworkConfirmRequest) (bool, error) {
	bytes, err := serializeNetworkReq(req)

	if err != nil {
		return false, err
	}

	resolutionMeta, holderDoc, _ := did.Resolve(req.Did, serviceModels.CreateResolutionOptions())
	if resolutionMeta.Error != "" {
		return false, errors.New(resolutionMeta.Error)
	}

	targetVM := holderDoc.VerificationMethod[0]

	hashedData := sha256.Sum256(bytes)
	pubData, err := base64.StdEncoding.DecodeString(req.PubKey)
	if err != nil {
		return false, err
	}

	pubKey, err := crypto.UnmarshalPubkey(pubData)
	if err != nil {
		return false, err
	}

	address := crypto.PubkeyToAddress(*pubKey)
	accountId := "eip155:1666600000:" + address.Hex()

	if accountId != targetVM.BlockchainAccountId {
		return false, errors.New("pubkey and document mismatch")
	}

	sig, err := base64.StdEncoding.DecodeString(req.Signature)
	if err != nil {
		return false, err
	}
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:])

	return ecdsa.Verify(pubKey, hashedData[:], r, s), nil
}

func serializeNetworkReq(req *models.NetworkConfirmRequest) ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString(req.Did)
	buffer.WriteString(req.Target)
	buffer.WriteString(req.LastBlockHash)
	buffer.WriteString(req.Quality)
	buffer.WriteString(req.PubKey)
	buffer.WriteString(req.Challenge)

	return buffer.Bytes(), nil
}

func verifyNetworkResult(resp *models.NetworkConfirmResult) (bool, error) {
	lbytes, err := serializeNetworkResult(resp)

	if err != nil {
		return false, err
	}

	resolutionMeta, holderDoc, _ := did.Resolve(resp.Did, serviceModels.CreateResolutionOptions())
	if resolutionMeta.Error != "" {
		return false, errors.New(resolutionMeta.Error)
	}

	targetVM := holderDoc.VerificationMethod[0]

	hashedData := sha256.Sum256(lbytes)
	pubData, err := base64.StdEncoding.DecodeString(resp.PubKey)
	if err != nil {
		return false, err
	}

	pubKey, err := crypto.UnmarshalPubkey(pubData)
	if err != nil {
		return false, err
	}

	address := crypto.PubkeyToAddress(*pubKey)
	accountId := "eip155:1666600000:" + address.Hex()

	if accountId != targetVM.BlockchainAccountId {
		return false, errors.New("pubkey and document mismatch")
	}

	sig, err := base64.StdEncoding.DecodeString(resp.Signature)
	if err != nil {
		return false, err
	}
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:])

	return ecdsa.Verify(pubKey, hashedData[:], r, s), nil
}

func serializeNetworkResult(result *models.NetworkConfirmResult) ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString(result.Did)
	buffer.WriteString(result.Target)
	buffer.WriteString(result.LastBlockHash)
	buffer.WriteString(result.PubKey)
	buffer.WriteString(result.Challenge)

	return buffer.Bytes(), nil
}

func ExchangeSeed(c *gin.Context) (*models.SeedExchangeOutput, error) {
	input := models.CreateSeedExchangeInput()
	err := c.BindJSON(input)
	if err != nil {
		return nil, err
	}

	err = ValidateDID(input.DID)
	if err != nil {
		return nil, err
	}

	_, err = presentations.VerifyVP(&input.SeedPresentation) //going to fail at the moment as we don't have all the info to do this verification
	if err != nil {                                          //skip this error check to avoid failures until we can properly verify seed presentations
		//return nil, error
	}

	targetAddress := common.HexToAddress(input.WalletAddress)

	seedVC := input.SeedPresentation.VerifiableCredential[0]
	splitID := strings.Split(seedVC.ID, "/")
	if len(splitID) != 5 {
		return nil, errval.ErrVCIDFormat
	}
	models.ConvertCredentialSubject(&seedVC)
	seedInfo := seedVC.CredentialSubject.(models.SeedInfo)
	exchangeValue := seedInfo.Amount * placeholderExchangeRate //todo: may have to change calculation method

	txHash, err := contract.TransferTokens(targetAddress, int(exchangeValue)) //todo: need a proper method of converting exchangeValue into an int
	if err != nil {
		return nil, err
	}

	vcID := splitID[4] //should equal numerical ID
	amount := exchangeValue

	exchange := models.NewSeedExchange(vcID, seedInfo.ID, placeholderExchangeRate, amount)

	err = dao.UploadSeedExchange(exchange)
	if err != nil {
		return nil, err
	}

	txTime := strconv.FormatFloat(float64(time.Now().UnixNano())/float64(time.Second), 'f', 3, 64)
	output := models.NewSeedExchangeOutput(exchange.Amount, txHash, txTime, exchange.ExchangeRate)

	return output, nil
}
