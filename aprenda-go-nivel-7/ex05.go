package main

import "fmt"

func zeroval(ival int)  {
	ival = 0
}

func zeroptr(iptr *int)  {
	*iptr = 0
}

func main()  {
	i := 1
	fmt.Println("inicial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)
	// Passa o caminho de memória da variavel i
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	
	
	fmt.Println("ponteiro:", &i)

}