package main

import (
	"log"
	"net/http"

	"github.com/Trinergy/fib_api/actions"
	"github.com/Trinergy/fib_api/datastore"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Setup DB
	db, err := datastore.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	err = db.Seed()
	if err != nil {
		log.Fatal(err)
	}

	// Setup Routes
	a := actions.Action{Store: db}

	router := httprouter.New()
	router.GET("/current/", a.Current)
	router.GET("/next", a.Next)
	router.GET("/previous", a.Previous)

	log.Fatal(http.ListenAndServe(":8080", router))
}
