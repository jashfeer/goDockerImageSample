package main

import (
	"html/template"
	"net/http"
)


func index(res http.ResponseWriter, req *http.Request) {
	tmp, _ := template.ParseFiles("index.gohtml")
	tmp.Execute(res, nil)}

func main(){
	http.HandleFunc("/", index)

	http.ListenAndServe(":9090",nil)
	
}