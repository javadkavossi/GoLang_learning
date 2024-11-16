package externalServices

import (
	"fmt"
	// "notification/entities"
)

type SmsService struct{}

// func (e *SmsService) SendMessage(order *entities.Order) {
// 	fmt.Printf("sms Sent : %v \n ", order)
// }

func (e *SmsService) SendNotify(receiver string, message string) {
	fmt.Printf("sms Sent : %v Message: %v \n ", receiver, message)
}

func NewSmsService() *SmsService {
	return &SmsService{}
}
