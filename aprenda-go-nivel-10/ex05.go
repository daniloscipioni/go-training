package main

import "fmt"

func main()  {
	c := make(chan int, 1)

	c <- 42

	
	v, ok := <-c
	// Saida v = 42 e ok = true
	fmt.Println(v, ok)
	
	
	close(c)
	
	// V, ok depois de fechar o canal
	// Saida v = 0 e ok = false
	v, ok = <-c
	fmt.Println(v, ok)
	
}