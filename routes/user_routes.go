package routes

import (
	"github.com/gorilla/mux"
	"github.com/nisalsp9/Play Maths HTML5 Game/api"
)

func UserRoutes()  {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user/create", api.CreateUser)



}


