package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prtf-server/internal/models"
)

type Question struct {
	Id int32 `json:"id"`
	Title string `json:"title"`
	Answer string `json:"answer"`
}

type Quiz struct {
	Title string `json:"title"`
	Questions []Question `json:"questions"`
}

func (app *Applicaiton) AddQuiz(w http.ResponseWriter, r *http.Request) {
	var quiz Quiz

	err := json.NewDecoder(r.Body).Decode(&quiz)
	if err != nil{
		app.errorLog.Println(err)
		return
	}

	q := models.Quiz{
		Title: quiz.Title,
	}
	id, err := app.DB.SaveQuiz(q)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	// app.infoLog.Println(id)

	
	for _, q :=range quiz.Questions {
		question := models.Question{
			Rf_quiz_id: id,
			Title: q.Title,
			Answer: q.Answer,
		}
		err := app.DB.SaveQuestion(question)
		if err != nil {
			app.errorLog.Println(err)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	// app.infoLog.Println("HERE")
	// app.infoLog.Println("HERE")
	// app.infoLog.Println("TITLE", quiz.Title)
}

func (app *Applicaiton) GetAllQuiz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
