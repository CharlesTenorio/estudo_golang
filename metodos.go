package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome       string
	sobre_nome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobre_nome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobre_nome = partes[1]

}

func main() {
	p1 := pessoa{"Lidane", "Tenorio"}
	fmt.Println(p1.getNomeCompleto())
}
