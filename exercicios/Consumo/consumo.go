package main

import "fmt"

func consumoMedio(X int, Y float64) float64 {
	var consumo = float64(X) / Y
	return consumo
}

func main() {
	var X int
	var Y float64

	fmt.Scanf("%d \n %f", &X, &Y)

	resultado := (consumoMedio(X, Y))

	fmt.Printf("%.3f km/l \n", resultado)
}