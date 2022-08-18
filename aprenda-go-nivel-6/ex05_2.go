package main

import (
	"fmt"
	   )


type imposto interface {
	impostoGeral() float64
	devolveValor(i float64) float64
}

type sudeste struct{
	icms, iptu, ipva, cofins float64 
}

type centroOeste struct{
	iva, ipi float64
}

type norte struct{
	icms, iptu float64
}

type nordeste struct{
	ipva, cofins, iptu float64
}

type sul struct{
	iva, iptu float64
}

func (s sudeste) impostoGeral() float64  {
	fator1 := s.icms * s.iptu
	fator2 := s.ipva * s.cofins
	fmt.Print("Sudeste: ")
	return fator1 + fator2
}


func (c centroOeste) impostoGeral() float64  {
	fmt.Print("Centro Oeste: ")
	return c.iva * c.ipi * 2
}

func (n norte) impostoGeral() float64  {
	fmt.Print("Norte: ")
	return n.icms + n.iptu
}
func (n nordeste) impostoGeral() float64  {
	fator1 := n.ipva * n.iptu
	fator2 := fator1 * n.cofins
	fmt.Print("Nordeste: ")
	return fator1 * fator2
}
func (s sul) impostoGeral() float64  {
	fmt.Print("Sul: ")
	return s.iva * 2 * (s.iptu/2)
}
/////////////////////////////////////////
func (c sudeste) devolveValor(i float64) float64  {
	fmt.Print("Sudeste Devolve: ")
	return i * 0.20
}

func (c centroOeste) devolveValor(i float64) float64  {
	fmt.Print("Centro Oeste Devolve: ")
	return i * 0.20
}


func (c norte) devolveValor(i float64) float64  {
	fmt.Print("Norte Devolve: ")
	return i * 0.20
}

func (c nordeste) devolveValor(i float64) float64  {
	fmt.Print("Nordeste Devolve: ")
	return i * 0.20
}

func (c sul) devolveValor(i float64) float64  {
	fmt.Print("Sul Devolve: ")
	return i * 0.20
}

func calcImposto(i imposto) float64  {
	fmt.Println(i)
	fmt.Println(i.impostoGeral())
	return i.impostoGeral()
}
func calcDevolucao(i imposto, valor float64)  {
	
	fmt.Println(i.devolveValor(valor))
}



func main()  {
	
	sudeste := sudeste{icms: 5.25, iptu:18.5, ipva: 3.65 ,cofins: 5.98}
	centrooeste := centroOeste{iva: 5.25, ipi:18.5}
	norte := norte{icms: 5.25, iptu:18.5}
	nordeste := nordeste{ipva: 5.25, cofins:18.5, iptu: 3.85}
	sul := sul{iva: 5.25, iptu:18.5}
	
	calcDevolucao(sudeste, calcImposto(sudeste))
	calcDevolucao(centrooeste, calcImposto(centrooeste))
	calcDevolucao(norte, calcImposto(norte))
	calcDevolucao(nordeste, calcImposto(nordeste))
	calcDevolucao(sul, calcImposto(sul))

}