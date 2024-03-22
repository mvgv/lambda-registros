package dto

type PontoEntidade struct {
	Email     string `json:"email"`
	Timestamp string `json:"timestamp"`
	Evento    string `json:"evento"`
}

func NewPontoEntidade(email, timestamp, evento string) *PontoEntidade {
	return &PontoEntidade{
		Email:     email,
		Timestamp: timestamp,
		Evento:    evento,
	}
}
