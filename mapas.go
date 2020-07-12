package main

import "fmt"

func main() {
	nuevoMapa := make(map[string]int)

	nuevoMapa["Intel"] = 15
	nuevoMapa["Google"] = 37

	otroMapa := make(map[string]Edades)

	otroMapa["juan"] = Edades{2001, 9, 31}

	fmt.Println(nuevoMapa, otroMapa)
}
