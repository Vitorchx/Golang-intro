package main

import "fmt"

func main() {
	idade := map[string]int{}
	idade["Vitor"] = 26
	idade["Bruno"] = 4
	fmt.Println(idade)
	fmt.Println(idade["Vitor"])
	fmt.Println(idade["Bruno"])

	anoNasc := map[string]int{
		"Vitor": 1999,
		"Bruno": 2022,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["Vitor"])
	fmt.Println(anoNasc["Bruno"])

	anoNasc["Flamengo"] = 1981
	fmt.Println(anoNasc)

}
