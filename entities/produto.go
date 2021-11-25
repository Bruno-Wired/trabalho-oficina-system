package entities

type Produto struct {
	Id             int64   `json:"id"`
	Nome           string  `json:"nome"`
	Categoria      string  `json:"categoria"`
	Fornecedor     string  `json:"fornecedor"`
	Descricao      string  `json:"descricao"`
	QuantidadeMin  int64   `json:"quantidademin"`
	Preco          float64 `json:"preco"`
	Lote           string  `json:"lote"`
	DataVencimento string  `json:"vencimento"`
}
