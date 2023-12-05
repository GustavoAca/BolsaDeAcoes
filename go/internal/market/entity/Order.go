package entity

type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PendingShares int
	Price         float64
	OrderType     string
	Status        string
	Transactions  []*Transaction
}

const (
	OrderTypeBuy  string = "BUY"
	OrderTypeSell string = "SELL"
)

const (StatusOpen = "OPEN"
StatusClosed string = "CLOSED")


func NewOrder(orderID string, investor *Investor, asset *Asset, shares int, price float64, orderType string) *Order {
	return &Order{
		ID:            orderID,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PendingShares: shares,
		Price:         price,
		OrderType:     orderType,
		Status:        "OPEN",
		Transactions:  []*Transaction{},
	}
}

func (o *Order) AddTransaction(transaction *Transaction) {
	o.Transactions = append(o.Transactions, transaction)
}

func (o *Order) AddOrderPendingShares(minShares int) {
	o.PendingShares += minShares
}

func (o *Order) IsBuy() bool {
	return o.Status == OrderTypeBuy
}

func (o *Order) IsSell() bool {
	return o.Status == OrderTypeSell
}

func (o *Order) IsClose() {
	if o.PendingShares == 0{
		o.Status = StatusClosed
	}
}

func (o *Order) GetStatus() StatusOrder {
	statusMap := map[string]StatusOrder{
		OrderTypeBuy:  Buy{},
		OrderTypeSell: Sell{},
	}

	return statusMap[o.Status]
}