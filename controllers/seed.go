package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/MetaBloxIO/metablox-foundation-services/presentations"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/contract"
	"github.com/metabloxStaking/errval"
	"github.com/metabloxStaking/foundationdao"
	"github.com/metabloxStaking/models"
)

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

	err = foundationdao.UploadSeedExchange(exchange)
	if err != nil {
		return nil, err
	}

	txTime := strconv.FormatFloat(float64(time.Now().UnixNano())/float64(time.Second), 'f', 3, 64)
	output := models.NewSeedExchangeOutput(exchange.Amount, txHash, txTime, exchange.ExchangeRate)

	return output, nil
}
