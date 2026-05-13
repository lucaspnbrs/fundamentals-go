package main

import "fmt"

//The defer function looks like a defer in js, the serve to same purpose 
//Defer can be use in requests from the systems, inside status and error messages
//Function defer is useful in work with databases

func writeScreen() {
	fmt.Println("Writing in the screen 1!")
}

func writeScreen2() {
	fmt.Println("Writing in the screen 2!")
}

func voteIsValid(vote1, vote2 float32) bool{
	defer fmt.Println("Calculating your last results from last elections ")
	result := (vote1 + vote2)/2

	if result > 6 {
		return true
	}
	return false 
}

func main() {
	fmt.Println("New Defer Code")

    defer writeScreen()
	writeScreen2()

	fmt.Println("The candidate is valid!", voteIsValid(7, 13))
}