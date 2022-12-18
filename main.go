package main

import "fmt"

func main() {
	var dateOfBirth int
	switch {
	case dateOfBirth >= 1946 && dateOfBirth < 1965:
		fmt.Println("Привет, бумер!")
	case dateOfBirth >= 1965 && dateOfBirth < 1981:
		fmt.Println("Привет, представитель X!")
	case dateOfBirth >= 1981 && dateOfBirth < 1997:
		fmt.Println("Привет, миллениал!")
	case dateOfBirth >= 1997 && dateOfBirth < 2013:
		fmt.Println("Привет, зумер!")
	case dateOfBirth >= 2013:
		fmt.Println("Привет, альфа!")
	default:
		fmt.Println("Приветствую, о великий старец!")
	}
}
