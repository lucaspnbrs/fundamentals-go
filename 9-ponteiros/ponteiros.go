package main

import "fmt"

//Pointer is a memory reference 
//Pointer is form to modify values in memory, when don´t use copys 

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	variavel2++
	fmt.Println("Variable is",variavel1, variavel2)

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	variavel3 = 150
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro)

}