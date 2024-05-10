package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/tranquilitodev/gambituser/awsgo"
	"github.com/tranquilitodev/gambituser/db"
	"github.com/tranquilitodev/gambituser/models"
)

func main() {
	lambda.Start(EjecutoLambda)
}
func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializoAWS()
	if !ValidoParametros() {
		fmt.Println("Error en los parametros.debe enviar 'SecretManager'")
		err := errors.New("Error en los parametros debe enviar SecretName")
		return event, err

	}
	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)
		}
	}
	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el Secret " + err.Error())
		return event, err

	}

}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
