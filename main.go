package main

import (
	"fmt"
	"log"
	"myapp/app"
	"myapp/db"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.SetupRouter()
	fmt.Println("finished setting up router")
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
