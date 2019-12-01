package main

import "fmt"

type esportivo interface {
	ligaTurbo()
}

type ferrari struct {
	modelo          string
	turboligado     bool
	velocidadeAtual int16
}

func (f *ferrari) ligaTurbo() {
	f.turboligado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligaTurbo()

	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligaTurbo()

	fmt.Println(carro1, carro2)
}
