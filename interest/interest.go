package interest

import (
	"github.com/metabloxStaking/dao"
	"github.com/metabloxStaking/models"
	logger "github.com/sirupsen/logrus"
	"sort"
	"strconv"
	"time"
)

func CalculateTotalInterest(orderID int) float32 { //placeholder
	return 12.34
}

func calculatePeriodInterest(product *models.StakingProduct) float32 {
	MB := product.TopUpLimit
	R := product.DefaultAPY
	N := float32(product.LockUpPeriod)
	return (MB * R) / (360.0 / N)
}

func calculateMaxAPY(product *models.StakingProduct) float32 {
	A := calculatePeriodInterest(product)
	Z := float32(product.MinOrderValue)
	N := float32(product.LockUpPeriod)
	return (A / Z) * (360.0 / N)
}

func calculateCurrentAPY(product *models.StakingProduct, totalPrincipal float32) float32 {
	A := calculatePeriodInterest(product)
	N := float32(product.LockUpPeriod)
	return (A / totalPrincipal) * (360.0 / N)
}

func GetOrderInterestList(orderID int, until time.Time) ([]*models.OrderInterest, error) {
	targetTime := TruncateToHour(until.In(time.UTC))
	order, err := dao.GetOrderByID(strconv.Itoa(orderID))
	if err != nil {
		return nil, err
	}
	product, err := dao.GetProductInfoByID(strconv.Itoa(order.ProductID))
	if err != nil {
		return nil, err
	}

	principalUpdates, err := dao.GetPrincipalUpdates(order.ProductID)
	if err != nil {
		return nil, err
	}
	principalIndex := 0

	interestList, err := dao.GetSortedOrderInterestList(orderID)
	if err != nil {
		return nil, err
	}

	var latestTime time.Time
	if len(interestList) == 0 {
		startTimeStr, err := dao.GetOrderCreateDate(strconv.Itoa(orderID))
		if err != nil {
			return nil, err
		}
		latestTime, _ = time.Parse("2006-01-02 15:04:05", startTimeStr)
	} else {
		latestOrderInterest := interestList[len(interestList)-1]
		latestTime, err = time.Parse("2006-01-02 15:04:05", latestOrderInterest.Time)
		if err != nil {
			return nil, err
		}
	}

	latestTime = TruncateToHour(latestTime.In(time.UTC))
	if latestTime.Sub(targetTime) >= 0 {
		return interestList, nil
	}

	// generate orderInterest until it reaches the current hour
	interestToAdd := models.NewOrderInterestList()
	for until.Sub(latestTime) > 0 {
		latestTime = latestTime.Add(time.Hour)
		// calculate numbers at given time
		totalPrincipal := float32(0)
		principalIndex = findMostRecentPrincipalUpdate(principalUpdates, latestTime)
		if principalIndex >= 0 {
			totalPrincipal = principalUpdates[principalIndex].TotalPrincipal
		}
		interest, err := calculateOrderInterest(order, product, latestTime, totalPrincipal)
		if err != nil {
			return nil, err
		}
		interestToAdd = append(interestToAdd, interest)
	}
	err = dao.InsertOrderInterestList(interestToAdd)
	if err != nil {
		return nil, err
	}
	interestList = append(interestList, interestToAdd...)
	return interestList, nil
}

func calculateOrderInterest(order *models.Order, product *models.StakingProduct, when time.Time, totalPrincipal float32) (*models.OrderInterest, error) {
	principal, err := dao.GetOrderBuyInPrincipal(strconv.Itoa(order.OrderID))
	if err != nil {
		return nil, err
	}

	interest := models.NewOrderInterest()
	interest.OrderID = order.OrderID
	interest.APY = calculateCurrentAPY(product, totalPrincipal)
	interest.Time = when.Format("2006-01-02 15:04:05")

	N := float32(product.LockUpPeriod)
	interest.InterestGain = (interest.APY / (360.0/N)) * principal * (1/(N*24))
	return interest, nil
}

func findMostRecentPrincipalUpdate(principalUpdates []*models.PrincipalUpdate, now time.Time) int {
	index := sort.Search(len(principalUpdates), func(i int) bool {
		updateTime, _ := time.Parse("2006-01-02 15:04:05",principalUpdates[i].Time)
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
