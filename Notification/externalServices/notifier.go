package externalServices


type Notifier interface {
	SendNotify(receiver string,message string)
}


