package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://v2.jokeapi.dev/joke/Programming?blacklistFlags=political&sample=sample1"

func main() {
	fmt.Println("Welcome to handling URLs in golang")

	fmt.Println(myurl)

	//parsing
	result, err := url.Parse(myurl)
	checkNilErr(err)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())
	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n",qparams)
	fmt.Println(qparams["blacklistFlags"])

	for _,val := range qparams{
		fmt.Println("Parameter is: ",val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "google.com",
		Path: "/learn",
		RawPath: "user=omkar",
	}
	fmt.Println(partsOfUrl.String())
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
