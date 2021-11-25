package entities

type Fornecedor struct {
	Id        int64  `json:"id"`
	Nome      string `json:"nome"`
	Categoria string `json:"categoria"`
}
