package main

import (
	"encoding/json"
	"fmt"
)

// É possível alterar o valor da chave retornada para o type	
type Pessoa struct {
	First   string   `json:"FirstChanged"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	var pessoas []Pessoa
	s := []byte(`[{"FirstChanged":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"FirstChanged":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"FirstChanged":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`)
	err := json.Unmarshal(s, &pessoas)

	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Printf("%+v\n", pessoas)
	fmt.Println("Quantidade =", len(pessoas))

	for _, value := range pessoas {
		fmt.Println(value.First)
	}

}