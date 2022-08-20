// Documentação do package Sort
// https://pkg.go.dev/sort

package main

import (
	"fmt"
	"sort"
)


// type sortAsc []int
// func (a sortAsc)  Len() int { return len(a)}
// func (a sortAsc)  Swap(i, j int) { a[i], a[j] =  a[j], a[i]}
// func (a sortAsc)  Less(i, j int) bool { return a[i] < a[j]}

// type sortAlpha []string
// func (a sortAlpha)  Len() int { return len(a)}
// func (a sortAlpha)  Swap(i, j int) { a[i], a[j] =  a[j], a[i]}
// func (a sortAlpha)  Less(i, j int) bool { return a[i] < a[j]}

func main() {



	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	//Ordena um type customizado
	//sort.Sort(sortAsc(xi))
		
	// Ordena uma slice de Ints  
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	//Ordena um type customizado
	//sort.Sort(sortAlpha(xs))
	fmt.Println("Strings estão ordenadas", sort.StringsAreSorted(xs))
	// Ordena uma slice de Strings  
	sort.Strings(xs)
	fmt.Println(xs)
	fmt.Println("Strings estão ordenadas", sort.StringsAreSorted(xs))
}
