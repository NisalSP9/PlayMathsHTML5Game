package main

import (
	"log"
	"net/http"
	"github.com/nisalsp9/PlayMathsHTML5Game/routes"
)

func main() {
	fs := http.FileServer(http.Dir("server/webapps/play_maths"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":3000", routes.UserRoutes())


}