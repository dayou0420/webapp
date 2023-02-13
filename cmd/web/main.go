package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"webapp/pkg/db"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	DSN     string
	DB      db.PostgresConn
	Session *scs.SessionManager
}

func main() {
	// set up an app config
	app := application{}

	// get "github.com/joho/godotenv"
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env := os.Getenv("GO_ENV")
	flag.StringVar(&app.DSN, "dsn", env, "Posgtres connection")

	//flag.StringVar(&app.DSN, "dsn", "host=postgres port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Posgtres connection")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = db.PostgresConn{DB: conn}

	// get a session manager
	app.Session = getSession()

	// print out a message
	log.Println("Starting server on port 8080...")

	// start the server
	err = http.ListenAndServe(":8080", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
