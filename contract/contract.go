package contract

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metabloxStaking/stakingContract"
)

const deployedContract = "0xc70A4185af369cfF34507Fe14b651fbEe53fed88"
const network = "wss://ws.s0.b.hmny.io"

var client *ethclient.Client

var instance *stakingContract.StakingContract
var contractAddress common.Address

var ownerKey *ecdsa.PrivateKey

func Init() error {
	var err error
	client, err = ethclient.Dial(network)
	if err != nil {
		return err
	}
	contractAddress = common.HexToAddress(deployedContract)
	instance, err = stakingContract.NewStakingContract(contractAddress, client)
	if err != nil {
		return err
	}

	ownerKey, _ = crypto.HexToECDSA("dbbd9634560466ac9713e0cf10a575456c8b55388bce0c044f33fc6074dc5ae6")

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

	tx, err := instance.Transfer(auth, toAddress, bigValue)
	if err != nil {
		return "", nil
	}

	return tx.Hash().Hex(), nil
}

func CheckIfTransactionCompleted(txHash string) (bool, error) { //todo: full implementation
	return true, nil
}

func RedeemOrder() string { //todo: full implementation
	return "placeholderHash"
}

func RedeemInterest() string { //todo: full implementation
	return "placeholderHash"
}
