package main

import "fmt"

type profissao struct {
	profi string
}

type endereco struct {
	cep    string
	rua    string
	bairro string
	cidade string
}

type pessoa struct {
	nome  string
	idade int32
	endereco
}

func main() {
	p := pessoa{}
	p.nome = "Charles"
	p.idade = 46
	p.bairro = "centro"

	p2 := pessoa{}
	p2.nome = "Kyara"
	p2.idade = 5
	p2.bairro = "sao jose"

	fmt.Println(p.nome, p.bairro)
	fmt.Println(p2.nome, p2.bairro)
	fmt.Println(p)
	fmt.Println(p2)
}
