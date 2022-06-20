package contract

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/common"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	"github.com/metabloxStaking/comm/abiutil"
	"github.com/metabloxStaking/contract/erc20"
	"github.com/metabloxStaking/contract/tokenutil"
	"github.com/metabloxStaking/errval"
	"github.com/metabloxStaking/models"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	RpcUrl       = "https://api.s0.ps.hmny.io/"
	TokenAddress = "0x7d9A974db87320Dd92fE74D31e2c4c32b3c4CB0C"
)

func TestCheckIfTransactionMatchesOrder(t *testing.T) {
	txHash := "0xd7147ee886e349e767ed7751ce4f3070479b87eeddac3fd9ae294e3bd702143e"
	_client, _ := ethclient.Dial(RpcUrl)
	order := &models.Order{
		PaymentAddress: "0x0F3f0B1B914FF55988E7168890dcF70CAB378EA6",
		Amount:         decimal.NewFromInt(25300000),
		UserAddress:    "0x3aea49553Ce2E478f1c0c5ACC304a84F5F4d1f98",
	}

	// === Mock start
	patch0 := gomonkey.ApplyFunc(tokenutil.MBLXTokenAddress, func() common.Address {
		return common.HexToAddress(TokenAddress)
	})
	defer patch0.Reset()

	patch1 := gomonkey.ApplyFunc(tokenutil.DecodeData, func(data string) (abiutil.MethodData, error) {
		return abiutil.NewABIDecoder(erc20.Erc20ABI).DecodeMethod(data)
	})
	defer patch1.Reset()

	patch2 := gomonkey.ApplyGlobalVar(&client, _client)
	defer patch2.Reset()
	// === Mock end

	err := CheckIfTransactionMatchesOrder(txHash, order)
	assert.NoError(t, err)

	order.PaymentAddress = common.Address{}.Hex()
	err = CheckIfTransactionMatchesOrder(txHash, order)
	assert.Error(t, errval.ErrWalletAddress, err)
}
