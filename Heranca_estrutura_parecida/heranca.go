// estrutura em go que mais chega próximo a herença, de uma linguagem com orientação a objetos
package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    int
}

type estudante struct {
	pessoa    //unico caso que não preciso especificar o tipo
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	people1 := pessoa{"João", "Pedro", 20, 189}
	fmt.Println(people1)

	estudante1 := estudante{people1, "ADS", "USP"}
	fmt.Println(estudante1)
	fmt.Println(estudante1.nome)//acesso mais rápido, invés de people1.people.nome
	fmt.Println(estudante1.altura)
}
