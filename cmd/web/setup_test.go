package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
	"webapp/pkg/db"
)

var app application

func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates/"

	app.Session = getSession()
	err := godotenv.Load("../../.test_env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env := os.Getenv("GO_ENV")
	app.DSN = env

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = db.PostgresConn{DB: conn}

	os.Exit(m.Run())
}
