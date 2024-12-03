// main -> OrderService -> emailService || smsService

package main

import (
	"notification/entities"
	"notification/externalServices"
	"notification/services"
)

func main() {

	order := entities.Order{
		ID:           1,
		UserFullName: "javadkavossi",
		UserEmail:    "jkavossi@gmail.com",
		UserPhone:    "0910505050",
		Price:        1000,
		Status:       true,
	}
	orderService := services.NewOrderService(externalServices.NewEmailService())
	orderService.CreateOrder(&order)
	// fmt.Println(order)

}
