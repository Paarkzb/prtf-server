package main

import (
	"fmt"
	"net/http"
)

func (app *Applicaiton) Quiz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
