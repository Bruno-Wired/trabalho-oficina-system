package models

import (
	"database/sql"
	"fmt"
	"trabalhogocopia/config"
	"trabalhogocopia/entities"
)

type EstoqueModel struct {
}

//Listar estoque
func (*EstoqueModel) FindAllEstoque() ([]entities.Estoque, error) {
	db, err := config.GetDB()

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT e.id, e.lote, DATE_FORMAT(e.dataEntrada, '%d-%m-%Y %H:%i:%s') AS data_hora_entrada, DATE_FORMAT(e.dataVencimento, '%d-%m-%Y') AS data_vencimento, e.quantidade, e.precoUnit AS preco_unitario, e.total, p.nome AS produto, f.nome AS fornecedor FROM estoque e, produtos p, fornecedores f WHERE e.produto_id = p.id AND p.fornecedor_id = f.id;")
		db.Close()

		if err2 != nil {
			return nil, err2
		} else {
			var produtos []entities.Estoque

			for rows.Next() {
				var estoque entities.Estoque

				rows.Scan(&estoque.Id, &estoque.Lote, &estoque.DataEntrada, &estoque.DataVencimento, &estoque.Quantidade, &estoque.PrecoUnit, &estoque.Total, &estoque.Produto, &estoque.Fornecedor)
				produtos = append(produtos, estoque)
			}
			rows.Close()
			return produtos, nil
		}
	}
}

//buscar estoque
func (*EstoqueModel) EstoqueFind(id int64) (entities.Estoque, error) {
	db, err := config.GetDB()

	if err != nil {
		return entities.Estoque{}, err

	} else {
		var estoque entities.Estoque
		rows := db.QueryRow("SELECT e.id, e.lote, DATE_FORMAT(e.dataEntrada, '%d-%m-%Y %H:%i:%s') AS data_hora_entrada, DATE_FORMAT(e.dataVencimento, '%d-%m-%Y') AS data_vencimento, e.quantidade, e.precoUnit AS preco_unitario, e.total, p.nome AS produto, f.nome AS fornecedor FROM estoque e, produtos p, fornecedores f WHERE e.produto_id = p.id AND p.fornecedor_id = f.id AND p.id = ? ORDER BY p.nome;", id)
		db.Close()
		err2 := rows.Scan(&estoque.Id, &estoque.Lote, &estoque.DataEntrada, &estoque.DataVencimento, &estoque.Quantidade, &estoque.PrecoUnit, &estoque.Total, &estoque.Produto, &estoque.Fornecedor)

		if err2 == sql.ErrNoRows {
			return entities.Estoque{}, err2

		} else {
			return estoque, nil
		}
	}
}

//Adicionar estoque
func (*EstoqueModel) CreateEstoque(estoque *entities.Estoque) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows, err2 := db.Exec("INSERT INTO estoque(lote, dataVencimento, quantidade) VALUES(?, ?, ?);", estoque.Lote, estoque.DataVencimento, estoque.Quantidade)
		id := db.QueryRow("SELECT MAX(id) FROM estoque;")
		_ = id.Scan(&estoque.Id)
		rows2 := db.QueryRow("SELECT e.lote, DATE_FORMAT(e.dataEntrada, '%d-%m-%Y %H:%i:%s') AS data_hora_entrada, DATE_FORMAT(e.dataVencimento, '%d-%m-%Y') AS data_vencimento, e.quantidade, e.precoUnit AS preco_unitario, e.total, p.nome AS produto, f.nome AS fornecedor FROM estoque e, produtos p, fornecedores f WHERE e.produto_id = p.id AND p.fornecedor_id = f.id  AND e.id = ?;", estoque.Id)
		err3 := rows2.Scan(&estoque.Lote, &estoque.DataEntrada, &estoque.DataVencimento, &estoque.Quantidade, &estoque.PrecoUnit, &estoque.Total, &estoque.Produto, &estoque.Fornecedor)
		db.Close()

		if err2 != nil {
			return false

		} else if err3 == sql.ErrNoRows {
			return false

		} else {
			rowsAffected, _ := rows.RowsAffected()
			return rowsAffected > 0
		}
	}
}

//Deletar estoque
func (*EstoqueModel) EstoqueDelete(id int64) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows, err2 := db.Exec("DELETE FROM estoque WHERE produto_id = ?", id)

		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := rows.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*EstoqueModel) PreencheEstoque(estoque *entities.Estoque) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows := db.QueryRow("SELECT p.nome, e.lote, e.DataEntrada, e.dataVencimento, e.quantidade, e.precoUnit, e.total, f.nome FROM produtos p, estoque e, fornecedores f WHERE p.id = e.produto_id AND e.produto_id AND p.fornecedor_id = f.id AND e.id = ?;", estoque.Id)
		err3 := rows.Scan(&estoque.Produto, &estoque.Lote, &estoque.DataEntrada, &estoque.DataVencimento, &estoque.Quantidade, &estoque.PrecoUnit, &estoque.Total, &estoque.Fornecedor)

		if err3 != nil {
			return false

		} else {

			return true
		}
	}
}

//Update estoque
func (*EstoqueModel) EstoqueUpdate(estoque *entities.Estoque) bool {
	db, err := config.GetDB()

	if err != nil {
		return false

	} else {
		rows, err2 := db.Exec("UPDATE estoque SET lote = ?, dataVencimento = ?, quantidade = ? WHERE id = ?;", estoque.Lote, estoque.DataVencimento, estoque.Quantidade, estoque.Id)
		_ = db.QueryRow("SELECT DATE_FORMAT(dataEntrada, '%d-%m-%Y %H:%i:%s'), DATE_FORMAT(dataVencimento, '%d-%m-%Y') FROM estoque WHERE id = ?", estoque.Id).Scan(&estoque.DataEntrada, &estoque.DataVencimento)

		if err2 != nil {
			fmt.Println(err2)
			return false

		} else {
			rowsAffected, _ := rows.RowsAffected()
			return rowsAffected > 0
		}
	}
}
