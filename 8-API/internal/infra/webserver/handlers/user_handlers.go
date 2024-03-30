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

type Error struct {
	Message string `json:"message"`
}

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

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateUserInput  true  "user request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /user [post]
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
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return 
	}
	err = userHandler.UserDB.Create(userNormalized)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return 
	}
	w.WriteHeader(http.StatusCreated)
}

// GetJWT godoc
// @Summary      Get a user JWT
// @Description  Get a user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body     dto.GetJWTInput  true  "user credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /user/generate_token [post]
func (userHandler *UserHandler)GetJWT(w http.ResponseWriter, r *http.Request){
	userPayload := dto.GetJWTInput{}

	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error() }
		json.NewEncoder(w).Encode(error)
		return 
	}

	userFound, err := userHandler.UserDB.FindByEmail(userPayload.Email)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error() }
		json.NewEncoder(w).Encode(error)
		return 
	}
	if !userFound.ValidatePassword(userPayload.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return 
	}
	// info que vai voltar do jwt
	jwtAuth := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)
	// gera o token pelo metodo encode. para escolher qual info que vai dentro do token, usa o map
	_, tokenString, _ := jwtAuth.Encode(map[string]interface{}{
		"sub": userFound.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	token, _ := jwtAuth.Decode(tokenString)
	fmt.Printf("token decoded -> expiraiton: %s",token.Expiration())
	// monta a struct aqui pra poder serializar os dados do JSON e facilitar na resposta
	accessToken := dto.GetJWTOutput{
		AccessToken: tokenString,
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
		error := Error{Message: err.Error() }
		json.NewEncoder(w).Encode(error)
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
