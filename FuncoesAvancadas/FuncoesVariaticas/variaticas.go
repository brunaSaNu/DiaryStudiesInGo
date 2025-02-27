package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros{
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int){ //função que combina parâmetros fixos com variáticos
	for _, numero := range numeros{
		fmt.Println(texto,numero)
	}
}


func main() {
	totalDasoma := soma(1, 2, 10, 20, 40)
	fmt.Println(totalDasoma)

	escrever("Caju e Castanha", 10, 20)
}