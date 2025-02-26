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

	//const n float64 = 3.14159
	//var area, raio float64

	//fmt.Scanf("%f", &raio)

	//area = n * (raio * raio)

	//fmt.Printf("A=%.4f\n", area)

	//var A, B, media float64

	//fmt.Scanf("%f\n%f", &A, &B)
	
	//media = ((A * 3.5) + (B * 7.5)) / 11.0

	//fmt.Printf("MEDIA = %.5f\n", media)

	//var A, B, C, media float64

	//fmt.Scanf("%f\n%f\n%f\n", &A, &B, &C)

	//media = (((A * 2) + (B * 3) + (C * 5)) / 10.0)

	//fmt.Printf("MEDIA = %.1f\n", media)

	//var A, B, C, D, diferenca int

	//fmt.Scan(&A, &B, &C, &D)

	//diferenca = (A * B) - (C * D)

	//fmt.Println("DIFERENCA =", diferenca)

	//var number, hours int
	//var rececipe float64
	//var salary float64

	//fmt.Scan(&number, &hours)
	//fmt.Scan(&rececipe)

	//salary = float64(hours) * rececipe

	//fmt.Println("NUMBER =", number)
	//fmt.Printf("SALARY = U$ %.2f\n", salary)

	//var name string
	//var salary, sales, total float64
	//const percentage = 15.00

	//fmt.Scan(&name)
	//fmt.Scan(&salary, &sales)

	//total = ((percentage * sales) / 100)
	//total = total + salary

	//fmt.Printf("TOTAL = R$ %.2f\n", total)

	var codigoPeca, numeroPeca int
	var codigoPeca1, numeroPeca1 int
	var valorPeca float64
	var valorPeca1 float64
	var valorAPagar float64
	var valorAPagar1 float64
	var valorTotal float64

	fmt.Scanln(&codigoPeca, &numeroPeca, &valorPeca)
	fmt.Scanln(&codigoPeca1, &numeroPeca1, &valorPeca1)

	valorAPagar = float64(numeroPeca) * valorPeca
	valorAPagar1 = float64(numeroPeca1) * valorPeca1
	valorTotal = valorAPagar + valorAPagar1

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", valorTotal)

}
