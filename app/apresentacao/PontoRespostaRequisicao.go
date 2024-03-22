package apresentacao

type PontoRespostaRequisicao struct {
	Ponto string `json:"ponto"`
}

func NewPontoRespostaRequisicao(ponto string) *PontoRespostaRequisicao {
	return &PontoRespostaRequisicao{
		Ponto: ponto,
	}
}
