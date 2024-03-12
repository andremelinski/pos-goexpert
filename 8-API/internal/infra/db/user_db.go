package db

import (
	"github.com/andremelinski/pos-goexpert/8-API/internal/entity"
	"gorm.io/gorm"
)

// layer db -> struct da funcao init precisa ter a conexao com o banco
type UserDB struct {
	DB *gorm.DB
}

func UserInitDB(db *gorm.DB)*UserDB{
	return &UserDB{db}
}

// atrelando as interfaces com o db
func (u *UserDB) Create(user *entity.User) error{
	return u.DB.Create(user).Error
}

func (u *UserDB )FindByEmail(email string) (*entity.User, error){
	user := entity.User{}
	if err := u.DB.Where("email =?", email).First(&user).Error; err!=nil{
		return nil, err
	}
	return &user, nil
}