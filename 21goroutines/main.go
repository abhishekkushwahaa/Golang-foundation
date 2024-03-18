package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string {"test"}

var wg sync.WaitGroup // WaitGroup is used to wait for the program to finish the go routines

var mut sync.Mutex // Mutex is used to lock the resource and unlock the resource

func main() {
	// // goroutine means a function that runs concurrently with other functions
	// go greeter("Hello G")

	// greeter("Sir")

	websiteList := []string{
		"https://golang.org",
		"https://go.dev",
		"https://lco.dev",
		"https://fb.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("All go routines are done")
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		// It is not better approach please use sync package for better approach
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string)  {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPs in endpoint")
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	
	fmt.Println("Status code is", res.StatusCode)
}