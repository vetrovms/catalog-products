package tests

import (
	"catalog-products/internal/helpers"
	"testing"

	"gotest.tools/assert"
)

func TestSearchQuery(t *testing.T) {
	tests := []struct {
		input map[string]string
		want1 string
		want2 []string
	}{
		{
			input: map[string]string{
				"price_min": "10",
				"price_max": "100",
			},
			want1: "price >= ? and price <= ?",
			want2: []string{"10", "100"},
		},
		{
			input: map[string]string{
				"price_min":   "23",
				"price_max":   "32",
				"search_type": "or",
			},
			want1: "price >= ? or price <= ?",
			want2: []string{"23", "32"},
		},
	}

	for _, test := range tests {
		got1, got2 := helpers.SearchQuery(test.input)
		assert.Equal(t, test.want1, got1)
		for i, want := range test.want2 {
			assert.Equal(t, want, got2[i])
		}
	}
}

func TestOrderQuery(t *testing.T) {
	tests := []struct {
		input map[string]string
		want  string
	}{
		{
			input: map[string]string{
				"price_min": "10",
				"price_max": "100",
				"sort":      "smthwrong,title_asc,price_desc",
			},
			want: "title ASC, price DESC",
		},
	}

	for _, test := range tests {
		got := helpers.OrderQuery(test.input)
		assert.Equal(t, test.want, got)
	}
}
