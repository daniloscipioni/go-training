package main

import "fmt"

func main()  {
	canal := make(chan int)
	quit := make(chan int)
	go enviaParaCanal(canal, quit)
	recebeQuit(canal, quit)
}

func recebeQuit(canal, quit chan int )  {
	for i := 0; i < 50; i++ {
		fmt.Println("Recebido", <-canal)
	}
	quit <- 0
}
func enviaParaCanal(canal, quit chan int)  {
	qq := 1
	for{
		select{
		case canal <- qq:
			qq++
		case <-quit:
			// Se houver coisa no canal QUIT, para o select
			return
		}
	}
}