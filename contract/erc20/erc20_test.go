package erc20

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metabloxStaking/comm"
	logger "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var (
	rpcUrl = "https://api.s0.b.hmny.io"
	// MBLX contract (testnet)
	tokenAddress = common.HexToAddress("0x7aEC610a4a3287B7b210bC04bDA781B2afd7538c")
	wallet, _    = crypto.HexToECDSA("9fd8f6049129527c63830aead266bcf7c53aa82109422da9335aac0c0a36a968")
	chainId      = big.NewInt(1666700000)
)

func TestErc20(t *testing.T) {
	// 1.start connecting...
	client, err := ethclient.Dial(rpcUrl)
	defer client.Close()
	if err != nil {
		logger.Panic("rpc connect failed")
		return
	}
	// 2.build token instance
	token, err := NewErc20(tokenAddress, client)
	if err != nil {
		logger.Panic("contract bind failed")
		return
	}
	// 3. check wallet address
	fromAddress := crypto.PubkeyToAddress(wallet.PublicKey)
	// 4.build signer instance
	ctx := context.Background()
	signer, _ := bind.NewKeyedTransactorWithChainID(wallet, chainId)
	nonce, _ := client.PendingNonceAt(ctx, fromAddress)
	price, _ := client.SuggestGasPrice(ctx)
	signer.Nonce = new(big.Int).SetUint64(nonce)
	signer.Value = new(big.Int)
	signer.GasPrice = price
	signer.GasLimit = 100000

	// 5.transfer token to user
	tx, err := token.Transfer(signer, common.HexToAddress("0x58883f2b3F59EaA85aBa8D33b7f97b6f8c7f495e"), comm.ToWei(100000, comm.Ether))
	fmt.Println(comm.ToWei(1000000, comm.Ether))
	if err != nil {
		logger.Error("send tx failed: ", err)
		return
	}
	logger.Infof("交易hash=%s\n", tx.Hash().Hex())
	// 6. check contact call
	symbol, _ := token.Symbol(&bind.CallOpts{})
	assert.Equal(t, "MBLX", symbol, "token symbol error")
}

func TestQueryTxReceipt(t *testing.T) {
	client, err := ethclient.Dial(rpcUrl)
	defer client.Close()
	if err != nil {
		logger.Panic("rpc connect failed")
		return
	}
	hash := common.HexToHash("0xa52be23b5e73c5e0c4514cc790febc5d54aea66d211fd57b0ff3e78e24f300bd")
	ctx := context.Background()
	receipt, _ := client.TransactionReceipt(ctx, hash)
	assert.Equal(t, receipt.Status, uint64(1))
	b, _ := receipt.MarshalJSON()
	logger.Infoln("tx detail: " + string(b))

}
