package main

import (
	"log"
	"net/http"

	"github.com/Trinergy/fib_api/actions"
	"github.com/Trinergy/fib_api/datastore"
	"github.com/julienschmidt/httprouter"
)

func recoverPanicHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Println(r.URL.Path, err)
	log.Println("Redirecting from panic to /current")
	http.Redirect(w, r, "/current", 301)
}

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
	router.PanicHandler = recoverPanicHandler

	log.Fatal(http.ListenAndServe(":8080", router))
}
