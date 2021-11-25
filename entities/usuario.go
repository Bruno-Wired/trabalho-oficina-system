package entities

type Usuario struct {
	Id    int64  `json:"id"`
	Nome  string `json:"nome"`
	Senha string `json:"senha"`
}
