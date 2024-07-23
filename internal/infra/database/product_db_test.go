package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/savioafs/apiWithGo/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	as := assert.New(t)

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		as.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Nike Shox", 10)
	as.Nil(err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	as.NoError(err)
	as.NotEmpty(product.ID)
}

func TestFindAllProducts(t *testing.T) {
	as := assert.New(t)

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		as.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		as.NoError(err)

		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	as.NoError(err)
	as.Len(products, 10)
	as.Equal("Product 1", products[0].Name)
	as.Equal("Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	as.NoError(err)
	as.Len(products, 10)
	as.Equal("Product 11", products[0].Name)
	as.Equal("Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	as.NoError(err)
	as.Len(products, 3)
	as.Equal("Product 21", products[0].Name)
	as.Equal("Product 23", products[2].Name)
}

func TestFindByID(t *testing.T) {
	as := assert.New(t)

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		as.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 20.00)
	as.NoError(err)
	db.Create(product)

	productDB := NewProduct(db)
	product, err = productDB.FindByID(product.ID.String())
	as.NoError(err)
	as.Equal("Product 1", product.Name)
}

func TestUpdateProduct(t *testing.T) {
	as := assert.New(t)

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		as.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 20.50)
	as.NoError(err)
	db.Create(product)

	productDB := NewProduct(db)
	product.Name = "Product 2"

	err = productDB.Update(product)
	as.NoError(err)

	product, err = productDB.FindByID(product.ID.String())
	as.NoError(err)
	as.Equal("Product 2", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	as := assert.New(t)

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		as.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 20.50)
	as.NoError(err)
	db.Create(product)

	productDB := NewProduct(db)

	err = productDB.Delete(product.ID.String())
	as.NoError(err)

	_, err = productDB.FindByID(product.ID.String())
	as.Error(err)

}
