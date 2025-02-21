package main

import "fmt"

func adicionarDez(adc *int){
	*adc += 10
}

func main() {
	//variavelInteira := 10

	//var ponteiro * int
	//ponteiro = &variavelInteira

	//fmt.Println(*ponteiro)

	//variavelInteira++

	//*ponteiro = 30

	//fmt.Println(*ponteiro)

	//idade := 25
	//var ponteiro * int
	//ponteiro = &idade

	//fmt.Println(*ponteiro)

	//*ponteiro = 26

	//fmt.Println(*ponteiro)

	//mensagem := "Olá, sou a Bruna, uma desenvolvedora de Go excelente! Maga do Go"
	//var ponteiro = &mensagem

	//fmt.Println(mensagem)

	//*ponteiro = "Olá, sou a Bruna, uma desenvolvedora de Go excelente! Aka bruxona do Go"

	//fmt.Println(mensagem)

	numero := 30

	fmt.Println("Valor antes:", numero)

	adicionarDez(&numero)

	fmt.Println("Valor depois: ", numero)

}