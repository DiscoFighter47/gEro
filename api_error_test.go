package gero_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/DiscoFighter47/gero"
	"github.com/stretchr/testify/assert"
)

func TestAPIError(t *testing.T) {
	t.Run("api error", func(t *testing.T) {
		err := gero.NewAPIerror("api error", http.StatusInternalServerError, fmt.Errorf("demo api error"))
		assert.JSONEq(t, `{"title":"api error","detail":"demo api error"}`, err.Error())
	})

	t.Run("validation error", func(t *testing.T) {
		errV := gero.ValidationError{}
		errV.Add("field1", "required")
		errV.Add("field1", "invalid")
		errV.Add("field2", "required")

		err := gero.NewAPIerror("api error", http.StatusInternalServerError, errV)
		assert.JSONEq(t, `{"title":"api error","detail":{"field1":["required","invalid"],"field2":["required"]}}`, err.Error())
		t.Log(err)
	})
}
