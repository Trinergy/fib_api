package datastore

import (
	"testing"
)

// This is an integration test that will test NewDB, Seed, Get, and Set
func TestDatastore_Integration(t *testing.T) {
	db, err := NewDB("/tmp/badger/test")
	if err != nil {
		t.Error(err)
	}
	db.Seed()
	defer db.DropAll()
	defer db.DB.Close()

	// Test Get
	v, err := db.Get("index")
	if err != nil {
		t.Error(err)
	}
	if v != "0" {
		t.Error("Seed failed")
	}

	err = db.Set("index", "1")
	if err != nil {
		t.Error(err)
	}

	setValue, err := db.Get("index")
	if err != nil {
		t.Error(err)
	}
	if setValue != "1" {
		t.Error("Set failed")
	}
}
