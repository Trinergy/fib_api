package datastore

import (
	"fmt"

	badger "github.com/dgraph-io/badger/v2"
)

// DB represents the database to be used for the datastore
type DB struct {
	DB *badger.DB
}

// NewDB creates a brand new badger db
func NewDB() (*DB, error) {
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
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

// Used for teardown in test development and test cases
func (store *DB) dropAll() error {
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
