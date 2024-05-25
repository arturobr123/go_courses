package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	} else {
		fmt.Fprint(w, "POST request successful \n")
		name := r.FormValue("name")
		address := r.FormValue("address")

		fmt.Fprint(w, "Name: ", name, "\n")
		fmt.Fprint(w, "Address: ", address, "\n")

		fmt.Fprint(w, "Lets gooo!!!")
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Hello!!!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 8080 \n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

/// run ====>  go run main.go

// COURSE https://www.youtube.com/watch?v=ASBUp7stqjo&list=PLQeMe1h4vVmPjVtcnsnVRHJamW7zw3cse
