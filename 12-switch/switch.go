package main

import "fmt"

//The name of variable come before the type, for example in func (n1 int)
//The variable fallthrough exits to jump cases in switch case struct
//Don't need to use the break clause

func weekend(n1 int) string{
		switch n1 {
		case 1:
			return "segunda"
		case 2:
			return "thursday"
		default:
			return "Invalid Number"
		}
	}

func weekend2(n1 int) string{
		switch {
		case n1 ==1:
			return "segunda"
		case n1 == 2:
			return "thursday"
		default:
			return "Invalid Number"
		}
	}

func main() {
	fmt.Println("How to make a switch case in golang")

	daysWeek := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	switch daysWeek[6] {
	case "Monday":
		fmt.Println("Segunda Feira")
		fallthrough
	case "Tuesday":
		fmt.Println("Terça feira")
	case "Wednesday":
		fmt.Println("Quarta feira")
	case "Thursday":
		fmt.Println("Quinta feira")
	case "Friday":
		fmt.Println("Sexta feira")
	case "Saturday":
		fmt.Println("Sábado")
	case "Sunday":
		fmt.Println("Domingo")
	}

	

	day := weekend(2)
	fmt.Println(day)

	day2 := weekend2(1)
	fmt.Println(day2)
}