// main -> OrderService -> emailService || smsService

package main

import (
	"notification/constant"
	"notification/entities"
	"notification/services"
)

func main() {

	order1 := entities.Order{
		ID:               1,
		UserFullName:     "javadkavossi",
		UserID:           "0910505050",
		Price:            1000,
		Status:           true,
		NotificationType: constant.Email,
	}
	order2 := entities.Order{
		ID:               2,
		UserFullName:     "javadkavossi2",
		UserID:           "0910505050",
		Price:            1000,
		Status:           true,
		NotificationType: constant.Sms,
	}



	orderService := services.NewOrderService()
	orderService.CreateOrder(&order1)

	orderService2 := services.NewOrderService()
	orderService2.CreateOrder(&order2)
	// fmt.Println(order)
}
