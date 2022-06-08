package comm

import (
	logger "github.com/sirupsen/logrus"
	"math/big"
)

type Unit int64

const (
	Wei   Unit = 1
	KWei  Unit = 1e3
	MWei  Unit = 1e6
	GWei  Unit = 1e9
	Ether Unit = 1e18
)

func (c Unit) IntValue() *big.Int {
	return big.NewInt(int64(c))
}

func (c Unit) FloatValue() *big.Float {
	return new(big.Float).SetInt64(int64(c))
}

// FromWei unit conversion
func FromWei(number *big.Int, unit Unit) *big.Float {
	var value = new(big.Float)
	switch unit {
	case Wei:
		value = value.SetInt(number)
	case KWei:
		value = value.SetInt(number).Quo(value, KWei.FloatValue())
	case MWei:
		value = value.SetInt(number).Quo(value, MWei.FloatValue())
	case GWei:
		value = value.SetInt(number).Quo(value, GWei.FloatValue())
	case Ether:
		value = value.SetInt(number).Quo(value, Ether.FloatValue())
	default:
		logger.Panicf("illegal unit param")
	}
	return value
}

// ToWei unit conversion
func ToWei(number int64, unit Unit) *big.Int {
	var value = new(big.Int)
	switch unit {
	case Wei:
		value = value.SetInt64(number)
	case KWei:
		value = value.SetInt64(number).Mul(value, KWei.IntValue())
	case MWei:
		value = value.SetInt64(number).Mul(value, MWei.IntValue())
	case GWei:
		value = value.SetInt64(number).Mul(value, GWei.IntValue())
	case Ether:
		value = value.SetInt64(number).Mul(value, Ether.IntValue())
	default:
		logger.Panicf("illegal unit param")
	}
	return value
}
