package main

import "fmt"

func main() {
	var variavel string = "tipo string"
	variavel2 := "inferência de tipo" /*Uma forma reduzida de deixar implícito o tipo da variável, inferência*/
	fmt.Println(variavel, variavel2)

    //As a object of strings 
	var (
		variavel3 string = "variable of number three"
		variavel4 string = "variable of number four"
	)

	variable5, variable6 := "variable of number five", "variable of number six"//Another form of declaration with inference of type

	fmt.Println(variavel3, variavel4)
	fmt.Println(variable5, variable6)
	

	const constante string = "variavel constante"
	fmt.Println(constante)
	
}