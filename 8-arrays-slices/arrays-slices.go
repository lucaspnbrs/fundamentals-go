package main

import "fmt"

//Arrays and Slices in go are used to store multiple values, they are similar but they have some differences
//reflect.TypeOf() returns the type of the variable
//Slices born from instances to arrays or intern arrays

func main() {
	fmt.Println("Arrays and Slices in Go") 

	var array1 [5]string
	array1[0] = "Data"
	fmt.Println(array1)

	array2 := [5]string{"Data", "Atendimentos", "Demandas", "Projetos", "Ações"}
	fmt.Println(array2[3])

	array3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array3)
	fmt.Println(array3[2])

	slice := []int{
		1, 10, 22, 44, 5, 64, 64,
	}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:4]
	fmt.Println(slice2)

	array2[2] = "Altered Position"
	fmt.Println(slice2)

	// Intern Arrays
	// Make is native function

	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
    fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

   // arrays is a finite list and slice is a infinite length
}