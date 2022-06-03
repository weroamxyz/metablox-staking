package contract

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/MetaBloxIO/metablox-foundation-services/registry"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/errval"
	"github.com/metabloxStaking/models"
	CappedCashier "github.com/metabloxStaking/stakingContract"
	logger "github.com/sirupsen/logrus"
)

const deployedCashierContract = "0x3F4b2552e04c0EeDd1054272F46Bf90cAA143d51"

const deployedRegistryContract = "0xf880b97Be7c402Cc441895bF397c3f865BfE1Cb2"
const network = "wss://ws.s0.b.hmny.io"

var client *ethclient.Client

var cashierInstance *CappedCashier.CappedCashier
var registryInstance *registry.Registry

var ownerKey *ecdsa.PrivateKey

var domainSeparator []byte

func Init() error {
	var err error
	client, err = ethclient.Dial(network)
	if err != nil {
		return err
	}
	cashierContractAddress := common.HexToAddress(deployedCashierContract)
	cashierInstance, err = CappedCashier.NewCappedCashier(cashierContractAddress, client)
	if err != nil {
		return err
	}

	registryContractAddress := common.HexToAddress(deployedRegistryContract)
	registryInstance, err = registry.NewRegistry(registryContractAddress, client)
	if err != nil {
		return err
	}

	ownerKey, _ = crypto.HexToECDSA("9fd8f6049129527c63830aead266bcf7c53aa82109422da9335aac0c0a36a968")

	typeHash := crypto.Keccak256([]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"))
	nameHash := crypto.Keccak256([]byte("Metablox"))
	verHash := crypto.Keccak256([]byte("1"))
	chainID := big.NewInt(1666700000)
	domainSeparator = crypto.Keccak256(typeHash, nameHash, verHash, chainID.Bytes(), cashierContractAddress.Bytes())

	go EventListener()
	return nil
}

func CheckForRegisteredDID(did string) error {
	didAccount, err := registryInstance.Dids(nil, did)
	if err != nil {
		return err
	}
	if didAccount.Hex() == "0x0000000000000000000000000000000000000000" {
		return errval.ErrDIDNotRegistered
	}
	return nil
}

func generateAuth(privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1666700000))
	if err != nil {
		return nil, err
	}
	authNonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(authNonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth, nil
}

func TransferTokens(toAddress common.Address, value int) (string, error) {
	//balance, err := instance.TokenBalance(nil)
	//if err != nil {
	//	return err
	//}

	/*bigValue := big.NewInt(int64(value))

	auth, err := generateAuth(ownerKey)
	if err != nil {
		return "", err
	}

	tx, err := cashierInstance.Stake(auth, bigValue)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil*/
	return "", nil
}

func Stake(value int) (string, error) {
	bigValue := big.NewInt(int64(value))

	auth, err := generateAuth(ownerKey)
	if err != nil {
		return "", err
	}

	tx, err := cashierInstance.Stake(auth, bigValue)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func CheckIfTransactionMatchesOrder(txHash string, order *models.Order) error {
	tx, pending, err := client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		return err
	}
	if pending {
		return errval.ErrTransactionPending
	}

	msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()), big.NewInt(0))
	if err != nil {
		return err
	}

	if msg.From().Hex() != order.UserAddress {
		return errval.ErrAddressComparisonFail
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		return err
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(CappedCashier.CappedCashierABI)))
	if err != nil {
		return err
	}

	events, err := contractAbi.Unpack("Staked", receipt.Logs[0].Data)
	if err != nil {
		return err
	}

	amount := events[0].(*big.Int)
	if amount.Int64() != int64(order.Amount) { //todo: proper type conversion
		return errval.ErrAmountComparisonFail
	}

	return nil
}

func RedeemOrder() string { //todo: full implementation
	return "placeholderHash"
}

func RedeemInterest(amount, deadline int) (*models.RedemptionSignature, error) { //todo: make sure that the signature is valid, has not been properly tested
	hashedPermit := crypto.Keccak256([]byte("Permit(address staker,uint256 value,uint256 nonce,uint256 deadline)"))
	senderAddress := crypto.PubkeyToAddress(ownerKey.PublicKey)
	fmt.Println(senderAddress)
	nonce, err := cashierInstance.Nonces(nil, senderAddress)
	if err != nil {
		return nil, err
	}
	bigAmount := big.NewInt(int64(amount))
	bigDeadline := big.NewInt(int64(deadline))

	messageBytes := bytes.Join([][]byte{hashedPermit, senderAddress[:], bigAmount.Bytes(), nonce.Bytes(), bigDeadline.Bytes()}, nil)

	structHash := crypto.Keccak256(messageBytes)

	v4MessageBytes := bytes.Join([][]byte{[]byte("\x19\x01"), domainSeparator, structHash}, nil)
	v4Hash := crypto.Keccak256(v4MessageBytes)

	finalMessageBytes := bytes.Join([][]byte{[]byte("\x19Ethereum Signed Message:\n32"), v4Hash}, nil)
	finalHash := crypto.Keccak256(finalMessageBytes)
	signature, err := crypto.Sign(finalHash[:], ownerKey)
	if err != nil {
		return nil, err
	}
	var r [32]byte
	var s [32]byte
	var v uint8

	copy(r[:], signature[:32])
	copy(s[:], signature[32:64])
	v = signature[64] + 27 //have to increment this manually as the smart contract expects v to be 27 or 28, while the crypto package generates it as 0 or 1

	output := models.NewRedemptionSignature(amount, deadline, int(v), r, s)

	return output, nil
}

func EventListener() {

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(deployedCashierContract)},
	}
	logs := make(chan types.Log)

	contractAbi, err := abi.JSON(strings.NewReader(string(CappedCashier.CappedCashierABI)))
	if err != nil {
		logger.Error(err)
	}

	stakedSignature := crypto.Keccak256Hash([]byte("Staked(address,uint256)"))

RestartSubscription:
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		logger.Error(err)
	}

	for {
		select {
		case _ = <-sub.Err():
			goto RestartSubscription //if connection is broken, reconnect
		case vLog := <-logs:
			switch vLog.Topics[0] {
			case stakedSignature:
				events, err := contractAbi.Unpack("Staked", vLog.Data)
				if err != nil {
					logger.Error(err)
				}

				block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
				if err != nil {
					logger.Error(err)
				}
				eventStruct := models.NewStakeEvent(vLog.Topics[1].Hex(), int(events[0].(*big.Int).Int64()), block.Time())
				err = dao.UploadStakeEvent(eventStruct)
				if err != nil {
					logger.Error(err)
				}
			default:
				fmt.Println("unrecognized signature: ", vLog.Topics[0])
			}
		}
	}
}
