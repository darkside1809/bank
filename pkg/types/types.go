package types


// Money is amount of money in minimal currency amount like: panny, dirams etc.
type Money int64

// Category is a category that has payment (cars, restaurants and others)
type PaymentCategory string

// Currency is currency codes
type Currency string

// Status is status of payment
type PaymentStatus string

const (
	PaymentStatusOk 				PaymentStatus = "OK"
	PaymentStatusFail 			PaymentStatus = "FAIL"
	PaymentStatusInProgress 	PaymentStatus = "INPROGRESS"
)

// currency codes
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN is number of card, like "5555-xxxx-xxxx-8888"
type PAN string

// Card is information about payment card, balance,id etc.
type Card struct {
	ID				int
	PAN			PAN
	Balance		Money
	MinBalance 	Money
	Currency 	Currency
	Color			string
	Name			string
	Active		bool
}


//Payment is information about payment, withdraw, deposit
type Payment struct {
	ID 			string
	Amount 		Money
	Category 	PaymentCategory
	Status		PaymentStatus
}

type Phone string

// Account is person's account information
type Account struct {
	ID int64
	Phone Phone
	Balance Money
}

// PaymentSource info about card
type PaymentSource struct {
	Type string
	Number string
	Balance Money
}