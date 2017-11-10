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

const GET_ALL_USERS  = `SELECT * FROM playmaths.pm_users`

func GetAllUsers() ([]models.USER,error)  {

	println("sssssssss")

	var users []models.USER

	var username,password string
	var id,level int

	db := DBConnection.GetConnection()

	rst,err:=db.Query(GET_ALL_USERS)

	if err != nil {
		println(err.Error())
	}

	for(rst.Next()){

		err =rst.Scan(&id,&username,&password,&level)

		users = append(users, models.USER{UserID: id, Username: username,Password:password,Level:level})

	}



	if err != nil {
		println(err.Error())
	}



	return users,err
}