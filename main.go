package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("This is reuqest %s", r.URL.Path)

	if r.URL.Path != "/hello" {
		http.Error(w, "404 Page Not Found ", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello I'm serving at 8080 \n")

}

func formhandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Post request successful \n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {

	fmt.Printf("Starting web server at port 8080 \n")

	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/form", formhandler)

	err := http.ListenAndServe(":8080", nil) //this method starts server and serves request at specofied port

	if err != nil {
		log.Fatal(err)
	}

}
