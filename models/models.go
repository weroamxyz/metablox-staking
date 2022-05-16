package models

type Order struct {
	OrderID      int    `db:"OrderID"`
	ProductID    int    `db:"ProductID"`
	UserDID      string `db:"UserDID"`
	Type         bool   `db:"Type"`
	RedeemStatus bool   `db:"RedeemStatus"`
	Term         *int   `db:"Term"`
}

type StakingProduct struct {
	ID             int     `db:"ID"`
	MinOrderValue  int     `db:"MinOrderValue"`
	TopUpLimit     string  `db:"TopUpLimit"`
	MinRedeemValue int     `db:"MinRedeemValue"`
	LockUpPeriod   int     `db:"LockUpPeriod"`
	DefaultAPY     float32 `db:"DefaultAPY"`
	CreateDate     string  `db:"CreateDate"`
	StartDate      string  `db:"StartDate"`
	Term           int     `db:"Term"`
	TotalPrincipal string  `db:"TotalPrincipal"`
	Status         bool    `db:"Status"`
}

type User struct {
	DID        string `db:"DID"`
	Currency   string `db:"Currency"`
	CreateDate string `db:"CreateDate"`
}

type TXInfo struct {
	PaymentNo      int     `db:"PaymentNo"`
	UserDID        string  `db:"UserDID"`
	OrderID        int     `db:"OrderID"`
	TXCurrencyType string  `db:"TXCurrencyTYPE"`
	TXType         string  `db:"TXType"`
	TXHash         *string `db:"TXHash"`
	Amount         string  `db:"Amount"`
	UserAddress    string  `db:"UserAddress"`
	CreateDate     string  `db:"CreateDate"`
}

type OrderInterest struct {
	ID                int     `db:"ID"`
	OrderID           int     `db:"OrderID"`
	Time              string  `db:"Time"`
	APY               float32 `db:"APY"`
	InterestGain      float32 `db:"InterestGain"`
	TotalInterestGain float32 `db:"TotalInterestGain"`
}

type StakingRecord struct {
	OrderID           int     `db:"OrderID"`
	Term              *int    `db:"Term"`
	CreateDate        string  `db:"CreateDate"`
	TotalInterestGain float32 `db:"TotalInterestGain"`
	RedeemAll         bool
	Type              bool `db:"Type"`
}

type RedeemInput struct {
	OrderID        string `db:"OrderID"`
	TXCurrencyType string `db:"TXCurrencyTYPE"`
	Amount         string `db:"Amount"`
	UserAddress    string `db:"UserAddress"`
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

func NewStakingRecord() *StakingRecord {
	return &StakingRecord{}
}

func NewRedeemInput() *RedeemInput {
	return &RedeemInput{}
}
