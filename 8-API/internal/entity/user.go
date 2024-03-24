package entity

import (
	"fmt"

	"github.com/andremelinski/pos-goexpert/8-API/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

//  como ID eh gerado pelo uuid e essa lib pode ser acessadaspor outras camadas da aplicacao,
// faz sentido colocar ela na pasta pkg(package)
type User struct{
	ID entity.ID `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"-"`
	// Password string `json:"password"`
}
// inicia um novo usuario
func NewUser(name, email, password string) (*User, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err!=nil {
		return nil, err
	}
	return &User{
		ID: entity.NewID(),
		Name: name,
		Email: email,
		Password: string(hash),
	}, nil
}

// metodo pertencente unicamente a struct User
func (u *User) ValidatePassword(password string) bool{
	fmt.Println(u.Password)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(password))
	return err == nil
}