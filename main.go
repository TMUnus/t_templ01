package main

import (
	"fmt"
	"net/http"
)

func main() {
	//fmt.Println("Hello World")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Hello("Beast Titan").Render(r.Context(), w)
	})

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3001", nil)
}
