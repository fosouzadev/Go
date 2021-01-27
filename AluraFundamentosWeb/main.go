package main

import (
	"Go/AluraFundamentosWeb/routes"
	"net/http"
)

func main() {
	routes.ConfiguraRotas()
	http.ListenAndServe(":8000", nil)
}
