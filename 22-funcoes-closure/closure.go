package main 

import "fmt"

//functions closure are functions that make reference a variables out of context

func closure() func() {
	txt := "Inside the closure function"
	fmt.Println(txt)

	function := func() {
		fmt.Println(txt)
	}

	return function
}

func main() {
	fmt.Println("Closure functions")
	txt := "Inside the main function"
	fmt.Println("New function", txt)

	newFunction := closure() 
	newFunction()
}