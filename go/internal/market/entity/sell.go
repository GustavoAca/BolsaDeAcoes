package entity

type Sell struct{
}

func (sell Sell) Trade(order *Order, b *Book, buyOrders *OrderQueue, sellOrders *OrderQueue) {

	sellOrders.Push(order)
	if buyOrders.Len() > 0 && buyOrders.Orders[0].Price >= order.Price {
		buyOrder := buyOrders.Pop().(*Order)
		if buyOrder.PendingShares > 0 {
			transaction := NewTransaction(buyOrder, order, order.Shares, buyOrder.Price)
			b.AddTransaction(transaction, b.wg)
			buyOrder.AddTransaction(transaction)
			order.AddTransaction(transaction)
			b.OrdersChanOut <- buyOrder
			b.OrdersChanOut <- order
			if buyOrder.PendingShares > 0 {
				sellOrders.Push(buyOrder)
			}
		}
	}
}
