package main

import (

	"net/http"
	"github.com/nisalsp9/PlayMathsHTML5Game/routes"
	"log"
)

func main() {

	//myMux := http.NewServeMux()
	//myMux.HandleFunc("/", voidHandler)
	//
	//server := &http.Server{
	//	Addr:    ":3000",
	//	Handler: myMux,
	//}
	//
	//go server.ListenAndServe()
	//
	//gorillaMux := http.NewServeMux()
	//gorillaMux.HandleFunc("/", voidHandler2)
	//
	//server2 := &http.Server{
	//	Addr:    ":3001",
	//	Handler: gorillaMux,
	//}
	//
	//server2.ListenAndServe()





























	fs := http.FileServer(http.Dir("server/webapps/play_maths"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":3000", routes.UserRoutes())


}

//func voidHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Println(r.URL.Path)
//	//fs := http.FileServer(http.Dir("server/webapps/play_maths"))
//	http.ServeFile(w, r, r.URL.Path[2:])
//}
//
//func voidHandler2(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "server 2")
//}