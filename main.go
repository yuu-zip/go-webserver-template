package go-webserver-template

import (
	"fmt"
	"net/http"
	"text/template"
)

type Page struct {
	title string
}

func Handler(w http.ResponseWriter, r *http.Request){
	page := Page{title"Test replace"}
	// load a HTML file from the directory directly below -> "../"
	// warning: The file extension must .html
	tem, err := template.ParseFiles("index.html")
	if err != nil { panic(err) }
	err = tem.Execute(w, page)
	if err != nil { panic(err) }
}

func main(){
	http.HandleFunc("/", Handler)
	http.ListenAndServe("localhost:8080", nil)
}