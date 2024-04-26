package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBModel struct {
	DB *pgxpool.Pool
}

type Quiz struct {
	ID int `json:"id,omitempty"`
	Title string `json:"title"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Question struct {
	ID int `json:"id,omitempty"`
	Rf_quiz_id int `json:"rf_quiz_id,omitempty"`
	Title string `json:"title"`
	Answer string `json:"answer"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (m *DBModel) SaveQuiz(quiz Quiz) (int, error) {
	stmt := `
		insert into public.quiz 
			(title)
		values($1)
		returning id
	`

	var id int
	err := m.DB.QueryRow(context.Background(), stmt, quiz.Title).Scan(&id)
	if(err != nil) {
		return 0, err
	}

	return id, nil
}

func (m *DBModel) SaveQuestion(question Question) error {
	stmt := `
		insert into public.question 
			(rf_quiz_id, title, answer)
		values($1, $2, $3)
	`

	_, err := m.DB.Exec(context.Background(), stmt, question.Rf_quiz_id, question.Title, question.Answer)
	if(err != nil) {
		return err
	}

	return err
}