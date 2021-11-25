package models

import (
	"database/sql"
	"trabalhogocopia/config"
	"trabalhogocopia/entities"
)

type ClienteModel struct {
}

//Listar clientes
func (*ClienteModel) FindAllCliente() ([]entities.Cliente, error) {
	db, err := config.GetDB()

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, nome, nomeMeio, sobreNome, RG FROM clientes ORDER BY nome;")
		db.Close()

		if err2 != nil {
			return nil, err2
		} else {
			var clientes []entities.Cliente

			for rows.Next() {
				var cliente entities.Cliente

				rows.Scan(&cliente.Id, &cliente.Nome, &cliente.NomeMeio, &cliente.SobreNome, &cliente.Rg)
				clientes = append(clientes, cliente)
			}
			rows.Close()
			return clientes, nil
		}
	}
}

//Achar por ID
func (*ClienteModel) FindCliente(id int64) (entities.Cliente, error) {
	db, err := config.GetDB()

	if err != nil {
		return entities.Cliente{}, err

	} else {
		var cliente entities.Cliente
		rows := db.QueryRow("SELECT id, nome, nomeMeio, sobreNome, RG FROM clientes WHERE id = ?;", id)
		db.Close()
		err2 := rows.Scan(&cliente.Id, &cliente.Nome, &cliente.NomeMeio, &cliente.SobreNome, &cliente.Rg)

		if err2 == sql.ErrNoRows {
			return entities.Cliente{}, err2

		} else {
			return cliente, nil
		}
	}
}

//Adicionar cliente
func (*ClienteModel) CreateCliente(cliente *entities.Cliente) (bool, error) {
	db, err := config.GetDB()

	if err != nil {
		return false, err
	} else {

		rows, err2 := db.Exec("INSERT INTO clientes(nome, nomeMeio, sobreNome, RG) VALUES(?, ?, ?, ?);", cliente.Nome, cliente.NomeMeio, cliente.SobreNome, cliente.Rg)

		id := db.QueryRow("SELECT MAX(id) FROM clientes;")
		_ = id.Scan(&cliente.Id)

		db.Close()

		if err2 != nil {
			return false, err2

		} else {
			rowsAffected, _ := rows.RowsAffected()
			return rowsAffected > 0, nil
		}
	}
}

func (*ClienteModel) PreencheCliente(cliente *entities.Cliente) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows := db.QueryRow("SELECT nome, nomeMeio, sobreNome, RG FROM clientes WHERE id = ?;", cliente.Id)
		err3 := rows.Scan(&cliente.Nome, &cliente.NomeMeio, &cliente.SobreNome, &cliente.Rg)

		if err3 != nil {
			return false

		} else {

			return true
		}
	}
}

//Update cliente
func (*ClienteModel) UpdateCliente(cliente *entities.Cliente) bool {
	db, err := config.GetDB()

	if err != nil {
		return false

	} else {
		rows2, err2 := db.Exec("UPDATE clientes SET nome = ?, nomeMeio = ?, sobreNome = ?, RG = ? WHERE id = ?;", cliente.Nome, cliente.NomeMeio, cliente.SobreNome, cliente.Rg, cliente.Id)

		if err2 != nil {
			return false

		} else {
			rowsAffected, _ := rows2.RowsAffected()
			return rowsAffected > 0
		}
	}
}

//Deletar cliente
func (*ClienteModel) DeleteCliente(id int64) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows, err2 := db.Exec("DELETE FROM clientes WHERE id = ?", id)

		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := rows.RowsAffected()
			return rowsAffected > 0
		}
	}
}
