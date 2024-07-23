package database

import (
	"fmt"
	"testing"

	"github.com/savioafs/apiWithGo/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	as := assert.New(t)

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		as.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Marcos Cunha", "marquinhos@gmail.com", "12334")
	userDB := NewUser(db)

	err = userDB.Create(user)
	as.Nil(err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	as.Nil(err)
	as.Equal(userFound.Name, user.Name)
	as.Equal(userFound.Email, user.Email)
	as.NotNil(userFound.Password, user.Password)

}

func TestFindByEmail(t *testing.T) {
	as := assert.New(t)

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		as.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Julia", "julinha@gmail.com", "30312")
	userDB := NewUser(db)

	err = userDB.Create(user)
	as.Nil(err)

	userFound, err := userDB.FindByEmail(user.Email)
	as.Nil(err)
	as.Equal(user.ID, userFound.ID)
	as.Equal(user.Name, userFound.Name)
	as.Equal(user.Email, userFound.Email)
	as.NotNil(userFound.Password)

	fmt.Println(user)
}
