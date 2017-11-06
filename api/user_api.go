package api

import (
	"net/http"
	"encoding/json"
	"github.com/nisalsp9/Play Maths HTML5 Game/models"
	"github.com/nisalsp9/Play Maths HTML5 Game/controller"
)

func CreateUser(w http.ResponseWriter, r *http.Request) error  {

	println("dddddddddddddddddddddddddd")

	decoder := json.NewDecoder(r.Body)

	var user models.USER

	err := decoder.Decode(&user)

	if err != nil{

		println(err.Error())

	}

	controller.CreateUser(user)




	return nil
}
