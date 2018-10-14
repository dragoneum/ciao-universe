package main

import (
	"net/http"
	"fmt"
)

func main()  {
	http.HandleFunc("/ciao", ciao)
	http.ListenAndServe(":8080", nil)
}

func ciao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ciao universe!")
}