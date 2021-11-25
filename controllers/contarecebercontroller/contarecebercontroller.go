package contarecebercontroller

import (
	"net/http"
	"strconv"
	"trabalhogocopia/entities"
	"trabalhogocopia/models"

	"github.com/gin-gonic/gin"
)

//Adicionar contas a receber
func ContaReceberAdd(c *gin.Context) {
	var conta entities.ContasReceber

	conta.Cliente = c.PostForm("cliente")
	conta.Produto = c.PostForm("produto")
	conta.Forma = c.PostForm("forma")
	conta.Valor, _ = strconv.ParseFloat(c.PostForm("valor"), 64)
	conta.Situacao = c.PostForm("situacao")

	var contaReceberModel models.ContaReceberModel

	err := c.ShouldBind(&conta)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao converter conta"})
		return
	}

	err2 := contaReceberModel.CreateContaReceber(&conta)
	if !err2 {
		c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao adicionar conta"})
	} else {
		c.JSON(http.StatusOK, gin.H{"Conta adicionada com sucesso": conta})
	}
}

//Todas as contas
func ContaReceberIndex(c *gin.Context) {
	var contaReceberModel models.ContaReceberModel

	contas, _ := contaReceberModel.FindAllContasReceber()
	c.JSON(http.StatusOK, gin.H{"Contas a receber": contas})
}

//Contas por id
func ContaReceberId(c *gin.Context) {
	var contaReceberModel models.ContaReceberModel

	strId, _ := c.Params.Get("id")
	id, _ := strconv.ParseInt(strId, 10, 64)

	contas, err := contaReceberModel.FindContaReceber(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"ERRO": "conteúdo não encontrado"})
	}

	c.JSON(http.StatusOK, gin.H{"Contas a receber": contas})
}

//Deletando conta
func ContaReceberDel(c *gin.Context) {
	strId, _ := c.Params.Get("id")
	id, err := strconv.ParseInt(strId, 10, 64)

	if err != nil {
		c.JSON(404, gin.H{"erro": "Falha ao converter string"})
	} else {

		var contaReceberModel models.ContaReceberModel

		err2 := contaReceberModel.DeleteContaReceber(id)
		if !err2 {
			c.JSON(http.StatusBadRequest, gin.H{"mensagem": "falha ao deletar a conta"})
		} else {
			c.JSON(http.StatusOK, gin.H{"mensagem": "Conta deletada com sucesso"})
		}
	}
}

//Editando conta
func ContaReceberEdit(c *gin.Context) {
	var contas entities.ContasReceber
	var contaReceberModel models.ContaReceberModel

	strId, _ := c.Params.Get("id")
	id, _ := strconv.ParseInt(strId, 10, 64)
	contas.Id = id

	err2 := contaReceberModel.PreencheContaReceber(&contas)
	if !err2 {
		c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "falha ao PREENCHER a conta"})

	} else {
		cliente := c.PostForm("cliente")
		produto := c.PostForm("produto")
		forma := c.PostForm("pagamento")
		valor, _ := strconv.ParseFloat(c.PostForm("preco"), 64)
		situacao := c.PostForm("situacao")
		hasChanged := false

		switch {
		case (cliente != "" && cliente != contas.Cliente):
			contas.Cliente = cliente
			hasChanged = true

		case produto != "" && produto != contas.Produto:
			contas.Produto = produto
			hasChanged = true

		case forma != "" && forma != contas.Forma:
			contas.Situacao = forma
			hasChanged = true

		case valor != 0 && valor != contas.Valor:
			contas.Valor = valor
			hasChanged = true

		case situacao != contas.Situacao:
			contas.Situacao = situacao
			hasChanged = true

		default:
			c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "Não houve modificações na conta!"})
		}

		if hasChanged {
			UpdateCont(c, contas)
		}
	}
}

//Funções auxiliares
func UpdateCont(c *gin.Context, contas entities.ContasReceber) (teste bool) {
	var contasModel models.ContaReceberModel
	teste = contasModel.Update(&contas)

	if !teste {
		c.JSON(http.StatusBadRequest, gin.H{"UPDATE ERRO": "Falha ao modificar conta!"})
		return teste

	} else {
		c.JSON(http.StatusOK, gin.H{"mensagem": "Conta atualizada com sucesso", "Conta": contas})
		return
	}
}
