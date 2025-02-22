package main

import "fmt"

func main() {

	numero := -10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if otherNumber := numero; otherNumber > 0{ //variável que se inicia dentro do if, if init(limitada ao escopo do if)
		fmt.Println("É maior que zero")
	} else if numero < -10 {
		fmt.Println("É menor que -10")
	} else {
		fmt.Println("Está entre 0 e -10")
	}

}