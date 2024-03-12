package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/golang/learn?course=reactjs&paymentid=1234&price=0"

func main() {
	fmt.Println("Welcome to handling url!")
	fmt.Println("URL is:", myurl)

	// Parsing the URL
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// Query Parameters
	queryParams := result.Query()
	fmt.Printf("Type of Query Params: %T\n", queryParams)

	fmt.Println(queryParams["course"])

	for key, value := range queryParams {
		fmt.Println("Key:", key, "Value:", value)
	}

	// Creating a URL
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=abc",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}