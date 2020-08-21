package main

import "fmt"

func main() {
	materiais := [...]string{"Mochila", "Estojo", "Lápis", "Borracha", "Fita"}

	for i := 0; i < len(materiais); i++ {
		fmt.Printf("materiais[%d]: %s\n", i, materiais[i])
	}

	for i, material := range materiais {
		fmt.Printf("materiais[%d]: %s\n", i, material)
	}

	contador := 1
	for contador <= 10 {
		fmt.Printf("Contador: %d\n", contador)
		contador++
	}
}
