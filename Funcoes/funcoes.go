package main

import "fmt"

func somar(n1 int8, n2 int8) uint8 {
	return uint8(n1) + uint8(n2)
}

func calculosMatematicos (n1, n2 int8)(int8, int8)  {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(100, 20) //uma variável para guardar a função. Funções em go são tipos, por isso posso fazer isso
	fmt.Println(soma)//posso fazer com que o parâmetro de uma função seja outra função, ou retorno

	var f = func ()  {
		fmt.Println("Função f")
	}

	var v = func (txt string)  {
		fmt.Println(txt)
	}

	var w = func (txt string) string {
		fmt.Println(txt)
		return txt
	}

	f()//confuso//estudar mais
	v("Texto da função v")
	resultado := w("Texto da função w, com retorno")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10,50)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSoma1, _ := calculosMatematicos(10,30)//o _ faz com que o programa imprima o resultado mesmo tendo 2 parâmetros de retorno
	fmt.Println(resultadoSoma1)
}