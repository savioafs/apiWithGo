package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {

	tests := []struct {
		testName string
		name     string
		email    string
		password string

		expectName  string
		expectEmail string
	}{
		{
			testName: "expect user not nil",
			name:     "Fulano da Cunha",
			email:    "fulanodacunha@live.com",
			password: "12345",

			expectName:  "Fulano da Cunha",
			expectEmail: "fulanodacunha@live.com",
		},
		{
			testName: "expect user not nil",
			name:     "Marcelo Ferreira",
			email:    "marcelinhoferreira@live.com",
			password: "12345",

			expectName:  "Marcelo Ferreira",
			expectEmail: "marcelinhoferreira@live.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := assert.New(t)

			user, err := NewUser(tt.name, tt.email, tt.password)
			as.Nil(err)
			as.NotNil(user)
			as.NotEmpty(user.ID)
			as.Equal(tt.expectName, user.Name)
			as.Equal(tt.expectEmail, user.Email)
		})
	}
}

func TestUser_ValidatePassword(t *testing.T) {

	tests := []struct {
		testName        string
		name            string
		email           string
		password        string
		invalidPassword string
	}{
		{
			testName:        "expect user not nil",
			name:            "Fulano da Cunha",
			email:           "fulanodacunha@live.com",
			password:        "12345",
			invalidPassword: "304050",
		},
		{
			testName:        "expect user not nil",
			name:            "Marcelo Ferreira",
			email:           "marcelinhoferreira@live.com",
			password:        "12345",
			invalidPassword: "309223",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := assert.New(t)

			user, err := NewUser(tt.name, tt.email, tt.password)
			as.Nil(err)
			as.True(user.ValidatePassword(tt.password))
			as.False(user.ValidatePassword(tt.invalidPassword))
			as.NotEqual(tt.password, user.Password)
		})
	}
}
