package main

import "fmt"

// Structs
// Forma de criar sua própria estrutura de dados
// Personalizar de acordo com sua necessidade
// Podemos usar vários tipos diferentes

// type <nome da nossa estrutura> struct {<campos>}
type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	anoNasc   int
}

type Profissao struct {
	Pessoa
	Tipo string
}

func main() {
	fmt.Println(Pessoa{"Vitor", " Coelho", 26, 1999})
	fmt.Println(Pessoa{Nome: "Vitor", Sobrenome: "Coelho", Idade: 26, anoNasc: 1999})

	p1 := Pessoa{Nome: "Tel", Idade: 2}
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)

	p1.Idade = 3
	fmt.Println(p1.Idade)

	p2 := Pessoa{Nome: "Gabriel", Idade: 45}

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)

	// Strucs + map
	alunos := map[string][]Pessoa{}
	alunos["Programação"] = pessoas
	fmt.Println(alunos)

	// struct herdando campos de outra struct
	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Pessoa.Idade)
	fmt.Println(prof.Tipo)
}
