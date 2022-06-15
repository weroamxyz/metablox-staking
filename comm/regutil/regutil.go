package regutil

import (
	"regexp"
)

const (
	RegETHAddress        = "^0[xX][0-9a-zA-Z]{40}$"
	RegPositiveIntNumber = "^[1-9]\\d*"
)

func IsETHAddress(address string) bool {
	if address == "" {
		return false
	}
	flag, err := regexp.MatchString(RegETHAddress, address)
	return err == nil && flag
}

func IsPositiveIntNumber(s string) bool {
	flag, err := regexp.MatchString(RegPositiveIntNumber, s)
	return err == nil && flag
}
