package interest

import (
	"database/sql"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/models"
	logger "github.com/sirupsen/logrus"
	"sort"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

func calculatePeriodInterest(product *models.StakingProduct) float64 {
	MB := product.TopUpLimit
	R := product.DefaultAPY
	N := float64(product.LockUpPeriod)
	return (MB * R) / (360.0 / N)
}

func calculateMaxAPY(product *models.StakingProduct) float64 {
	A := calculatePeriodInterest(product)
	Z := float64(product.MinOrderValue)
	N := float64(product.LockUpPeriod)
	return (A / Z) * (360.0 / N)
}

func CalculateCurrentAPY(product *models.StakingProduct, totalPrincipal float64) float64 {
	A := calculatePeriodInterest(product)
	N := float64(product.LockUpPeriod)
	return (A / totalPrincipal) * (360.0 / N)
}

func UpdateOrderInterest(orderID string, product *models.StakingProduct, principalUpdates []*models.PrincipalUpdate, until time.Time) error {
	targetTime := TruncateToHour(until.UTC())
	orderPrincipal, err := dao.GetOrderBuyInPrincipal(orderID)
	if err != nil {
		return err
	}

	// latestTime will be set to either the latest orderInterest + 1 hour, or order date + 1 hour if there is no orderInterest yet
	var latestTime time.Time
	var latestSum float64

	latestInterest, err := dao.GetMostRecentOrderInterestUntilDate(orderID, targetTime.Format(TimeFormat))
	if err == nil {
		latestTime, _ = time.Parse(TimeFormat, latestInterest.Time)
		latestSum = latestInterest.TotalInterestGain
	} else if err == sql.ErrNoRows {
		orderCreateDateStr, err := dao.GetOrderCreateDate(orderID)
		if err != nil {
			return err
		}
		latestTime, _ = time.Parse(TimeFormat, orderCreateDateStr)
		latestSum = 0.0
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
		totalPrincipal := 0.0
		principalIndex = findMostRecentPrincipalUpdate(principalUpdates, latestTime)
		if principalIndex >= 0 {
			totalPrincipal = principalUpdates[principalIndex].TotalPrincipal
		}
		interest := calculateOrderInterest(orderID, orderPrincipal, product, latestTime, totalPrincipal)

		interestToAdd = append(interestToAdd, interest)
		latestTime = latestTime.Add(time.Hour)
	}

	if len(interestToAdd) > 0 {
		for _, interest := range interestToAdd {
			latestSum += interest.InterestGain
			interest.TotalInterestGain = latestSum
		}

		err = dao.InsertOrderInterestList(interestToAdd)
		if err != nil {
			return err
		}

		err = dao.UpdateOrderAccumulatedInterest(orderID, latestSum)
		if err != nil {
			return err
		}
	}
	return nil
}

func calculateOrderInterest(orderID string, orderPrincipal float64, product *models.StakingProduct, when time.Time, totalPrincipal float64) *models.OrderInterest {
	interest := models.CreateOrderInterest()
	interest.OrderID = orderID
	interest.APY = CalculateCurrentAPY(product, totalPrincipal)
	interest.Time = when.Format(TimeFormat)

	N := float64(product.LockUpPeriod)
	interest.InterestGain = (interest.APY / (360.0 / N)) * orderPrincipal * (1 / (N * 24))
	return interest
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
			updateAllOrderInterest(currentTime)

			nextHour := currentTime.Add(time.Hour)
			timer := time.NewTimer(nextHour.Sub(time.Now()))
			<-timer.C
		}
	}()
	return
}

func updateAllOrderInterest(currentTime time.Time) {
	products, err := dao.GetAllProductInfo()
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
			err = dao.InsertPrincipalUpdate(*product.NextProductID, principalUpdate.TotalPrincipal)
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
			err = UpdateOrderInterest(id, product, principalUpdates, currentTime)
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
