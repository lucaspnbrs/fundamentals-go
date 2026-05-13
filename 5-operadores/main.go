package main 

import "fmt"

//The operators in go dont have variances of another languages, they go don't accept results from different types, for example, you can't sum an int with a float, you need to convert one of them to the other type, and then do the operation.
// Logical operators: &&, ||, ! 
// Comparison operators: ==, !=, <, >, <=, >=
// Arithmetic operators: +, -, *, /, %
// Assignment operators: =, +=, -=, *=, /=, %=
//The golang has a unique operators, for example, the ++, and += or --
//No exists the ternary operator in go, but you can use an if statement to achieve the same result.


func main () {

	rightValue, negativeValue := true, false

	fmt.Println(rightValue && negativeValue)
	fmt.Println(rightValue || negativeValue)
	fmt.Println(!rightValue)

	numero := 10

	//Operadores unários
	numero++
	fmt.Println(numero)


	numero +=30
	fmt.Println(numero)

	numero--
	fmt.Println(numero)

	numero -= 20
	fmt.Println(numero)


}