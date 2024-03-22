package apresentacao

type PontoDoDiaRespostaRequisicao struct {
	PontoRespostaRequisicao []*PontoRespostaRequisicao `json:"ponto"`
}

func NewPontoDoDiaRespostaRequisicao(ponto []*PontoRespostaRequisicao) *PontoDoDiaRespostaRequisicao {
	return &PontoDoDiaRespostaRequisicao{
		PontoRespostaRequisicao: ponto,
	}
}
