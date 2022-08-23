package main

import ( 
	"fmt"
	"sync"
	"runtime"	
)

var wg sync.WaitGroup
var contador int

func main()  {
	goroutine(10)
	wg.Wait()
	fmt.Println("Total de goroutines\t", 10, "\nTotal do contador:\t", contador)
}

func goroutine(i int)  {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
			}()
			
	}
}