package models

import (
	"database/sql"
	"trabalhogocopia/config"
	"trabalhogocopia/entities"
)

type ProdutoModel struct {
}

//Listar produtos
func (*ProdutoModel) FindAll() ([]entities.Produto, error) {
	db, err := config.GetDB()

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT p.id, p.nome, c.nome AS categoria, f.nome AS fornecedor, p.descricao, p.quantidadeMin AS prateleira, p.preco, e.lote, DATE_FORMAT(e.dataVencimento, '%d-%m-%Y') FROM produtos p, fornecedores f, categorias c, estoque e WHERE p.fornecedor_id = f.id AND p.categoria_id = c.id AND p.id = e.produto_id;")
		db.Close()

		if err2 != nil {
			return nil, err2
		} else {
			var produtos []entities.Produto

			for rows.Next() {
				var produto entities.Produto

				rows.Scan(&produto.Id, &produto.Nome, &produto.Categoria, &produto.Fornecedor, &produto.Descricao, &produto.QuantidadeMin, &produto.Preco, &produto.Lote, &produto.DataVencimento)
				produtos = append(produtos, produto)
			}
			rows.Close()
			return produtos, nil
		}
	}
}

//Achar por ID
func (*ProdutoModel) Find(id int64) (entities.Produto, error) {
	db, err := config.GetDB()

	if err != nil {
		return entities.Produto{}, err

	} else {
		var produto entities.Produto
		rows := db.QueryRow("SELECT p.id, p.nome, c.nome AS categoria, f.nome AS fornecedor, p.descricao, p.quantidadeMin AS prateleira, p.preco, e.lote, DATE_FORMAT(e.dataVencimento, '%d-%m-%Y %H:%i:%s') FROM produtos p, fornecedores f, categorias c, estoque e WHERE p.fornecedor_id = f.id AND p.categoria_id = c.id AND p.id = ? AND p.id = e.produto_id ORDER BY p.nome;", id)
		db.Close()
		err2 := rows.Scan(&produto.Id, &produto.Nome, &produto.Categoria, &produto.Fornecedor, &produto.Descricao, &produto.QuantidadeMin, &produto.Preco, &produto.Lote, &produto.DataVencimento)

		if err2 == sql.ErrNoRows {
			return entities.Produto{}, err2

		} else {
			return produto, nil
		}
	}
}

//Adicionar produto
func (*ProdutoModel) Create(produto *entities.Produto) (bool, error) {
	db, err := config.GetDB()

	if err != nil {
		return false, err
	} else {

		rows, err2 := db.Exec("INSERT INTO produtos(nome, categoria_id, fornecedor_id, descricao, quantidadeMin, preco) VALUES(?, ?, ?, ?, ?, ?);", produto.Nome, produto.Categoria, produto.Fornecedor, produto.Descricao, produto.QuantidadeMin, produto.Preco)
		_ = db.QueryRow("SELECT MAX(id) FROM produtos;").Scan(&produto.Id)
		_ = db.QueryRow("SELECT nome FROM categorias WHERE id = ?;", produto.Categoria).Scan(&produto.Categoria)
		_ = db.QueryRow("SELECT nome FROM fornecedores WHERE id = ?;", produto.Fornecedor).Scan(&produto.Fornecedor)

		if err2 != nil {
			return false, err2

		} else {
			rowsAffected, _ := rows.RowsAffected()
			return rowsAffected > 0, nil
		}
	}
}

func (*ProdutoModel) PreencheUpdate(produto *entities.Produto) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows := db.QueryRow("SELECT p.nome, p.categoria_id, p.fornecedor_id, p.descricao, p.quantidademin, p.preco, e.lote FROM produtos p, estoque e WHERE p.id = e.produto_id AND p.id = ?;", produto.Id)
		err3 := rows.Scan(&produto.Nome, &produto.Categoria, &produto.Fornecedor, &produto.Descricao, &produto.QuantidadeMin, &produto.Preco, &produto.Lote)

		if err3 != nil {
			return false

		} else {
			return true
		}
	}
}

//Update produto
func (*ProdutoModel) Update(produto *entities.Produto) bool {
	db, err := config.GetDB()

	if err != nil {
		return false

	} else {
		rows2, err2 := db.Exec("UPDATE produtos SET nome = ?, categoria_id = ?, fornecedor_id = ?, descricao = ?, quantidadeMin = quantidadeMin + ?, preco = ?, id = ? WHERE id = ?;", produto.Nome, produto.Categoria, produto.Fornecedor, produto.Descricao, produto.QuantidadeMin, produto.Preco, produto.Id, produto.Id)
		_ = db.QueryRow("SELECT quantidadeMin FROM produtos WHERE id = ?;", produto.Id).Scan(&produto.QuantidadeMin)
		_ = db.QueryRow("SELECT MAX(id) FROM produtos;").Scan(&produto.Id)
		_ = db.QueryRow("SELECT nome FROM categorias WHERE id = ?;", produto.Categoria).Scan(&produto.Categoria)
		_ = db.QueryRow("SELECT nome FROM fornecedores WHERE id = ?;", produto.Fornecedor).Scan(&produto.Fornecedor)

		if err2 != nil {
			return false

		} else {
			rowsAffected, _ := rows2.RowsAffected()
			return rowsAffected > 0
		}
	}
}

//Deletar produto
func (*ProdutoModel) Delete(id int64) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows, err2 := db.Exec("DELETE FROM estoque WHERE id = ?", id)

		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := rows.RowsAffected()
			return rowsAffected > 0
		}
	}
}
