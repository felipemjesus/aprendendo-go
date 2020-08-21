package main

import "fmt"

func main() {
	animais := [...]string{"Cachoro", "Gato", "Vaca", "Bode", "Cavalo"}

	for i, animal := range animais {
		fmt.Printf("animais[%d]: %s\n", i, animal)

		if animal == "Vaca" {
			break
		}
	}

	contador := 0
	for contador < 10 {
		contador++

		if contador == 5 {
			continue
		}

		fmt.Printf("Contador: %d\n", contador)
	}
}
