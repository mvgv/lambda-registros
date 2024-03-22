package dto

type SolicitacaoRelatorio struct {
	Email string `json:"email"`
	Mes   string `json:"mes"`
}

func NewSolicitacaoRelatorio(email, mes string) *SolicitacaoRelatorio {
	return &SolicitacaoRelatorio{
		Email: email,
		Mes:   mes,
	}
}
