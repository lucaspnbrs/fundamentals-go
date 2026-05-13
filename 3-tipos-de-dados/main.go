package main

import (
	"errors"
	"fmt"
)

//if you not specify the type of the variable, it will be inferred by the value based on the your machine
//stack is concept of LAST IN, FIRST OUT
//uint not requires a sign, so it can only be positive, but it can store a larger range of values than int
//int32 can be called is rune, for example var number rune = 1000000 
//int8 can be called is byte, for example var number byte = 255 
//accept alias 
//Not exists the type char, but you can use string with only one character, for example var char string = "a"
//In the go exists a zero value, that is the default value of a variable when it is declared but not initialized, for example var number int will have the zero value of 0, var text string will have the zero value of "" (empty string), var boolean bool will have the zero value of false


func main() {
	/* int8, int16, int32, int64 */

	var number int16 = 10000
	fmt.Println(number)

/* 	float32, float64
 */
    var number2 float32 = 3.14
	fmt.Println(number2) 
    
	var floatNumber float64 = 3.1415926535897932384626433
	fmt.Println(floatNumber)

	var erro error = errors.New("This is an error")
	fmt.Println(erro)

	var boolean bool 
	fmt.Println(boolean)

}