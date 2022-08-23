
// Corrigindo race conditions com Mutex
package main

import ( 
	"fmt"
	"sync"
	"runtime"	
)

var wg sync.WaitGroup
// Race conditions
var mutex sync.Mutex
var contador int

func main()  {
	goroutine(1000)
	wg.Wait()
	fmt.Println("Total de goroutines\t", 1000, "\nTotal do contador:\t", contador)
}

func goroutine(i int)  {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			// Locka as variaveis até que a go routine termine utilizar
			// Inicio do Lock
			mutex.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mutex.Unlock()
			// Fim do lock
			wg.Done()
			}()
			
	}
}