package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type sortByAge []user
func (a sortByAge)  Len() int { return len(a)}
func (a sortByAge)  Swap(i, j int) { a[i], a[j] =  a[j], a[i]}
func (a sortByAge)  Less(i, j int) bool { return a[i].Age < a[j].Age}


type sortByLast []user
func (a sortByLast)  Len() int { return len(a)}
func (a sortByLast)  Swap(i, j int) { a[i], a[j] =  a[j], a[i]}
func (a sortByLast)  Less(i, j int) bool { return a[i].Last < a[j].Last}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	for _, v := range users {
			sort.Strings(v.Sayings)
	}

	fmt.Println(users)
	sort.Sort(sortByAge(users))
	//fmt.Println("Ordenado por Idade:\n", users)

	sort.Sort(sortByLast(users))
	//fmt.Println("Ordenado por sobrenome:\n", users)
	harmoniosa(users)
}

func harmoniosa(x []user)  {
		for i, v := range x {
			fmt.Println(i, "\tIdade", v.Age, "\tNome Completo:", v.First, v.Last, "\n")
			for _, c := range v.Sayings {
				fmt.Println("\t", c)
			}
			fmt.Println("\n")
		}
}