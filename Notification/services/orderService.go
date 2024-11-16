package services

import (
	"fmt"
	"notification/entities"
	externalServices "notification/externalServices"
)

type OrderService struct {
	Notifier externalServices.Notifier
}

func (o *OrderService) CreateOrder(order *entities.Order) *entities.Order {

	fmt.Printf("Create order: %v\n", order)
	o.Notifier = externalServices.NewNotifier(order.NotificationType)
	o.Notifier.SendNotify(order.UserID, "order created")

	//Send sms or email

	return order
}

func  NewOrderService() *OrderService {
	return &OrderService{}
}
