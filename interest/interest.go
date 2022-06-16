package interest

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"math/big"
	"sort"
	"time"

	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/models"
	logger "github.com/sirupsen/logrus"
)

const TimeFormat = "2006-01-02 15:04:05"

func calculatePeriodInterest(product *models.StakingProduct) *big.Float {
	MB := new(big.Float).SetInt(product.TopUpLimit)
	R := big.NewFloat(product.DefaultAPY)
	N := float64(product.LockUpPeriod)
	// result = (MB * R) / (360.0 / N)
	result := MB.Mul(MB, R)
	return result.Quo(result, big.NewFloat(360.0/N))
}

func calculateMaxAPY(product *models.StakingProduct) *big.Float {
	A := calculatePeriodInterest(product)
	Z := big.NewFloat(float64(product.MinOrderValue))
	N := float64(product.LockUpPeriod)
	// result = (A / Z) * (360.0 / N)
	result := A.Quo(A, Z)
	return result.Mul(result, big.NewFloat(360.0/N))
}

func CalculateCurrentAPY(product *models.StakingProduct, totalPrincipal *big.Int) *big.Float {
	A := calculatePeriodInterest(product)
	N := float64(product.LockUpPeriod)
	TP := new(big.Float).SetInt(totalPrincipal)
	// result = (A / TP) * (360.0 / N)
	result := A.Quo(A, TP)
	return result.Mul(result, big.NewFloat(360.0/N))
}

func updateOrderInterest(orderID string, product *models.StakingProduct, principalUpdates []*models.PrincipalUpdate,
	until time.Time, getPrincipalStmt *sqlx.Stmt, getRecentInterestStmt *sqlx.Stmt) error {

	targetTime := TruncateToHour(until.UTC())
	orderPrincipal, err := dao.ExecuteGetOrderBuyInPrincipal(getPrincipalStmt, orderID)
	if err != nil {
		return err
	}

	// latestTime will be set to either the latest orderInterest + 1 hour, or order date + 1 hour if there is no orderInterest yet
	var latestTime time.Time
	var latestSum *big.Int

	latestInterest, err := dao.ExecuteGetMostRecentOrderInterestUntilDate(getRecentInterestStmt, orderID, targetTime.Format(TimeFormat))
	if err == nil {
		latestTime, _ = time.Parse(TimeFormat, latestInterest.Time)
		latestSum = latestInterest.TotalInterestGain
	} else if err == sql.ErrNoRows {
		orderCreateDateStr, err := dao.GetOrderCreateDate(orderID)
		if err != nil {
			return err
		}
		latestTime, _ = time.Parse(TimeFormat, orderCreateDateStr)
		latestSum = big.NewInt(0)
	} else {
		return err
	}
	latestTime = TruncateToHour(latestTime.UTC())
	latestTime = latestTime.Add(time.Hour)

	// generate orderInterest until it reaches the current hour
	interestToAdd := models.CreateOrderInterestList()
	for !targetTime.Before(latestTime) {
		if isProductExpired(product, latestTime) {
			if product.NextProductID == nil {
				logger.Warn("tried to get nil NextProductID in product ", product.ID)
				break
			}
			err = dao.UpdateOrderNewProductID(orderID, *product.NextProductID)
			if err != nil {
				logger.Warn(err)
				break
			}
			product, err = dao.GetProductInfoByID(*product.NextProductID)
			if err != nil {
				logger.Warn(err)
				break
			}
		}
		// calculate numbers at given time
		principalIndex := 0
		totalPrincipal := big.NewInt(0)
		principalIndex = findMostRecentPrincipalUpdate(principalUpdates, latestTime)
		if principalIndex >= 0 {
			totalPrincipal = principalUpdates[principalIndex].TotalPrincipal
		}
		interest, err := calculateOrderInterest(orderID, orderPrincipal, product, latestTime, totalPrincipal)
		if err != nil {
			return err
		}

		interest.StringInterestGain = interest.InterestGain.String()
		interestToAdd = append(interestToAdd, interest)
		latestTime = latestTime.Add(time.Hour)
	}

	if len(interestToAdd) > 0 {
		for _, interest := range interestToAdd {
			latestSum.Add(latestSum, interest.InterestGain)
			interest.TotalInterestGain = latestSum
			interest.StringTotalInterestGain = latestSum.String()
		}

		err = dao.InsertOrderInterestList(interestToAdd)
		if err != nil {
			return err
		}

		err = dao.UpdateOrderAccumulatedInterest(orderID, latestSum.String())
		if err != nil {
			return err
		}
	}
	return nil
}

func calculateOrderInterest(orderID string, orderPrincipal *big.Int, product *models.StakingProduct, when time.Time, totalPrincipal *big.Int) (*models.OrderInterest, error) {
	CAPY := CalculateCurrentAPY(product, totalPrincipal)
	// interestGain = (interest.APY / (360.0 / N)) * principal * (1 / (N * 24))
	N := float64(product.LockUpPeriod)
	interestGain := new(big.Float).Quo(CAPY, big.NewFloat(360.0/N))
	interestGain.Mul(interestGain, new(big.Float).SetInt(orderPrincipal))
	interestGain.Mul(interestGain, big.NewFloat(1/(N*24)))

	interest := models.CreateOrderInterest()
	interest.OrderID = orderID
	interest.APY, _ = CAPY.Float64()
	interest.Time = when.Format(TimeFormat)
	interest.InterestGain, _ = interestGain.Int(nil)

	return interest, nil
}

func findMostRecentPrincipalUpdate(principalUpdates []*models.PrincipalUpdate, now time.Time) int {
	index := sort.Search(len(principalUpdates), func(i int) bool {
		updateTime, _ := time.Parse(TimeFormat, principalUpdates[i].Time)
		return updateTime.After(now)
	})
	return index - 1 // the last index for which updateTime is before now
}

func TruncateToHour(dateTime time.Time) time.Time {
	result, err := time.Parse("2006-01-02 15", dateTime.Format("2006-01-02 15"))
	if err != nil {
		logger.Warn(err.Error())
	}
	return result
}

func StartHourlyTimer() {
	go func() {
		for {
			currentTime := TruncateToHour(time.Now().UTC())
			UpdateAllOrderInterest(currentTime)

			nextHour := currentTime.Add(time.Hour)
			timer := time.NewTimer(nextHour.Sub(time.Now()))
			<-timer.C
		}
	}()
	return
}

func UpdateAllOrderInterest(currentTime time.Time) {
	products, err := dao.GetAllProductInfo()
	if err != nil {
		logger.Warn(err)
		return
	}
	getRecentInterestStmt, err := dao.PrepareGetMostRecentOrderInterestUntilDate()
	if err != nil {
		logger.Warn(err)
		return
	}
	getPrincipalStmt, err := dao.PrepareGetOrderBuyInPrincipal()
	if err != nil {
		logger.Warn(err)
		return
	}

	for _, product := range products {
		// 1. carry over the expired products' total principal (individual products will update their ids as they calculate)
		if product.Status && isProductExpired(product, currentTime) {
			if product.NextProductID == nil {
				logger.Warn("tried to get nil NextProductID in product ", product.ID)
				continue
			}
			principalUpdate, err := dao.GetLatestPrincipalUpdate(product.ID)
			if err != nil {
				if err != sql.ErrNoRows {
					logger.Warn(err)
				}
				continue
			}
			err = dao.InsertPrincipalUpdate(*product.NextProductID, principalUpdate.TotalPrincipal.String())
			if err != nil {
				logger.Warn(err)
				continue
			}
			err = dao.UpdateProductStatus(product.ID, false)
			if err != nil {
				logger.Warn(err)
			}
		}

		// 2. update order interest for all active orders
		orderIDs, err := dao.GetHoldingOrderIDsForProduct(product.ID)
		if err != nil {
			logger.Warn(err)
			return
		}
		principalUpdates, err := dao.GetPrincipalUpdates(product.ID)
		if err != nil {
			logger.Warn(err)
			return
		}
		for _, id := range orderIDs {
			err = updateOrderInterest(id, product, principalUpdates, currentTime, getPrincipalStmt, getRecentInterestStmt)
			if err != nil {
				logger.Warn(err)
			}
		}
	}
}

func isProductExpired(product *models.StakingProduct, currentTime time.Time) bool {
	startTime, _ := time.Parse(TimeFormat, product.StartDate)
	startTime = TruncateToHour(startTime.UTC())
	endTime := startTime.Add(time.Hour * 24 * time.Duration(product.LockUpPeriod))
	return !currentTime.Before(endTime)
}
