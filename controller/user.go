package controller

import (
	"github.com/nisalsp9/PlayMathsHTML5Game/DBConnection"
	"github.com/nisalsp9/PlayMathsHTML5Game/models"
)

const CREATE_USER  = `INSERT INTO pm_users (username, password, level) VALUES (?,?,?)`

func CreateUser(user models.USER) error {

	db := DBConnection.GetConnection()

	Stmt,err := db.Prepare(CREATE_USER)

	if err != nil{
		println(err.Error())
	}

	defer Stmt.Close()

	_,err=Stmt.Exec(
		user.Username,
		user.Password,
		user.Level,

	)

	if err != nil{

		println(err.Error())

	}

	return nil
}
