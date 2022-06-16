package models

import (
	"math/big"
	"time"

	foundationModels "github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/metabloxStaking/errval"
)

const TxTypeBuyIn = "BuyIn"
const TxTypeInterestOnly = "InterestOnly"
const TxTypeOrderClosure = "OrderClosure"

const OrderTypePending = "Pending"
const OrderTypeHolding = "Holding"
const OrderTypeComplete = "Complete"

const CurrencyTypeMBLX = "MBLX"

const MinimumUnitToMBLX = 1000000

type Order struct {
	OrderID                   string `db:"OrderID"`
	ProductID                 string `db:"ProductID"`
	UserDID                   string `db:"UserDID"`
	Type                      string `db:"Type" validate:"required,oneof=Pending Holding Complete"`
	Term                      int    `db:"Term"`
	AccumulatedInterest       *big.Int
	StringAccumulatedInterest string `db:"AccumulatedInterest" json:"-"`
	TotalInterestGained       *big.Int
	StringTotalInterestGained string `db:"TotalInterestGained" json:"-"`
	PaymentAddress            string `db:"PaymentAddress"`
	Amount                    *big.Int
	StringAmount              string `db:"Amount" json:"-"`
	UserAddress               string `db:"UserAddress"`
}

type StakingProduct struct {
	ID                   string `db:"ID" json:"id"`
	ProductName          string `db:"ProductName" json:"productName"`
	MinOrderValue        int    `db:"MinOrderValue" json:"minOrderValue"`
	TopUpLimit           *big.Int
	StringTopUpLimit     string  `db:"TopUpLimit" json:"topUpLimit"`
	MinRedeemValue       int     `db:"MinRedeemValue" json:"minRedeemValue"`
	LockUpPeriod         int     `db:"LockUpPeriod" json:"lockUpPeriod"`
	DefaultAPY           float64 `db:"DefaultAPY" json:"-"`
	CurrentAPY           float64 `json:"currentAPY" json:"-"`
	CreateDate           string  `db:"CreateDate" json:"-" validate:"required,datetime=2006-01-02 15:04:05"`
	StartDate            string  `db:"StartDate" json:"-" validate:"required,datetime=2006-01-02 15:04:05"`
	Term                 int     `db:"Term" json:"-"`
	BurnedInterest       *big.Int
	StringBurnedInterest string  `db:"BurnedInterest" json:"-"`
	NextProductID        *string `db:"NextProductID" json:"-"`
	PaymentAddress       string  `db:"PaymentAddress" json:"-"`
	CurrencyType         string  `db:"CurrencyType" json:"-"`
	Network              string  `db:"Network" json:"-"`
	Status               bool    `db:"Status" json:"status"`
}

type User struct {
	DID        string `db:"DID"`
	Currency   string `db:"Currency"`
	CreateDate string `db:"CreateDate"`
}

type TXInfo struct {
	PaymentNo       string  `db:"PaymentNo"`
	OrderID         string  `db:"OrderID"`
	TXCurrencyType  string  `db:"TXCurrencyType"`
	TXType          string  `db:"TXType" validate:"required,oneof=BuyIn InterestOnly OrderClosure"`
	TXHash          *string `db:"TXHash"`
	Principal       *big.Int
	StringPrincipal string `db:"Principal" json:"-"`
	Interest        *big.Int
	StringInterest  string `db:"Interest" json:"-"`
	UserAddress     string `db:"UserAddress"`
	CreateDate      string `db:"CreateDate" validate:"omitempty,datetime=2006-01-02 15:04:05"`
	RedeemableTime  string `db:"RedeemableTime" validate:"required,datetime=2006-01-02 15:04:05"`
}

type OrderInterest struct {
	ID                      string  `db:"ID"`
	OrderID                 string  `db:"OrderID"`
	Time                    string  `db:"Time" validate:"required,datetime=2006-01-02 15:04:05"`
	APY                     float64 `db:"APY"`
	InterestGain            *big.Int
	StringInterestGain      string `db:"InterestGain" json:"-"`
	TotalInterestGain       *big.Int
	StringTotalInterestGain string `db:"TotalInterestGain" json:"-"`
}

type PaymentInfo struct {
	PaymentAddress string `db:"PaymentAddress"`
	Tag            string `db:"Tag"`
	CurrencyType   string `db:"CurrencyType"`
	Network        string `db:"Network"`
}

type MinerInfo struct {
	ID           string   `db:"ID"`
	Name         string   `db:"Name"`
	SSID         *string  `db:"SSID"`
	BSSID        *string  `db:"BSSID"`
	CreateTime   string   `db:"CreateTime"`
	Longitude    *float64 `db:"Longitude"`
	Latitude     *float64 `db:"Latitude"`
	OnlineStatus bool     `db:"OnlineStatus"`
	MiningPower  *float64 `db:"MiningPower"`
	IsMinable    bool     `db:"IsMinable"`
	DID          string   `db:"DID"`
	Host         string   `db:"Host"`
	IsVirtual    bool     `db:"IsVirtual"`
}

type SeedExchange struct {
	UserDID      string `db:"UserDID"`
	TargetDID    string `db:"TargetDID"`
	Challenge    string `db:"Challenge"`
	ExchangeRate string `db:"ExchangeRate"`
	Amount       string `db:"Amount"`
	CreateTime   string `db:"CreateTime"`
}

type StakingRecord struct {
	OrderID               string `db:"OrderID"`
	ProductID             string `db:"ProductID"`
	OrderStatus           string `db:"Type"`
	Term                  *int   `db:"Term"`
	PurchaseTime          string `db:"CreateDate"`
	PrincipalAmount       *big.Int
	StringPrincipalAmount string `db:"Amount" json:"-"`
	TXCurrencyType        string `db:"TXCurrencyType"`
	InterestGain          string
	TotalAmount           string
	RedeemableTime        string `db:"RedeemableTime"`
	IsInClosureWindow     bool
}

type ProductDetails struct {
	ID             string  `db:"ID"`
	ProductName    string  `db:"ProductName"`
	MinOrderValue  int     `db:"MinOrderValue"`
	TopUpLimit     float64 `db:"TopUpLimit"`
	MinRedeemValue int     `db:"MinRedeemValue"`
	LockUpPeriod   int     `db:"LockUpPeriod"`
	CurrentAPY     float64
	Status         bool `db:"Status"`
}

type SeedInfo struct {
	ID     string
	Amount float64
}

type OrderInterestInfo struct {
	AccumulatedInterest       *big.Int
	StringAccumulatedInterest string `db:"AccumulatedInterest" json:"-"`
	TotalInterestGained       *big.Int
	StringTotalInterestGained string `db:"TotalInterestGained" json:"-"`
}

type OrderInput struct {
	Amount      string
	UserAddress string
	UserDID     string
	ProductID   string
}

type OrderOutput struct {
	OrderID        string
	PaymentAddress string
}

type SubmitBuyinInput struct {
	OrderID string
	TxHash  string
}

type SubmitBuyinOutput struct {
	ProductName    string
	Amount         string
	Time           string
	UserAddress    string
	TXCurrencyType string
}

type RedeemOrderOuput struct {
	ProductName    string
	Amount         string
	Time           string
	ToAddress      string
	TXCurrencyType string
	TXHash         string
}

type PrincipalUpdate struct {
	ID                   string `db:"ID"`
	ProductID            string `db:"ProductID"`
	Time                 string `db:"Time" validate:"required,datetime=2006-01-02 15:04:05"`
	TotalPrincipal       *big.Int
	StringTotalPrincipal string `db:"TotalPrincipal" json:"-"`
}

type SeedExchangeInput struct {
	DID              string //placeholder
	WalletAddress    string
	SeedPresentation foundationModels.VerifiablePresentation
	PublicKeyString  []byte
}

type SeedExchangeOutput struct {
	Amount       string
	TxHash       string
	TxTime       string
	ExchangeRate string
}

type VpNonceInput struct {
	Session string
}

type VpNonceOutput struct {
	Session string
	Nonce   uint64
}

type VpNonce struct {
	Session string
	Nonce   uint64
	Time    time.Time
}

type MiningRole struct {
	DID           string `db:"DID"`
	WalletAddress string `db:"WalletAddress"`
	Type          string `db:"Type"`
}

type MiningRoleInput struct {
	WalletAddress    string
	SeedPresentation foundationModels.VerifiablePresentation
}

type NetworkConfirmRequest struct {
	Did           string `json:"did"`
	Target        string `json:"target"`
	LastBlockHash string `json:"lastBlockHash"`
	Quality       string `json:"quality"`
	PubKey        string `json:"pubKey"`
	Challenge     string `json:"challenge"`
	Signature     string `json:"signature"`
}

type NetworkConfirmResult struct {
	Did           string `json:"did"`
	Target        string `json:"target"`
	LastBlockHash string `json:"lastBlockHash"`
	PubKey        string `json:"pubKey"`
	Challenge     string `json:"challenge"`
	Signature     string `json:"signature"`
}

type NewSeedExchangeInput struct {
	WalletAddress string
	Seeds         []Seed
}

type Seed struct {
	Confirm NetworkConfirmRequest
	Result  NetworkConfirmResult
}

type SeedHistoryKeys struct {
	DID       string
	Target    string
	Challenge string
}

func CreateOrder() *Order {
	return &Order{}
}

func NewOrder(
	productID string, userDID string, orderType string, paymentAddress string, amount *big.Int, userAddress string) *Order {
	return &Order{
		"",
		productID,
		userDID,
		orderType,
		1,
		big.NewInt(0),
		"0",
		big.NewInt(0),
		"0",
		paymentAddress,
		amount,
		amount.String(),
		userAddress,
	}
}

func CreateStakingProduct() *StakingProduct {
	return &StakingProduct{}
}

func CreateUser() *User {
	return &User{}
}

func CreateTXInfo() *TXInfo {
	return &TXInfo{}
}

func NewTXInfo(orderID, txCurrencyType, txType, txHash string, principal, interest *big.Int, userAddress, redeemableTime string) *TXInfo {
	return &TXInfo{
		"",
		orderID,
		txCurrencyType,
		txType,
		&txHash,
		principal,
		principal.String(),
		interest,
		interest.String(),
		userAddress,
		"",
		redeemableTime,
	}
}

func CreateOrderInterest() *OrderInterest {
	return &OrderInterest{}
}

func CreateOrderInterestList() []*OrderInterest {
	return []*OrderInterest{}
}

func CreateMinerInfo() *MinerInfo {
	return &MinerInfo{}
}

func CreateSeedExchange() *SeedExchange {
	return &SeedExchange{}
}

func NewSeedExchange(userDID, targetDID, challenge, exchangeRate, amount string) *SeedExchange {
	return &SeedExchange{
		userDID,
		targetDID,
		challenge,
		exchangeRate,
		amount,
		"",
	}
}

func CreateStakingRecord() *StakingRecord {
	return &StakingRecord{}
}

func CreateProductDetails() *ProductDetails {
	return &ProductDetails{}
}

func CreateSeedInfo() *SeedInfo {
	return &SeedInfo{}
}

func CreateOrderInput() *OrderInput {
	return &OrderInput{}
}

func CreateOrderOutput() *OrderOutput {
	return &OrderOutput{}
}

func NewOrderOutput(orderID, paymentAddress string) *OrderOutput {
	return &OrderOutput{
		orderID,
		paymentAddress,
	}
}

func CreateSubmitBuyinInput() *SubmitBuyinInput {
	return &SubmitBuyinInput{}
}

func CreateSubmitBuyinOutput() *SubmitBuyinOutput {
	return &SubmitBuyinOutput{}
}

func NewSubmitBuyinOutput(productName, amount, time, userAddress, txCurrencyType string) *SubmitBuyinOutput {
	return &SubmitBuyinOutput{
		productName,
		amount,
		time,
		userAddress,
		txCurrencyType,
	}
}

func CreateOrderInterestInfo() *OrderInterestInfo {
	return &OrderInterestInfo{}
}

func CreateRedeemOrderOutput() *RedeemOrderOuput {
	return &RedeemOrderOuput{}
}

func NewPrincipalUpdate() *PrincipalUpdate {
	return &PrincipalUpdate{}
}

func NewPrincipalUpdateList() []*PrincipalUpdate {
	return []*PrincipalUpdate{}
}

func NewRedeemOrderOutput(productName string, amount string, time, toAddress, txCurrencyType, txHash string) *RedeemOrderOuput {
	return &RedeemOrderOuput{
		productName,
		amount,
		time,
		toAddress,
		txCurrencyType,
		txHash,
	}
}

func CreateSeedExchangeInput() *SeedExchangeInput {
	return &SeedExchangeInput{}
}

func CreateSeedExchangeOutput() *SeedExchangeOutput {
	return &SeedExchangeOutput{}
}

func NewSeedExchangeOutput(amount, txHash, txTime, exchangeRate string) *SeedExchangeOutput {
	return &SeedExchangeOutput{
		amount,
		txHash,
		txTime,
		exchangeRate,
	}
}

func NewSeedHistoryKeys(did, target, challenge string) *SeedHistoryKeys {
	return &SeedHistoryKeys{
		did,
		target,
		challenge,
	}
}

//need to convert SeedInfo portion of presentation from a map to a struct.
//This should most likely be done in foundation service with the rest of the conversions,
//but I implemented it here to make the system work. In the future, this can be
//removed once the foundation service has implemented SeedInfo VCs
func ConvertCredentialSubject(vc *foundationModels.VerifiableCredential) {
	subjectMap := vc.CredentialSubject.(map[string]interface{})
	seedInfo := CreateSeedInfo()
	seedInfo.ID = subjectMap["id"].(string)
	seedInfo.Amount = subjectMap["amount"].(float64)
	vc.CredentialSubject = *seedInfo
}

func SetOrderBigFields(order *Order) error {
	var success bool
	order.AccumulatedInterest, success = big.NewInt(0).SetString(order.StringAccumulatedInterest, 10)
	if !success {
		return errval.ErrAccumulatedInterestNotNumber
	}
	order.TotalInterestGained, success = big.NewInt(0).SetString(order.StringTotalInterestGained, 10)
	if !success {
		return errval.ErrTotalInterestGainedNotNumber
	}
	order.Amount, success = big.NewInt(0).SetString(order.StringAmount, 10)
	if !success {
		return errval.ErrAmountNotNumber
	}
	return nil
}

func SetStakingProductBigFields(product *StakingProduct) error {
	var success bool
	product.TopUpLimit, success = new(big.Int).SetString(product.StringTopUpLimit, 10)
	if !success {
		return errval.ErrTopUpLimitNotNumber
	}
	product.BurnedInterest, success = new(big.Int).SetString(product.StringBurnedInterest, 10)
	if !success {
		return errval.ErrBurnedInterestNotNumber
	}
	return nil
}

func SetTXInfoBigFields(txinfo *TXInfo) error {
	var success bool
	txinfo.Principal, success = big.NewInt(0).SetString(txinfo.StringPrincipal, 10)
	if !success {
		return errval.ErrPrincipalNotNumber
	}
	txinfo.Interest, success = big.NewInt(0).SetString(txinfo.StringInterest, 10)
	if !success {
		return errval.ErrInterestNotNumber
	}
	return nil
}

func SetOrderInterestBigFields(interest *OrderInterest) error {
	var success bool
	interest.InterestGain, success = big.NewInt(0).SetString(interest.StringInterestGain, 10)
	if !success {
		return errval.ErrInterestGainNotNumber
	}
	interest.TotalInterestGain, success = big.NewInt(0).SetString(interest.StringTotalInterestGain, 10)
	if !success {
		return errval.ErrTotalInterestGainNotNumber
	}
	return nil
}

func SetOrderInterestInfoBigFields(info *OrderInterestInfo) error {
	var success bool
	info.AccumulatedInterest, success = big.NewInt(0).SetString(info.StringAccumulatedInterest, 10)
	if !success {
		return errval.ErrAccumulatedInterestNotNumber
	}
	info.TotalInterestGained, success = big.NewInt(0).SetString(info.StringTotalInterestGained, 10)
	if !success {
		return errval.ErrTotalInterestGainedNotNumber
	}
	return nil
}

func SetPrincipalUpdateBigFields(update *PrincipalUpdate) error {
	var success bool
	update.TotalPrincipal, success = big.NewInt(0).SetString(update.StringTotalPrincipal, 10)
	if !success {
		return errval.ErrTotalPrincipalNotNumber
	}
	return nil
}

func SetStakingRecordBigFields(record *StakingRecord) error {
	var success bool
	record.PrincipalAmount, success = big.NewInt(0).SetString(record.StringPrincipalAmount, 10)
	if !success {
		return errval.ErrAmountNotNumber
	}
	return nil
}
