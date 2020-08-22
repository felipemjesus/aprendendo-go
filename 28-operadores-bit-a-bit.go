package main

import (
	"fmt"
)

func main() {
	var x int = 60
	var y int = 13

	//0 0 = 0
	//0 1 = 0
	//1 1 = 1
	fmt.Println("Resultado do operador &:", x&y)

	//0 0 = 0
	//0 1 = 1
	//1 1 = 1
	fmt.Println("Resultado do operador |:", x|y)

	//deslocar 2 bit a esquerda
	fmt.Println("Resultado do operador <<:", x<<2)

	//deslocar 2 bit a direita
	fmt.Println("Resultado do operador >>:", x>>2)
}
