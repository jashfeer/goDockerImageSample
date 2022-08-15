package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	
	tmp, _ := template.ParseFiles("index.gohtml")
	tmp.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Started..")
	http.ListenAndServe(":8080", nil)

}
