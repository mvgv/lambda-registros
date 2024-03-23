package apresentacao

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type ValidaToken struct {
}

func NewValidaToken() *ValidaToken {
	return &ValidaToken{}
}

func (uc *ValidaToken) AutorizarCliente(tokenString string, reqUser string) (bool, string) {

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})

	if err != nil {
		fmt.Println("Erro ao fazer parse do token: ", err)
		return false, ""
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	if claims["iss"] != "hackathoncompany.com.br" {
		// Se os requisitos não forem atendidos, bloqueie a requisição]
		fmt.Println("Bloqueando requisição")
		return false, ""

	}

	if claims["user"] != reqUser {
		fmt.Println("Usuário não autorizado")
		return false, ""
	}

	return true, claims["user"].(string)
}
