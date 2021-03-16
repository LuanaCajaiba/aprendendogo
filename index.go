package main

import "fmt"

func main() {

	var a, b, x int

	//a = 10
	//b = 9

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	x = int(a) + int(b)

	fmt.Println("x =", x)

}
