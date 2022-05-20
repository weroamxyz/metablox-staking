package models

import (
	foundationModels "github.com/MetaBloxIO/metablox-foundation-services/models"
)

const OrderTypePending = "Pending"
const OrderTypeHolding = "Holding"
const OrderTypeComplete = "Complete"

type Order struct {
	OrderID             string  `db:"OrderID"`
	ProductID           string  `db:"ProductID"`
	UserDID             string  `db:"UserDID"`
	Type                string  `db:"Type"`
	Term                *int    `db:"Term"`
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

type PaymentInfo struct {
	PaymentAddress string `db:"PaymentAddress"`
	Tag            string `db:"Tag"`
	CurrencyType   string `db:"CurrencyType"`
	Network        string `db:"Network"`
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

type CreateOrderInput struct {
	Amount      float64
	UserAddress string
	UserDID     string
	ProductID   string
}

type CreateOrderOutput struct {
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

type MinerListInput struct {
	Latitude  *float64
	Longitude *float64
}

type SeedExchangeInput struct {
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

func NewOrder() *Order {
	return &Order{}
}

func NewStakingProduct() *StakingProduct {
	return &StakingProduct{}
}

func NewUser() *User {
	return &User{}
}

func NewTXInfo() *TXInfo {
	return &TXInfo{}
}

func NewOrderInterest() *OrderInterest {
	return &OrderInterest{}
}

func NewMinerInfo() *MinerInfo {
	return &MinerInfo{}
}

func NewSeedExchange() *SeedExchange {
	return &SeedExchange{}
}

func NewStakingRecord() *StakingRecord {
	return &StakingRecord{}
}

func NewProductDetails() *ProductDetails {
	return &ProductDetails{}
}

func NewCreateOrderInput() *CreateOrderInput {
	return &CreateOrderInput{}
}

func NewCreateOrderOutput() *CreateOrderOutput {
	return &CreateOrderOutput{}
}

func NewSubmitBuyinInput() *SubmitBuyinInput {
	return &SubmitBuyinInput{}
}

func NewSubmitBuyinOutput() *SubmitBuyinOutput {
	return &SubmitBuyinOutput{}
}

func NewOrderInterestInfo() *OrderInterestInfo {
	return &OrderInterestInfo{}
}

func NewRedeemOrderOutput() *RedeemOrderOuput {
	return &RedeemOrderOuput{}
}

func NewMinerListInput() *MinerListInput {
	return &MinerListInput{}
}

func NewSeedExchangeInput() *SeedExchangeInput {
	return &SeedExchangeInput{}
}

func NewSeedExchangeOutput() *SeedExchangeOutput {
	return &SeedExchangeOutput{}
}
