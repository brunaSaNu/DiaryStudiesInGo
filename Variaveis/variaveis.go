package main

import "fmt"

func main() {
	var variavel1 string = "Eu sou a variável 1 e estou explícita"
	variavel2 := "Eu sou a variável 2 e estou implícita"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var(
		variavel3 string = "Eu sou a variável 3 e fui declarada em conjunto"
		variavel4 string = "eu sou a variável 4 e também fui declarada em conjunto"
	)

	fmt.Println(variavel3,variavel4)

	variavel5, variavel6 := "Sou a variável 5 e fui declarada com inferência de tipos e em conjunto","Sou a variável 6 e também fui declarada com inferência de tipos e em conjunto"

	fmt.Println(variavel5,variavel6)

	const constante1 string = "sou a constante 1, toda forma de me declarar é parecido com uma variável"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5,variavel6)
}
