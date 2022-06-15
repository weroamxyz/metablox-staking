package abiutil

import (
	"encoding/hex"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Param struct {
	Name  string
	Value string
	Type  string
}
type MethodData struct {
	Name   string
	Params []Param
}

// ABIDecoder ethereum transaction data decoder
type ABIDecoder struct {
	abi.ABI
}

func NewABIDecoder(abiStr string) *ABIDecoder {
	abi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		logger.Fatal(err)
	}
	return &ABIDecoder{
		abi,
	}
}

func (d *ABIDecoder) DecodeMethod(txData string) (MethodData, error) {
	if strings.HasPrefix(txData, "0x") {
		txData = txData[2:]
	}

	decodedSig, err := hex.DecodeString(txData[:8])
	if err != nil {
		return MethodData{}, err
	}

	method, err := d.MethodById(decodedSig)
	if err != nil {
		return MethodData{}, err
	}

	decodedData, err := hex.DecodeString(txData[8:])
	if err != nil {
		return MethodData{}, err
	}

	inputs, err := method.Inputs.Unpack(decodedData)
	if err != nil {
		return MethodData{}, err
	}

	nonIndexedArgs := method.Inputs.NonIndexed()

	retData := MethodData{}
	retData.Name = method.Name
	for i, input := range inputs {
		arg := nonIndexedArgs[i]
		param := Param{
			Name:  arg.Name,
			Value: fmt.Sprintf("%v", input),
			Type:  arg.Type.String(),
		}
		retData.Params = append(retData.Params, param)
	}

	return retData, nil
}

func (d *ABIDecoder) GetABI() abi.ABI {
	return d.ABI
}
