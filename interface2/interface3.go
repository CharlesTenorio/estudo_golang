package main

import "fmt"

type esportivo interface {
	ligaTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	ligarGps()
}

type bmw7 struct{}

func (b bmw7) ligaTurbo() {
	fmt.Println("Turbo...")

}

func (b bmw7) fazerBaliza() {
	fmt.Println("FAZER baliza...")
}

func (b bmw7) ligarGps() {
	fmt.Println("ligar gps....")
}

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.ligarGps()
	b.ligaTurbo()
	b.fazerBaliza()
}
