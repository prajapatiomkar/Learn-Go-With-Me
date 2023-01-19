package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://v2.jokeapi.dev/joke/Any"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is of type: %T", response)

	defer response.Body.Close() //caller's responsibility to close the connection

	databyte, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println("Response: ", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
