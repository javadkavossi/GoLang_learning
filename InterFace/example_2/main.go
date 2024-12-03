package main

import "fmt"

type Runner interface {
	Run()
}

type Walker interface {
	Walk()
}

type Shooter interface {
	Shoot()
}

type Player struct {
	Name     string
	Age      int
	Gender   string
	Height   int
	Position string
}

func main() {

	player1 := &Player{
		Name:     "Alireza",
		Age:      20,
		Gender:   "Male",
		Height:   180,
		Position: "Midfielder",
	}

	player1.Run(10)
	player1.Walk()
	player1.Shoot()

}

func (player *Player) Run(speed int) {
	fmt.Println("Player is Running at", speed, "km/h")
}

func (player *Player) Walk() {
	fmt.Println("Player is Walking")
}

func (player *Player) Shoot() {
	fmt.Println("Player is Shooting")
}
