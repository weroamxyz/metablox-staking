package models

import (
	foundationModels "github.com/MetaBloxIO/metablox-foundation-services/models"
)

const OrderTypePending = "Pending"
const OrderTypeHolding = "Holding"
const OrderTypeComplete = "Complete"

const CurrencyTypeMBLX = "MBLX"

const TXTypeBuyin = "BuyIn"
const TXTypeInterest = "InterestOnly"
const TXTypeClosure = "OrderClosure"

type Order struct {
	OrderID             string  `db:"OrderID"`
	ProductID           string  `db:"ProductID"`
	UserDID             string  `db:"UserDID"`
	Type                string  `db:"Type"`
	Term                int     `db:"Term"`
	AccumulatedInterest float64 `db:"AccumulatedInterest"`
	TotalInterestGained float64 `db:"TotalInterestGained"`
	PaymentAddress      string  `db:"PaymentAddress"`
	Amount              float64 `db:"Amount"`
	UserAddress         string  `db:"UserAddress"`
}

type StakingProduct struct {
	ID             string  `db:"ID"`
	ProductName    string  `db:"ProductName"`
	MinOrderValue  int     `db:"MinOrderValue"`
	TopUpLimit     float64 `db:"TopUpLimit"`
	MinRedeemValue int     `db:"MinRedeemValue"`
	LockUpPeriod   int     `db:"LockUpPeriod"`
	DefaultAPY     float64 `db:"DefaultAPY"`
	CreateDate     string  `db:"CreateDate"`
	StartDate      string  `db:"StartDate"`
	Term           int     `db:"Term"`
	BurnedInterest float64 `db:"BurnedInterest"`
	Status         bool    `db:"Status"`
}

type User struct {
	DID        string `db:"DID"`
	Currency   string `db:"Currency"`
	CreateDate string `db:"CreateDate"`
}

type TXInfo struct {
	PaymentNo      string  `db:"PaymentNo"`
	OrderID        string  `db:"OrderID"`
	TXCurrencyType string  `db:"TXCurrencyType"`
	TXType         string  `db:"TXType"`
	TXHash         *string `db:"TXHash"`
	Principal      float64 `db:"Principal"`
	Interest       float64 `db:"Interest"`
	UserAddress    string  `db:"UserAddress"`
	CreateDate     string  `db:"CreateDate"`
	RedeemableTime string  `db:"RedeemableTime"`
}

type OrderInterest struct {
	ID                string  `db:"ID"`
	OrderID           string  `db:"OrderID"`
	Time              string  `db:"Time"`
	APY               float64 `db:"APY"`
	InterestGain      float64 `db:"InterestGain"`
	TotalInterestGain float64 `db:"TotalInterestGain"`
}

type PrincipalUpdates struct {
	ID             string  `db:"ID"`
	ProductID      string  `db:"ProductID"`
	Time           string  `db:"Time"`
	TotalPrincipal float64 `db:"TotalPrincipal"`
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
}

type SeedExchange struct {
	VcID         string  `db:"VcID"`
	UserDID      string  `db:"UserDID"`
	ExchangeRate float64 `db:"ExchangeRate"`
	Amount       float64 `db:"Amount"`
	CreateTime   string  `db:"CreateTime"`
}

type StakingRecord struct {
	OrderID           string  `db:"OrderID"`
	ProductID         string  `db:"ProductID"`
	OrderStatus       string  `db:"Type"`
	Term              *int    `db:"Term"`
	PurchaseTime      string  `db:"CreateDate"`
	PrincipalAmount   float64 `db:"Amount"`
	TXCurrencyType    string  `db:"TXCurrencyType"`
	InterestGain      float64
	TotalAmount       float64
	RedeemableTime    string `db:"RedeemableTime"`
	IsInClosureWindow bool
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
	AccumulatedInterest float64 `db:"AccumulatedInterest"`
	TotalInterestGained float64 `db:"TotalInterestGained"`
}

type OrderInput struct {
	Amount      float64
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
	Amount         float64
	Time           string
	UserAddress    string
	TXCurrencyType string
}

type RedeemOrderOuput struct {
	ProductName    string
	Amount         float64
	Time           string
	ToAddress      string
	TXCurrencyType string
	TXHash         string
}

type SeedExchangeInput struct {
	DID              string //placeholder
	WalletAddress    string
	SeedPresentation foundationModels.VerifiablePresentation
	PublicKeyString  []byte
}

type SeedExchangeOutput struct {
	Amount       float64
	TxHash       string
	TxTime       string
	ExchangeRate float64
}

func CreateOrder() *Order {
	return &Order{}
}

func NewOrder(
	productID string, userDID string, orderType string, paymentAddress string, amount float64, userAddress string) *Order {
	return &Order{
		"",
		productID,
		userDID,
		orderType,
		1,
		0,
		0,
		paymentAddress,
		amount,
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

func NewTXInfo(orderID, txCurrencyType, txType, txHash string, principal, interest float64, userAddress, redeemableTime string) *TXInfo {
	return &TXInfo{
		"",
		orderID,
		txCurrencyType,
		txType,
		&txHash,
		principal,
		interest,
		userAddress,
		"",
		redeemableTime,
	}
}

func CreateOrderInterest() *OrderInterest {
	return &OrderInterest{}
}

func CreateMinerInfo() *MinerInfo {
	return &MinerInfo{}
}

func CreateSeedExchange() *SeedExchange {
	return &SeedExchange{}
}

func NewSeedExchange(vcID, userDID string, exchangeRate, amount float64) *SeedExchange {
	return &SeedExchange{
		vcID,
		userDID,
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

func NewSubmitBuyinOutput(productName string, amount float64, time, userAddress, txCurrencyType string) *SubmitBuyinOutput {
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

func NewRedeemOrderOutput(productName string, amount float64, time, toAddress, txCurrencyType, txHash string) *RedeemOrderOuput {
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

func NewSeedExchangeOutput(amount float64, txHash, txTime string, exchangeRate float64) *SeedExchangeOutput {
	return &SeedExchangeOutput{
		amount,
		txHash,
		txTime,
		exchangeRate,
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
