package repositorio

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mvgv/lambda-registros/app/infraestrutura/dto"
)

type PontoRepositorioDynamoImpl struct {
	svc *dynamodb.DynamoDB
}

func NewPontoRepositorioDynamoImpl() *PontoRepositorioDynamoImpl {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return &PontoRepositorioDynamoImpl{svc: dynamodb.New(sess)}
}

func (p *PontoRepositorioDynamoImpl) RegistrarPonto(email string, timestamp string, evento string) error {
	pontoEntidade := dto.NewPontoEntidade(email, timestamp, evento)
	av, err := dynamodbattribute.MarshalMap(pontoEntidade)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("RegistrosPonto"),
	}

	_, err = p.svc.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}

func (p *PontoRepositorioDynamoImpl) ConsultarPontoDoDia(email string) ([][]string, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String("RegistrosPonto"),
	}
	result, err := p.svc.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}
	pontoDodia := &dto.PontoDoDiaEntidade{}
	err = dynamodbattribute.UnmarshalMap(result.Item, pontoDodia)
	if err != nil {
		return nil, err
	}

	registros := make([][]string, len(pontoDodia.Registros))
	for i, registro := range pontoDodia.Registros {
		registros[i] = []string{registro.Timestamp, registro.Evento}
	}

	return registros, nil

}
