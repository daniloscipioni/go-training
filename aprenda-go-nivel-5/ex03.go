package main

import "fmt"

type veiculo struct{
	portas int
	cor string
}

type caminhonete struct{
	veiculo
	traçãoNasQuatro bool
}

type sedan struct{
	veiculo
	modeloLuxo bool
}


func main()  {
	
	veiculoPadrao:= veiculo{
		portas: 4,
		cor: "Preta",
	}

	caminhonete := caminhonete{
    	veiculo: veiculoPadrao,
		traçãoNasQuatro: true,
	}

	sedan := sedan{
    	veiculo: veiculo{
			portas: 2,
			cor: "Branco",
		},
		modeloLuxo: false,
	}


	fmt.Println(caminhonete)	
	fmt.Println(sedan)	
	fmt.Println(caminhonete.portas)	
	fmt.Println(sedan.cor)	



}