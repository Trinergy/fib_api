package actions

import (
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Next returns the value of the next number in the fibonacci sequence
func (action *Action) Next(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	value, err := action.Store.Next()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	io.WriteString(w, value)
}
