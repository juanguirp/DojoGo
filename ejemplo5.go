package main

import "fmt"

type Persona struct {
	nombre   string
	estatura float64
}

func (p *Persona) correr() string {
	return "corriendo"
}
func main() {
	p := Persona{"Johan", 9.9}
	fmt.Println(p.nombre, p.correr())
}