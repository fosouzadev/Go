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

func ObterProdutoPorId(id string) structs.Produto {
	conn := connect()
	defer conn.Close()

	produtoResult, err := conn.Query(
		"SELECT id, nome, descricao, preco, quantidade FROM produtos "+
			"WHERE id = $1", id,
	)

	if err != nil {
		panic(err.Error())
	}

	produto := structs.Produto{}

	if produtoResult.Next() {
		err = produtoResult.Scan(
			&produto.Id,
			&produto.Nome,
			&produto.Descricao,
			&produto.Preco,
			&produto.Quantidade,
		)

		if err != nil {
			panic(err.Error())
		}
	}

	return produto
}

func ObterProdutos() []structs.Produto {
	conn := connect()
	defer conn.Close()

	produtosResult, err := conn.Query(
		"SELECT id, nome, descricao, preco, quantidade FROM produtos " +
			"ORDER BY id",
	)

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

func InserirProduto(produto structs.Produto) {
	conn := connect()
	defer conn.Close()

	insertResult, err := conn.Prepare(
		"INSERT INTO produtos (nome, descricao, preco, quantidade) " +
			"VALUES ($1, $2, $3, $4)",
	)

	if err != nil {
		panic(err.Error())
	}

	insertResult.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)
}

func ExcluirProduto(id string) {
	conn := connect()
	defer conn.Close()

	deleteResult, err := conn.Prepare("DELETE FROM produtos WHERE id = $1")

	if err != nil {
		panic(err.Error())
	}

	deleteResult.Exec(id)
}

func EditarProduto(produto structs.Produto) {
	conn := connect()
	defer conn.Close()

	insertResult, err := conn.Prepare(
		"UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 " +
			"WHERE id = $5",
	)

	if err != nil {
		panic(err.Error())
	}

	insertResult.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.Id)
}
