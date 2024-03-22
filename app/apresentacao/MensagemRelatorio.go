package apresentacao

type MensagemRelatorio struct {
	Mensagem string `json:"mensagem"`
}

func NewMensagemRelatorio(mensagem string) *MensagemRelatorio {
	return &MensagemRelatorio{
		Mensagem: mensagem,
	}
}
