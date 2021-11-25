package fornecedorcontroller

import (
	"net/http"
	"strconv"
	"trabalhogocopia/entities"
	"trabalhogocopia/models"

	"github.com/gin-gonic/gin"
)

//Todos os fornecedor
func FornecedorIndex(c *gin.Context) {
	var fornecedorModel models.FornecedorModel

	fornecedor, _ := fornecedorModel.FindAllFornecedor()
	c.JSON(http.StatusOK, gin.H{"Fornecedor": fornecedor})
}

//Fornecedor por id
func FornecedorId(c *gin.Context) {
	var fornecedorModel models.FornecedorModel
	strId, _ := c.Params.Get("id")
	id, _ := strconv.ParseInt(strId, 10, 64)

	fornecedor, err := fornecedorModel.FindFornecedor(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"ERRO": "conteúdo não encontrado"})
	}

	c.JSON(http.StatusOK, gin.H{"Fornecedor": fornecedor})
}

//Adicionar fornecedor
func FornecedorAdd(c *gin.Context) {
	var fornecedor entities.Fornecedor

	fornecedor.Nome = c.PostForm("nome")
	fornecedor.Categoria = c.PostForm("categoria")

	var fornecedorModel models.FornecedorModel

	err := c.ShouldBind(&fornecedor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao converter fornecedor"})
		return
	}

	_, err2 := fornecedorModel.CreateFornecedor(&fornecedor)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao adicionar fornecedor", "ERRO": err2})
	} else {
		c.JSON(http.StatusOK, gin.H{"Fornecedor adicionado": fornecedor})
	}
}

//Deletando fornecedor
func FornecedorDelete(c *gin.Context) {
	strId, _ := c.Params.Get("id")
	id, err := strconv.ParseInt(strId, 10, 64)

	if err != nil {
		c.JSON(404, gin.H{"erro": "Falha ao converter string"})
	} else {

		var fornecedorModel models.FornecedorModel

		err2 := fornecedorModel.DeleteFornecedor(id)
		if !err2 {
			c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao deletar fornecedor"})
		} else {
			c.JSON(http.StatusOK, gin.H{"mensagem": "Fornecedor deletado com sucesso"})
		}
	}
}

//Editando fornecedor
func FornecedorEdit(c *gin.Context) {
	var fornecedor entities.Fornecedor
	var fornecedorModel models.FornecedorModel

	strId, _ := c.Params.Get("id")
	id, _ := strconv.ParseInt(strId, 10, 64)
	fornecedor.Id = id

	err2 := fornecedorModel.PreencheFornecedor(&fornecedor)
	if !err2 {
		c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "falha ao PREENCHER o fornecedor"})

	} else {
		nome := c.PostForm("nome")
		categoria := c.PostForm("categoria")
		hasChanged := false

		switch {
		case (nome != "" && nome != fornecedor.Nome):
			fornecedor.Nome = nome
			hasChanged = true

		case categoria != "" && categoria != fornecedor.Categoria:
			fornecedor.Categoria = categoria
			hasChanged = true

		default:
			c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "Não houve modificações no fornecedor!"})
		}

		if hasChanged {
			UpdateFornecedor(c, fornecedor)
		}
	}
}

//Funções auxiliares
func UpdateFornecedor(c *gin.Context, fornecedor entities.Fornecedor) (teste bool) {
	var fornecedorModel models.FornecedorModel
	teste = fornecedorModel.UpdateFornecedor(&fornecedor)

	if !teste {
		c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "Falha ao modificar fornecedor!"})
		return teste

	} else {
		c.JSON(http.StatusOK, gin.H{"mensagem": "Fornecedor atualizado com sucesso", "Fornecedor": fornecedor})
		return
	}
}
