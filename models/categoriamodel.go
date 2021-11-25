package models

import (
	"database/sql"
	"trabalhogocopia/config"
	"trabalhogocopia/entities"
)

type CategoriaModel struct {
}

//Listar categorias
func (*CategoriaModel) FindAllCategoria() ([]entities.Categoria, error) {
	db, err := config.GetDB()

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, nome FROM categorias ORDER BY nome;")
		db.Close()

		if err2 != nil {
			return nil, err2
		} else {
			var categorias []entities.Categoria

			for rows.Next() {
				var categoria entities.Categoria

				rows.Scan(&categoria.Id, &categoria.Nome)
				categorias = append(categorias, categoria)
			}
			rows.Close()
			return categorias, nil
		}
	}
}

//Achar por ID
func (*CategoriaModel) FindCategoria(id int64) (entities.Categoria, error) {
	db, err := config.GetDB()

	if err != nil {
		return entities.Categoria{}, err

	} else {
		var categoria entities.Categoria
		rows := db.QueryRow("SELECT id, nome FROM categorias WHERE id = ?;", id)
		db.Close()
		err2 := rows.Scan(&categoria.Id, &categoria.Nome)

		if err2 == sql.ErrNoRows {
			return entities.Categoria{}, err2

		} else {
			return categoria, nil
		}
	}
}

//Adicionar categoria
func (*CategoriaModel) CreateCategoria(categoria *entities.Categoria) (bool, error) {
	db, err := config.GetDB()

	if err != nil {
		return false, err
	} else {

		rows, err2 := db.Exec("INSERT INTO categorias(nome) VALUES(?);", categoria.Nome)

		id := db.QueryRow("SELECT MAX(id) FROM categorias;")
		_ = id.Scan(&categoria.Id)

		db.Close()

		if err2 != nil {
			return false, err2

		} else {
			rowsAffected, _ := rows.RowsAffected()
			return rowsAffected > 0, nil
		}
	}
}

func (*CategoriaModel) PreencheCategoria(categoria *entities.Categoria) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows := db.QueryRow("SELECT nome FROM categorias WHERE id = ?;", categoria.Id)
		err3 := rows.Scan(&categoria.Nome)

		if err3 != nil {
			return false

		} else {
			return true
		}
	}
}

//Update categoria
func (*CategoriaModel) UpdateCategoria(categoria *entities.Categoria) bool {
	db, err := config.GetDB()

	if err != nil {
		return false

	} else {
		rows2, err2 := db.Exec("UPDATE categorias SET nome = ? WHERE id = ?;", categoria.Nome, categoria.Id)

		if err2 != nil {
			return false

		} else {
			rowsAffected, _ := rows2.RowsAffected()
			return rowsAffected > 0
		}
	}
}

//Deletar categoria
func (*CategoriaModel) DeleteCategoria(id int64) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows, err2 := db.Exec("DELETE FROM categorias WHERE id = ?", id)

		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := rows.RowsAffected()
			return rowsAffected > 0
		}
	}
}
