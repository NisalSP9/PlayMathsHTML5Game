package api

import (
	"net/http"
	"encoding/json"
	"github.com/nisalsp9/PlayMathsHTML5Game/models"
	"github.com/nisalsp9/PlayMathsHTML5Game/controller"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var user models.USER
	err := decoder.Decode(&user)
	if err != nil{
		println(err.Error())
	}
	controller.CreateUser(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request)  {

	users,err := controller.GetAllUsers()

	if err != nil {

		println(err.Error())

	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(users);

	if err != nil {
		println(err.Error())
	}





}
