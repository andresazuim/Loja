package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Vermelha", Preco: 50, Quantidade: 5},
		{"Tenis", "confortavel", 90, 4},
		{"fone", "muito bom", 59, 2},
	}
	temp.ExecuteTemplate(w, "index", produtos)
}
