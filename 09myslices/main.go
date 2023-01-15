package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{}
	fmt.Printf("Type of fruits list %T", fruitList)

	fruitList = append(fruitList, "fruit1", "fruit2", "fruit3", "fruit4", "fruit5")

	fruitList = append(fruitList[1:4])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 456
	highScore[1] = 852
	highScore[2] = 258
	highScore[3] = 321

	highScore = append(highScore, 741, 951)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	// how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)	
	var index int = 2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)
}
