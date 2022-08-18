package main

import (
	"fmt"
	"math"
	   )


type geometria interface{
	area() float64
	perim() float64
}

type retangulo struct{
	largura, altura float64
}	   

type circulo struct{
	raio float64
}

func (r retangulo) area() float64  {
	return r.largura * r.altura
}

func (r retangulo) perim() float64 {
	return 2*r.largura + 2*r.altura
}

func (c circulo) area() float64 {
	return math.Pi *  c.raio * c.raio
}

func (c circulo) perim() float64 {
	return 2 * math.Pi * c.raio
}

func medir(g geometria)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main()  {

	r := retangulo{largura:9, altura:3.65}
	c := circulo{raio: 25.8}

	medir(r)
	medir(c)
	
}

