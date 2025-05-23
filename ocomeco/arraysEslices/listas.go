package main

import "fmt"

func main() {

	//Array - tamanho fixo
	var array [2]string
	array[0] = "Hello"
	array[1] = "Flamengo"
	fmt.Println(array[0], array[1])
	fmt.Println(array)

}

// 1 - Arrays e Slices: Homogeneos
// todos os elementos tem oo mesmo tipo
// [1, 2, 3, 4, , 6] - [] int
// ["strph", "vvitor", "golang"] - []string

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
