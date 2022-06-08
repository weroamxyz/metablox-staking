package tokenutil

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metabloxStaking/contract/erc20"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/big"
)

const DefaultGasLimit = 100000

var (
	client       *ethclient.Client
	tokenAddress common.Address
	wallet       *ecdsa.PrivateKey
	fromAddress  common.Address
	chainId      *big.Int
)

func Init() {
	logger.Debugln("Initializing tokenutil...")

	rpcUrl := viper.GetString("metablox.rpcUrl")
	tokenStr := viper.GetString("metablox.tokenAddress")
	walletStr := viper.GetString("metablox.walletPrivateKey")

	rpc, err := ethclient.Dial(rpcUrl)
	if err != nil {
		logger.Panicf("connecting to rpc node failed")
	}
	client = rpc

	id, err := rpc.ChainID(context.Background())
	if err != nil {
		logger.Panicf("query current chainID failed")
	}
	chainId = id

	tokenAddress = common.HexToAddress(tokenStr)
	wallet, _ = crypto.HexToECDSA(walletStr)

	fromAddress = crypto.PubkeyToAddress(wallet.PublicKey)

	logger.Debugln("tokenutil initialized completed")

}

func NewToken() (*erc20.Erc20, error) {
	token, err := erc20.NewErc20(tokenAddress, client)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func NewSigner(gasLimit uint64) (*bind.TransactOpts, error) {
	ctx := context.Background()
	signer, err := bind.NewKeyedTransactorWithChainID(wallet, chainId)
	if err != nil {
		return nil, err
	}
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}
	price, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	signer.Nonce = new(big.Int).SetUint64(nonce)
	signer.Value = new(big.Int)
	signer.GasPrice = price
	signer.GasLimit = gasLimit
	return signer, nil
}

func EthPendingBalance(address common.Address) (*big.Int, error) {
	return client.PendingBalanceAt(context.Background(), address)
}

func BalanceOf(address common.Address) (*big.Int, error) {
	token, err := NewToken()
	if err != nil {
		return nil, err
	}
	return token.BalanceOf(&bind.CallOpts{}, address)
}

func Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	// 1.build token instance
	token, err := NewToken()
	if err != nil {
		return nil, err
	}
	// 2. check token balance
	balance, _ := BalanceOf(fromAddress)
	if balance.Cmp(amount) < 0 {
		return nil, errors.New("balance of the admin wallet is not enough")
	}
	// 3. build signer instance
	signer, err := NewSigner(DefaultGasLimit)
	if err != nil {
		return nil, err
	}
	// 4. check eth Balance
	ethBalance, _ := EthPendingBalance(fromAddress)
	if ethBalance.Cmp(amount) < 0 {
		return nil, errors.New("token balance is not enough")
	}
	if !checkEthBalance(ethBalance, signer.GasPrice, signer.GasLimit) {
		return nil, errors.New("eth balance is not enough")
	}
	// 5. return raw response
	return token.Transfer(signer, to, amount)
}

func checkEthBalance(balance *big.Int, gasPrice *big.Int, gasLimit uint64) bool {
	return balance.Cmp(new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(gasLimit))) > 0
}
