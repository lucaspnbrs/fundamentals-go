package main

import "fmt"

//it's possible to cause a system crash, that's nice!!

func recoverExecution()  {
	if r := recover(); r != nil {
		fmt.Println("Function recovered with success")
	}
}

func voteIsValid(vote1, vote2 float32) bool {
	defer fmt.Println("Calculating your last results from last elections ")
	defer recoverExecution()
	result := (vote1 + vote2) / 2

	if result > 6 {
		return true
	} else if result < 6 {
		return false
	}

	panic("Fuck, The result é exactly 6!!")
}

func main() {
	fmt.Println("Learn functions panic and recover")

	fmt.Println(voteIsValid(6, 6))
}
