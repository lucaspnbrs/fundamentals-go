package main

import "fmt"

//way to save code

func mathCalculators(n1, n2 int) (sum int, subtration int) {
	sum = n1 + n2
	subtration = n1 - n2
	return 
}


func main() {
	fmt.Println("Funções de retorno nomeado")


	sum, subtration := mathCalculators(10, 20)
	fmt.Println(sum, subtration)
}