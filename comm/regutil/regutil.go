package regutil

import (
	"github.com/ethereum/go-ethereum/common"
	"regexp"
)

const (
	RegPositiveIntNumber = "^[1-9]\\d*"
)

func IsETHAddress(address string) bool {
	if address == "" {
		return false
	}
	return common.IsHexAddress(address)
}

func IsPositiveIntNumber(s string) bool {
	flag, err := regexp.MatchString(RegPositiveIntNumber, s)
	return err == nil && flag
}
