package main

import "fmt"

// if / else
// se / se não

func main() {

	// numero := 5
	// if (condicional) {<ação>}
	// if numero == 1 { // true
	// 	fmt.Println("Valor é igual a 1")
	// } else {
	// 	fmt.Println("O valor não é igual a 1!")
	// }

	// if numero == 1 {
	// 	fmt.Println("O valor é igual a 1")
	// } else if numero == 2 {
	// 	fmt.Println("Valor igual a 2")
	// } else {
	// 	fmt.Println("Valor é diferente de 1 e 2 ")
	// }

	numero := 10
	if numero%2 == 0 {
		fmt.Printf("%d é par", numero)
	} else {
		fmt.Printf("%d e impar", numero)
	}

}
