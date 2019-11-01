package gero

import "encoding/json"

// ValidationError holds multiple errors for multiple keys
type ValidationError map[string]string

func (err ValidationError) Error() string {
	buf, _ := json.Marshal(err)
	return string(buf)
}

// Add a new error for the given key
func (err ValidationError) Add(key, msg string) {
	err[key] = msg
}
