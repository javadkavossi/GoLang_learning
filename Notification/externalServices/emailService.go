package externalServices

import (
	"fmt"
	"notification/entities"
)

type EmailService struct{}

func (e *EmailService) SendMessage(order *entities.Order) {
	fmt.Printf("email Sent : %v \n ", order)
}

func (e *EmailService) SendNotify(receiver string, message string) {
	fmt.Printf("sms Sent : %v Message: %v \n ", receiver, message)
}

func NewEmailService() *EmailService {
	return &EmailService{}
}
