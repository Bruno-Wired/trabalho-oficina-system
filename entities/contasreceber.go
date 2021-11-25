package entities

type ContasReceber struct {
	Id       int64   `json:"id"`
	Cliente  string  `json:"cliente"`
	Produto  string  `json:"produto"`
	Forma    string  `json:"forma"`
	Valor    float64 `json:"valor"`
	DataAtt  string  `json:"dataatt"`
	DataNova string  `json:"datanova"`
	Situacao string  `json:"situacao"`
}
