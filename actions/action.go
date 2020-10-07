package actions

import (
	"github.com/Trinergy/fib_api/datastore"
)

// Action represents a server action that can be taken
type Action struct {
	Store datastore.DataStore
}
