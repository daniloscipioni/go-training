package main

import (
	"fmt"
)

func main() {
	// Com go func
	// c := make(chan int)
	// go func() {
	// 	c <- 42
	// }()
	// fmt.Println(<-c)

	// Com buffer
	c := make(chan int, 5)
	
		c <- 42
		c <- 42
		c <- 42
		c <- 42
		c <- 42
	
	fmt.Println(<-c)

}