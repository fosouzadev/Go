package routes

import (
	"go/AluraFundamentosWeb/controllers"
	"net/http"
)

func ConfiguraRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/novoProduto", controllers.NovoProduto)
}
