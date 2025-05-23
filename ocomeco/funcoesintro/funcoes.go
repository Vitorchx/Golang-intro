package main

import "fmt"

// Funçao começando com letra minúscula:
// Função é Privada exemplo PrintName = Publica, printName = Privada
// Função ela só pode ser utilizada no próprio package
// Como utilizaria ela fora do package: main.PrintName(name)

func main() {
	//soma(42, 13)
	fmt.Println(soma(42, 13))

	sub := subtracao(10, 5)
	fmt.Println(sub)

	name := printName("Vitor")
	fmt.Println(name)
}

func subtracao(x int, y int) int {
	return x - y
}

func soma(x int, y int) int {
	return x + y
}

func printName(name string) string {
	return name
}
