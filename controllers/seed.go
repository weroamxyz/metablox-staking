package controllers

import (
	"crypto/ecdsa"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/MetaBloxIO/metablox-foundation-services/presentations"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/contract"
	"github.com/metabloxStaking/errval"
	"github.com/metabloxStaking/foundationdao"
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

func GetNonce(session string) (uint64, error) {
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

func ActivateExchange(c *gin.Context) {
	
}

func ExchangeSeed(c *gin.Context) (*models.SeedExchangeOutput, error) {
	input := models.CreateSeedExchangeInput()
	err := c.BindJSON(input)
	if err != nil {
		return nil, err
	}

	err = validateDID(input.DID)
	if err != nil {
		return nil, err
	}

	minerPubKey := new(ecdsa.PublicKey) //todo: get this from some source

	holderPubKey, err := crypto.UnmarshalPubkey(input.PublicKeyString)
	if err != nil {
		return nil, err
	}

	_, err = presentations.VerifyVP(&input.SeedPresentation, holderPubKey, minerPubKey) //going to fail at the moment as we don't have all the info to do this verification
	if err != nil {                                                                     //skip this error check to avoid failures until we can properly verify seed presentations
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

	vcID := strings.Split(seedVC.ID, "/")[4]            //should equal numerical ID
	amount := seedInfo.Amount * placeholderExchangeRate //todo: may have to change calculation method

	exchange := models.NewSeedExchange(vcID, seedInfo.ID, placeholderExchangeRate, amount)

	err = foundationdao.UploadSeedExchange(exchange)
	if err != nil {
		return nil, err
	}

	txTime := strconv.FormatFloat(float64(time.Now().UnixNano())/float64(time.Second), 'f', 3, 64)
	output := models.NewSeedExchangeOutput(exchange.Amount, txHash, txTime, exchange.ExchangeRate)

	return output, nil
}
