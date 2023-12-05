package entity

type StatusOrder interface{
	Trade(order *Order, b *Book, buyOrders *OrderQueue, sellOrders *OrderQueue)
}