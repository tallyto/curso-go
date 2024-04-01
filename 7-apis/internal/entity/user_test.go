package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "johndoe@example.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)

	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "johndoe@example.com", user.Email)

}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "johndoe@example.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.True(t, user.ValidatePassword("123456"))

	assert.False(t, user.ValidatePassword("invalidpassword"))

	assert.NotEqual(t, "123456", user.Password)
}
