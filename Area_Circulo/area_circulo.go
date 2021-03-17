package main

import (
	"fmt"
)

func main() {
	var n float32 = 3.14159
	var raio float32
	var area float32

	fmt.Scanln(&raio)

	area = (n * (raio * raio))

	fmt.Println(area)

}
