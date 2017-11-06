package routes

import (
	"github.com/gorilla/mux"
	"github.com/nisalsp9/PlayMathsHTML5Game/api"
)

func UserRoutes() *mux.Router  {
	var router = mux.NewRouter()
	router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user/create", api.CreateUser)

	return router

}


