package main
import "fmt"

// Ponteiros: 
//Quando passar o argumento, passar como &valor
//Quando passar como tipo no parametro, declarar como *valor
type carro struct{
	cor string
	ano int
}

func main()  {
		
	carro := carro{
		"branco",
		1989	,
	}
    
	// Variavel original
	fmt.Println(carro)
	// Tentando alterar o valor
    alterTest(carro)
	fmt.Println(carro)

	// Alterando o valor do local de memória
	alterTestPtr(&carro)
	fmt.Println(carro)
}

// Não altera pois está fora do escopo da variavel
func alterTest(c carro)  {
	c.ano = 1958
}

// Altera o valor do endereço de memória da variavel ano do tipo carro
func alterTestPtr(c *carro)  {
	(*c).ano = 2022
}