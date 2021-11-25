package usuariocontroller

import (
	"net/http"
	"strconv"
	"trabalhogocopia/entities"
	"trabalhogocopia/models"

	"github.com/gin-gonic/gin"
)

//Todos os usuarios
func UsuarioIndex(c *gin.Context) {
	var usuarioModel models.UsuarioModel

	usuarios, _ := usuarioModel.FindAllUsuario()
	c.JSON(http.StatusOK, gin.H{"Usuarios": usuarios})
}

//Editando usuario
func UsuarioEdit(c *gin.Context) {
	var usuario entities.Usuario
	var usuarioModel models.UsuarioModel

	strId, _ := c.Params.Get("id")
	id, _ := strconv.ParseInt(strId, 10, 64)
	usuario.Id = id

	err2 := usuarioModel.PreencheUsuario(&usuario)
	if !err2 {
		c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "falha ao PREENCHER o usuario"})

	} else {
		nome := c.PostForm("nome")
		senha := c.PostForm("senha")
		hasChanged := false

		switch {
		case (nome != "" && nome != usuario.Nome):
			usuario.Nome = nome
			hasChanged = true

		case senha != "" && senha != usuario.Senha:
			usuario.Senha = senha
			hasChanged = true

		default:
			c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "Não houve modificações no item!"})
		}

		if hasChanged {
			UpdateUsu(c, usuario)
		}
	}
}

//Deletando usuario
func UsuarioDel(c *gin.Context) {
	strId, _ := c.Params.Get("id")
	id, err := strconv.ParseInt(strId, 10, 64)

	if err != nil {
		c.JSON(404, gin.H{"erro": "Falha ao converter string"})
	} else {

		var usuarioModel models.UsuarioModel

		err2 := usuarioModel.UsuarioDelete(id)
		if !err2 {
			c.JSON(http.StatusBadRequest, gin.H{"mensagem": "Falha ao deletar usuario"})
		} else {
			c.JSON(http.StatusOK, gin.H{"mensagem": "Usuario deletado com sucesso"})
		}
	}
}

func UsuarioAdd(c *gin.Context) {
	var usuario entities.Usuario

	usuario.Nome = c.PostForm("nome")
	usuario.Senha = c.PostForm("senha")

	var usuarioModel models.UsuarioModel

	err := c.ShouldBind(&usuario)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao adicionar usuario"})
		return
	}

	usuarioModel.NovoCadastro(&usuario)
	c.JSON(http.StatusOK, gin.H{"Usuario adicionado com sucesso": usuario})
}

func UsuarioLogin(c *gin.Context) {
	var usuario entities.Usuario

	usuario.Nome = c.PostForm("nome")
	usuario.Senha = c.PostForm("senha")

	var usuarioModel models.UsuarioModel

	err := c.ShouldBind(&usuario)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao adicionar usuario"})
		return
	}

	err2 := usuarioModel.Autenticacao(&usuario)

	if !err2 {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha na autenticação"})
	} else {
		c.JSON(http.StatusOK, gin.H{"mensagem": "Logado com sucesso"})
	}
}

//Funções auxiliares
func UpdateUsu(c *gin.Context, usuario entities.Usuario) (teste bool) {
	var usuarioModel models.UsuarioModel
	teste = usuarioModel.UsuarioUpdate(&usuario)

	if !teste {
		c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "Falha ao modificar usuário!"})
		return teste

	} else {
		c.JSON(http.StatusOK, gin.H{"mensagem": "Usuário atualizado com sucesso", "usuário": usuario})
		return
	}
}
