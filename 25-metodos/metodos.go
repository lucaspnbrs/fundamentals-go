package main

import "fmt"

type user struct {
	name string
	age uint8
}

//Reference to struct, for simulate a method of OOP
//Go not have OOP, but have support to paradigm 

func (u user) salvar(){
	fmt.Println("Account created with success!", u.name)
}

func (u user) majority() bool {
	return u.age >= 18
}

func main() {
	fmt.Println("Another way, Methods in golang it's possible")

	user1 := user{ "Josh", 23}
	fmt.Println(user1)
	user1.salvar()
	maioridade := user1.majority()
	fmt.Println(maioridade)

}
