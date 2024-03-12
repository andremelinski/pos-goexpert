package entity_test

import (
	"testing"

	"github.com/andremelinski/pos-goexpert/8-API/internal/entity"
	"github.com/stretchr/testify/assert"
)


func TestNewUSer(t *testing.T){
	user, err := entity.NewUser("andre", "email", "andre123")
	assert.Nil(t, err)
	assert.Equal(t, user.Name, "andre")
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
}

func TestUser_ValidatePassword(t *testing.T){
	user, err := entity.NewUser("andre", "email", "andre123")
	assert.Nil(t, err)

	assert.True(t, user.ValidatePassword("andre123"))
	assert.False(t, user.ValidatePassword("123"))
	assert.NotEqual(t,"123" ,user.Password)

}