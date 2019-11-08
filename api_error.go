package gero

import "encoding/json"

// APIerror ...
type APIerror struct {
	Status int             `json:"-"`
	Title  string          `json:"title"`
	Detail json.RawMessage `json:"detail,omitempty"`
	Tags   []string        `json:"tags,omitempty"`
}

// NewAPIerror ...
func NewAPIerror(title string, status int, src error, tags ...string) *APIerror {
	err := &APIerror{
		Status: status,
		Title:  title,
		Tags:   tags,
	}
	if src != nil {
		if _, ok := src.(ValidationError); ok {
			err.Detail = json.RawMessage(src.Error())
		} else {
			err.Detail, _ = json.Marshal(src.Error())
		}
	}
	return err
}

func (err *APIerror) Error() string {
	buf, _ := json.Marshal(err)
	return string(buf)
}
