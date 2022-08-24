package main

import "fmt"

func main()  {
	canal := make(chan int)

	go colocaValoresNoCanal(canal)

	for v := range canal {
		fmt.Println(v)
	}
}

func colocaValoresNoCanal(canal chan int)  {
	for i := 1; i <= 100; i++ {
		canal <- i
	}
	close(canal)
	
}