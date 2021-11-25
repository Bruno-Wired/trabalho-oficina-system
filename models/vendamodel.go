package models

import (
	"database/sql"
	"fmt"
	"trabalhogocopia/config"
	"trabalhogocopia/entities"
)

type VendasModel struct {
}

//Adicionar venda
func (*VendasModel) CreateVenda(venda *entities.Venda) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows, err2 := db.Exec("INSERT INTO vendas(cliente_id, produto_id, quantidade) VALUES(?, ?, ?);", venda.Cliente, venda.Produto, venda.Quantidade)
		venda.Id, _ = rows.LastInsertId()
		rows2 := db.QueryRow("SELECT c.nome, DATE_FORMAT(v.dataCompra, '%d-%m-%Y %H:%i:%s'), DATE_FORMAT(e.dataVencimento, '%d-%m-%Y') AS Vencimento, p.nome, v.quantidade, v.precoUnit, v.total FROM clientes c, vendas v, produtos p, estoque e WHERE v.cliente_id = c.id AND v.produto_id = p.id AND v.produto_id = e.produto_id AND v.id = ?;", venda.Id)
		err3 := rows2.Scan(&venda.Cliente, &venda.DataVenda, &venda.Vencimento, &venda.Produto, &venda.Quantidade, &venda.PrecoUnit, &venda.Total)
		db.Close()
		fmt.Println(err3)
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

//Listar vendas
func (*VendasModel) FindAllVendas() ([]entities.Venda, error) {
	db, err := config.GetDB()

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT v.id, CONCAT(c.nome,' ', c.nomemeio,' ', c.sobrenome) AS cliente, DATE_FORMAT(v.dataCompra, '%d-%m-%Y %H:%i:%s') AS data_Venda, DATE_FORMAT(e.dataVencimento, '%d-%m-%Y') AS Vencimento, p.nome AS produto, v.quantidade, v.precoUnit, v.total FROM clientes c, vendas v, produtos p, estoque e WHERE c.id = v.cliente_id AND p.id = v.produto_id AND p.id = e.produto_id;")
		db.Close()

		if err2 != nil {
			return nil, err2
		} else {
			var vendas []entities.Venda

			for rows.Next() {
				var venda entities.Venda

				rows.Scan(&venda.Id, &venda.Cliente, &venda.DataVenda, &venda.Vencimento, &venda.Produto, &venda.Quantidade, &venda.PrecoUnit, &venda.Total)
				vendas = append(vendas, venda)
			}
			rows.Close()
			return vendas, nil
		}
	}
}

//buscar venda
func (*VendasModel) VendaFind(id int64) (entities.Venda, error) {
	db, err := config.GetDB()

	if err != nil {
		return entities.Venda{}, err

	} else {
		var venda entities.Venda
		rows := db.QueryRow("SELECT v.id, CONCAT(c.nome,' ', c.nomemeio,' ', c.sobrenome) AS cliente, DATE_FORMAT(v.dataCompra, '%d-%m-%Y %H:%i:%s') AS data_Venda, DATE_FORMAT(e.dataVencimento, '%d-%m-%Y') AS Vencimento, p.nome AS produto, v.quantidade, v.precoUnit, v.total FROM clientes c, vendas v, produtos p, estoque e WHERE c.id = v.cliente_id AND p.id = v.produto_id AND p.id = e.produto_id AND v.id = ?;", id)
		db.Close()
		err2 := rows.Scan(&venda.Id, &venda.Cliente, &venda.DataVenda, &venda.Vencimento, &venda.Produto, &venda.Quantidade, &venda.PrecoUnit, &venda.Total)

		if err2 == sql.ErrNoRows {
			return entities.Venda{}, err2

		} else {
			return venda, nil
		}
	}
}
