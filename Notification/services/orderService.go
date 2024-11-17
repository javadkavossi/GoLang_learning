package services

import (
	"fmt"
	"notification/entities"
	externalServices "notification/externalServices"
)

type OrderService struct {
	Notifier externalServices.Notifier
}


func (orderService *OrderService) CreateOrder(order *entities.Order) *entities.Order {

	fmt.Printf("Create order: %v\n", order)
	orderService.Notifier = externalServices.NewNotifier(order.NotificationType)
	orderService.Notifier.SendNotify(order.UserID, "order created")

	//Send sms or email

	return order
}

func  NewOrderService() *OrderService {
	return &OrderService{}
}
