package models

import (
	"database/sql"
	"trabalhogocopia/config"
	"trabalhogocopia/entities"
)

type FornecedorModel struct {
}

//Listar fornecedor
func (*FornecedorModel) FindAllFornecedor() ([]entities.Fornecedor, error) {
	db, err := config.GetDB()

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT f.id, f.nome AS fornecedor, c.nome AS categoria FROM fornecedores f, categorias c WHERE f.categoria_id = c.id;")
		db.Close()

		if err2 != nil {
			return nil, err2
		} else {
			var fornecedores []entities.Fornecedor

			for rows.Next() {
				var fornecedor entities.Fornecedor

				rows.Scan(&fornecedor.Id, &fornecedor.Nome, &fornecedor.Categoria)
				fornecedores = append(fornecedores, fornecedor)
			}
			rows.Close()
			return fornecedores, nil
		}
	}
}

//Achar por ID
func (*FornecedorModel) FindFornecedor(id int64) (entities.Fornecedor, error) {
	db, err := config.GetDB()

	if err != nil {
		return entities.Fornecedor{}, err

	} else {
		var fornecedor entities.Fornecedor
		rows := db.QueryRow("SELECT f.id, f.nome AS fornecedor, c.nome AS categoria FROM fornecedores f, categorias c WHERE f.categoria_id = c.id AND f.id = ?;", id)
		db.Close()
		err2 := rows.Scan(&fornecedor.Id, &fornecedor.Nome, &fornecedor.Categoria)

		if err2 == sql.ErrNoRows {
			return entities.Fornecedor{}, err2

		} else {
			return fornecedor, nil
		}
	}
}

//Adicionar fornecedor
func (*FornecedorModel) CreateFornecedor(fornecedor *entities.Fornecedor) (bool, error) {
	db, err := config.GetDB()

	if err != nil {
		return false, err
	} else {

		rows, err2 := db.Exec("INSERT INTO fornecedores (nome, categoria_id) VALUES(?, ?);", fornecedor.Nome, fornecedor.Categoria)

		id := db.QueryRow("SELECT MAX(id) FROM fornecedores;")
		_ = id.Scan(&fornecedor.Id)

		db.Close()

		if err2 != nil {
			return false, err2

		} else {
			rowsAffected, _ := rows.RowsAffected()
			return rowsAffected > 0, nil
		}
	}
}

func (*FornecedorModel) PreencheFornecedor(fornecedor *entities.Fornecedor) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows := db.QueryRow("SELECT nome, categoria_id FROM fornecedores WHERE id = ?;", fornecedor.Id)
		err3 := rows.Scan(&fornecedor.Nome, &fornecedor.Categoria)
		db.Close()

		if err3 != nil {
			return false

		} else {

			return true
		}
	}
}

//Update fornecedor
func (*FornecedorModel) UpdateFornecedor(fornecedor *entities.Fornecedor) bool {
	db, err := config.GetDB()

	if err != nil {
		return false

	} else {
		rows2, err2 := db.Exec("UPDATE fornecedores SET nome = ?, categoria_id = ? WHERE id = ?;", fornecedor.Nome, fornecedor.Categoria, fornecedor.Id)
		db.Close()

		if err2 != nil {
			return false

		} else {
			rowsAffected, _ := rows2.RowsAffected()
			return rowsAffected > 0
		}
	}
}

//Deletar fornecedor
func (*FornecedorModel) DeleteFornecedor(id int64) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows, err2 := db.Exec("DELETE FROM fornecedores WHERE id = ?", id)
		db.Close()

		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := rows.RowsAffected()
			return rowsAffected > 0
		}
	}
}
