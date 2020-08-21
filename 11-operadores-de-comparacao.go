package main

import "fmt"

func main() {
	x := 10
	y := 20

	if x == y {
		fmt.Println("Está condição é verdadeira")
	} else {
		fmt.Println("Está condição é falsa")
	}

	if x != y {
		fmt.Println("Está condição é verdadeira")
	} else {
		fmt.Println("Está condição é falsa")
	}

	if x < y {
		fmt.Println("Está condição é verdadeira")
	} else {
		fmt.Println("Está condição é falsa")
	}

	if x > y {
		fmt.Println("Está condição é verdadeira")
	} else {
		fmt.Println("Está condição é falsa")
	}

	if x <= y {
		fmt.Println("Está condição é verdadeira")
	} else {
		fmt.Println("Está condição é falsa")
	}

	if x >= y {
		fmt.Println("Está condição é verdadeira")
	} else {
		fmt.Println("Está condição é falsa")
	}

	if x <= y && x == y {
		fmt.Println("Está condição é verdadeira")
	} else {
		fmt.Println("Está condição é falsa")
	}

	if x <= y || x == y {
		fmt.Println("Está condição é verdadeira")
	} else {
		fmt.Println("Está condição é falsa")
	}
}
