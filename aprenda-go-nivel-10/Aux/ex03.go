package main

import "fmt"

func main()  {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go mandaNumero(par, impar, quit)
	recebe(par, impar, quit)
}


func mandaNumero(par,impar chan int, quit chan bool)  {
	number := 50
	for i := 0; i < number; i++ {
		if i%2 ==0 {
			par<-i
		}else{
			impar<-i
		}
	}
	close(par)
	close(impar)
	quit <- true
}

func recebe(par,impar chan int, quit chan bool)  {
for{
	select{
		// Se o elemento estiver no canal Par
	case v:= <-par:
		fmt.Println("Par:", v)
		// Se o elemento estiver no canal Impar
	case v:= <-impar:
		fmt.Println("Impar:", v)
	case <-quit:
		return
	}
}
}