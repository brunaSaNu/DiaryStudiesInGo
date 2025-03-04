package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int){ //não é necessário retorno pois faço a alteração diretamente na variável
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 600
	println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	println(novoNumero)
}