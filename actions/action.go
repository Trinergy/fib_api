package actions

import (
	"github.com/Trinergy/fib_api/fibwithdb"
)

// Action represents a server action that can be taken
type Action struct {
	Store fibwithdb.FibWithDB
}
