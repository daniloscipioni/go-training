package main

import "fmt"

func main()  {
	// inciando um canal multidirecional
	c := make(chan int)

	go meuloop(10, c)

	prints(c)
}

// meuloop recebe um valor inteiro e um canal de envio (send)
func meuloop(t int, s chan<- int)  {
	for i := 0; i < t; i++ {
			s <- i
			 }

	close(s)
}

// prints recebe um canal de recebimento (receive)
func prints(r <-chan int)  {
	for v := range r {
		fmt.Println("channel", v)
	}
}