package main

import "fmt"

func main() {
	soma := 1 + 2
	sub := 1 - 9
	div := 16 / 3
	multipliq := 8 * 2
	resto := 4 % 2

	fmt.Println(soma, sub, div, multipliq, resto)

    var n1 int16 = 10
    var n2 int16 = 30 // não é possível a soma caso seja um int de bits diferentes
    soma2 := n1 + n2
    fmt.Println(soma2)

    //fim dos aritméticos

    //Atribuição de variáveis
    var variavel string = "Sou uma sting com a atribuição string"
    variavel2 := "Sou uma varável com inferência de tipos"
    fmt.Println(variavel, variavel2)
    //fim dos operadores de atribuição

    //operadores relacionais

    fmt.Println(1 > 2)
    fmt.Println(1 >= 2)
    fmt.Println(1 == 2)
    fmt.Println(1 <= 2)
    fmt.Println(1 < 2)
    fmt.Println(1 != 2)

    // fim dos relacionais

    //operadores lógicos
    fmt.Println("_____________")
    verdadeiro, falso := true, false
    fmt.Println(verdadeiro && falso)
    fmt.Println(verdadeiro || falso)
    fmt.Println(!verdadeiro)
    fmt.Println(!falso)

    //fim dos operadores lógicos

    //Operadores unários
    numero := 80
    numero++;
    numero += 50 // numero = numero + 50
    fmt.Println(numero)

    //decremento

    numero--
    numero -= 10 //numero = numero - 10
    numero *= 2 // numero = numero * 2
    numero /= 5 //numero = numero / 5
    numero %= 2
    fmt.Println(numero)

    //fim dos operadores unários

    // não tem ternário em GO(estou genuinamente feliz)
}