package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup 

func main()  {
	newGoroutine(100)
	wg.Wait()
}

func newGoroutine(i int)  {
	wg.Add(i)

	for j := 0; j < i; j++ {
		x := j
		go func(i int)  {
			fmt.Println("Go Routine number:", i)
			wg.Done()
		}(x)
	}

}