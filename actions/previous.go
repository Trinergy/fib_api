package actions

import (
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Previous returns the value of the previous fibonacci sequence
func (action *Action) Previous(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	value, err := action.Store.Previous()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	io.WriteString(w, value)
}
