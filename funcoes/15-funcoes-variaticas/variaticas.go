package main

import "fmt"

//Variatic functions, you have a choice in amount numbers to use this function, or don't pass numbers if you want
//The range and the numbers come on to the variables in "for", for = variable create, range = variable instance
// Allows a one variatic pattern for function and he must be the last parameter

func variatics(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number	
	}
	return total
}

func main() {
	fmt.Println("Variatics functions")

	total := variatics(40, 50, 533, 55, 43, 85, 12, 9)
	fmt.Println(total)


}