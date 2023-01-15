package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in goLang")
	rand.Seed(time.Now().UnixNano())
	diceNumer := rand.Intn(6) + 1
	// fmt.Println("Value of dice is ", diceNumer)

	switch diceNumer {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move to 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
	case 4:
		fmt.Println("You can move to 4 spot")
		fallthrough
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	default:
		fmt.Println("What was that!")
	}
}
