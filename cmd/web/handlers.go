package main

import(
	"net/http"
	"strconv"
	"html/template"
	"log"
)

func home (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"{
		http.NotFound(w,r)
		return	
	}
	
	files:= []string{
		"./ui/html/pages/base.html",
		"./ui/html/pages/home.html",
		"./ui/html/pages/nav.html",
	}

	ts, err:= template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil { 
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func snippetCreate(w http.ResponseWriter, r *http.Request) { 
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Unauthorized request...", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a snippet..."))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err:= strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 { 
		http.NotFound(w,r)
		return
	}
	w.Write([]byte("View all snippets..."))
}