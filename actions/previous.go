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
		log.Panic(err)
	}
	previousIndex := index

	if intIndex, err := strconv.Atoi(index); intIndex > 0 {
		previousIndex = strDecrement(index)
		if err != nil {
			log.Panic(err)
		}
	}

	previousValue, _ := action.Store.Get(previousIndex)

	err = action.Store.Set("index", previousIndex)
	if err != nil {
		log.Panic(err)
	}

	io.WriteString(w, previousValue)
}
