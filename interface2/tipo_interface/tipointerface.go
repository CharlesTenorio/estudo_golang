package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)
	coisa = 3
	fmt.Println(coisa)
	type dinamico interface{}
	var coisa2 dinamico = "ola povo"
	fmt.Println(coisa2)
	coisa2 = curso{"go lang explorando a linguagem"}
	fmt.Println(coisa2)
}
