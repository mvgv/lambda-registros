package dto

type PontoDoDiaEntidade struct {
	Email     string          `json:"email"`
	Registros []PontoEntidade `json:"registros"`
}

func NewPontoDoDiaEntidade(email string, registros []PontoEntidade) *PontoDoDiaEntidade {
	return &PontoDoDiaEntidade{
		Email:     email,
		Registros: registros,
	}
}
