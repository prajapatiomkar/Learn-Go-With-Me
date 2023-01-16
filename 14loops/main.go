package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to loops in golang ")

	days := []string{"Sunday","Tuesday","Webnesday","Friday","Saturday"}

	fmt.Println(days)

	for d:=0;d<len(days);d++{
		fmt.Println(days[d])
	}

	for i:=range days{
		fmt.Println(days[i])
	}

	for index,value := range days{
		fmt.Printf("index = %v value = %v\n",index,value)
	}

	for _,value := range days{
		fmt.Printf("index = none value = %v\n",value)
	}

	rougueValue := 1

	for rougueValue <10{
		if rougueValue == 2{
			goto loc
		}
		if rougueValue == 6 {
			rougueValue++
			continue
		}
		if rougueValue == 6 {
			break
		}

		fmt.Println("value is :",rougueValue)
		// ++rougueValue invalid 
		rougueValue++
	}

	loc:
	fmt.Println("Jumping at sample.co")
}