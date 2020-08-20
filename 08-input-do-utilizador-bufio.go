package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)

	fmt.Print("Informe seu nome: ")
	nome, _ := input.ReadString('\n')

	fmt.Print("Informe seu sobrenome: ")
	sobrenome, _ := input.ReadString('\n')

	fmt.Print("Informe sua idade: ")
	idade, _ := input.ReadString('\n')

	fmt.Println("Nome:", nome, sobrenome, "Idade:", idade)
}