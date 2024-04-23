package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Question struct {
	Id int32 `json:"id"`
	Question string `json:"question"`
	Answer string `json:"answer"`
}

type Quiz struct {
	Questions []Question `json:"questions"`
}

func (app *Applicaiton) AddQuiz(w http.ResponseWriter, r *http.Request) {
	var quiz Quiz

	err := json.NewDecoder(r.Body).Decode(&quiz)
	if err != nil{
		app.errorLog.Println(err)
		return
	}

	app.infoLog.Println("HERE")
	for _, q :=range quiz.Questions {
		app.infoLog.Println(q.Answer)
	}
	app.infoLog.Println("HERE")
}

func (app *Applicaiton) GetAllQuiz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
