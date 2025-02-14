package main

import (
	"errors"
	"fmt"
)

//tipos básicos de dados no go
func main() {
	var numero int = 100
	numeroComInferencia := 12544
	fmt.Println(numero)
	fmt.Println(numeroComInferencia)

	var numero2 uint = 1000
	fmt.Println(numero2)

	//alias = "apelido"
	//rune == int32
	var numero3 rune = 1234
	fmt.Println(numero3)

	//byte == uint8
	var numero4 byte = 44
	fmt.Println(numero4)

	//numeros reais

	var numeroReal float32 = 1235.65
	fmt.Println(numeroReal)

	var numeroReal2 float32 = 14544445235.65
	fmt.Println(numeroReal2)

	numeroReal3 := 4645454.645
	fmt.Println(numeroReal3)

	//Fim números reais

	//Stings

	var str string = "Text"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	//Fim strings

	//Valor zero

	var texto int
	fmt.Println(texto)

	//fim valor zero

	//boolean

	var boolean1 bool = true
	fmt.Println(boolean1)

	//fim do boolean

	//tipo erro

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
