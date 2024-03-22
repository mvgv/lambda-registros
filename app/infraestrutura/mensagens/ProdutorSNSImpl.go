package mensagens

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/mvgv/lambda-registros/app/infraestrutura/dto"
)

type ProdutorSNSImpl struct {
	snsClient *sns.SNS
	TopicArn  string
}

func NewProdutorSNSImpl() *ProdutorSNSImpl {
	sess := session.Must(session.NewSession())
	snsClient := sns.New(sess)
	return &ProdutorSNSImpl{snsClient: snsClient,
		TopicArn: "arn:aws:sns:us-east-1:101478099523:solicitar-relatorio"}
}

func (p *ProdutorSNSImpl) EnviarMensagem(mensagem *dto.SolicitacaoRelatorio) error {
	msg, err := json.Marshal(mensagem)
	if err != nil {
		return err
	}

	_, err = p.snsClient.Publish(&sns.PublishInput{
		Message:  aws.String(string(msg)),
		TopicArn: aws.String(p.TopicArn),
	})

	return err
}
