package fibwithdb

import (
	"testing"
)

// This is an integration test that will overall functionality of Database:
// NewDB - initialization
// Seed - seeding data
// Get - fetching value at key
// Set - setting value at key
func TestFibWithDB_Integration(t *testing.T) {
	db, err := NewDB("/tmp/badger/test")
	if err != nil {
		t.Error(err)
	}
	db.DropAll()
	db.Seed()
	defer db.DropAll()
	defer db.DB.Close()

	// Test Seed + Get
	v, err := db.Get("index")
	if err != nil {
		t.Error(err)
	}
	if v != "0" {
		t.Error("Seed failed")
	}

	// Test Set
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

func TestFibWithDB_Current(t *testing.T) {
	db, err := NewDB("/tmp/badger/test")
	if err != nil {
		t.Error(err)
	}
	db.DropAll()
	db.Seed()
	defer db.DropAll()
	defer db.DB.Close()

	value, err := db.Current()
	if err != nil {
		t.Error(err)
	}
	if value != "0" {
		t.Error("Current is broken")
	}
}

func TestFibWithDB_Next(t *testing.T) {
	db, err := NewDB("/tmp/badger/test")
	if err != nil {
		t.Error(err)
	}
	db.DropAll()
	db.Seed()
	defer db.DropAll()
	defer db.DB.Close()

	value, err := db.Next()
	if err != nil {
		t.Error(err)
	}
	if value != "1" {
		t.Error("Next is broken")
	}
}

func TestFibWithDB_Previous(t *testing.T) {
	db, err := NewDB("/tmp/badger/test")
	if err != nil {
		t.Error(err)
	}
	db.DropAll()
	db.Seed()
	defer db.DropAll()
	defer db.DB.Close()

	value, err := db.Previous()
	if err != nil {
		t.Error(err)
	}
	if value != "0" {
		t.Error("Previous is broken")
	}
}

// This tests the business logic of the fibonacci sequence
func TestFibWithDB_FibLogic(t *testing.T) {
	db, err := NewDB("/tmp/badger/test")
	if err != nil {
		t.Error(err)
	}
	db.DropAll()
	db.Seed()
	defer db.DropAll()
	defer db.DB.Close()

	_, _ = db.Next()
	_, _ = db.Next()
	value, _ := db.Next()

	if value != "2" {
		t.Error("FibLogic is broken - Next")
	}

	value, _ = db.Previous()
	if value != "1" {
		t.Error("FibLogic is broken - Previous")
	}
}
