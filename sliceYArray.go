package main

import "fmt"

func main() {
	//Array
	fruta := [5]string{"manzana", "pera", "guayaba", "sandia", "melon"}
	var precio [5]int
	precio = [5]int{25, 31, 45, 12, 16}

	fmt.Println(fruta, precio)

	precio[0] = 80

	fmt.Println(fruta, precio)

	//Slice
	verdura := []string{"zanahoria", "jicama"}

	fmt.Println(verdura)

	verdura = append(verdura, "rabano", "lechuga")

	fmt.Println(verdura)

	var costo []int

	costo = append(costo, precio[1:4]...)

	fmt.Println(verdura, costo)

	costo = append(costo, 21)

	fmt.Println(verdura, costo)
}
