package main

import (
	"comand-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("This is your ip!")

	application := app.Generate()

	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}