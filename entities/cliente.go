package entities

type Cliente struct {
	Id        int64  `json:"id"`
	Nome      string `json:"nome"`
	NomeMeio  string `json:"nomemeio"`
	SobreNome string `json:"sobrenome"`
	Rg        string `json:"rg"`
}
