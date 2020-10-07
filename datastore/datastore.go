package datastore

// DataStore is a key value store which has two methods:
// Get() - for fetching value from key
// Set() - for setting keys with value
// temporarily put here
type DataStore interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}
