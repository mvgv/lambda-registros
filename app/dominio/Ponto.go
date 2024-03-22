package dominio

type Ponto struct {
	Horario string `json:"horario"`
	Tipo    string `json:"tipo"`
}

func NewPonto(horario, tipo string) *Ponto {
	return &Ponto{
		Horario: horario,
		Tipo:    tipo,
	}
}
