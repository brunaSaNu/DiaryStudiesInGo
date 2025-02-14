package main

import (
	"fmt"
	"modulo/packages/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo da main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("castanha@gmail.com")
	fmt.Println(erro)
}
