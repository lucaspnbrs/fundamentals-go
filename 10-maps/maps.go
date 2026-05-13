package main

import "fmt"

func main() {

	fmt.Println("Who write in go give the ass, This a lesson of Maps")

	//inside the [] to be called a type of keys and outside, the type of values
	//keep the same types when your works with maps in golang

	usuario := map[string]string {
		"nome": "Bunderus",
		"sobrenome": "Serotonios",
	}

	//alinhated map

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro":"Martelo",
			"segundo":"Do Thor",
		},
		"categoria": {
			"cidade":"asgard",
			"reino":"asgardianos",
		},
	}

    //Delete keys
	/* delete(usuario2, "nome") */

	fmt.Println(usuario)
	fmt.Println(usuario2)
}
