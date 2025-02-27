package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		fmt.Println(texto)
			return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando valor do parâmetro solicitado") // como usar

	fmt.Println(retorno)
}