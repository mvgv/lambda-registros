package mensagens

import (
	"github.com/mvgv/lambda-registros/app/infraestrutura/dto"
)

type Produtor interface {
	EnviarMensagem(mensagem *dto.SolicitacaoRelatorio) error
}
