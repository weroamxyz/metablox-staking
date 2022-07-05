package controllers

import (
	"strconv"
	"time"

	"github.com/MetaBloxIO/metablox-foundation-services/did"
	"github.com/gin-gonic/gin"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/errval"
	"github.com/metabloxStaking/models"
)

//returns all staking records (info about orders and their buy-in transactions) associated with a specified did
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
		record.IsInClosureWindow = 0 < timeElapsed.Hours() && timeElapsed.Hours() < 24

		interestInfo, err := dao.ExecuteGetInterestStmt(record.OrderID, stmt)
		if err != nil {
			return nil, err
		}
		bigInterestGain := interestInfo.AccumulatedInterest.Sub(interestInfo.TotalInterestGained)
		convertedInterestGain := models.MinimumUnitToMBLX(bigInterestGain)
		record.InterestGain = strconv.FormatFloat(convertedInterestGain, 'f', -1, 64)

		bigTotalAmount := bigInterestGain.Add(record.PrincipalAmount)
		convertedTotalAmount := models.MinimumUnitToMBLX(bigTotalAmount)
		record.TotalAmount = strconv.FormatFloat(convertedTotalAmount, 'f', -1, 64)
	}
	return records, nil
}
