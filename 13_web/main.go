package main

import (
	"fmt" //Format Package
	"net/http"
)

func index(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w,"<h1>Hello World</h1>")
}
func about(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w,"<h1>About us</h1>")
}
func main() {
	//route 
	http.HandleFunc("/" , index)
	http.HandleFunc("/about" , about)

	fmt.Println("Server is starting.....")
	http.ListenAndServe(":3000" , nil)
}