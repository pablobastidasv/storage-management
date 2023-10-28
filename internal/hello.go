package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, _ *http.Request) {
	redirect := `
	<!DOCTYPE html>
	<html>
	   <head>
		  <title>Bastriguez</title>
		  <meta http-equiv = "refresh" content = "0; url = statics/" />
	   </head>
	   <body>
		  <p>Redirecting</p>
	   </body>
	</html>
	`
	_, err := fmt.Fprint(w, redirect)
	if err != nil {
		log.Print(err.Error())
	}
}

func RunServer(addr string) {
	r := mux.NewRouter()

	r.HandleFunc("/", index)

	http.Handle("/", r)

	// serving static files
	fs := http.FileServer(http.Dir("statics"))
	http.Handle("/statics/", http.StripPrefix("/statics/", fs))

	log.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(addr, nil))
}
