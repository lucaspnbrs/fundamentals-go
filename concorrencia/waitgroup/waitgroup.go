package main

import (
	"fmt"
	"sync"
	"time"
)

//WaitGroup is a form to sync goroutines in code for execute at the same time

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		write("Hello world with wait groups")
		waitGroup.Done()  //1

	}()

	go func() {
		write("Programming with golang")
		waitGroup.Done() //1
	}()

	waitGroup.Wait()

}

func write(txt string) {
	for i := 0; i < 5; i++ {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}