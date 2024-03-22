package apresentacao

type PontoDoDiaRespostaRequisicao struct {
	PontoRespostaRequisicao []*PontoRespostaRequisicao `json:"ponto"`
	HorasTrabalhadas        string                     `json:"horas_trabalhadas"`
}

func NewPontoDoDiaRespostaRequisicao(ponto []*PontoRespostaRequisicao, horasTrabalhadas string) *PontoDoDiaRespostaRequisicao {
	return &PontoDoDiaRespostaRequisicao{
		PontoRespostaRequisicao: ponto,
		HorasTrabalhadas:        horasTrabalhadas,
	}
}
