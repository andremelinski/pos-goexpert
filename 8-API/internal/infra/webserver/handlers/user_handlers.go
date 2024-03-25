package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/andremelinski/pos-goexpert/8-API/internal/dto"
	"github.com/andremelinski/pos-goexpert/8-API/internal/entity"
	"github.com/andremelinski/pos-goexpert/8-API/internal/infra/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct{
	UserDB db.UserInterface
	// JwtExpiresIn int
	// Jwt *jwtauth.JWTAuth
}

func UserHandlerInit(userDB db.UserInterface) *UserHandler{
	return &UserHandler{
		userDB,
		// expiresIn,
		// jwtAuth,
	}
}

func (userHandler *UserHandler)CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	userPayload := dto.CreateUserInput{}
	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	userNormalized, err := entity.NewUser(userPayload.Name, userPayload.Email, userPayload.Password)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}
	err = userHandler.UserDB.Create(userNormalized)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}
	w.WriteHeader(http.StatusCreated)
}

func (userHandler *UserHandler)GetJWT(w http.ResponseWriter, r *http.Request){
	userPayload := dto.GetJWTInput{}

	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	userFound, err := userHandler.UserDB.FindByEmail(userPayload.Email)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}
	if !userFound.ValidatePassword(userPayload.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return 
	}
	// info que vai voltar do jwt
	jwtAuth := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)
	_, stringToken, _ := jwtAuth.Encode(map[string]interface{}{
		"sub": userFound.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	token, _ := jwtAuth.Decode(stringToken)
	fmt.Printf("token decoded -> expiraiton: %s",token.Expiration())

	accessToken := struct{
		AccessToken string `json:"access_token"`
	}{
		AccessToken: stringToken,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

}

func (userHandler *UserHandler)GetUserByMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	email := chi.URLParam(r, "email")

	if email == ""{
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	user, err := userHandler.UserDB.FindByEmail(email)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
