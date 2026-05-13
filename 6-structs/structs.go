package main

import "fmt"

type atendimento struct {
	cidadao string
	eleitor string
	lideranca string
	servidor string
	apoiador apoiador
	municipe bool
	canal string
	status bool
	descricao string
}
//Is a polimorfism model? No, it's not, because the struct atendimento doesn't have a method, it's just a data structure that holds information about an atendimento, and the struct apoiador is another data structure that holds information about an apoiador, and the struct atendimento has a field do tipo apoiador, that is another struct, but they don't have any methods, so they are not polimorfism models, they are just data structures that hold information.
//This type how it works like a foregn key in other languages, you can use it to create a relationship between two structs, in this case, the struct atendimento has a field do tipo apoiador, that is another struct, and you can access the fields of the struct apoiador through the struct atendimento, for example, u.apoiador.nome, u.apoiador.partido, u.apoiador.idade.
type apoiador struct {
	nome string
	partido string
	idade int8
}



func main() {
	fmt.Println("Structs in golang")
	var u atendimento // like a instance of a class in other languages, but in go we call it struct
	u.cidadao = "Lucas"
	u.eleitor = "Sim"
	u.lideranca = "Não"
	u.servidor = "Sim"
	u.apoiador = apoiador{ nome: "João", partido: "UNIÃO", idade: 35}
	u.municipe = true
	u.canal = "Telefone"
	u.status = true
	u.descricao = "Atendimento realizado com sucesso"

	fmt.Println(u)

	usuario2 := atendimento{
		"Maria",
		"Sim",
		"Não",
		"Sim",
		apoiador{ nome: "Jacobina", partido: "DORES", idade: 28},
		true,
		"E-mail",
		true,
		"Atendimento realizado com sucesso",
	}
	fmt.Println(usuario2)

	usuario3 := atendimento{
		cidadao: "Joana",
	}
	fmt.Println(usuario3)
}
