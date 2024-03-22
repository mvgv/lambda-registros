package repositorio

import (
	"fmt"
	"time"

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

func (p *PontoRepositorioDynamoImpl) ConsultarPontoDoDia(email string) (*dto.PontoDoDiaEntidade, error) {
	hojeInicio := time.Now().Format("2006-01-02") + "T00:00:00"
	hojeFim := time.Now().Format("2006-01-02") + "T23:59:59"

	input := &dynamodb.QueryInput{
		TableName: aws.String("RegistrosPonto"),
		KeyConditions: map[string]*dynamodb.Condition{
			"email": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(email),
					},
				},
			},
			"timestamp": {
				ComparisonOperator: aws.String("BETWEEN"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(hojeInicio),
					},
					{
						S: aws.String(hojeFim),
					},
				},
			},
		},
	}

	result, err := p.svc.Query(input)
	if err != nil {
		return nil, err
	}

	registros := make([]dto.PontoEntidade, len(result.Items))

	for i, item := range result.Items {
		registros[i] = *dto.NewPontoEntidade(*item["email"].S, *item["timestamp"].S, *item["evento"].S)
	}

	pontoDodia := dto.NewPontoDoDiaEntidade(*result.Items[0]["email"].S, registros)
	fmt.Println("Ponto do dia: ", pontoDodia)
	return pontoDodia, nil

}
