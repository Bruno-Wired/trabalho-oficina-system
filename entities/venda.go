package entities

type Venda struct {
	Id         int64   `json:"id"`
	Cliente    string  `json:"cliente"`
	DataVenda  string  `json:"datavenda"`
	Vencimento string  `json:"vencimento"`
	Produto    string  `json:"produto"`
	Quantidade int64   `json:"quantidade"`
	PrecoUnit  float64 `json:"preco"`
	Total      float64 `json:"total"`
}
