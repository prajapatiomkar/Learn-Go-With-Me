package main

import "fmt"

func main() {
	fmt.Println("If else in goLang")
	logInCount := 9
	var result string

	if logInCount < 10 {
		result = "Regular User"
	} else if logInCount > 10 {
		result = "Watch Out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// if err != nil{
		
	// }
}
