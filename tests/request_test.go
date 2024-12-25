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
		"Price: значення має бути більше 0",
		"Color: значення має відповідати формату кольора",
		"Currency: значення має відповідати формату iso4217",
	}

	err := test.Validate()
	for i, v := range expected {
		assert.Equal(t, v, err[i])
	}
}
