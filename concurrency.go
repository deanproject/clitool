package main

import (
	"fmt"
	"sync"
)

func call(from string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	for i := range 3 {
		fmt.Println(from, i+1)
	}
}

func main() {

	var wg sync.WaitGroup

	call("Direct : ", nil)
	wg.Add(1)
	go call("Go routine : ", &wg)

	go func(msg string) {
        fmt.Println(msg)
    }("going")

	wg.Wait()

	//	time.Sleep(10 * time.Millisecond)

}
