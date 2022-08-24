package main

import "fmt"

func main()  {
	
	canal := make(chan int)

	c := colocaValoresNoCanal(canal)

	demonstraValores(c)

	fmt.Println("Terminou!")

}

func colocaValoresNoCanal(canal chan<- int) <-chan int   {
	c:= make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
		close(c)
		canal <- 0
	}()
	    return c
}

func demonstraValores(canal <-chan int)  {
	for v := range canal {
		fmt.Println(v)
	}
}