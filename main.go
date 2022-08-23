package main

import (
	"fmt"
)

func main() {

	nama := [10]string{"Reza", "Jery", "Burhann", "Andri", "Faisal", "Fian", "Alwi", "Warman", "Nugraha", "Muhendra"}

	for index, element := range nama {

		fmt.Println(index)
		fmt.Println(element)

	}

}
