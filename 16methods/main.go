package main

import "fmt"

func main() {
	fmt.Println("Structs in goLang")
	// No inheritance in goLang; No Super or
	omkar := User{"omkar","omkar@9833gmail.com",true,19}
	fmt.Printf("Omkar user %+v\n",omkar)
	fmt.Printf("Name is %v and email is %v",omkar.Name,omkar.Email)

	omkar.GetStatus()
	omkar.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("Is user active",u.Status)
}

func (u User) NewMail(){
	u.Email = "sample@gmail.com"
	fmt.Println("Email of this user is: ",u.Email)
}