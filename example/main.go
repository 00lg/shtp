package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("New request! ")
		fmt.Fprint(w, "Hello World!")
	})
	fmt.Printf("Starging server...")
	http.ListenAndServe(":8081", nil)
}
