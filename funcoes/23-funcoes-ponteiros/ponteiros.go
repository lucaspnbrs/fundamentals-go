package main 

import "fmt"

//Exists a way to manipulate pointers inside functions, this code proves that, with the function invertedSignWithPointer()
// HTTP requests in go requires a pointer functions, as shown in the package

func invertedSign(number int) int {
	return number * - 1
}

func inverterSignWithPointer(number * int) {
	*number = *number * -1
}

func main() {

	
	fmt.Println("Pointers with functions")
	number := 20
	fmt.Println(number)
	fmt.Println(invertedSign(number))

	number2 := 40
	fmt.Println(number2)
	inverterSignWithPointer(&number2)
	fmt.Println(number2)
}