package interest

import (
	"errors"
	"fmt"
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/models"
	logger "github.com/sirupsen/logrus"
	"sort"
	"time"
)

func CalculateInterest() float64 { //placeholder
	return 12.34
}

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

func GetOrderInterestList(orderID string, until time.Time) ([]*models.OrderInterest, error) {
	targetTime := TruncateToHour(until.In(time.UTC))

	order, err := dao.GetOrderByID(orderID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to get order, %s", err.Error()))
	}

	product, err := dao.GetProductInfoByID(order.ProductID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to get product, %s", err.Error()))
	}

	principalUpdates, err := dao.GetPrincipalUpdates(product.ID)
	if err != nil {
		return nil, err
	}
	principalIndex := 0

	interestList, err := dao.GetSortedOrderInterestListUntilDate(orderID, targetTime.Format("2006-01-02 15:04:05"))
	if err != nil {
		return nil, err
	}

	// latestTime will be set to either the latest orderInterest + 1 hour, or order date + 1 hour if there is no orderInterest yet
	var latestTime time.Time
	if len(interestList) == 0 {
		orderCreateDateStr, err := dao.GetOrderCreateDate(orderID)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to get txInfo, %s", err.Error()))
		}
		latestTime, _ = time.Parse("2006-01-02 15:04:05", orderCreateDateStr)
	} else {
		latestOrderInterest := interestList[len(interestList)-1]
		latestTime, _ = time.Parse("2006-01-02 15:04:05", latestOrderInterest.Time)
	}
	latestTime = TruncateToHour(latestTime.In(time.UTC))
	latestTime = latestTime.Add(time.Hour)

	// generate orderInterest until it reaches the current hour
	interestToAdd := models.NewOrderInterestList()
	for targetTime.After(latestTime) {
		if isProductExpired(product, latestTime) {
			if product.NextProductID == nil {
				break
			}
			product, err = dao.GetProductInfoByID(*product.NextProductID)
			if err != nil {
				logger.Warn(err)
				break
			}
		}
		// calculate numbers at given time
		totalPrincipal := 0.0
		principalIndex = findMostRecentPrincipalUpdate(principalUpdates, latestTime)
		if principalIndex >= 0 {
			totalPrincipal = principalUpdates[principalIndex].TotalPrincipal
		}
		interest, err := calculateOrderInterest(order, product, latestTime, totalPrincipal)
		if err != nil {
			return nil, err
		}

		interestToAdd = append(interestToAdd, interest)
		latestTime = latestTime.Add(time.Hour)
	}

	if len(interestToAdd) > 0 {
		sum := 0.0
		for _, interest := range interestList {
			sum += interest.InterestGain
		}
		for _, interest := range interestToAdd {
			sum += interest.InterestGain
			interest.TotalInterestGain = sum
		}

		err = dao.InsertOrderInterestList(interestToAdd)
		if err != nil {
			return nil, err
		}

		err = dao.UpdateOrderAccumulatedInterest(orderID, sum)
		if err != nil {
			return nil, err
		}
		interestList = append(interestList, interestToAdd...)
	}
	return interestList, nil
}

func calculateOrderInterest(order *models.Order, product *models.StakingProduct, when time.Time, totalPrincipal float64) (*models.OrderInterest, error) {
	principal, err := dao.GetOrderBuyInPrincipal(order.OrderID)
	if err != nil {
		return nil, err
	}

	interest := models.NewOrderInterest()
	interest.OrderID = order.OrderID
	interest.APY = CalculateCurrentAPY(product, totalPrincipal)
	interest.Time = when.Format("2006-01-02 15:04:05")

	N := float64(product.LockUpPeriod)
	interest.InterestGain = (interest.APY / (360.0 / N)) * principal * (1 / (N * 24))
	return interest, nil
}

func findMostRecentPrincipalUpdate(principalUpdates []*models.PrincipalUpdate, now time.Time) int {
	index := sort.Search(len(principalUpdates), func(i int) bool {
		updateTime, _ := time.Parse("2006-01-02 15:04:05", principalUpdates[i].Time)
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
			currentTime := TruncateToHour(time.Now().In(time.UTC))
			ok := true

			err := updateExpiredProducts(currentTime)
			if err != nil {
				logger.Warn(err)
				ok = false
			}

			// 2. update order interest for all active orders
			if ok {
				orders, err := dao.GetHoldingOrders()
				if err != nil {
					logger.Warn(err)
				}
				for _, order := range orders {
					_, err = GetOrderInterestList(order.OrderID, currentTime)
					if err != nil {
						logger.Warn(err)
					}
				}
			}

			nextHour := currentTime.Add(time.Hour)
			timer := time.NewTimer(nextHour.Sub(time.Now()))
			<-timer.C
		}
	}()
	return
}

// updateExpiredProducts updates the ProductHistory table, then the ProductID column of associated Orders
func updateExpiredProducts(currentTime time.Time) error {
	productIDs, err := dao.GetActiveOrdersProductIDs()
	if err != nil {
		return err
	}
	for _, productID := range productIDs {
		product, err := dao.GetProductInfoByID(productID)
		if err != nil {
			return err
		}
		if isProductExpired(product, currentTime) {
			err = dao.UpdateActiveOrdersProductID(productID, *product.NextProductID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func isProductExpired(product *models.StakingProduct, currentTime time.Time) bool {
	startTime, _ := time.Parse("2006-01-02 15:04:05", product.StartDate)
	startTime = TruncateToHour(startTime.In(time.UTC))
	endTime := startTime.Add(time.Hour * 24 * time.Duration(product.LockUpPeriod))
	return !currentTime.Before(endTime)
}
