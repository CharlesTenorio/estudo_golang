package main

import (
	"fmt"
)

type item struct {
	produtoID int
	nome      string
	valor     float64
	qtd       int
}

type pedido struct {
	usarioID int
	intens   []item
}

// funcao com reciver q e opedio

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.intens {
		total += item.valor * float64(item.qtd)
	}
	return total

}

func main() {

	pedido := pedido{

		usarioID: 12,
		intens: []item{
			item{89, "notbook dell", 1500.00, 9},
			item{8, "notbook IBM", 1500.00, 9},
			item{9, "notbook AA", 15000.00, 9},
		},
	}
	fmt.Println("valor total do pedio e %.2f", pedido.valorTotal())
}
