package main

import (
	"fmt"
	"os"
)

var caminho = "arquivo.txt"

func main() {
	var _, arquivo = os.Stat(caminho)

	if !os.IsNotExist(arquivo) {
		lerArquivo()
	}
}

func lerArquivo() {
	var arquivo, _ = os.OpenFile(caminho, os.O_RDONLY, 0644)
	defer arquivo.Close()

	var conteudo = make([]byte, 1024)

	arquivo.Read(conteudo)
	fmt.Println(string(conteudo))
}
