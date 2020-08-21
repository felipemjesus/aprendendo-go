package main

import "fmt"

func main() {
	// var cores [5]string

	// cores[0] = "Azul"
	// cores[1] = "Verde"
	// cores[2] = "Amarelo"
	// cores[3] = "Vermelho"
	// cores[4] = "Laranja"

	cores := [...]string{"Azul", "Verde", "Amarelo", "Vermelho", "Laranja", "Preto"}

	fmt.Println("Número de cores:", len(cores))
	fmt.Println("Primeira cor:", cores[0])
	fmt.Println("Última cor:", cores[len(cores)-1])
}
