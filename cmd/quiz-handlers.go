package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Quiz struct {
	Id int32 `json:"id"`
	Question string `json:"question"`
	Answer string `json:"answer"`
}

func (app *Applicaiton) AddQuiz(w http.ResponseWriter, r *http.Request) {
	var quizes = []Quiz{}

	err := json.NewDecoder(r.Body).Decode(&quizes)
	if err != nil{
		app.errorLog.Println(err)
		return
	}

	for quiz :=range quizes {
		fmt.Println(quiz)
	}
}

func (app *Applicaiton) GetAllQuiz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
