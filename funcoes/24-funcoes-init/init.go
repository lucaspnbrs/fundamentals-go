package main 

import "fmt"

//it's possible just one init function for file, as shown in the file bellow
//comes before main

func main() {
	fmt.Println("Main function being executed")
}

func init() {
	fmt.Println("Main function being executed")
}