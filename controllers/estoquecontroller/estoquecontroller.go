package estoquecontroller

import (
	"net/http"
	"strconv"
	"trabalhogocopia/entities"
	"trabalhogocopia/models"

	"github.com/gin-gonic/gin"
)

//Todos o estoque
func EstoqueIndex(c *gin.Context) {
	var estoqueModel models.EstoqueModel

	produtos, _ := estoqueModel.FindAllEstoque()
	c.JSON(http.StatusOK, gin.H{"Produtos em estoque": produtos})
}

//Estoque por id
func EstoqueId(c *gin.Context) {
	var estoqueModel models.EstoqueModel
	strId, _ := c.Params.Get("id")
	id, _ := strconv.ParseInt(strId, 10, 64)

	estoque, err := estoqueModel.EstoqueFind(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"ERRO": "conteúdo não encontrado"})
	}

	c.JSON(http.StatusOK, gin.H{"Estoque do produto": estoque})
}

//Adicionar estoque
func EstoqueAdd(c *gin.Context) {
	var estoque entities.Estoque

	estoque.Lote = c.PostForm("lote")
	estoque.DataVencimento = c.PostForm("vencimento")

	if estoque.DataVencimento == "" {
		estoque.DataVencimento = "00-00-0000"
	}

	strQuantidade := c.PostForm("quantidade")
	estoque.Quantidade, _ = strconv.ParseInt(strQuantidade, 10, 64)

	var estoqueModel models.EstoqueModel

	err := c.ShouldBind(&estoque)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao converter estoque"})
		return
	}

	err2 := estoqueModel.CreateEstoque(&estoque)
	if !err2 {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao adicionar produto ao estoque"})
	} else {
		c.JSON(http.StatusOK, gin.H{"Produto adicionado com sucesso ao estoque": estoque})
	}
}

//Deletando estoque
func EstoqueDel(c *gin.Context) {
	strId, _ := c.Params.Get("id")
	id, err := strconv.ParseInt(strId, 10, 64)

	if err != nil {
		c.JSON(404, gin.H{"erro": "Falha ao converter string"})
	} else {

		var estoqueModel models.EstoqueModel

		err2 := estoqueModel.EstoqueDelete(id)
		if !err2 {
			c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao deletar estoque"})
		} else {
			c.JSON(http.StatusOK, gin.H{"mensagem": "Estoque deletado com sucesso"})
		}
	}
}

//Editando estoque
func EstoqueEdit(c *gin.Context) {
	var estoque entities.Estoque
	var estoqueModel models.EstoqueModel

	strId, _ := c.Params.Get("id")
	id, _ := strconv.ParseInt(strId, 10, 64)
	estoque.Id = id

	err2 := estoqueModel.PreencheEstoque(&estoque)
	if !err2 {
		c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "falha ao PREENCHER o estoque"})

	} else {
		lote := c.PostForm("lote")
		vencimento := c.PostForm("vencimento")
		quantidade, _ := strconv.ParseInt(c.PostForm("quantidade"), 10, 64)
		preco, _ := strconv.ParseFloat(c.PostForm("preco"), 64)
		fornecedor := c.PostForm("fornecedor")
		hasChanged := false

		switch {
		case (lote != "" && lote != estoque.Lote):
			estoque.Lote = lote
			hasChanged = true

		case (vencimento != "" && vencimento != estoque.DataVencimento):
			estoque.DataVencimento = vencimento
			hasChanged = true

		case quantidade != 0 && quantidade != estoque.Quantidade:
			estoque.Quantidade = quantidade
			hasChanged = true

		case preco != 0 && preco != estoque.PrecoUnit:
			estoque.PrecoUnit = preco
			hasChanged = true

		case fornecedor != "" && fornecedor != estoque.Fornecedor:
			estoque.Fornecedor = fornecedor
			hasChanged = true

		default:
			c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "Não houve modificações no estoque!"})
		}

		if hasChanged {
			UpdateEst(c, estoque)
		}
	}

}

//Funções auxiliares
func UpdateEst(c *gin.Context, estoque entities.Estoque) (teste bool) {
	var estoqueModel models.EstoqueModel
	teste = estoqueModel.EstoqueUpdate(&estoque)

	if !teste {
		c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "Falha ao modificar estoque!"})
		return teste

	} else {
		c.JSON(http.StatusOK, gin.H{"mensagem": "Estoque atualizado com sucesso", "Estoque": estoque})
		return
	}
}
