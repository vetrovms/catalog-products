package tests

import (
	"catalog-products/internal/request"
	"testing"

	"gotest.tools/assert"
)

func TestValidate(t *testing.T) {
	test := request.ProductRequest{
		Price:    -2,
		Color:    "#FFFFFF123",
		Currency: "test",
	}
	expected := []string{
		"[Price]: не відповідає правилу 'gte' 0",
		"[Color]: не відповідає правилу 'iscolor' ",
		"[Currency]: не відповідає правилу 'iso4217' ",
	}

	err := test.Validate()
	for i, v := range expected {
		assert.Equal(t, v, err[i])
	}
}
