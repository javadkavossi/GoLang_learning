package entities

import "notification/constant"

// type NotificationType string

// const (
// 	Email NotificationType = "email"
// 	Sms   NotificationType = "sms"
// )

type Order struct {
	ID               int
	UserFullName     string
	UserID           string
	Price            float64
	Status           bool
	NotificationType constant.NotificationType
}
