package main

import "fmt"

type erroEspecial struct{
	test string
}

func (e erroEspecial) Error() string  {
	return fmt.Sprintf("erro aqui! e olha isso: %v", e.test)
}

func errAsParameter(e error)  {
	 fmt.Println("Passaram como argumento pra mim, o seguinte: ", e.(erroEspecial).test,
	"E no método Error eu tenho:", e)
}

func main()  {
	a := erroEspecial{"Testando"}
	errAsParameter(a)
}
