package main

import (
	"os"
)

var caminho = "arquivo.txt"
var paises = [...]string{"Brasil", "Argentina", "Chile", "Paraguai", "Uruguai"}

func main() {
	var _, arquivo = os.Stat(caminho)

	if !os.IsNotExist(arquivo) {
		escreverNoArquivo()
	}
}

func escreverNoArquivo() {
	var arquivo, _ = os.OpenFile(caminho, os.O_WRONLY, 0644)
	defer arquivo.Close()

	for _, pais := range paises {
		arquivo.WriteString(pais + "\n")
	}

	arquivo.Sync()
}
