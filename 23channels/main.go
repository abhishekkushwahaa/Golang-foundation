package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in GoLang")

	myChannel := make(chan int, 1)
	wg := &sync.WaitGroup{}

	// myChannel <- 5 
	// fmt.Println(<-myChannel)

	wg.Add(2)

	// Receive only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(myChannel)
		val, isOpen := <-myChannel
		fmt.Println(val, isOpen)

	    fmt.Println(<-myChannel)
	    fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel, wg)

	// Send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChannel <- 5
		myChannel <- 6
		close(myChannel)
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}