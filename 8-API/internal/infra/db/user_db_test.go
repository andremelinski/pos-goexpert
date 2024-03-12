package db

import (
	"testing"

	"github.com/andremelinski/pos-goexpert/8-API/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T){
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("andre", "email@.com", "123")
	userDB := UserInitDB(db)
	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound := entity.User{}
	err = db.Find(&userFound, "id = ? ", user.ID).Error

	assert.Nil(t, err)

	assert.Equal(t, user.ID, userFound.ID,)
	assert.Equal(t, userFound.Name, "andre")

	assert.True(t, userFound.ValidatePassword("123"))
}

func TestFindByEmail(t *testing.T){
	db, err := gorm.Open(sqlite.Open("file:memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("andre", "email2@.com", "123")
	userDB := UserInitDB(db)
	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)

	assert.Equal(t, user.ID, userFound.ID,)
	assert.Equal(t, userFound.Name, "andre")

	assert.True(t, userFound.ValidatePassword("123"))
}