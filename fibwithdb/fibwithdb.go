package fibwithdb

// FibWithDB returns values in the fibonacci sequence
// Current returns the value of current sequence
// Next returns the value of next sequence and increments the index
// Previous returns the values of the previous sequence and decrements the index
type FibWithDB interface {
	Current() (string, error)
	Next() (string, error)
	Previous() (string, error)
}
