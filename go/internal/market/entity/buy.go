package entity

type Buy struct{

}

func (buy Buy) Trade(order *Order, b *Book, buyOrders *OrderQueue, sellOrders *OrderQueue) {
	buyOrders.Push(order)
	if sellOrders.Len() > 0 && sellOrders.Orders[0].Price <= order.Price {
		sellOrder := sellOrders.Pop().(*Order)
		if sellOrder.PendingShares > 0 {
			transaction := NewTransaction(sellOrder, order, order.Shares, sellOrder.Price)
			b.AddTransaction(transaction, b.wg)
			sellOrder.AddTransaction(transaction)
			order.AddTransaction(transaction)
			b.OrdersChanOut <- sellOrder
			b.OrdersChanOut <- order
			if sellOrder.PendingShares > 0 {
				sellOrders.Push(sellOrder)
			}
		}
	}
}
