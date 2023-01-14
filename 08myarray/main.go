package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var sampleList [5]string

	sampleList[0] = "sample 0"
	sampleList[1] = "sample 1"
	sampleList[3] = "sample 3"
	fmt.Println(sampleList)
	fmt.Println(len(sampleList))

	var numList = [3]string{"one", "two", "three"}
	fmt.Println(numList)
	fmt.Println(len(numList))
}
