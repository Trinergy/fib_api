package actions

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Trinergy/fib_api/datastore"
	"github.com/julienschmidt/httprouter"
)

func TestAction_Current(t *testing.T) {
	db, err := datastore.NewDB("/tmp/badger/test")
	if err != nil {
		t.Error(err)
	}
	db.DropAll()
	db.Seed()
	defer db.DB.Close()

	a := Action{Store: db}

	router := httprouter.New()
	router.GET("/current", a.Current)

	req, _ := http.NewRequest("GET", "/current", nil)
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
