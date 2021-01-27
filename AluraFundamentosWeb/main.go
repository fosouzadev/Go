package main

import (
	"Go/AluraFundamentosWeb/postgresdatabase"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := postgresdatabase.GetProdutos()

	// produtos := []structs.Produto{
	// 	{Nome: "Camiseta", Descricao: "Bem bonita", Preco: 29, Quantidade: 10},
	// 	{"Notebook", "Muito rápido", 1999, 2},
	// 	{"Tenis", "Comfortável", 89, 3},
	// 	{"Fone", "Muito bom", 59, 2},
	// }

	temp.ExecuteTemplate(w, "Index", produtos)
}
