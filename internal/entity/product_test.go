package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {

	tests := []struct {
		testName string
		name     string
		price    int

		expectName  string
		expectPrice int
	}{
		{
			testName: "create product Nike Air Force One without error",
			name:     "Nike Air Force 1",
			price:    3500,

			expectName:  "Nike Air Force 1",
			expectPrice: 3500,
		},
		{
			testName: "create product Stanley Cup without error",
			name:     "Stanley Cup",
			price:    200,

			expectName:  "Stanley Cup",
			expectPrice: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := assert.New(t)

			product, err := NewProduct(tt.name, tt.price)
			as.Nil(err)
			as.NotNil(product)
			as.NotEmpty(product.ID)
			as.Equal(tt.expectName, product.Name)
			as.Equal(tt.expectPrice, product.Price)
		})
	}
}

func TestProductWhenNameIsRequired(t *testing.T) {
	as := assert.New(t)
	product, err := NewProduct("", 10)
	as.Nil(product)
	as.Equal(ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	as := assert.New(t)
	product, err := NewProduct("Product 1", 0)
	as.Nil(product)
	as.Equal(ErrPriceIsRequired, err)
}

func TestProductWhenPriceInvalid(t *testing.T) {
	as := assert.New(t)
	product, err := NewProduct("Product 1", -10)
	as.Nil(product)
	as.Equal(ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	as := assert.New(t)
	product, err := NewProduct("Product 1", 10)
	as.NotNil(product)
	as.Nil(err)
	as.Nil(product.Validate())
}
