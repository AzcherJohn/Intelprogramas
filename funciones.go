package main

import "fmt"

func main() {
	simple()
	recibeValor("la suma fue: ", 5, 15)
	texto := regresaValor1()
	fmt.Println(texto)
	texto = regresaValor2()
	fmt.Println(texto)
	valores, texto := regresaValores()
	fmt.Println(valores, texto)
}

func simple() {
	fmt.Println("Está es una función simple")
}

func recibeValor(texto string, a, s int) {
	fmt.Println("Está es una función que recibe valores")
	resul := a + s

	fmt.Print(texto)
	fmt.Println(resul)
}

func regresaValor1() string {
	return "Esta función regresa un valor"
}

func regresaValor2() (texto string) {
	texto = "esta funcion regresa un valor desde una variable declarada en la funcion"
	return texto
}

func regresaValores() (suma int, texto string) {
	suma = 2
	texto = " valores es lo que retorno está función"
	return suma, texto
}
