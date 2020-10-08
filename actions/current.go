package actions

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Current returns the value of the current fibonacci sequence
func (action *Action) Current(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	value, err := action.Store.Current()
	if err != nil {
		log.Println(err)
		s := fmt.Sprintf("500 Something bad happened: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(s))
		return
	}

	io.WriteString(w, value)
}
