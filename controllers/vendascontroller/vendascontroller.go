package vendascontroller

import (
	"net/http"
	"strconv"
	"trabalhogocopia/entities"
	"trabalhogocopia/models"

	"github.com/gin-gonic/gin"
)

//Todas as vendas
func VendaIndex(c *gin.Context) {
	var vandaModel models.VendasModel

	vendas, _ := vandaModel.FindAllVendas()
	c.JSON(http.StatusOK, gin.H{"vendas efetuadas": vendas})
}

//Vendas por id
func VendaId(c *gin.Context) {
	var vandaModel models.VendasModel

	strId, _ := c.Params.Get("id")
	id, _ := strconv.ParseInt(strId, 10, 64)

	venda, err := vandaModel.VendaFind(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"ERRO": "conteúdo não encontrado"})
	}

	c.JSON(http.StatusOK, gin.H{"venda": venda})
}

//Adicionar venda
func VendaAdd(c *gin.Context) {
	var venda entities.Venda

	venda.Cliente = c.PostForm("cliente")
	venda.Produto = c.PostForm("produto")
	venda.Quantidade, _ = strconv.ParseInt(c.PostForm("quantidade"), 10, 64)

	var vendaModel models.VendasModel

	err := c.ShouldBind(&venda)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao converter venda"})
		return
	}

	err2 := vendaModel.CreateVenda(&venda)
	if !err2 {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao adicionar venda"})
	} else {
		c.JSON(http.StatusOK, gin.H{"venda adicionada com sucesso": venda})
	}
}
