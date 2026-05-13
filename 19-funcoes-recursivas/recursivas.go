package main


import "fmt"

//Recursive is a concept when one function returns another function 

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position - 2) + fibonacci(position - 1)
}

func main() {
	fmt.Println("Recursive functions")

	position := uint(15)
	/* fmt.Println("This number occuped fibonacci position is", fibonacci(position)) */

	for i := uint(1); i <= position; i++ {
		fmt.Println(fibonacci(i))
	}
	
}