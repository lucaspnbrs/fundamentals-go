package main

import (
	"fmt"
	"time"
)

//Concurrency != Paralelism
//The usefull of "go" property increments in the code a use of concurrency

func main() {
	go write("Hello world in go!")
    write("Programming in golang")
}

func write(txt string) {
	for {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}

}