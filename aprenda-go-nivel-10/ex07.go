package main

import "fmt"

func main() {

	canal := make(chan int)

	for i := 1; i <= 10; i++ {
		go func() {
			for k := 1; k <= 10; k++ {
				canal <- k
			}
		}()

	}

	for t := 1; t <= 100; t++ {
		fmt.Println("Valor:", <-canal)
	}

	fmt.Println("Terminou!")

}

