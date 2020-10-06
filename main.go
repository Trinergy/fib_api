package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/julienschmidt/httprouter"
)

// DB is a global variable for all handlers to access
var DB *bolt.DB

const (
	dbName = "my.db"
	bucket = "fibonnaciBucket"
)

// Current returns the value of the current fibonacci sequence
func Current(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		idx := b.Get([]byte("index"))
		v := b.Get([]byte(idx))

		message := fmt.Sprintf("%s", v)
		io.WriteString(w, message)
		return nil
	})
}

// Next returns the value of the next number in the fibonacci sequence
func Next(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

// Previous returns the value of the previous fibonacci sequence
func Previous(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
}

func main() {
	// Setup Routes
	router := httprouter.New()
	router.GET("/current/", Current)
	router.PUT("/next", Next)
	router.GET("/previous", Previous)

	// Setup DB
	db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b.Put([]byte("index"), []byte("0"))
		b.Put([]byte("0"), []byte("0"))
		b.Put([]byte("1"), []byte("1"))
		b.Put([]byte("2"), []byte("1"))

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	log.Fatal(http.ListenAndServe(":8080", router))
}
