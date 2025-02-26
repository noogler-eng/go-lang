package main

import (
	"fmt"  		// for simply printing  
	"log"  		// for console.log ...
	"net/http" 	// for server and routing
)

// what we are building ?
// server 
// 	"/" 		index.html
// 	"/hello" 	hello func
// 	"/form"     form func -> form.html 

// r is pointer basically pointing to the incomming request
// w is respose writer
func formHandler(w http.ResponseWriter, r *http.Request){
	// Fprintf used to write into the frontend
	// ParseForm, getting the data from frontend
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parse form error %v", err);
		return;
	}

	fmt.Fprintf(w, "post request succesfull\n");
	email := r.FormValue("email");
	password := r.FormValue("password");
	fmt.Fprintf(w, "email := %s\n", email);
	fmt.Fprintf(w, "password := %s\n", password);
}


func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found!", http.StatusNotFound)
		return;
	}

	if r.Method != "GET" {
		http.Error(w, "Method not found!", http.StatusNotFound)
		return;
	}

	fmt.Fprintf(w, "hello world!!!");
}

func main(){
	// directly assignment can be possibgle using :=
	// creating an http server
	fileServer := http.FileServer(http.Dir("./static"));
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler);
	http.HandleFunc("/hello", helloHandler);

	
	// listening the server
	fmt.Println("server started at port: ", 8080);
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}