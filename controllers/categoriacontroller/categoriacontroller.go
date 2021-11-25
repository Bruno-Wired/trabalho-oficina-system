package categoriacontroller

import (
	"net/http"
	"strconv"
	"trabalhogocopia/entities"
	"trabalhogocopia/models"

	"github.com/gin-gonic/gin"
)

//Todos as categorias
func CategoriaIndex(c *gin.Context) {
	var categoriaModel models.CategoriaModel

	categorias, _ := categoriaModel.FindAllCategoria()
	c.JSON(http.StatusOK, gin.H{"Categorias": categorias})
}

//categoria por id
func CategoriaId(c *gin.Context) {
	var categoriaModel models.CategoriaModel
	strId, _ := c.Params.Get("id")
	id, _ := strconv.ParseInt(strId, 10, 64)

	categoria, err := categoriaModel.FindCategoria(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"ERRO": "conteúdo não encontrado"})
	}

	c.JSON(http.StatusOK, gin.H{"Categoria": categoria})
}

//Adicionar categoria
func CategoriaAdd(c *gin.Context) {
	var categoria entities.Categoria

	categoria.Nome = c.PostForm("nome")

	var categoriaModel models.CategoriaModel

	err := c.ShouldBind(&categoria)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao converter categoria"})
		return
	}

	_, err2 := categoriaModel.CreateCategoria(&categoria)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao adicionar categoria", "ERRO": err2})
	} else {
		c.JSON(http.StatusOK, gin.H{"Categoria adicionada": categoria})
	}
}

//Deletando categoria
func CategoriaDelete(c *gin.Context) {
	strId, _ := c.Params.Get("id")
	id, err := strconv.ParseInt(strId, 10, 64)

	if err != nil {
		c.JSON(404, gin.H{"erro": "Falha ao converter string"})
	} else {

		var categoriaModel models.CategoriaModel

		err2 := categoriaModel.DeleteCategoria(id)
		if !err2 {
			c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao deletar categoria"})
		} else {
			c.JSON(http.StatusOK, gin.H{"mensagem": "Categoria deletada com sucesso"})
		}
	}
}

//Editando categoria
func CategoriaEdit(c *gin.Context) {
	var categoria entities.Categoria
	var categoriaModel models.CategoriaModel

	strId, _ := c.Params.Get("id")
	id, _ := strconv.ParseInt(strId, 10, 64)
	categoria.Id = id

	err2 := categoriaModel.PreencheCategoria(&categoria)
	if !err2 {
		c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "falha ao PREENCHER a categoria"})

	} else {
		nome := c.PostForm("nome")

		switch {
		case (nome != "" && nome != categoria.Nome):
			categoria.Nome = c.PostForm("nome")
			UpdateCategoria(c, categoria)

		default:
			c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "Não houve modificações na categoria!"})
		}
	}
}

//Funções auxiliares
func UpdateCategoria(c *gin.Context, categoria entities.Categoria) (teste bool) {
	var categoriaModel models.CategoriaModel
	teste = categoriaModel.UpdateCategoria(&categoria)

	if !teste {
		c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "Falha ao modificar categoria!"})
		return teste

	} else {
		c.JSON(http.StatusOK, gin.H{"mensagem": "Categoria atualizada com sucesso", "categoria": categoria})
		return
	}
}
