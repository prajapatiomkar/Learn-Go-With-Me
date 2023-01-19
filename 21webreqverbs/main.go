package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:1111/get"

	response, err := http.Get(myurl)
	checkErrNil(err)

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content lenght is: ", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
	// fmt.Println(content)
	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:1111/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Sample course",
			"price":100,
			"platform":"LCO.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	checkErrNil(err)

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:1111/postform"
	data := url.Values{}
	data.Add("fname", "omkar")
	data.Add("lname", "prajapati")
	data.Add("email", "om@123.com")

	response, err := http.PostForm(myurl, data)
	checkErrNil(err)
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	checkErrNil(err)
	fmt.Println(string(content))

}

func checkErrNil(err error) {
	if err != nil {
		panic(err)
	}
}
