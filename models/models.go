package models

const TxTypeBuyIn = "BuyIn"
const TxTypeInterestOnly = "InterestOnly"
const TxTypeOrderClosure = "OrderClosure"

const OrderTypePending = "Pending"
const OrderTypeHolding = "Holding"
const OrderTypeComplete = "Complete"

type Order struct {
	OrderID             string  `db:"OrderID"`
	ProductID           string  `db:"ProductID"`
	UserDID             string  `db:"UserDID"`
	Type                string  `db:"Type" validate:"required,oneof=Pending Holding Complete"`
	Term                *int    `db:"Term"`
	AccumulatedInterest float64 `db:"AccumulatedInterest"`
	TotalInterestGained float64 `db:"TotalInterestGained"`
	PaymentAddress      string  `db:"PaymentAddress"`
	Amount              float64 `db:"Amount"`
	UserAddress         string  `db:"UserAddress"`
}

type StakingProduct struct {
	ID             string  `db:"ID" json:"id"`
	ProductName    string  `db:"ProductName" json:"productName"`
	MinOrderValue  int     `db:"MinOrderValue" json:"minOrderValue"`
	TopUpLimit     float64 `db:"TopUpLimit" json:"topUpLimit"`
	MinRedeemValue int     `db:"MinRedeemValue" json:"minRedeemValue"`
	LockUpPeriod   int     `db:"LockUpPeriod" json:"lockUpPeriod"`
	DefaultAPY     float64 `db:"DefaultAPY"`
	CurrentAPY     float64 `json:"currentAPY"`
	CreateDate     string  `db:"CreateDate"`
	StartDate      string  `db:"StartDate"`
	Term           int     `db:"Term"`
	BurnedInterest float64 `db:"BurnedInterest"`
	Status         bool    `db:"Status" json:"status"`
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
	TXType         string  `db:"TXType" validate:"required,oneof=BuyIn InterestOnly OrderClosure"`
	TXHash         *string `db:"TXHash"`
	Principal      float64 `db:"Principal"`
	Interest       float64 `db:"Interest"`
	UserAddress    string  `db:"UserAddress"`
	CreateDate     string  `db:"CreateDate"`
	RedeemableTime string  `db:"RedeemableTime"`
}

type OrderInterest struct {
	ID           string  `db:"ID"`
	OrderID      string  `db:"OrderID"`
	Time         string  `db:"Time" validate:"required,datetime=2006-01-02 15:04:05"`
	APY          float64 `db:"APY"`
	InterestGain float64 `db:"InterestGain"`
}

type PaymentInfo struct {
	PaymentAddress string `db:"PaymentAddress"`
	Tag            string `db:"Tag"`
	CurrencyType   string `db:"CurrencyType"`
	Network        string `db:"Network"`
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

type PrincipalUpdate struct {
	ID             string  `db:"ID"`
	ProductID      string  `db:"ProductID"`
	Time           string  `db:"Time" validate:"required,datetime=2006-01-02 15:04:05"`
	TotalPrincipal float64 `db:"TotalPrincipal"`
}

type ProductHistory struct {
	OrderID    string `db:"ID"`
	ProductID  string `db:"ProductID"`
	CreateDate string `db:"CreateDate"`
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

func NewOrderInterestList() []*OrderInterest {
	return []*OrderInterest{}
}

func NewStakingRecord() *StakingRecord {
	return &StakingRecord{}
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

func NewPrincipalUpdate() *PrincipalUpdate {
	return &PrincipalUpdate{}
}

func NewPrincipalUpdateList() []*PrincipalUpdate {
	return []*PrincipalUpdate{}
}

func NewProductHistory() *ProductHistory {
	return &ProductHistory{}
}

func NewProductHistoryList() []*ProductHistory {
	return []*ProductHistory{}
}
