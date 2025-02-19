package main

import "fmt"

type user struct {
	nome string
	altura int
	casinha adress
}

type adress struct{
	logradouro string
	numero int
}

func main() {
	fmt.Println("Arquivo structs")

	var usuario1 user
	usuario1.nome = "Caio"
	usuario1.altura = 178
	fmt.Println(usuario1)

	casaDeNego := adress{"Av. Aiai uiui", 0}
	usuario2 := user{"Caiozinho", 200, casaDeNego} //declaração de struct com inferência de tipos
	fmt.Println(usuario2)

	//formas de delarar/instanciar com inferência de tipos quando não tenho todas as informações do objeto
	usuario3 := user {altura: 200} 
	fmt.Println(usuario3)

	usuario4 := user {nome: "Caio tunado"}
	fmt.Println(usuario4)
}