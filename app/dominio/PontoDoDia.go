package dominio

type PontoDoDia struct {
	Email     string  `json:"email"`
	Registros []Ponto `json:"registros"`
}

func NewPontoDoDia(email string, registros []Ponto) *PontoDoDia {
	return &PontoDoDia{
		Email:     email,
		Registros: registros,
	}

}

func AdicionaPontoAoDia(pontoDoDia *PontoDoDia, ponto *Ponto) {
	if len(pontoDoDia.Registros) > 0 {
		ultimoPonto := pontoDoDia.Registros[len(pontoDoDia.Registros)-1]
		if ultimoPonto.Tipo == ponto.Tipo {
			return
		}
	}
	pontoDoDia.Registros = append(pontoDoDia.Registros, *ponto)
}
