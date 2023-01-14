package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "time"
	// "crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 2

	// fmt.Println("The Sum is: ",mynumberOne+mynumberTwo)
	// fmt.Println("The Sum is: ",mynumberOne+int(mynumberTwo))

	// Random No
	// rand.Seed(30)
	// rand.Seed(time.Now().Unix())
	// fmt.Println(rand.Intn(5))

	// Random from Crypto
	myRandomNum,_ :=rand.Int(rand.Reader,big.NewInt(50))

	fmt.Println(myRandomNum)
}
