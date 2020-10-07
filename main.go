package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/julienschmidt/httprouter"
)

// DataStore is a key value store which has two methods:
// Get() - for fetching keys with value
// Set() - for setting keys with value
type DataStore interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type dbStore struct {
	db *badger.DB
}

var store DataStore

// Current returns the value of the current fibonacci sequence
func Current(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	index, err := store.Get("index")
	if err != nil {
		log.Panic(err)
	}
	value, err := store.Get(index)
	if err != nil {
		log.Panic(err)
	}

	io.WriteString(w, value)
}

// Next returns the value of the next number in the fibonacci sequence
func Next(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	index, err := store.Get("index")
	if err != nil {
		log.Panic(err)
	}
	nextIndex := strIncrement(index)

	nextValue, err := store.Get(nextIndex)

	// next value needs to be created because key was not found
	if err != nil {
		previousIndex := strDecrement(index)
		currentValue, _ := store.Get(index)
		previousValue, _ := store.Get(previousIndex)

		nextValue = addStr(currentValue, previousValue)
	}

	err = store.Set(nextIndex, nextValue)
	if err != nil {
		log.Panic(err)
	}

	err = store.Set("index", nextIndex)
	if err != nil {
		log.Panic(err)
	}

	io.WriteString(w, nextValue)
}

// Previous returns the value of the previous fibonacci sequence
func Previous(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	index, err := store.Get("index")
	if err != nil {
		log.Panic(err)
	}
	previousIndex := index

	if intIndex, err := strconv.Atoi(index); intIndex > 0 {
		previousIndex = strDecrement(index)
		if err != nil {
			log.Panic(err)
		}
	}

	previousValue, _ := store.Get(previousIndex)

	err = store.Set("index", previousIndex)
	if err != nil {
		log.Panic(err)
	}

	io.WriteString(w, previousValue)
}

func addStr(s string, s2 string) string {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	n2, err := strconv.Atoi(s2)
	if err != nil {
		log.Panic(err)
	}

	return strconv.Itoa(n + n2)
}

func strIncrement(s string) string {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}

	n++
	return strconv.Itoa(n)
}

// decrement can be smart and never decrement past 0 to handle base cases
func strDecrement(s string) string {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	n--
	return strconv.Itoa(n)
}

func recoverPanicHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Println(r.URL.Path, err)
	log.Println("Redirecting from panic to /current")
	http.Redirect(w, r, "/current", 301)
}

func main() {
	// Setup Routes
	router := httprouter.New()
	router.GET("/current/", Current)
	router.GET("/next", Next)
	router.GET("/previous", Previous)
	router.PanicHandler = recoverPanicHandler

	// Setup DB
	db, err := newDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.db.Close()

	err = db.configure()
	if err != nil {
		log.Fatal(err)
	}
	store = db

	log.Fatal(http.ListenAndServe(":8080", router))
}

func newDB() (*dbStore, error) {
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	return &dbStore{db: db}, err
}

func (store *dbStore) configure() error {
	return store.db.Update(func(tx *badger.Txn) error {
		// Seed DB if index key is not found
		if _, err := tx.Get([]byte("index")); err != nil {
			tx.Set([]byte("index"), []byte("0"))
			tx.Set([]byte("0"), []byte("0"))
			tx.Set([]byte("1"), []byte("1"))
			tx.Set([]byte("2"), []byte("1"))
		}

		return nil
	})
}

// Used for teardown in test development and test cases
func (store *dbStore) dropAll() error {
	return store.db.DropAll()
}

func (store *dbStore) Get(key string) (string, error) {
	var valCopy []byte

	err := store.db.View(func(tx *badger.Txn) error {
		item, err := tx.Get([]byte(key))
		if err != nil {
			return err
		}

		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	})

	return fmt.Sprintf("%s", valCopy), err
}

func (store *dbStore) Set(key string, value string) error {
	return store.db.Update(func(tx *badger.Txn) error {
		err := tx.Set([]byte(key), []byte(value))
		if err != nil {
			return fmt.Errorf("set key: %s", err)
		}

		return nil
	})
}
