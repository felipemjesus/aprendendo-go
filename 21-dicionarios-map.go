package main

import (
	"fmt"
)

var pessoas = map[string]int{
	"Felipe": 29,
	"João":   35,
}

func main() {
	pessoas["Pedro"] = 50
	pessoas["Maria"] = 40

	delete(pessoas, "Maria") // remover do array

	pessoas = make(map[string]int) // limpar array

	fmt.Println("Número de pessoas:", len(pessoas))

	for chave, valor := range pessoas {
		fmt.Println("Nome:", chave)
		fmt.Println("Idade:", valor)
	}
}
