package casodeuso

import (
	"fmt"
	"sort"
	"time"

	"github.com/mvgv/lambda-registros/app/dominio"
)

type CalcularHorasTrabalhadasImpl struct {
}

func NewCalcularHorasTrabalhadasImpl() *CalcularHorasTrabalhadasImpl {
	return &CalcularHorasTrabalhadasImpl{}
}

func (c *CalcularHorasTrabalhadasImpl) CalcularHorasTrabalhadasNoDia(listaPontos []dominio.Ponto) (string, error) {

	const layout = "2006-01-02T15:04:05"
	times := make([]time.Time, len(listaPontos))

	for i, ponto := range listaPontos {
		t, err := time.Parse(layout, ponto.Horario)
		if err != nil {
			return "", err
		}
		times[i] = t
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i].Before(times[j])
	})

	var total time.Duration
	for i := 0; i < len(times); i += 2 {
		if i+1 < len(times) {
			total += times[i+1].Sub(times[i])
		}
	}

	hours := int(total.Hours())
	minutes := int(total.Minutes()) % 60
	seconds := int(total.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds), nil
}
