package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"prtf-server/internal/driver"
	"prtf-server/internal/models"
	"time"
)

const version = "1.0.0"

type Config struct {
	port int
	env  string
	db string
}

type Applicaiton struct {
	config   Config
	infoLog  *log.Logger
	errorLog *log.Logger
	DB models.DBModel
	version  string
}

func (app *Applicaiton) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf(fmt.Sprintf("Starting HTTP server in %s mode on port %d", app.config.env, app.config.port))

	return srv.ListenAndServe()
}

func main() {
	cfg := Config{
		port: 8001,
		env:  "development",
		db: "postgresql://postgres:postgres@localhost:5432/prtfdb?sslmode=disable",
	}

	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := driver.OpenDB(cfg.db)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	app := &Applicaiton{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		DB: models.DBModel{DB: conn},
		version:  version,
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
