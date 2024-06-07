package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/andremelinski/pos-goexpert/16-sqlc/internal/db"
	"github.com/andremelinski/pos-goexpert/16-sqlc/internal/interfaces"
	"github.com/andremelinski/pos-goexpert/16-sqlc/internal/web/dto"
	"github.com/google/uuid"
)

type CourseHandler struct{
	courseDB interfaces.CourseInterface
}


func CourseHandlerInit(courseDB interfaces.CourseInterface) *CourseHandler{
	return &CourseHandler{
		courseDB,
	}
}

func (ch *CourseHandler)CreateCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	coursePayload := dto.CreateCourseInput{}
	err := json.NewDecoder(r.Body).Decode(&coursePayload)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	categoryPayloadNormalized := db.CategoryParams{
		ID: uuid.New().String(),
		Name: coursePayload.Name, 
		Description: sql.NullString{String: coursePayload.Description, Valid: true},
	}

	coursePayloadNormalized := db.CourseParams{
		ID: uuid.New().String(),
		Name: coursePayload.Name, 
		Description: sql.NullString{String: coursePayload.Description, Valid: true},
		Price: coursePayload.Price,
	}

	 err = ch.courseDB.CreateCourseAndCategory(context.Background(), categoryPayloadNormalized, coursePayloadNormalized)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return 
	}
	w.WriteHeader(http.StatusCreated)
}

func(ch *CourseHandler) GetCourses( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	categories, err := ch.courseDB.ListCourses(context.Background())
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

