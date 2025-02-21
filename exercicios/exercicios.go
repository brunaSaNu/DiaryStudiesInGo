package main

import "fmt"

func main() {
	variavelInteira := 10

	var ponteiro * int
	ponteiro = &variavelInteira

	fmt.Println(*ponteiro)

	variavelInteira++

	*ponteiro = 30

	fmt.Println(*ponteiro)

}