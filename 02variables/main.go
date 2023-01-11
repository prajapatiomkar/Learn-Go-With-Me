package main

import "fmt"
// First character Capital indicate that it's Public in scope 
const LoginToken string = "asdffdsa"

// jwtToken := 300000 //-> throw error
// var jwtToken = 300000 // -> Allowed
// var jwtToken int = 300000 // -> Allowed
func main() {
	var username string = "omkar"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var small uint8 = 255
	// var small uint8 = 256  Error Due to Overflow value
	// var small int = 256  At any use-case you can use
	fmt.Println(small)
	fmt.Printf("Variable is of type: %T \n", small)

	var smallFloat float32 = 255.00455665666 // > 255.00456
	// var smallFloat float64 = 255.00455665666 // > 255.00455665666
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// Default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// Implicit	type
	var twitter = "https://twitter.com/omkartwts"
	fmt.Println(twitter)
	// twitter = 3  //-> Once it declare by its own then we can't change it datatype

	//No var style

	numberOfUser := 300000 // You are allowed to declare in this syntax, But you can't declare it in Global or Outside the method it will give error
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
