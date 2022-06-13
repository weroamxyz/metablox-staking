package controllers

import (
	"math/big"
	"strconv"
	"time"

	"github.com/MetaBloxIO/metablox-foundation-services/did"
	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/errval"
	"github.com/metabloxStaking/interest"
	"github.com/metabloxStaking/models"
)

func GetStakingRecords(c *gin.Context) ([]*models.StakingRecord, error) {
	userDID := c.Param("did")
	valid := did.IsDIDValid(did.SplitDIDString(userDID))
	if !valid {
		return nil, errval.ErrBadDID
	}

	records, err := dao.GetStakingRecords(userDID)
	if err != nil {
		return nil, err
	}

	stmt, err := dao.PrepareGetInterestByOrderID()
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	for _, record := range records {

		purchaseTime, err := time.Parse("2006-01-02 15:04:05", record.PurchaseTime)
		if err != nil {
			return nil, err
		}
		record.PurchaseTime = strconv.FormatFloat(float64(purchaseTime.UnixNano())/float64(time.Second), 'f', 3, 64)

		redeemDate, err := time.Parse("2006-01-02 15:04:05", record.RedeemableTime)
		if err != nil {
			return nil, err
		}
		record.RedeemableTime = strconv.FormatFloat(float64(redeemDate.UnixNano())/float64(time.Second), 'f', 3, 64)

		timeElapsed := time.Since(redeemDate)
		record.IsInClosureWindow = (0 < timeElapsed.Hours() && timeElapsed.Hours() < 24)

		if record.OrderStatus == models.OrderTypeHolding {
			interest.CalculateInterest() //query Colin's code to update interest value in db
		}

		interestInfo, err := dao.ExecuteGetInterestStmt(record.OrderID, stmt)
		if err != nil {
			return nil, err
		}
		bigInterestGain := big.NewInt(0).Sub(interestInfo.AccumulatedInterest, interestInfo.TotalInterestGained)
		record.InterestGain = bigInterestGain.String()
		record.TotalAmount = big.NewInt(0).Add(bigInterestGain, record.PrincipalAmount).String()
	}
	return records, nil
}
