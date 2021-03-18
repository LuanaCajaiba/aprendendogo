package main

import (
	"fmt"
)

type ninja int

var x ninja

func main() {
	//demonstra o valor da variável x
	fmt.Println(x)
	//demonstra o tipo da variável x
	fmt.Printf("%T\n", x)
	//atribui 42 a variável x utilizando o operador de atribuição
	x = 42
	fmt.Println(x)

}
