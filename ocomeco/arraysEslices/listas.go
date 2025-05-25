package main

import "fmt"

func main() {

	//Array - tamanho fixo
	var array [2]string
	array[0] = "Hello"
	array[1] = "Flamengo"
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[0:4])
	fmt.Println(numPrimos[1])

	numPares := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numPares)

	numPares = append(numPares, 14, 16)
	fmt.Println(numPares)

	//Slices

	//var slice []string
	slice := make([]string, 5)
	slice[0] = "Hello"
	slice[1] = "World"
	fmt.Println(slice[0], slice[1])
	fmt.Println(slice)

	fmt.Println(len(slice))

}

// LISTAS

// 1 - Arrays e Slices: Homogeneos
// todos os elementos tem oo mesmo tipo
// [1, 2, 3, 4, 5, 6] - [] int
// ["strph", "vitor", "golang"] - []string

// 2 - Maps: Heterogeneo
// pode misturar tipos
// estrutura chave - valor
// [key]  value
// chave tem tipo, e o valor pode ter outro
// map[istring]int
// {"steph": 28, "vitor": 4}
// map [string]string
// {"steph": "cardoso", "Vitor": "Coelho"}

// Array

// Tamanho fixo, de zero ou mais elementos do mesmo tipo
// Acessamo os valoores com indice: a[0], a[1]...
// Função embutida len() retorna o tamanho do array
// por conta do tamanho fixo, não é tão usado. Só em casos especificos

//Slice

// Tipo o array, mas sem tamnaho fixo
// Acessamos os valores coom indice: a[0], a[1]...
// Função embutida len() retorna o tamanho
// função append() usada para adiicionar valores
