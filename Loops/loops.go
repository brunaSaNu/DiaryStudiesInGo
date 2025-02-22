package main

import (
	"fmt"
)

func main() {

	//i := 0

	//for i < 10 {
		//i++
		//fmt.Println("Incrementando i")
		//time.Sleep(time.Second) //estudar essa biblioteca
		
	//}

	//fmt.Println(i)

	//for j := 0; j < 10; j+=2{
		//fmt.Println("Incrementando j", j)
		//time.Sleep(time.Second)
	//}

	nomes := [3]string{"Xena", "Jenna", "Caiena"}

	for indice, nome := range nomes{
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes{ //usa-se o underline para trazer apenas o dado que eu quero
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA"{
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{//iterando em um map
		"nome": "Leonardo",
		"sobrenome": "Xico",
	}

	for chave, valor := range usuario{
		fmt.Println(chave, valor)
	}
}