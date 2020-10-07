package actions

import (
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Current returns the value of the current fibonacci sequence
func (action *Action) Current(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	index, err := action.Store.Get("index")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}
	value, err := action.Store.Get(index)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	io.WriteString(w, value)
}
