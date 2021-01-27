package controllers

import (
	"Go/AluraFundamentosWeb/postgresdatabase"
	"Go/AluraFundamentosWeb/structs"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := postgresdatabase.GetProdutos()

	temp.ExecuteTemplate(w, "Index", produtos)
}

func NovoProduto(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NovoProduto", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		produto := structs.Produto{}

		produto.Nome = r.FormValue("nome")
		produto.Descricao = r.FormValue("descricao")
		produto.Preco = float64(r.FormValue("preco"))
		produto.Quantidade = int(r.FormValue("quantidade"))
	} 
}    
