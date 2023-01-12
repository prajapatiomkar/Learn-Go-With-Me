package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our service:")

	// comma ok || err ok

	inupt, _ := reader.ReadString('\n')
	// inupt, err := reader.ReadString('\n') // err acts a catch

	fmt.Println("Thanks for rating, ", inupt)
	fmt.Printf("Types of this rating is %T", inupt)
}
