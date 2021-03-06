package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type pageData struct {
	Title     string
	FirstName string
}

var handlers = []struct{
	pattern string
	msg string
}{
	{"/", "Default"},
}

type stringHandler string

func (s stringHandler) ServeHTTP(w http.ResponseWriter, req *http.Request){
	log.Println("LOGGED", req.RequestURI)
	fmt.Println(req.RequestURI)
	w.Header().Set("Result", string(s))
	w.Write([]byte(s))
}

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	mux := http.NewServeMux()
	for _, h := range handlers {
		mux.Handle(h.pattern, stringHandler(h.msg))
	}
	//mux.HandleFunc("/", index)
	http.ListenAndServe(":8080", mux)
	//http.HandleFunc("/", index)
	//http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, req *http.Request) {
	pd := pageData{
		Title: "Index Page",
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", pd)

	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal serverrrrrr error", http.StatusInternalServerError)
		return
	}


	log.Println("LOGGED", req.RequestURI)
}
