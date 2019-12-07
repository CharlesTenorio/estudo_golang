package main

import (
	"fmt"
	"time"
)

func fala(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Println("%s: %s(iterecao %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	go fala("Lidiane gostosa", "quero ir ao cinema", 10)
	fala("Charles", "oba", 5)

}
