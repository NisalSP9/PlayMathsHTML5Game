package api

import (
	"net/http"
	"encoding/json"
	"github.com/nisalsp9/PlayMathsHTML5Game/models"
	"github.com/nisalsp9/PlayMathsHTML5Game/controller"
)

func CreateUser(w http.ResponseWriter, r *http.Request){

	println("dddddddddddddddddddddddddd")

	decoder := json.NewDecoder(r.Body)

	var user models.USER

	err := decoder.Decode(&user)

	if err != nil{

		println(err.Error())

	}

	controller.CreateUser(user)

}
