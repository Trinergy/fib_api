package actions

/*
Store returns values in the fibonacci sequence and keeps track of the step in the sequence
Current returns the value of current sequence
Next returns the value of next sequence and increments the index
Previous returns the values of the previous sequence and decrements the index
*/
type Store interface {
	Current() (string, error)
	Next() (string, error)
	Previous() (string, error)
}

// Action represents a server action that can be taken
type Action struct {
	Store Store
}
