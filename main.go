package main

import (
	"log"
	"net/http"
	"github.com/nisalsp9/Play Maths HTML5 Game/routes"
)

func main() {
	fs := http.FileServer(http.Dir("server/webapps/play_maths"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
	routes.UserRoutes()

}