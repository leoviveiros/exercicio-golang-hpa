package main

import (
	"fmt"
	"net/http"
)

func main() {
	const greeting = "Code.education Rocks!"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		CpuStresser()
		fmt.Fprintf(w, greeting)
	})

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		panic(err)
	}
}