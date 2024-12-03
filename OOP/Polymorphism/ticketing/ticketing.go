package main

import "fmt"

type BusTicket struct {
	Id                int
	DepartureCity     string
	ArrivalCity       string
	DepartureTime     string
	BusType           string
	PassengerName     string
	DepartureTerminal string
	ArrivalTerminal   string
	NationalCode      string
	Price             int
}

type PlaneTicket struct {
	Id               int
	DepartureAirport string
	ArrivalAirport   string
	DepartureTime    string
	ArrivalTime      string
	PlaneType        string
	PassengerName    string
	PassportId       string
	PassengerType    string
	Price            int
}

type TrainTicket struct {
	Id                int
	DepartureCity     string
	ArrivalCity       string
	DepartureTime     string
	TrainType         string
	PassengerName     string
	DepartureTerminal string
	ArrivalTerminal   string
	NationalCode      string
	Price             int
}

func main() {

	busTicket := BusTicket{
		Id:                1,
		DepartureCity:     "Tehran",
		ArrivalCity:       "Shiraz",
		DepartureTime:     "10:00",
		BusType:           "Bus",
		PassengerName:     "Javad",
		DepartureTerminal: "Terminal1",
		ArrivalTerminal:   "Terminal2",
		NationalCode:      "123456789",
		Price:             10000,
	}

	planeTicket := PlaneTicket{
		Id:               2,
		DepartureAirport: "Shiraz",
		ArrivalAirport:   "Tehran",
		DepartureTime:    "10:00",
		ArrivalTime:      "12:00",
		PlaneType:        "Boing",
		PassengerName:    "Ali",
		PassportId:       "123456789",
		PassengerType:    "Business",
		Price:            2000000,
	}

	TrainTicket := TrainTicket{
		Id:                3,
		DepartureCity:     "Tehran",
		ArrivalCity:       "Shiraz",
		DepartureTime:     "10:00",
		TrainType:         "Train",
		PassengerName:     "Javad",
		DepartureTerminal: "Terminal1",
		ArrivalTerminal:   "Terminal2",
		NationalCode:      "123456789",
		Price:             10000,
	}

	// busTicket.PrintTicket()
	// planeTicket.PrintTicket()

	printer := []TicketPrinter{busTicket, planeTicket, TrainTicket}  

	for _, printer := range printer {
		printer.PrintTicket()
	}
}

type TicketPrinter interface {
	PrintTicket()
}

func (ticket BusTicket) PrintTicket() {
	fmt.Println("\nBus Ticket : \n" + "PassengerName :" + ticket.PassengerName + "\n DepartureTime :" + ticket.DepartureTime + " \n DepartureCity:" + ticket.DepartureCity + "\nArrivalCity : " + ticket.ArrivalCity + "\n")
}

func (ticket PlaneTicket) PrintTicket() {
	fmt.Println("\n Plane Ticket : \n  " + "PassengerName :" + ticket.PassengerName + "\n DepartureTime :" + ticket.DepartureTime + "\n ArrivalTime:" + ticket.ArrivalTime + " \n DepartureAirport:" + ticket.DepartureAirport + "\nArrivalAirport : " + ticket.ArrivalAirport + "\n")
}

func (ticket TrainTicket) PrintTicket() {
	fmt.Println("\n  Train Ticket : \n " + "PassengerName :" + ticket.PassengerName + "\n DepartureTime :" + ticket.DepartureTime + "\n DepartureCity:" + ticket.DepartureCity + "\nArrivalCity : " + ticket.ArrivalCity + "\n")
}
