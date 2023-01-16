package main

import (
	"fmt"
)

func main() {
	greeter()
	fmt.Println("Welcome to function in golang")
	// greeterTwo()
	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proResult := proAdder(5,6,8,4,4)
	fmt.Println("Pro Result ",proResult)
	num,str := returnTwoVal()
	fmt.Printf("num %v str %v",num,str)
}

func greeter() {
	fmt.Println("Namstey from golang")
}
func greeterTwo() {
	fmt.Println("Another method")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0

	for _,value := range values{
		total+= value
	}
	return total
}

func returnTwoVal()(int,string){
	return 5,"sample"
}