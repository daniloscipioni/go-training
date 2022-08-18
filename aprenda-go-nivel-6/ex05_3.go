package main

import "fmt"


type animal interface{
	comer() 
}

type cachorro struct{
	porte string
	raca string	
}

type gato struct{
	cor string
	siames bool
}

type dinossauro struct{
	carnivoro bool
	cor string
}

func (c cachorro) comer()  {
	fmt.Println("Cachorro é", c.porte, "da raça", c.raca)

}
func (g gato) comer()  {
	fmt.Println("Gato é", g.cor, "e é siames?", g.siames)


}
func (d dinossauro) comer()  {
	fmt.Println("Dinossauro é Carnivoro:", d.carnivoro, "da cor", d.cor)

}

func faz(a animal)  {
	a.comer()
}


func main()  {

	cachorro := cachorro{
		porte: "Grande",
		raca: "Poodle",
	}
	gato := gato{
		cor: "Branco",
		siames: true,
	}
	dinossauro := dinossauro{
		carnivoro: true,
		cor: "Verde",
	}

	faz(cachorro)
	faz(gato)
	faz(dinossauro)
}