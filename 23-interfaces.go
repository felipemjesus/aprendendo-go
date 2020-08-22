package main

import (
	"fmt"
	"math"
)

type FiguraGeometrica interface {
	area() float64
	perimetro() float64
}

type Retangulo struct {
	largura, altura float64
}

type Circulo struct {
	raio float64
}

func (retangulo Retangulo) area() float64 {
	return retangulo.largura * retangulo.altura
}

func (retangulo Retangulo) perimetro() float64 {
	return (2 * retangulo.largura) + (2 * retangulo.altura)
}

func (circulo Circulo) area() float64 {
	return math.Pi * circulo.raio * circulo.raio
}

func (circulo Circulo) perimetro() float64 {
	return 2 * math.Pi * circulo.raio
}

func calcularAreaEPerimetro(figuraGeometrica FiguraGeometrica) {
	fmt.Println("Área:", figuraGeometrica.area())
	fmt.Println("Perímetro", figuraGeometrica.perimetro())
}

func main() {
	retangulo := Retangulo{3, 4}
	circulo := Circulo{5}

	fmt.Println("Retângulo")
	calcularAreaEPerimetro(retangulo)

	fmt.Println("Círculo")
	calcularAreaEPerimetro(circulo)
}
