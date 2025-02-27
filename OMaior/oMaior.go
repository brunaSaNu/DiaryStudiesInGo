package main

import (
	"fmt"
	"math"
)

func maiorNumero(A, B int) int {
	var maior = (A + B + int(math.Abs(float64(A - B)))) /2
	return maior
}

func maiorNumeroDeTres(A,B,C int) int{
	return maiorNumero(maiorNumero(A, B), C)
}

func main() {
	var A, B, C int
	fmt.Scanln(&A, &B, &C)

	fmt.Println(maiorNumeroDeTres(A, B, C), "eh o maior")
}