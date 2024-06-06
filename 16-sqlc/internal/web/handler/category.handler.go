package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/andremelinski/pos-goexpert/16-sqlc/internal/db"
	"github.com/andremelinski/pos-goexpert/16-sqlc/internal/interfaces"
	"github.com/andremelinski/pos-goexpert/16-sqlc/internal/web/dto"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)
type Error struct {
	Message string `json:"message"`
}

type CategoryHandler struct{
	categoryDB interfaces.CategoryInterface
}


func CategoryHandlerInit(categoryDB interfaces.CategoryInterface) *CategoryHandler{
	return &CategoryHandler{
		categoryDB,
	}
}

func (ch *CategoryHandler)CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	categoryPayload := dto.CreateCategoryInput{}
	err := json.NewDecoder(r.Body).Decode(&categoryPayload)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	userNormalized := db.CreateCategoryParams{
		ID: uuid.New().String(),
		Name: categoryPayload.Name, 
		Description: sql.NullString{String: categoryPayload.Description, Valid: true},
	}
	// if err != nil{
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	error := Error{Message: err.Error()}
	// 	json.NewEncoder(w).Encode(error)
	// 	return 
	// }
	_, err = ch.categoryDB.CreateCategory(context.Background(), userNormalized)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return 
	}
	w.WriteHeader(http.StatusCreated)
}

func (ch *CategoryHandler)GetCategory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	categoryId := chi.URLParam(r,"id")

	if categoryId=="" {
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	category, err := ch.categoryDB.GetCategory(context.Background(), categoryId)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}


func(ch *CategoryHandler) GetCategories( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	categories, err := ch.categoryDB.ListCategories(context.Background())
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}