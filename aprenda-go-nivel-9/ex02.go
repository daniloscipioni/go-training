package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	idade int
}

type cachorro struct{
	nome string
	sobrenome string
	idade int
}

type humanos interface{
	falar()
}

func(p *pessoa) falar(){
	fmt.Println(p.nome, "Fala coisas")
}

func(c *cachorro) falar(){
	fmt.Println(c.nome, "Late para coisas")
}

func dizerAlgumaCoisa(h humanos)  {
	h.falar()
}

func main()  {
	pessoa:= pessoa{"Danilo", "Scipioni", 32}
	cachorro:= cachorro{"Frederico", "silva", 1}

     pessoa.falar() // <-- Shortcut para (&pessoa).falar() -- Referencia https://go.dev/ref/spec#Calls
	(&pessoa).falar() // <-- forma correta
	(&cachorro).falar()

	dizerAlgumaCoisa(&pessoa)
	dizerAlgumaCoisa(&cachorro)
     
	// Não funciona
	//dizerAlgumaCoisa(pessoa) <-- Isso não é atalho para dizerAlgumaCoisa(&pessoa)
}