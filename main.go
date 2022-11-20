package main

import (
	"fmt"
	"log"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/greet" {
		http.Error(w, "404 - NOT FOUND", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Ahoy there!!!")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/form" {
		http.Error(w, "404 - NOT FOUND", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method not supported", http.StatusBadRequest)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "Form parsing error: %v", err)
		return
	}
	fmt.Fprint(w, "POST request reach to server\n")
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username =="admin" && password == "admin" {
		fmt.Fprint(w, "You have successfully logged in")
		return
	} else {
		fmt.Fprint(w, "Username or password is wrong")
		return
	}
}

func main()  {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/greet", greetHandler)
	
	//Start go server
	fmt.Printf("Starting Go server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
