package main

import (
    "fmt"
)

//In the go, a function can be returned as more than one value, is a common practice to return an error as second value, for example func mathCalculator(n1, n2 int8) (int8, int8, error) that returns the sum and the multiplication of two numbers and an error if the numbers are negative
// The form to returns one result than two or more is to use the blank identifier, for example resultSum, _ := mathCalculator(10, 20) that returns only the sum and ignore the multiplication
func soma(number1 int8, number2 int8) int8 {
	return number1 + number2
}

func mathCalculator(n1, n2 int8) (int8, int8){
	sum := n1 + n2
	mult := n1 * n2

	return sum, mult
}
func main(){
    somado := soma(10, 20)
	fmt.Println(somado)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	var result string = f("Ass one two three")

	fmt.Println(result)
	f("Jadson da a bunda")

	resultSum, _ := mathCalculator(10, 20)
	fmt.Println(resultSum )
}