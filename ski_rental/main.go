package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("vim-go")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bananas hello!")
	})
	http.ListenAndServe(":8888", nil)
}
