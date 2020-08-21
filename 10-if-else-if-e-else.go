package main

import (
	"fmt"
)

func main() {
	x := 30

	if x == 10 {
		fmt.Println("O valor de x é igual a 10.")
	} else if x == 20 {
		fmt.Println("o valor de x é igual a 20.")
	} else {
		fmt.Println("O valor de x é diferente de 10 e de 20.")
	}
}
