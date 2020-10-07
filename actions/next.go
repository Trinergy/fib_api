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
		log.Panic(err)
	}
	nextIndex := strIncrement(index)

	nextValue, err := action.Store.Get(nextIndex)

	// next value needs to be created because key was not found
	if err != nil {
		previousIndex := strDecrement(index)
		currentValue, _ := action.Store.Get(index)
		previousValue, _ := action.Store.Get(previousIndex)

		nextValue = addStr(currentValue, previousValue)
	}

	err = action.Store.Set(nextIndex, nextValue)
	if err != nil {
		log.Panic(err)
	}

	err = action.Store.Set("index", nextIndex)
	if err != nil {
		log.Panic(err)
	}

	io.WriteString(w, nextValue)
}
