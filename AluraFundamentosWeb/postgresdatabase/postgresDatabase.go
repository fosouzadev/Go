package postgresdatabase

import (
	"Go/AluraFundamentosWeb/structs"
	"database/sql"

	_ "github.com/lib/pq"
)

const connectionString string = "user=postgres dbname=aluralojadb password=123456 host=localhost sslmode=disable"

func connect() *sql.DB {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func GetProdutos() []structs.Produto {
	conn := connect()
	defer conn.Close()

	produtosResult, err := conn.Query("SELECT id, nome, descricao, preco, quantidade FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	produtos := []structs.Produto{}

	for produtosResult.Next() {
		produto := structs.Produto{}

		err = produtosResult.Scan(
			&produto.Id,
			&produto.Nome,
			&produto.Descricao,
			&produto.Preco,
			&produto.Quantidade,
		)

		if err != nil {
			panic(err.Error())
		}

		produtos = append(produtos, produto)
	}

	return produtos
}
