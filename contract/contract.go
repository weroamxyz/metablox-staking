package contract

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/metabloxStaking/comm/regutil"
	"github.com/metabloxStaking/contract/tokenutil"

	"github.com/MetaBloxIO/metablox-foundation-services/registry"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metabloxStaking/errval"
	"github.com/metabloxStaking/models"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	registryContract common.Address
	rpcUrl           string
	client           *ethclient.Client
	registryInstance *registry.Registry
)

const transferMethodName = "Transfer"

func Init() error {
	var err error
	rpcUrl = viper.GetString("metablox.rpcUrl")
	registryStr := viper.GetString("metablox.registryContract")
	client, err = ethclient.Dial(rpcUrl)
	if err != nil {
		return err
	}
	registryContract = common.HexToAddress(registryStr)
	registryInstance, err = registry.NewRegistry(registryContract, client)
	if err != nil {
		return err
	}
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

	if msg.From() != common.HexToAddress(order.UserAddress) || *msg.To() != tokenutil.MBLXTokenAddress() {
		return errval.ErrAddressComparisonFail
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		return err
	}

	if receipt.Status != uint64(1) {
		return errval.ErrTransactionReverted
	}

	method, err := tokenutil.DecodeData(hexutil.Encode(tx.Data()))
	if err != nil {
		logger.Warn("parse contract input error: ", err.Error())
		return errval.ErrContractData
	}

	if method.Name != transferMethodName {
		return errval.ErrContractMethod
	}

	params := method.Params
	if params == nil || len(params) != 2 {
		return errval.ErrContractParam
	}

	toAddress := params[0].Value
	value := params[1].Value
	if !regutil.IsETHAddress(toAddress) || toAddress != order.PaymentAddress {
		return errval.ErrWalletAddress
	}

	if !regutil.IsPositiveIntNumber(value) {
		return errval.ErrContractParam
	}

	amount, ok := new(big.Int).SetString(value, 0)
	if !ok {
		return errors.New(value + " is not a correct amount")
	}

	result, _ := new(big.Float).SetInt(order.Amount).Int(nil)
	if amount.Cmp(result) < 0 {
		return errors.New(value + " is not enough")
	}

	//amount := events[0].(*big.Int)
	//conversionRate := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil) //db stores values in MBLX, need to convert to minimum units
	//fltConversionRate := new(big.Float).SetInt(conversionRate)
	//dbAmount := fltConversionRate.Mul(fltConversionRate, big.NewFloat(1.3))
	//intDBAmount, accuracy := dbAmount.Int(nil)
	//
	//if intDBAmount.Sub(intDBAmount, amount).Int64()*int64(accuracy) > 1000 { //need some margin of error to account for inaccuracy when converting between big.Float and big.Int
	//	return errval.ErrAmountComparisonFail
	//}

	return nil
}

func RedeemOrder(addressStr string, amount *big.Int) (*types.Transaction, error) {
	// verify eth address
	if !regutil.IsETHAddress(addressStr) {
		return nil, errval.ErrETHAddress
	}
	address := common.HexToAddress(addressStr)
	return tokenutil.Transfer(address, amount)
}
