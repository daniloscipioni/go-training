package main
import "fmt"

func main()  {
	x := 500
	a:= make(chan int)
	b:= make(chan int)

	go func(x int) {
		for i := 0; i < x; i++ {
			a <- i
		}
	}(x/2)
	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i
		}
	}(x/2)

	for i := 0; i < x; i++ {
		select{
			case v := <-a:
				fmt.Println("Canal A:", v)
			case v := <-b:
				fmt.Println("Canal B:", v)
		}
	}
}

