package main

import (
	"log"
	"net/http"
	"flag"
)

func main() {
	addr:= flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()


	router:= http.NewServeMux()
	fileServer:= http.FileServer(http.Dir("./ui/static/"))

	router.Handle("/static/", http.StripPrefix("/static", fileServer))
	router.HandleFunc("/",home)
	router.HandleFunc("/snippet/view", snippetView)
	router.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on localhost:%s", *addr)
	err:= http.ListenAndServe(*addr, router)
	log.Fatal(err)

}
