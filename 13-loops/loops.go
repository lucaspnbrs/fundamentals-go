package main

import (
	"fmt"
	"time"
)

//if i push the print to inside of scope, the variable accounts and displays
//The repeat loop allows that user iterates to maps, arrays, slices, numbers and variables

func main() {
	fmt.Println("Loops Structs ")

	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando I")
		time.Sleep(time.Second)
	}

	for j := 0; j < 10; j++ {
		fmt.Println("Implemented more one")
		time.Sleep(time.Second)
		
	}


	for k := 0; k < 10; k += 4 {
		fmt.Println("Implemented more one")
		time.Sleep(time.Second)
		fmt.Println(k)
	}

    //Form to apply for in arrays to iterate item to item
    alias := [3]string{ "Johns", "Malcom", "Silas"}

	//A form to use one variables than more two, apply _ in your code, this example show index, names, but returns just "names"
	for index, names := range alias {
		fmt.Println(index, names)
	}

	for _, palavra :=  range "PARALELEPIPEDO"  {
		fmt.Println(string(palavra))
	} 

	//Fibonacci in golang
	/* fibo1, fibo2 = 0, 1
	
	for fibo2 < 100 {
		fmt.Println(fibo2)
		fibo1, fibo2 = fibo1, fibo1 + fibo2
	} */

	/* for {
		fmt.Println("Infinite loop")
	} */
	
}