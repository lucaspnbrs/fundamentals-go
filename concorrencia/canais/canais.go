package main

import (
	"fmt"
	"time"
)

func write(txt string, channels chan string ) {

	for i := 0; i < 8; i++ {
	    channels <- txt
		time.Sleep(time.Second)
	}

	close(channels)
}

func main() {
	channel := make( chan string )
	go write("Hello World!", channel)

	fmt.Println("Later this write function being executed ")

	for {
		message, opened := <- channel
		if !opened {
			break
		}
		fmt.Println(message)
	}

	fmt.Println("End Program")
}
