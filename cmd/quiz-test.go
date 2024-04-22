package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestQuizHandler(t *testing.T) {
	cfg := Config{
		port: 8001,
		env:  "development",
	}

	app := &Applicaiton{
		config:   cfg,
		infoLog:  log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile),
		version:  version,
	}

	req, err := http.NewRequest("GET", "/quiz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.GetAllQuiz)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello World"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
