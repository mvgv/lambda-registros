package casodeuso

type GerarRelatorio interface {
	GerarRelatorioMensal(email, mes string) error
}
