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

const (
	dbName = "my.db"
	bucket = "fibonnaciBucket"
)

// DataStore is a key value store which has two methods:
// Get() - for fetching keys with value
// Set() - for setting keys with value
type DataStore interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type dbStore struct {
	bucket string
	db     *bolt.DB
}

var store DataStore

// Current returns the value of the current fibonacci sequence
func Current(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	idx, _ := store.Get("index")
	value, _ := store.Get(idx)

	io.WriteString(w, value)
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
	router.GET("/next", Next)
	router.GET("/previous", Previous)

	// Setup DB
	db := newDB()
	err := db.configure()
	if err != nil {
		log.Fatal(err)
	}
	store = db

	log.Fatal(http.ListenAndServe(":8080", router))
}

func newDB() *dbStore {
	db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return &dbStore{bucket: bucket, db: db}
}

func (store *dbStore) configure() error {
	return store.db.Update(func(tx *bolt.Tx) error {
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
}

func (store *dbStore) Get(key string) (string, error) {
	var returnValue []byte

	err := store.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(store.bucket))
		v := b.Get([]byte(key))

		returnValue = make([]byte, len(v))
		copy(returnValue, v)

		fmt.Printf("GET CALL: %s", returnValue)

		return nil
	})

	return fmt.Sprintf("%s", returnValue), err
}

func (store *dbStore) Set(key string, value string) error {
	return store.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(store.bucket))
		err := b.Put([]byte(key), []byte(value))
		if err != nil {
			return fmt.Errorf("set key: %s", err)
		}

		return nil
	})
}