package types

type Money int64

type Currency string

//Коды валют
const (
	TJS Currency = "TJS"
	USD Currency = "USD"
	RUB Currency = "RUB"
)

type PAN string

type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	
}


