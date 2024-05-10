package db

import (
	"os"

	"github.com/tranquilitodev/gambituser/models"
	"github.com/tranquilitodev/gambituser/secretm"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}
