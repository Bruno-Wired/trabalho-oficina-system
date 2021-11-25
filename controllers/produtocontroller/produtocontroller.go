package produtocontroller

import (
	"net/http"
	"strconv"
	"trabalhogocopia/entities"
	"trabalhogocopia/models"

	"github.com/gin-gonic/gin"
)

//Todos os Produtos
func ProdutoIndex(c *gin.Context) {
	var produtoModel models.ProdutoModel

	produtos, _ := produtoModel.FindAll()
	c.JSON(http.StatusOK, gin.H{"Produtos": produtos})
}

//Produto por id
func ProdutoId(c *gin.Context) {
	var produtoModel models.ProdutoModel
	strId, _ := c.Params.Get("id")
	id, _ := strconv.ParseInt(strId, 10, 64)

	produto, err := produtoModel.Find(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"ERRO": "conteúdo não encontrado"})
	}

	c.JSON(http.StatusOK, gin.H{"Produto": produto})
}

//Adicionar produto
func ProdutoAdd(c *gin.Context) {
	var produto entities.Produto

	produto.Nome = c.PostForm("nome")
	produto.Categoria = c.PostForm("categoria")
	produto.Fornecedor = c.PostForm("fornecedor")
	produto.Descricao = c.PostForm("descricao")
	strQuantidade := c.PostForm("quantidademin")
	produto.QuantidadeMin, _ = strconv.ParseInt(strQuantidade, 10, 64)
	strPreco := c.PostForm("preco")
	produto.Preco, _ = strconv.ParseFloat(strPreco, 64)
	produto.Lote = c.PostForm("lote")

	var produtoModel models.ProdutoModel

	err := c.ShouldBind(&produto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao converter produto"})
		return
	}

	_, err2 := produtoModel.Create(&produto)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao adicionar produto", "ERRO": err2})
	} else {
		c.JSON(http.StatusOK, gin.H{"Produto adicionado": produto})
	}
}

//Deletando produto
func ProdutoDelete(c *gin.Context) {
	strId, _ := c.Params.Get("id")
	id, err := strconv.ParseInt(strId, 10, 64)

	if err != nil {
		c.JSON(404, gin.H{"erro": "Falha ao converter string"})
	} else {

		var produtoModel models.ProdutoModel

		err2 := produtoModel.Delete(id)
		if !err2 {
			c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao deletar o item"})
		} else {
			c.JSON(http.StatusOK, gin.H{"mensagem": "Item deletado com sucesso"})
		}
	}
}

//Editando produto
func ProdutoEdit(c *gin.Context) {
	var produto entities.Produto
	var produtoModel models.ProdutoModel

	strId, _ := c.Params.Get("id")
	id, _ := strconv.ParseInt(strId, 10, 64)
	produto.Id = id

	err2 := produtoModel.PreencheUpdate(&produto)
	if !err2 {
		c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "falha ao PREENCHER o item"})

	} else {
		nome := c.PostForm("nome")
		categoria := c.PostForm("categoria")
		fornecedor := c.PostForm("fornecedor")
		descricao := c.PostForm("descricao")
		quantidademin, _ := strconv.ParseInt(c.PostForm("quantidademin"), 10, 64)
		preco, _ := strconv.ParseFloat(c.PostForm("preco"), 64)
		hasChanged := false

		switch {
		case (nome != "" && nome != produto.Nome):
			produto.Nome = c.PostForm("nome")
			hasChanged = true

		case categoria != "" && categoria != produto.Categoria:
			produto.Categoria = c.PostForm("categoria")
			hasChanged = true

		case fornecedor != "" && fornecedor != produto.Fornecedor:
			produto.Fornecedor = c.PostForm("fornecedor")
			hasChanged = true

		case descricao != "" && descricao != produto.Descricao:
			produto.Descricao = c.PostForm("descricao")
			hasChanged = true

		case quantidademin != 0 && quantidademin != produto.QuantidadeMin:
			produto.QuantidadeMin, _ = strconv.ParseInt(c.PostForm("quantidademin"), 10, 64)
			hasChanged = true

		case preco != 0 && preco != produto.Preco:
			strPreco := c.PostForm("preco")
			produto.Preco, _ = strconv.ParseFloat(strPreco, 64)
			hasChanged = true

		default:
			c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "Não houve modificações no item!"})
		}

		if hasChanged {
			UpdateProd(c, produto)
		}
	}
}

//Funções auxiliares
func UpdateProd(c *gin.Context, produto entities.Produto) (teste bool) {
	var produtoModel models.ProdutoModel
	teste = produtoModel.Update(&produto)

	if !teste {
		c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "Falha ao modificar item!"})
		return teste

	} else {
		c.JSON(http.StatusOK, gin.H{"mensagem": "Item atualizado com sucesso", "Produto": produto})
		return
	}
}
