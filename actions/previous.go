package actions

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Previous returns the value of the previous fibonacci sequence
func (action *Action) Previous(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	index, err := action.Store.Get("index")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}
	previousIndex := index

	if intIndex, err := strconv.Atoi(index); intIndex > 0 {
		previousIndex, err = strDecrement(index)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
			return
		}
	}

	previousValue, err := action.Store.Get(previousIndex)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	err = action.Store.Set("index", previousIndex)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	io.WriteString(w, previousValue)
}
