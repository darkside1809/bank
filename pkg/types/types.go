package types


// Money is amount of money in minimal currency amount like: panny, dirams etc.
type Money int

// Category is a category that has payment (cars, restaurants and others)
type Category string

// Currency is currency codes
type Currency string

// Status is status of payment
type Status string

const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// currency codes
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN is number of card
type PAN string


// Card is information about payment card
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

//Payment is info about payment
type Payment struct {
	ID 			int
	Amount 		Money
	Category 	Category
	Status		Status
}

// PaymentSource info about card
type PaymentSource struct {
	Type string
	Number string
	Balance Money
}