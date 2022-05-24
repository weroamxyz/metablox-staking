package interest

import (
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

func GetAllOrderInterest(orderID string, until time.Time) ([]*models.OrderInterest, error) {
	// TODO: maybe update product history?
	targetTime := TruncateToHour(until.In(time.UTC))
	historyList, err := dao.GetProductHistory(orderID)
	if err != nil {
		return nil, err
	}

	orderInterestList := make([]*models.OrderInterest, 0)
	for _, productHistory := range historyList {
		product, err := dao.GetProductInfoByID(productHistory.ProductID)
		if err != nil {
			return nil, err
		}
		startDate, _ := time.Parse("2006-01-02 15:04:05", productHistory.CreateDate)
		productStartDate, _ := time.Parse("2006-01-02 15:04:05", product.StartDate)
		productStartDate = TruncateToHour(productStartDate.In(time.UTC))
		endDate := productStartDate.Add(time.Hour * 24 * time.Duration(product.LockUpPeriod))
		if targetTime.Before(endDate) {
			endDate = targetTime
		}

		interestToAdd, err := GetOrderInterestList(orderID, product, startDate, endDate)
		if err != nil {
			return nil, err
		}
		orderInterestList = append(orderInterestList, interestToAdd...)
	}
	return orderInterestList, nil
}

func GetOrderInterestList(orderID string, product *models.StakingProduct, from time.Time, until time.Time) ([]*models.OrderInterest, error) {
	order, err := dao.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	principalUpdates, err := dao.GetPrincipalUpdates(product.ID)
	if err != nil {
		return nil, err
	}
	principalIndex := 0

	interestList, err := dao.GetSortedOrderInterestListUntilDate(orderID, until.Format("2006-01-02 15:04:05"))
	if err != nil {
		return nil, err
	}

	var latestTime time.Time
	if len(interestList) == 0 {
		latestTime = from
	} else {
		latestOrderInterest := interestList[len(interestList)-1]
		latestTime, err = time.Parse("2006-01-02 15:04:05", latestOrderInterest.Time)
		if err != nil {
			return nil, err
		}
	}

	latestTime = TruncateToHour(latestTime.In(time.UTC))
	if latestTime.Sub(until) >= 0 {
		return interestList, nil
	}

	// generate orderInterest until it reaches the current hour
	interestToAdd := models.NewOrderInterestList()
	for until.After(latestTime) {
		latestTime = latestTime.Add(time.Hour)
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
	}
	err = dao.InsertOrderInterestList(interestToAdd)
	if err != nil {
		return nil, err
	}
	interestList = append(interestList, interestToAdd...)

	// set new accumulated interest value
	sum := 0.0
	for _, interest := range interestList {
		sum += interest.InterestGain
	}
	err = dao.UpdateOrderAccumulatedInterest(orderID, sum)
	if err != nil {
		return nil, err
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
