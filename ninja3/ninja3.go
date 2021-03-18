package main

import (
	"fmt"
)

type ninja int

var x ninja

//utilizando a palavra chave var, criando uma variável com o identificador y
var y int

//package-leve scope
func main() {
	//demonstra o valor da variável x
	fmt.Println(x)
	//demonstra o tipo da variável x
	fmt.Printf("%T\n", x)
	//atribui 42 a variável x utilizando o operador de atribuição
	x = 42
	//demonstra o valor de x
	fmt.Println(x)
	//utilização de conversão para transformar o tipo do valor da variável "x" em seu subjacente e, utilizando o operador "="
	//atribui o valor y ao x
	y = int(x)
	//valor de x
	fmt.Println(x)
	//tipo de x
	fmt.Printf("%T\n", x)

}
