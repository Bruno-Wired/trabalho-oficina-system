package entities

type Estoque struct {
	Id             int64   `json:"id"`
	Lote           string  `json:"lote"`
	DataEntrada    string  `json:"dataentrada"`
	DataVencimento string  `json:"vencimento"`
	Quantidade     int64   `json:"quantidade"`
	PrecoUnit      float64 `json:"precounit"`
	Total          float64 `json:"total"`
	Produto        string  `json:"produto"`
	Fornecedor     string  `json:"fornecedor"` //Apenas usado no Find
}
