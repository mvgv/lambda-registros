package apresentacao

type PontoRespostaRequisicao struct {
	Ponto  string `json:"ponto"`
	Evento string `json:"evento"`
}

func NewPontoRespostaRequisicao(ponto, evento string) *PontoRespostaRequisicao {
	return &PontoRespostaRequisicao{
		Ponto:  ponto,
		Evento: evento,
	}
}
