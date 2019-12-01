package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json: "id"`
	Nome  string   `json: "nome`
	Preco float64  `json:preco`
	Tags  []string `json:tags`
}

func main() {
	//struc to json
	p1 := produto{10, "notebook", 23.78, []string{"Promocao", "Eletronico"}}

	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 produto
	jsonString := `{"id":19,"nome":"MAC PRO","preco":123.78:,"tags":["Livro","pc"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])

}
