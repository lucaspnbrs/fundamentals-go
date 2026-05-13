package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea (f forma)  {
	fmt.Println("The area of form is", f.area())
}

type retangulo struct {
	largura float64
	altura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2) //Function receive to number to elevate in potence and the number of potences
}

func main() {

	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)

}