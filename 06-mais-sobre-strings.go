package main

import (
	"fmt"
	"strings"
)

func main() {
	nome, sobrenome := "Felipe", "Martins"
	idade := 29

	fmt.Println("Nome:", nome, sobrenome, "\n\tIdade:", idade)

	fmt.Println("Nome:", len(nome), len(sobrenome), "\n\tIdade:", idade)

	fmt.Printf("Nome: %s %s\nIdade: %d\n", strings.ToUpper(nome), strings.ToLower(sobrenome), idade)
}