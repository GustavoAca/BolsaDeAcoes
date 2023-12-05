package entity

import (
	"math"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Price        float64
	Total        float64
	DateTime     time.Time
}

const (
	TransactionPending string = "CLOSED"
)

func NewTransaction(sellingOrder *Order, buyingOrder *Order, shares int, price float64) *Transaction {
	total := float64(shares) * price
	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
		Shares:       shares,
		Price:        price,
		Total:        total,
		DateTime:     time.Now(),
	}
}

func (t *Transaction) CalculateTotal(shares int, prices float64) {
	t.Total = float64(shares) * prices
}

func (t *Transaction) CloseSellOrder() {
	t.SellingOrder.IsClose()
}

func (t *Transaction) CloseBuyOrder() {
	t.BuyingOrder.IsClose()
}

func (t *Transaction) minPendingShares() int{
	return int(math.Min(float64(t.SellingOrder.PendingShares),float64(t.BuyingOrder.PendingShares)))
}
