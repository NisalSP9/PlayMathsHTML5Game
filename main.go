package main

import (
	"net/http"
	"github.com/nisalsp9/PlayMathsHTML5Game/routes"
	"log"
)

func main() {

	//Starting the API server
	router := routes.UserRoutes()
	http.Handle("/api/", router)

	//Starting the FileServer
	fs := http.FileServer(http.Dir("server/webapps/play_maths"))
	http.Handle("/", fs)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":3000", nil))

}


