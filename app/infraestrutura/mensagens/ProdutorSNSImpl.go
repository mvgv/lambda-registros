package mensagens

import (
	"github.com/mvgv/lambda-registros/app/infraestrutura/dto"
)

type ProdutorSNSImpl struct{}

func NewProdutorSNSImpl() *ProdutorSNSImpl {
	return &ProdutorSNSImpl{}
}

func (p *ProdutorSNSImpl) EnviarMensagem(mensagem *dto.SolicitacaoRelatorio) error {
	return nil
}
