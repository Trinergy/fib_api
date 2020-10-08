package actions

import (
	"fmt"
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
		s := fmt.Sprintf("500 Something bad happened: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(s))
		return
	}

	io.WriteString(w, value)
}
