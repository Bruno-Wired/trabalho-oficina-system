package models

import (
	"database/sql"
	"trabalhogocopia/config"
	"trabalhogocopia/entities"
)

type UsuarioModel struct {
}

//Update usuario
func (*UsuarioModel) UsuarioUpdate(usuario *entities.Usuario) bool {
	db, err := config.GetDB()

	if err != nil {
		return false

	} else {
		rows2, err2 := db.Exec("UPDATE usuarios SET login = ?, senha = ? WHERE id = ?;", usuario.Nome, usuario.Senha, usuario.Id)

		if err2 != nil {
			return false

		} else {
			rowsAffected, _ := rows2.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UsuarioModel) PreencheUsuario(usuario *entities.Usuario) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows := db.QueryRow("SELECT id, login, senha FROM usuarios WHERE id = ?;", usuario.Id)
		err3 := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Senha)

		if err3 != nil {
			return false

		} else {

			return true
		}
	}
}

//Deleta usuario
func (*UsuarioModel) UsuarioDelete(id int64) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		rows, err2 := db.Exec("DELETE FROM usuarios WHERE id = ?", id)

		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := rows.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UsuarioModel) NovoCadastro(usuario *entities.Usuario) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("INSERT usuarios(login, senha) VALUES(?, ?)", usuario.Nome, usuario.Senha)
		usuario.Id, _ = result.LastInsertId()

		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

func (*UsuarioModel) Autenticacao(usuario *entities.Usuario) bool {
	db, err := config.GetDB()

	if err != nil {
		return false
	}

	row := db.QueryRow("SELECT login, senha FROM usuarios WHERE login = ? AND senha = ?", usuario.Nome, usuario.Senha)
	err2 := row.Scan(&usuario.Nome, &usuario.Senha)

	return err2 != sql.ErrNoRows
}

//Listar usuarios
func (*UsuarioModel) FindAllUsuario() ([]entities.Usuario, error) {
	db, err := config.GetDB()

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT id, login, senha FROM usuarios")
		db.Close()

		if err2 != nil {
			return nil, err2
		} else {
			var usuarios []entities.Usuario

			for rows.Next() {
				var usuario entities.Usuario

				rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Senha)
				usuarios = append(usuarios, usuario)
			}
			rows.Close()
			return usuarios, nil
		}
	}
}
