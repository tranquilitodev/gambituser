package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tranquilitodev/gambituser/models"
	"github.com/tranquilitodev/gambituser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comenza Registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email , User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechamySQL() + "')"
	fmt.Println(sentencia)
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("SignUp > Ejecución Exitosa")
	return nil
}
