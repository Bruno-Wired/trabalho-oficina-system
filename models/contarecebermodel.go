package models

import (
	"database/sql"
	"fmt"
	"trabalhogocopia/config"
	"trabalhogocopia/entities"
)

type ContaReceberModel struct {
}

//Adicionar Conta a receber
func (*ContaReceberModel) CreateContaReceber(conta *entities.ContasReceber) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows, err2 := db.Exec("INSERT INTO contas_receber(cliente_id, produto_id, pag_forma_id, valor, situacao) VALUES(?, ?, ?, ?, ?);", conta.Cliente, conta.Produto, conta.Forma, conta.Valor, conta.Situacao)
		id, _ := rows.LastInsertId()
		conta.Id = id
		rows2 := db.QueryRow("SELECT c.nome, p.nome, pag.nome, contas.valor, DATE_FORMAT(contas.dataUpdate, '%d-%m-%Y %H:%i:%s'), DATE_FORMAT(contas.dataUpdated, '%d-%m-%Y %H:%i:%s'), if (contas.situacao = 1, 'A receber','Recebido') FROM clientes c, produtos p, pag_formas pag, contas_receber contas WHERE contas.cliente_id = c.id AND contas.produto_id = p.id AND contas.pag_forma_id = pag.id AND contas.id = ?;", id)
		err3 := rows2.Scan(&conta.Cliente, &conta.Produto, &conta.Forma, &conta.Valor, &conta.DataAtt, &conta.DataNova, &conta.Situacao)
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

//Listar conta
func (*ContaReceberModel) FindAllContasReceber() ([]entities.ContasReceber, error) {
	db, err := config.GetDB()

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT contas.id, c.nome AS cliente, p.nome AS produto, pag.nome AS forma, contas.valor, DATE_FORMAT(contas.dataUpdate, '%d-%m-%Y %H:%i:%s') AS ultima_atualizacao, DATE_FORMAT(contas.dataUpdated, '%d-%m-%Y %H:%i:%s'), if (contas.situacao = 1, 'A receber','Recebido') AS situacao FROM clientes c, produtos p, pag_formas pag, contas_receber contas WHERE contas.cliente_id = c.id AND contas.produto_id = p.id AND contas.pag_forma_id = pag.id;")

		if err2 != nil {
			return nil, err2
		} else {
			var contas []entities.ContasReceber

			for rows.Next() {
				var conta entities.ContasReceber

				err := rows.Scan(&conta.Id, &conta.Cliente, &conta.Produto, &conta.Forma, &conta.Valor, &conta.DataAtt, &conta.DataNova, &conta.Situacao)
				contas = append(contas, conta)
				fmt.Println(err)
			}
			db.Close()
			rows.Close()
			return contas, nil
		}
	}
}

//Achar por ID
func (*ContaReceberModel) FindContaReceber(id int64) (entities.ContasReceber, error) {
	db, err := config.GetDB()

	if err != nil {
		return entities.ContasReceber{}, err

	} else {
		var conta entities.ContasReceber
		rows := db.QueryRow("SELECT contas.id, c.nome AS cliente, p.nome AS produto, pag.nome AS forma, contas.valor, DATE_FORMAT(contas.dataUpdate, '%d-%m-%Y %H:%i:%s') AS ultima_atualizacao, DATE_FORMAT(contas.dataUpdated, '%d-%m-%Y %H:%i:%s'), if (contas.situacao = 1, 'A receber','Recebido') AS situacao FROM clientes c, produtos p, pag_formas pag, contas_receber contas WHERE contas.cliente_id = c.id AND contas.produto_id = p.id AND contas.pag_forma_id = pag.id AND contas.id = ?;", id)
		db.Close()
		err2 := rows.Scan(&conta.Id, &conta.Cliente, &conta.Produto, &conta.Forma, &conta.Valor, &conta.DataAtt, &conta.DataNova, &conta.Situacao)

		if err2 == sql.ErrNoRows {
			return entities.ContasReceber{}, err2

		} else {
			return conta, nil
		}
	}
}

//Deletar estoque
func (*ContaReceberModel) DeleteContaReceber(id int64) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows, err2 := db.Exec("DELETE FROM contas_receber WHERE id = ?", id)

		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := rows.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*ContaReceberModel) PreencheContaReceber(conta *entities.ContasReceber) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows := db.QueryRow("SELECT c.id AS cliente, p.id AS produto, pag.id AS forma, contas.valor, DATE_FORMAT(contas.dataUpdate, '%d-%m-%Y %H:%i:%s') AS ultima_atualizacao, DATE_FORMAT(contas.dataUpdated, '%d-%m-%Y %H:%i:%s'), if (contas.situacao = 1, 'A receber','Recebido') AS situacao FROM clientes c, produtos p, pag_formas pag, contas_receber contas WHERE contas.cliente_id = c.id AND contas.produto_id = p.id AND contas.pag_forma_id = pag.id AND contas.id = ?;", conta.Id)
		//rows := db.QueryRow("SELECT c.nome AS cliente, p.nome AS produto, pag.nome AS forma, contas.valor, DATE_FORMAT(contas.dataUpdate, '%d-%m-%Y %H:%i:%s') AS ultima_atualizacao, DATE_FORMAT(contas.dataUpdated, '%d-%m-%Y %H:%i:%s'), if (contas.situacao = 1, 'A receber','Recebido') AS situacao FROM clientes c, produtos p, pag_formas pag, contas_receber contas WHERE contas.cliente_id = c.id AND contas.produto_id = p.id AND contas.pag_forma_id = pag.id AND contas.id = ?;", conta.Id)

		err3 := rows.Scan(&conta.Cliente, &conta.Produto, &conta.Forma, &conta.Valor, &conta.DataAtt, &conta.DataNova, &conta.Situacao)

		if err3 == sql.ErrNoRows {
			return false

		} else {
			return true

		}
	}
}

//Update conta
func (*ContaReceberModel) Update(conta *entities.ContasReceber) bool {
	db, err := config.GetDB()

	if err != nil {
		return false

	} else {
		rows2, err2 := db.Exec("UPDATE contas_receber SET cliente_id = ?, produto_id = ?, pag_forma_id = ?, valor = ?, situacao = ? WHERE id = ?;", conta.Cliente, conta.Produto, conta.Forma, conta.Valor, conta.Situacao, conta.Id)
		rows := db.QueryRow("SELECT c.nome AS cliente, p.nome AS produto, pag.nome AS forma, contas.valor, DATE_FORMAT(contas.dataUpdate, '%d-%m-%Y %H:%i:%s') AS ultima_atualizacao, DATE_FORMAT(contas.dataUpdated, '%d-%m-%Y %H:%i:%s'), if (contas.situacao = 1, 'A receber','Recebido') AS situacao FROM clientes c, produtos p, pag_formas pag, contas_receber contas WHERE contas.cliente_id = c.id AND contas.produto_id = p.id AND contas.pag_forma_id = pag.id AND contas.id = ?;", conta.Id)
		_ = rows.Scan(&conta.Cliente, &conta.Produto, &conta.Forma, &conta.Valor, &conta.DataAtt, &conta.DataNova, &conta.Situacao)
		fmt.Println(err2)
		if err2 != nil {
			return false

		} else {
			rowsAffected, _ := rows2.RowsAffected()
			return rowsAffected > 0
		}
	}
}
