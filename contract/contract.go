package contract

import (
	"context"
	"crypto/ecdsa"
	"github.com/metabloxStaking/comm/regutil"
	"github.com/metabloxStaking/contract/tokenutil"
	"math/big"
	"strings"

	"github.com/MetaBloxIO/metablox-foundation-services/registry"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metabloxStaking/errval"
	"github.com/metabloxStaking/models"
	"github.com/metabloxStaking/stakingContract"
)

const deployedStakingContract = "0xc70A4185af369cfF34507Fe14b651fbEe53fed88"
const deployedRegistryContract = "0xf880b97Be7c402Cc441895bF397c3f865BfE1Cb2"
const network = "wss://ws.s0.b.hmny.io"

var client *ethclient.Client

var stakingInstance *stakingContract.StakingContract
var registryInstance *registry.Registry

var ownerKey *ecdsa.PrivateKey

func Init() error {
	var err error
	client, err = ethclient.Dial(network)
	if err != nil {
		return err
	}
	stakingContractAddress := common.HexToAddress(deployedStakingContract)
	stakingInstance, err = stakingContract.NewStakingContract(stakingContractAddress, client)
	if err != nil {
		return err
	}

	registryContractAddress := common.HexToAddress(deployedRegistryContract)
	registryInstance, err = registry.NewRegistry(registryContractAddress, client)
	if err != nil {
		return err
	}

	ownerKey, _ = crypto.HexToECDSA("dbbd9634560466ac9713e0cf10a575456c8b55388bce0c044f33fc6074dc5ae6")

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

	bigValue := big.NewInt(int64(value))

	auth, err := generateAuth(ownerKey)
	if err != nil {
		return "", err
	}

	tx, err := stakingInstance.Transfer(auth, toAddress, bigValue)
	if err != nil {
		return "", nil
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

	contractAbi, err := abi.JSON(strings.NewReader(string(stakingContract.StakingContractABI)))
	if err != nil {
		return err
	}

	events, err := contractAbi.Unpack("Transfer", receipt.Logs[0].Data)
	if err != nil {
		return err
	}

	amount := events[0].(*big.Int)
	conversionRate := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil) //db stores values in MBLX, need to convert to minimum units
	fltConversionRate := new(big.Float).SetInt(conversionRate)
	dbAmount := fltConversionRate.Mul(fltConversionRate, big.NewFloat(1.3))
	intDBAmount, accuracy := dbAmount.Int(nil)

	if intDBAmount.Sub(intDBAmount, amount).Int64()*int64(accuracy) > 1000 { //need some margin of error to account for inaccuracy when converting between big.Float and big.Int
		return errval.ErrAmountComparisonFail
	}

	return nil
}

func RedeemOrder(addressStr string, amountF float64) (*types.Transaction, error) {
	// verify eth address
	if !regutil.IsETHAddress(addressStr) {
		return nil, errval.ErrETHAddress
	}
	address := common.HexToAddress(addressStr)
	amount := new(big.Int).SetUint64(uint64(amountF))
	return tokenutil.Transfer(address, amount)
}
