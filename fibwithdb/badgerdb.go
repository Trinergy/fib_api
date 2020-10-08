package fibwithdb

import (
	"errors"
	"fmt"
	"strconv"

	badger "github.com/dgraph-io/badger/v2"
)

// DB represents the database to be used for the datastore
type DB struct {
	DB *badger.DB
}

// NewDB creates a brand new badger db
func NewDB(path string) (*DB, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	return &DB{DB: db}, err
}

// Seed will only create seed records if they do not exist yet
func (store *DB) Seed() error {
	return store.DB.Update(func(tx *badger.Txn) error {
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

// DropAll is used for teardown in development and test cases
func (store *DB) DropAll() error {
	return store.DB.DropAll()
}

// Get returns the value at given key
func (store *DB) Get(key string) (string, error) {
	var valCopy []byte

	err := store.DB.View(func(tx *badger.Txn) error {
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

// Set will create or update a key with a value
func (store *DB) Set(key string, value string) error {
	return store.DB.Update(func(tx *badger.Txn) error {
		err := tx.Set([]byte(key), []byte(value))
		if err != nil {
			return fmt.Errorf("set key: %s", err)
		}

		return nil
	})
}

// Current returns the value of fibonacci sequence at current index
func (store *DB) Current() (string, error) {
	index, err := store.Get("index")
	if err != nil {
		return "", err
	}
	value, err := store.Get(index)
	if err != nil {
		return "", err
	}

	return value, nil
}

// Next returns the value of the next step in the fibonacci sequence and increments the index by one
func (store *DB) Next() (string, error) {
	index, err := store.Get("index")
	if err != nil {
		return "", err
	}
	nextIndex, err := strIncrement(index)
	if err != nil {
		return "", err
	}

	nextValue, err := store.Get(nextIndex)
	// next value needs to be created because key was not found
	if err != nil {
		previousIndex, err := strDecrement(index)
		if err != nil {
			return "", err
		}
		currentValue, err := store.Get(index)
		if err != nil {
			return "", err
		}
		previousValue, err := store.Get(previousIndex)
		if err != nil {
			return "", err
		}

		nextValue, err = addStr(currentValue, previousValue)
		if err != nil {
			return "", err
		}

		// Check for integer overflow
		num, err := strconv.Atoi(nextValue)
		if err != nil {
			return "", err
		}
		if num < 0 {
			return "", errors.New("One does not simply fib - integer overflow")
		}
	}

	err = store.Set(nextIndex, nextValue)
	if err != nil {
		return "", err
	}

	err = store.Set("index", nextIndex)
	if err != nil {
		return "", err
	}

	return nextValue, nil
}

// Previous returns the values of the previous sequence and decrements the index
func (store *DB) Previous() (string, error) {
	index, err := store.Get("index")
	if err != nil {
		return "", err
	}
	previousIndex := index

	if intIndex, err := strconv.Atoi(index); intIndex > 0 {
		previousIndex, err = strDecrement(index)
		if err != nil {
			return "", err
		}
	}

	previousValue, err := store.Get(previousIndex)
	if err != nil {
		return "", err
	}

	err = store.Set("index", previousIndex)
	if err != nil {
		return "", err
	}

	return previousValue, nil
}
