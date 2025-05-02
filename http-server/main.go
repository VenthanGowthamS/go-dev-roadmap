package main

//provides HTTP client and server functionality. we need this to build web servers or clients.

import (
	"fmt"
	"net/http"
)

//w http.ResponseWriter – used to send a response back to the client (e.g., browser).

func handler(w http.ResponseWriter, r *http.Request) {
	//print method (get post etc)
	fmt.Println("Request method:", r.Method)

	//print path

	fmt.Println("request path", r.URL.Path)

	//send response back to browser

	fmt.Fprintln(w, "Thanks for checking venthan awesome work")

}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("🌐 Started web server at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
