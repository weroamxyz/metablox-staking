package models

import (
	"math"
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

const MinimumUnitsPerMBLX = 1000000

type Order struct {
	OrderID                   string   `db:"OrderID"`
	ProductID                 string   `db:"ProductID"`
	UserDID                   string   `db:"UserDID"`
	Type                      string   `db:"Type" validate:"required,oneof=Pending Holding Complete"`
	Term                      int      `db:"Term"`
	AccumulatedInterest       *big.Int `json:"-"`
	StringAccumulatedInterest string   `db:"AccumulatedInterest" json:"-"`
	MBLXAccumulatedInterest   float64  `json:"AccumulatedInterest"`
	TotalInterestGained       *big.Int `json:"-"`
	StringTotalInterestGained string   `db:"TotalInterestGained" json:"-"`
	MBLXTotalInterestGained   float64  `json:"TotalInterestGained"`
	PaymentAddress            string   `db:"PaymentAddress"`
	Amount                    *big.Int `json:"-"`
	StringAmount              string   `db:"Amount" json:"-"`
	MBLXAmount                float64  `json:"Amount"`
	UserAddress               string   `db:"UserAddress"`
}

type StakingProduct struct {
	ID                   string   `db:"ID" json:"id"`
	ProductName          string   `db:"ProductName" json:"productName"`
	MinOrderValue        int      `db:"MinOrderValue" json:"minOrderValue"`
	TopUpLimit           *big.Int `json:"-"`
	StringTopUpLimit     string   `db:"TopUpLimit"`
	MBLXTopUpLimit       float64  `json:"TopUpLimit"`
	MinRedeemValue       int      `db:"MinRedeemValue" json:"minRedeemValue"`
	LockUpPeriod         int      `db:"LockUpPeriod" json:"lockUpPeriod"`
	DefaultAPY           float64  `db:"DefaultAPY" json:"-"`
	CurrentAPY           float64  `json:"currentAPY" json:"-"`
	CreateDate           string   `db:"CreateDate" json:"-" validate:"required,datetime=2006-01-02 15:04:05"`
	StartDate            string   `db:"StartDate" json:"-" validate:"required,datetime=2006-01-02 15:04:05"`
	Term                 int      `db:"Term" json:"-"`
	BurnedInterest       *big.Int `json:"-"`
	StringBurnedInterest string   `db:"BurnedInterest"`
	MBLXBurnedInterest   float64  `json:"BurnedInterest"`
	NextProductID        *string  `db:"NextProductID" json:"-"`
	PaymentAddress       string   `db:"PaymentAddress"`
	CurrencyType         string   `db:"CurrencyType" json:"-"`
	Network              string   `db:"Network" json:"-"`
	Status               bool     `db:"Status" json:"status"`
}

type User struct {
	DID        string `db:"DID"`
	Currency   string `db:"Currency"`
	CreateDate string `db:"CreateDate"`
}

type TXInfo struct {
	PaymentNo       string   `db:"PaymentNo"`
	OrderID         string   `db:"OrderID"`
	TXCurrencyType  string   `db:"TXCurrencyType"`
	TXType          string   `db:"TXType" validate:"required,oneof=BuyIn InterestOnly OrderClosure"`
	TXHash          string   `db:"TXHash"`
	Principal       *big.Int `json:"-"`
	StringPrincipal string   `db:"Principal" json:"-"`
	MBLXPrincipal   float64  `json:"Principal"`
	Interest        *big.Int `json:"-"`
	StringInterest  string   `db:"Interest" json:"-"`
	MBLXInterest    float64  `json:"Interest"`
	UserAddress     string   `db:"UserAddress"`
	CreateDate      string   `db:"CreateDate" validate:"omitempty,datetime=2006-01-02 15:04:05"`
	RedeemableTime  string   `db:"RedeemableTime" validate:"required,datetime=2006-01-02 15:04:05"`
}

type OrderInterest struct {
	ID                      string   `db:"ID"`
	OrderID                 string   `db:"OrderID"`
	Time                    string   `db:"Time" validate:"required,datetime=2006-01-02 15:04:05"`
	APY                     float64  `db:"APY"`
	InterestGain            *big.Int `json:"-"`
	StringInterestGain      string   `db:"InterestGain" json:"-"`
	MBLXInterestGain        float64  `json:"InterestGain"`
	TotalInterestGain       *big.Int `json:"-"`
	StringTotalInterestGain string   `db:"TotalInterestGain" json:"-"`
	MBLXTotalInterestGain   float64  `db:"TotalInterestGain"`
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
	OrderID               string   `db:"OrderID"`
	ProductID             string   `db:"ProductID"`
	OrderStatus           string   `db:"Type"`
	Term                  *int     `db:"Term"`
	PurchaseTime          string   `db:"CreateDate"`
	PrincipalAmount       *big.Int `json:"-"`
	StringPrincipalAmount string   `db:"Amount" json:"-"`
	MBLXPrincipalAmount   float64  `json:"PrincipalAmount"`
	TXCurrencyType        string   `db:"TXCurrencyType"`
	InterestGain          string
	TotalAmount           string
	RedeemableTime        string `db:"RedeemableTime"`
	IsInClosureWindow     bool
}

type SeedInfo struct {
	ID     string
	Amount float64
}

type OrderInterestInfo struct {
	AccumulatedInterest       *big.Int `json:"-"`
	StringAccumulatedInterest string   `db:"AccumulatedInterest" json:"-"`
	MBLXAccumulatedInterest   float64  `json:"AccumulatedInterest"`
	TotalInterestGained       *big.Int `json:"-"`
	StringTotalInterestGained string   `db:"TotalInterestGained" json:"-"`
	MBLXTotalInterestGained   float64  `json:"TotalInterestGained"`
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
	ID                   string   `db:"ID"`
	ProductID            string   `db:"ProductID"`
	Time                 string   `db:"Time" validate:"required,datetime=2006-01-02 15:04:05"`
	TotalPrincipal       *big.Int `json:"-"`
	StringTotalPrincipal string   `db:"TotalPrincipal" json:"-"`
	MBLXTotalPrincipal   float64  `json:"TotalPrincipal"`
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
		0,
		big.NewInt(0),
		"0",
		0,
		paymentAddress,
		amount,
		amount.String(),
		MinimumUnitToMBLX(amount),
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
		txHash,
		principal,
		principal.String(),
		MinimumUnitToMBLX(principal),
		interest,
		interest.String(),
		MinimumUnitToMBLX(interest),
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
	f, _, err := big.ParseFloat(order.StringAmount, 10, 128, big.ToNearestEven)
	if err != nil {
		return errval.ErrAmountNotNumber
	}
	order.Amount, _ = f.Int(nil)

	f, _, err = big.ParseFloat(order.StringAccumulatedInterest, 10, 128, big.ToNearestEven)
	if err != nil {
		return errval.ErrAccumulatedInterestNotNumber
	}
	order.AccumulatedInterest, _ = f.Int(nil)

	f, _, err = big.ParseFloat(order.StringTotalInterestGained, 10, 128, big.ToNearestEven)
	if err != nil {
		return errval.ErrTotalInterestGainedNotNumber
	}
	order.TotalInterestGained, _ = f.Int(nil)
	return nil
}

func SetStakingProductBigFields(product *StakingProduct) error {
	f, _, err := big.ParseFloat(product.StringTopUpLimit, 10, 128, big.ToNearestEven)
	if err != nil {
		return errval.ErrTopUpLimitNotNumber
	}
	product.TopUpLimit, _ = f.Int(nil)

	f, _, err = big.ParseFloat(product.StringBurnedInterest, 10, 128, big.ToNearestEven)
	if err != nil {
		return errval.ErrBurnedInterestNotNumber
	}
	product.BurnedInterest, _ = f.Int(nil)
	return nil
}

func SetTXInfoBigFields(txinfo *TXInfo) error {
	f, _, err := big.ParseFloat(txinfo.StringPrincipal, 10, 128, big.ToNearestEven)
	if err != nil {
		return errval.ErrPrincipalNotNumber
	}
	txinfo.Principal, _ = f.Int(nil)

	f, _, err = big.ParseFloat(txinfo.StringInterest, 10, 128, big.ToNearestAway)
	if err != nil {
		return errval.ErrInterestNotNumber
	}
	txinfo.Interest, _ = f.Int(nil)
	return nil
}

func SetOrderInterestBigFields(interest *OrderInterest) error {
	f, _, err := big.ParseFloat(interest.StringInterestGain, 10, 128, big.ToNearestAway)
	if err != nil {
		return errval.ErrInterestGainNotNumber
	}
	interest.InterestGain, _ = f.Int(nil)

	f, _, err = big.ParseFloat(interest.StringTotalInterestGain, 10, 128, big.ToNearestEven)
	if err != nil {
		return errval.ErrTotalInterestGainNotNumber
	}
	interest.TotalInterestGain, _ = f.Int(nil)
	return nil
}

func SetOrderInterestInfoBigFields(info *OrderInterestInfo) error {
	f, _, err := big.ParseFloat(info.StringAccumulatedInterest, 10, 128, big.ToNearestEven)
	if err != nil {
		return errval.ErrAccumulatedInterestNotNumber
	}
	info.AccumulatedInterest, _ = f.Int(nil)

	f, _, err = big.ParseFloat(info.StringTotalInterestGained, 10, 128, big.ToNearestEven)
	if err != nil {
		return errval.ErrTotalInterestGainedNotNumber
	}
	info.TotalInterestGained, _ = f.Int(nil)
	return nil
}

func SetPrincipalUpdateBigFields(update *PrincipalUpdate) error {
	f, _, err := big.ParseFloat(update.StringTotalPrincipal, 10, 128, big.ToNearestEven)
	if err != nil {
		return errval.ErrTotalPrincipalNotNumber
	}
	update.TotalPrincipal, _ = f.Int(nil)
	return nil
}

func SetStakingRecordBigFields(record *StakingRecord) error {
	f, _, err := big.ParseFloat(record.StringPrincipalAmount, 10, 128, big.ToNearestEven)
	if err != nil {
		return errval.ErrPrincipalNotNumber
	}
	record.PrincipalAmount, _ = f.Int(nil)
	return nil
}

func (order *Order) SetMBLXValues() {
	order.MBLXAccumulatedInterest = MinimumUnitToMBLX(order.AccumulatedInterest)
	order.MBLXAmount = MinimumUnitToMBLX(order.Amount)
	order.MBLXTotalInterestGained = MinimumUnitToMBLX(order.TotalInterestGained)
}

func (product *StakingProduct) SetMBLXValues() {
	product.MBLXTopUpLimit = MinimumUnitToMBLX(product.TopUpLimit)
	product.MBLXBurnedInterest = MinimumUnitToMBLX(product.BurnedInterest)
}

func (info *TXInfo) SetMBLXValues() {
	info.MBLXPrincipal = MinimumUnitToMBLX(info.Principal)
	info.MBLXInterest = MinimumUnitToMBLX(info.Interest)
}

func (interest *OrderInterest) SetMBLXValues() {
	interest.MBLXTotalInterestGain = MinimumUnitToMBLX(interest.TotalInterestGain)
	interest.MBLXInterestGain = MinimumUnitToMBLX(interest.InterestGain)
}

func (record *StakingRecord) SetMBLXValues() {
	record.MBLXPrincipalAmount = MinimumUnitToMBLX(record.PrincipalAmount)
}

func (interestInfo *OrderInterestInfo) SetMBLXValues() {
	interestInfo.MBLXAccumulatedInterest = MinimumUnitToMBLX(interestInfo.AccumulatedInterest)
	interestInfo.MBLXTotalInterestGained = MinimumUnitToMBLX(interestInfo.TotalInterestGained)
}

func (principal *PrincipalUpdate) SetMBLXValues() {
	principal.MBLXTotalPrincipal = MinimumUnitToMBLX(principal.TotalPrincipal)
}

func MinimumUnitToMBLX(amount *big.Int) float64 {
	convertedAmount := new(big.Float).SetInt(amount)
	convertedAmount = convertedAmount.Quo(convertedAmount, big.NewFloat(MinimumUnitsPerMBLX))
	floatAmount, _ := convertedAmount.Float64()
	return floatAmount
}

func MBLXToMinimumUnit(amount float64) *big.Int {
	return new(big.Int).SetInt64(int64(math.Round(amount * MinimumUnitsPerMBLX)))
}
