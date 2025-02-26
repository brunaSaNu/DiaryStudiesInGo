package main

import "fmt"

const pi = 3.14159

func calcularAreaTriangulo(A float64, C float64) float64 {
	var triangulo = ((A * C) / 2)
	return triangulo
}

func calcularAreaCirculo(C float64) float64 {
	var circulo = (pi * (C * C))
	return circulo
}

func calcularAreaTrapezio(A float64, B float64, C float64) float64 {
	var trapezio = ((A + B) * C / 2)
	return trapezio
}

func calcularAreaQuadrado(B float64) float64 {
	var quadrado = B * B
	return quadrado
}

func calcularAreaRetangulo(A float64, B float64) float64 {
	var retangulo = A * B
	return retangulo
}

func main() {
	var A, B, C float64

	fmt.Scanln(&A, &B, &C)

	fmt.Printf("TRIANGULO: %.3f\n", calcularAreaTriangulo(A, C))
	fmt.Printf("CIRCULO: %.3f\n", calcularAreaCirculo(C))
	fmt.Printf("TRAPEZIO: %.3f\n", calcularAreaTrapezio(A, B, C))
	fmt.Printf("QUADRADO: %.3f\n", calcularAreaQuadrado(B))
	fmt.Printf("RETANGULO: %.3f\n", calcularAreaRetangulo(A, B))
}