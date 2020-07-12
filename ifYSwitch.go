package main

import "fmt"

func main() {
	promedio := 12.3
	excento := 8.50

	var texto string

	if promedio > 10.0 {
		texto = "Eso es imposible, reprobado por tramposo"
	} else if promedio >= excento {
		texto = "felicidades has aprobado"
	} else {
		texto = "no es que seas burro, pero esfuerzate más"
	}

	fmt.Println(texto)

	switch texto {
	case "Eso es imposible, reprobado por tramposo":
		fmt.Println("Deshonor a tu vaca")
	default:
		fmt.Println("Y lavate las manos :)")
	}
}
