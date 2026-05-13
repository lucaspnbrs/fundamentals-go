package main

import "fmt"
//Go requires all positional or all named fields when initializing a struct, you can't mix them, for example, you can't do this:
type pessoa struct {
	nome string
	sobrenome string
	idade int8
	altura float32
}

type cidadao struct {
	pessoa
	cpf string
	partido string
	numeroEleitoral int8
}

func main() {
	fmt.Println("Inheritance in Go is achieved through composition, not through traditional class-based inheritance.")

	p1 := pessoa{
		nome: "George Saraiva",
		sobrenome: "Dantas",
		idade: 34,
		altura: 1.80,
	}

	fmt.Println(p1)

	c1 := cidadao{
		pessoa: p1,
		cpf: "123.456.789-00",
		partido: "União Brasil",
		numeroEleitoral: 16,
	}

	fmt.Println(c1)

}