package main

import "fmt"
//its possible in go, declarate variables inside of controller structs

func main() {
	fmt.Println("Controller struct")

    number := 23

	if number > 18 {
		fmt.Println("You is old man")
	} else {
		fmt.Println("You is young")
	}

}