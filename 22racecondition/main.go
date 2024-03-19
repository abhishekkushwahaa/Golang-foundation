package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition is a situation where two or more threads or processes read or write a shared variable and the final result depends on the timing of how the threads are scheduled.")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(3)
	go func (wg *sync.WaitGroup, mut *sync.RWMutex)  {
		fmt.Println("First player to score")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()	
		}(wg, mut)
	go func (wg *sync.WaitGroup, mut *sync.RWMutex)  {
		fmt.Println("Second player to score")
		mut.Lock()
		score = append(score, 2)	
		mut.Unlock()
		wg.Done()	
	}(wg, mut)
	go func (wg *sync.WaitGroup, mut *sync.RWMutex)  {
		fmt.Println("Third player to score")
		mut.Lock()
		score = append(score, 3)	
		mut.Unlock()
		wg.Done()	
	}(wg, mut)

	wg.Wait()
	fmt.Println("Final score: ", score)
}