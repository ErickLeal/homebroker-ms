package entity

type OrderType string
type OrderStatus string

const (
	BUY  OrderType = "BUY"
	SELL OrderType = "SELL"

	OPEN   OrderStatus = "OPEN"
	CLOSED OrderStatus = "CLOSED"
)

type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PendingShares int
	Price         float64
	Type          OrderType
	Status        OrderStatus
	Transactions  []*Transaction
}

func NewOrder(orderID string, investor *Investor, asset *Asset, shares int, price float64, orderType OrderType) *Order {
	return &Order{
		ID:            orderID,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PendingShares: shares,
		Price:         price,
		Type:          orderType,
		Status:        OPEN,
		Transactions:  []*Transaction{},
	}
}

func (o *Order) ApplyTrade(tradedShares int) {
	if tradedShares > o.PendingShares {
		tradedShares = o.PendingShares
	}
	o.PendingShares -= tradedShares
	if o.PendingShares == 0 {
		o.Status = CLOSED
	}
}

func (o *Order) AddTransaction(transaction *Transaction) {
	o.Transactions = append(o.Transactions, transaction)
}
