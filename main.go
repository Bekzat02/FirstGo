package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name":"Bekzat"}`))
	ts,err:=template.ParseFiles("./ui/html/home.page.tmpl")
	if err!=nil{
		log.Println(err.Error())
		http.Error(w,"Internal error",500)
		return
	}
	err=ts.Execute(w,nil)
	if err!=nil{
		log.Println(err.Error())
		http.Error(w,"Internal error",500)
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Helllo with ID %d...", id)
}

func loremIpsum(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Lorem Ipsum is simply" +
		" dummy text of the printing and typesetting industry. "))
}
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		/*w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))*/

		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Write([]byte("Creating a snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/lorem", loremIpsum)
	mux.HandleFunc("/snippet/create", createSnippet)
	mux.HandleFunc("/snippet", showSnippet)
	log.Println("http:4000 is running")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// go mod init se07.com
