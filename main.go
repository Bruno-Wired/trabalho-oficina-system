package main

import (
	"trabalhogocopia/controllers/categoriacontroller"
	"trabalhogocopia/controllers/clientecontroller"
	"trabalhogocopia/controllers/contarecebercontroller"
	"trabalhogocopia/controllers/estoquecontroller"
	"trabalhogocopia/controllers/fornecedorcontroller"
	"trabalhogocopia/controllers/produtocontroller"
	"trabalhogocopia/controllers/usuariocontroller"
	"trabalhogocopia/controllers/vendascontroller"

	"github.com/gin-gonic/gin"
)

func main() {
	servidor := gin.Default()

	//Categorias End-Point -
	servidor.GET("/categorias", categoriacontroller.CategoriaIndex)
	servidor.GET("/categorias/:id", categoriacontroller.CategoriaId)
	servidor.POST("/categorias/add", categoriacontroller.CategoriaAdd)
	servidor.DELETE("/categorias/delete/:id", categoriacontroller.CategoriaDelete)
	servidor.PUT("/categorias/edit/:id", categoriacontroller.CategoriaEdit)

	//Estoque End-Point -
	servidor.GET("/estoque", estoquecontroller.EstoqueIndex)
	servidor.GET("/estoque/:id", estoquecontroller.EstoqueId)
	servidor.POST("/estoque/add", estoquecontroller.EstoqueAdd) //opcional
	servidor.DELETE("/estoque/delete/:id", estoquecontroller.EstoqueDel)
	servidor.PUT("/estoque/edit/:id", estoquecontroller.EstoqueEdit)

	//Produtos End-Point -
	servidor.GET("/produtos", produtocontroller.ProdutoIndex)
	servidor.GET("/produtos/:id", produtocontroller.ProdutoId)
	servidor.POST("/produtos/add", produtocontroller.ProdutoAdd, estoquecontroller.EstoqueAdd)
	servidor.DELETE("/produtos/delete/:id", produtocontroller.ProdutoDelete)
	servidor.PUT("/produtos/edit/:id", produtocontroller.ProdutoEdit)

	//Fornecedor End-Point -
	servidor.GET("/fornecedor", fornecedorcontroller.FornecedorIndex)
	servidor.GET("/fornecedor/:id", fornecedorcontroller.FornecedorId)
	servidor.POST("/fornecedor/add", fornecedorcontroller.FornecedorAdd)
	servidor.DELETE("/fornecedor/delete/:id", fornecedorcontroller.FornecedorDelete)
	servidor.PUT("/fornecedor/edit/:id", fornecedorcontroller.FornecedorEdit)

	//Vendas End-Point -
	servidor.GET("/vendas", vendascontroller.VendaIndex)
	servidor.GET("/vendas/:id", vendascontroller.VendaId)
	servidor.POST("/vendas/add", vendascontroller.VendaAdd)

	//Clientes End-Point -
	servidor.GET("/clientes", clientecontroller.ClienteIndex)
	servidor.GET("/clientes/:id", clientecontroller.ClienteId)
	servidor.POST("/clientes/add", clientecontroller.ClienteAdd)
	servidor.DELETE("/clientes/delete/:id", clientecontroller.ClienteDelete)
	servidor.PUT("/clientes/edit/:id", clientecontroller.ClienteEdit)

	//Contas End-Point
	servidor.GET("/contas/receber", contarecebercontroller.ContaReceberIndex)
	servidor.GET("/contas/receber/:id", contarecebercontroller.ContaReceberId)
	servidor.POST("/contas/receber/add", contarecebercontroller.ContaReceberAdd)
	servidor.DELETE("/contas/receber/delete/:id", contarecebercontroller.ContaReceberDel)
	servidor.PUT("/contas/receber/edit/:id", contarecebercontroller.ContaReceberEdit)

	//Usuario End-Point
	servidor.GET("/usuarios", usuariocontroller.UsuarioIndex)
	servidor.POST("/usuario/add", usuariocontroller.UsuarioAdd)
	servidor.POST("/usuario/login", usuariocontroller.UsuarioLogin)
	servidor.DELETE("/usuario/delete/:id", usuariocontroller.UsuarioDel)
	servidor.PUT("/usuario/edit/:id", usuariocontroller.UsuarioEdit)

	_ = servidor.Run(":3000")

}
