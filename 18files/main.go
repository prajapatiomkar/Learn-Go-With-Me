package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - Sample.in"
	file, err := os.Create("./mylcogofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	length, err := io.WriteString(file, content)

	checkNilErr(err)
	
	fmt.Println("length is: ", length)

	defer file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}
// To avoid repetition
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
