package clientecontroller

import (
	"net/http"
	"strconv"
	"trabalhogocopia/entities"
	"trabalhogocopia/models"

	"github.com/gin-gonic/gin"
)

//Todos os clientes
func ClienteIndex(c *gin.Context) {
	var clienteModel models.ClienteModel

	cliente, _ := clienteModel.FindAllCliente()
	c.JSON(http.StatusOK, gin.H{"Clientes": cliente})
}

//clientes por id
func ClienteId(c *gin.Context) {
	var clienteModel models.ClienteModel
	strId, _ := c.Params.Get("id")
	id, _ := strconv.ParseInt(strId, 10, 64)

	cliente, err := clienteModel.FindCliente(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"ERRO": "conteúdo não encontrado"})
	}

	c.JSON(http.StatusOK, gin.H{"Cliente": cliente})
}

//Adicionar clientes
func ClienteAdd(c *gin.Context) {
	var cliente entities.Cliente

	cliente.Nome = c.PostForm("nome")
	cliente.NomeMeio = c.PostForm("nomemeio")
	cliente.SobreNome = c.PostForm("sobrenome")
	cliente.Rg = c.PostForm("rg")

	var clienteModel models.ClienteModel

	err := c.ShouldBind(&cliente)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao converter cliente"})
		return
	}

	_, err2 := clienteModel.CreateCliente(&cliente)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao adicionar cliente", "ERRO": err2})
	} else {
		c.JSON(http.StatusOK, gin.H{"Cliente adicionado": cliente})
	}
}

//Deletando clientes
func ClienteDelete(c *gin.Context) {
	strId, _ := c.Params.Get("id")
	id, err := strconv.ParseInt(strId, 10, 64)

	if err != nil {
		c.JSON(404, gin.H{"erro": "Falha ao converter string"})
	} else {

		var clienteModel models.ClienteModel

		err2 := clienteModel.DeleteCliente(id)
		if !err2 {
			c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao deletar cliente"})
		} else {
			c.JSON(http.StatusOK, gin.H{"mensagem": "Cliente deletado com sucesso"})
		}
	}
}

//Editando clientes
func ClienteEdit(c *gin.Context) {
	var cliente entities.Cliente
	var clienteModel models.ClienteModel

	strId, _ := c.Params.Get("id")
	id, _ := strconv.ParseInt(strId, 10, 64)
	cliente.Id = id

	err2 := clienteModel.PreencheCliente(&cliente)
	if !err2 {
		c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "falha ao PREENCHER o cliente"})

	} else {
		nome := c.PostForm("nome")
		nomeMeio := c.PostForm("nomemeio")
		sobreNome := c.PostForm("sobrenome")
		rg := c.PostForm("rg")
		hasChanged := false

		switch {
		case (nome != "" && nome != cliente.Nome):
			cliente.Nome = nome
			hasChanged = true

		case nomeMeio != "" && nomeMeio != cliente.NomeMeio:
			cliente.NomeMeio = nomeMeio
			hasChanged = true

		case sobreNome != "" && sobreNome != cliente.SobreNome:
			cliente.SobreNome = sobreNome
			hasChanged = true

		case rg != "" && rg != cliente.Rg:
			cliente.Rg = rg
			hasChanged = true

		default:
			c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "Não houve modificações no cliente!"})
		}

		if hasChanged {
			UpdateCli(c, cliente)
		}
	}
}

//Funções auxiliares
func UpdateCli(c *gin.Context, cliente entities.Cliente) (teste bool) {
	var clienteModel models.ClienteModel
	teste = clienteModel.UpdateCliente(&cliente)

	if !teste {
		c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "Falha ao modificar cliente!"})
		return teste

	} else {
		c.JSON(http.StatusOK, gin.H{"mensagem": "Cliente atualizado com sucesso", "Cliente": cliente})
		return
	}
}
