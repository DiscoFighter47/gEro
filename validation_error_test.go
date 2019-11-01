package gero_test

import (
	"testing"

	"github.com/DiscoFighter47/gero"
	"github.com/stretchr/testify/assert"
)

func TestValidationError(t *testing.T) {
	err := gero.ValidationError{}
	err.Add("field1", "required")
	err.Add("filed2", "invalid")
	assert.JSONEq(t, `{"field1":"required", "filed2":"invalid"}`, err.Error())
}
