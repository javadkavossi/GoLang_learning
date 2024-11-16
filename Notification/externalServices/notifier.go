package externalServices

import (
	"notification/constant"
)

type Notifier interface {
	SendNotify(receiver string, message string)
}

func NewNotifier(notifierType constant.NotificationType) Notifier {
	switch notifierType {
	case constant.Email:
		return NewEmailService()
	case constant.Sms:
		return NewSmsService()
	default:
		return nil
	}
}
