package controllers

import (
	"Go/AluraFundamentosWeb/postgresdatabase"
	"Go/AluraFundamentosWeb/structs"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := postgresdatabase.ObterProdutos()

	temp.ExecuteTemplate(w, "Index", produtos)
}

func NovoProduto(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NovoProduto", nil)
}

func InserirProduto(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		produto := structs.Produto{}

		preco, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, _ := strconv.Atoi(r.FormValue("quantidade"))

		produto.Nome = r.FormValue("nome")
		produto.Descricao = r.FormValue("descricao")
		produto.Preco = preco
		produto.Quantidade = quantidade

		postgresdatabase.InserirProduto(produto)
	}

	http.Redirect(w, r, "/", 301)
}

func ExcluirProduto(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	postgresdatabase.ExcluirProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func EdicaoProduto(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := postgresdatabase.ObterProdutoPorId(idProduto)
	temp.ExecuteTemplate(w, "EditarProduto", produto)
}

func EditarProduto(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		produto := structs.Produto{}

		id, _ := strconv.Atoi(r.FormValue("id"))
		preco, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, _ := strconv.Atoi(r.FormValue("quantidade"))

		produto.Id = id
		produto.Nome = r.FormValue("nome")
		produto.Descricao = r.FormValue("descricao")
		produto.Preco = preco
		produto.Quantidade = quantidade

		postgresdatabase.EditarProduto(produto)
	}

	http.Redirect(w, r, "/", 301)
}
