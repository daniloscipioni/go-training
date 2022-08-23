// Corrigindo race conditions com Atomic
package main

import ( 
	"fmt"
	"sync"
	"sync/atomic"
	"runtime"	
)
var wg sync.WaitGroup
var contador int64

func main()  {
	goroutine(100000)
	wg.Wait()
	fmt.Println("Total de goroutines\t", 100000, "\nTotal do contador:\t", contador)
}

func goroutine(i int)  {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
		// Locka as variaveis até que a go routine termine utilizar
		atomic.AddInt64(&contador, 1)
			//	v := contador
			runtime.Gosched()
			// v++
			// contador = v

			wg.Done()
			}()
			
	}
}