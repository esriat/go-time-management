package main

import (
	"log"
	"net/http"

	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/datastores"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/handlers"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/schema"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	datastore datastores.IDatastore
	e         handlers.Env
)

func main() {
	var err error

	globals.Init()

	globals.Log.Info("Creating database")
	if datastore, err = datastores.NewDatabase("myDatabase.db"); err != nil {
		log.Fatal(err)
	}

	e = handlers.Env{
		DB: datastore,
	}

	globals.Log.Info("Creating the routes")

	r := mux.NewRouter()

	handlers.HandleRoutes(r, &e)

	http.ListenAndServe(":8080", r)
}
