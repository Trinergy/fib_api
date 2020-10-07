package actions

import (
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Next returns the value of the next number in the fibonacci sequence
func (action *Action) Next(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	index, err := action.Store.Get("index")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}
	nextIndex, err := strIncrement(index)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	nextValue, err := action.Store.Get(nextIndex)
	// next value needs to be created because key was not found
	if err != nil {
		previousIndex, err := strDecrement(index)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
			return
		}
		currentValue, err := action.Store.Get(index)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
			return
		}
		previousValue, err := action.Store.Get(previousIndex)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
			return
		}

		nextValue, err = addStr(currentValue, previousValue)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
			return
		}
	}

	err = action.Store.Set(nextIndex, nextValue)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	err = action.Store.Set("index", nextIndex)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	io.WriteString(w, nextValue)
}
