package main

import "fmt"

//func adicionarDez(adc *int){
	//*adc += 10
//}

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

	//numero := 30

	//fmt.Println("Valor antes:", numero)

	//adicionarDez(&numero)

	//fmt.Println("Valor depois: ", numero)

	//fmt.Println("Olá, mundo!")

	//var a, b int
	//fmt.Println("Digite dois números: ")
	//fmt.Scan(&a, &b)
	//fmt.Println("Você digitou:", a, "e", b)

	//var nome string
	//var idade int
	
	//fmt.Println("Digite seu nome e sua idade: ")
	//fmt.Scanln(&nome, &idade)
	//fmt.Println("Nome: ", nome, "- idade:", idade)

	//fmt.Println("Digite seu nome e sua idade: ")
	//fmt.Scanf("%s %d", &nome, &idade)
	//fmt.Printf("Nome: %s - idade: %d", nome, idade)

	var inteiro1, inteiro2, PROD int
	fmt.Scan(&inteiro1, &inteiro2)
	PROD = inteiro1 * inteiro2

	fmt.Println("PROD =", PROD)
}