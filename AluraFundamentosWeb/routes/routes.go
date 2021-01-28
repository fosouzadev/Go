package routes

import (
	"Go/AluraFundamentosWeb/controllers"
	"net/http"
)

func ConfiguraRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/novoProduto", controllers.NovoProduto)
	http.HandleFunc("/inserirProduto", controllers.InserirProduto)
	http.HandleFunc("/excluirProduto", controllers.ExcluirProduto)
	http.HandleFunc("/edicaoProduto", controllers.EdicaoProduto)
	http.HandleFunc("/editarProduto", controllers.EditarProduto)
}
