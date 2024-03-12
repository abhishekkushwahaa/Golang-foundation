package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang!!")

	content := "This needs to go in a file"

	file, err := os.Create("./file.txt")

	checkNilErr(err);

	defer file.Close();
	file.WriteString(content)

	readFile("./file.txt")
}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err);

	fmt.Println("Text data inside the file is:", string(databyte))
}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}