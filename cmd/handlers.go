package main

import "net/http"

func (app *applicaiton) quiz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}
