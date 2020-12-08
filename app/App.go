package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	app.Router.Methods("GET").Path("/endpoint/{id}").HandlerFunc(app.getFunction)
	app.Router.Methods("POST").Path("/endpoint").HandlerFunc(app.postFunction)
}

func (app *App) getFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	fmt.Println("vars: ", vars)

	fmt.Println("id: ", id)
	fmt.Println("ok: ", ok)

	if !ok {
		fmt.Println("No ID in the path. Things are not OK")
	}

	dbData := &DbData{}
	err := app.Database.QueryRow("SELECT id, name FROM `test` WHERE id = ?", id).Scan(&dbData.ID, &dbData.Name)

	if err != nil {
		log.Fatal("No ID in the path", err)
	}
	log.Println("You fetched a thing!")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dbData); err != nil {
		panic(err)
	}
}

func (app *App) postFunction(w http.ResponseWriter, r *http.Request) {
	_, err := app.Database.Exec("INSERT INTO `test` (name) VALUES ('myname')")
	if err != nil {
		log.Fatal("Database INSERT failed", err)
	}

	log.Println("You called a thing")
	w.WriteHeader(http.StatusOK)
}
