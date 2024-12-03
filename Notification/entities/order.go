package entities

type Order struct {
	ID           int
	UserFullName string
	UserEmail    string
	UserPhone    string
	Price        float64
	Status       bool
}
