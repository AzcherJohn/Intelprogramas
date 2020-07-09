package main

import (
	"fmt"
)

var nombreVar int = 56

type Alumnos struct {
	Nombre, ApellidMat, ApellidPat string
	Edad                           []Edades
	Hermanos                       bool
}

type Edades struct {
	Año, Mes, Dia int
}

func otros() {
	nuevaConexionDB()
	catalogoID("catalogo_familia", "Computadoras")
	resultadosQuery("catalogo_familia")
	/*var i int = 25
	var a, h, j float32 = 64.23, 59.21, 37.201*/

	/*	estado := "soltera"

		var alumno Alumnos

		const nombresss = "NOmbre"

		alumno.Edad = 12

		fmt.Println(estado)

		_, sumTotal := suma(25, 30)

		fmt.Println(sumTotal)*/
	//ciclos()
	//ejemploMaps()
}

/*
func suma(a, s int) (bool, int) {
	return a+s < 100, a + s
}*/

func Ciclos() {
	var edad Edades
	edad.Año = 1996
	edad.Mes = 12
	edad.Dia = 25

	var slice []Edades

	slice = append(slice, edad)

	edad.Año = 2002
	edad.Mes = 05
	edad.Dia = 12

	slice = append(slice, edad)
	fmt.Println(slice)
	/*var arreglo [10]string

	for i := 0; i < len(arreglo); i++ { //Ciclo for
		fmt.Println(arreglo[i])
	}
	i := 0
	for i < len(arreglo) { //Ciclo while
		fmt.Println(arreglo[i])
	}

	for i, valores := range arreglo { //ciclo range
		fmt.Println(i, valores)
	}*/
}

func ejemploMaps() {
	nuevoMapa := make(map[string]int)

	nuevoMapa["Intel"] = 15
	nuevoMapa["Google"] = 37

	otroMapa := make(map[string]Edades)

	otroMapa["juan"] = Edades{2001, 9, 31}

	fmt.Println(nuevoMapa, otroMapa)
}
