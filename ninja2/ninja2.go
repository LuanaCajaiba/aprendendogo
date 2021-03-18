package main

import "fmt"

var x int = 42 //aqui foram atribuidos os valores as variáveis
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("%v\n%v\n%v\n", x, y, z) //aqui foi atribuído utilizando operador curto de atribuição
	fmt.Println(s)

}
