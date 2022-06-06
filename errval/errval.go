package errval

import "errors"

var ErrBadDID = errors.New("did is improperly formatted")
var ErrDIDNotRegistered = errors.New("did is not registered in contract")
var ErrNotEnoughInterest = errors.New("order does not meet minimum interest required to redeem")
var ErrTransactionPending = errors.New("transaction is still pending")
var ErrAddressComparisonFail = errors.New("tx address does not match order address")
var ErrAmountComparisonFail = errors.New("tx amount does not match order amount")
var ErrExistingTXHash = errors.New("provided tx hash is already recorded in db")
var ErrEarlyOrderRedeem = errors.New("order can only be redeemed on final day of term")
var ErrVCIDFormat = errors.New("VC id is improperly formatted")
var ErrUpdateOrderStatus = errors.New("failed to update order status; it may not exist, or it may already be holding")
var ErrInvalidSession = errors.New("invalid session id")
var ErrNonceTimeout = errors.New("nonce timeout")
