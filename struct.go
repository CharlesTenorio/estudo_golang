package main

import (
	"fmt"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
	qtd      int32
}

func (prd produto) precoDesconto() float64 {
	return prd.preco * (1 - prd.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "faca",
		preco:    2.50,
		desconto: 0.05,
	}

	fmt.Print(produto1)
}
