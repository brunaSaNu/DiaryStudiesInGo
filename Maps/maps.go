package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Caio",
		"sobrenome": "zinho",
	}

	fmt.Println(usuario["sobrenome"])//acesso a campos

	usuario2 := map[string]map[string]string{//aninhamento de maps
		"nome": {
			"primeiro": "Jo",
			"ultimo": "Suado",
		},
		"curso": {
			"nome": "Ciência da computação",
			"campus": "campus cus cus",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{//adiciona campo
		"nome": "Escorpicão",
	}

	fmt.Println(usuario2)
}