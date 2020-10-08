package actions

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Trinergy/fib_api/badgerstore"
	"github.com/julienschmidt/httprouter"
)

func TestAction_Previous(t *testing.T) {
	db, err := badgerstore.NewDB("/tmp/badger/test")
	if err != nil {
		t.Error(err)
	}
	db.DropAll()
	db.Seed()
	defer db.DropAll()
	defer db.DB.Close()

	a := Action{Store: db}

	router := httprouter.New()
	router.GET("/previous", a.Previous)

	req, _ := http.NewRequest("GET", "/previous", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}

	resp := rr.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "0" {
		t.Errorf("Wrong response")
	}
}
