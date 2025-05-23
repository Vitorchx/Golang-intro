package main

import "fmt"

func main() {
	//Variáveis
	var name string
	name = "Vitor"
	fmt.Println(name)

	name = "steph"
	fmt.Println(name)

	var idade int
	idade = 4
	fmt.Println(idade)

	var a = "steph"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b)
	fmt.Println(c)

	var d = true
	fmt.Println(d)

	//exemplo de inferencia de tipos cooolocando :=
	f := "apple"
	fmt.Println(f)

	//Constante
	const idadeVitor = 4
	fmt.Println(idadeVitor)

}
