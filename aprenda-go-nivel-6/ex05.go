package main

import (
	"fmt"
	"math"
	)

type square struct{
	lado float64
}

type circle struct{
	raio float64
}

func (s circle) area(){
	fmt.Println(math.Pi * 2 * s.raio)
}

func (c square) area() {
	fmt.Println(c.lado * c.lado)
}


type info interface {
	area()
}


func medida(i info){
	i.area()
}

func main()  {
	c := circle{raio: 5.0}
	s := square{lado: 25.25}

	medida(c)
	medida(s)
}