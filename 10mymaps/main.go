package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)
	// fmt.Printf("Type is %T",languages)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	//loops are interesting in goLang
	for key, value := range languages {
		fmt.Printf("key %v value %v \n", key, value)
	}
}
