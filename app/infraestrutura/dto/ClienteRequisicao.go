package dto

type ClienteRequisicao struct {
	Email string `json:"email"`
}

func NewClienteRequisicao(email string) *ClienteRequisicao {
	return &ClienteRequisicao{
		Email: email,
	}
}
