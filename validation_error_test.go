package gero_test

import (
	"testing"

	"github.com/DiscoFighter47/gero"
	"github.com/stretchr/testify/assert"
)

func TestValidationError(t *testing.T) {
	err := gero.ValidationError{}

	t.Run("nil error", func(t *testing.T) {
		assert.Equal(t, 0, len(err))
	})

	t.Run("validation error", func(t *testing.T) {
		err.Add("field1", "required")
		err.Add("field2", "required")
		assert.JSONEq(t, `{"field1":["required"],"field2":["required"]}`, err.Error())
		assert.Equal(t, 2, len(err))
	})

	t.Run("multiple error", func(t *testing.T) {
		err.Add("field2", "invalid")
		assert.JSONEq(t, `{"field1":["required"],"field2":["required","invalid"]}`, err.Error())
		assert.Equal(t, 2, len(err))
	})

	t.Run("error extend", func(t *testing.T) {
		errV := gero.ValidationError{}
		errV.Add("field1", "invalid")
		errV.Add("field3", "required")

		err.Extend(errV)
		assert.JSONEq(t, `{"field1":["required","invalid"],"field2":["required","invalid"],"field3":["required"]}`, err.Error())
		assert.Equal(t, 3, len(err))
	})
}
