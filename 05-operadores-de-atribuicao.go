package main

import "fmt"

func main() {
	x := 2

	x += 3
	fmt.Println("x = x + 3:", x)

	x -= 2
	fmt.Println("x = x + 2:", x)

	x *= 2
	fmt.Println("x = x * 2:", x)

	x /= 2
	fmt.Println("x = x / 2:", x)

	x %= 3
	fmt.Println("x = x % 3:", x)
}