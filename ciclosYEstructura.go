package main

import "fmt"

type Edades struct {
	Año, Mes, Dia int
}

func main() {
	var edad = Edades{
		Año: 1996,
		Mes: 12,
		Dia: 25,
	}

	var slice []Edades

	slice = append(slice, edad)

	edad.Año = 2002
	edad.Mes = 05
	edad.Dia = 12

	slice = append(slice, edad)
	fmt.Println(slice)
	var arreglo = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Println("Esto es un 'for':")
	for i := 0; i < len(arreglo); i++ { //Ciclo for
		fmt.Println(arreglo[i])
	}
	fmt.Println("Esto es un 'while':")
	j := 0
	for j < len(arreglo) { //Ciclo while
		fmt.Println(arreglo[j])
		j++
	}

	fmt.Println("Esto es un 'range':")
	for i, valores := range arreglo { //ciclo range
		fmt.Println(i, valores)
	}
}
