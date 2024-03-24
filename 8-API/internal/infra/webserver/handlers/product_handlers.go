package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andremelinski/pos-goexpert/8-API/internal/dto"
	"github.com/andremelinski/pos-goexpert/8-API/internal/entity"
	"github.com/andremelinski/pos-goexpert/8-API/internal/infra/db"
	entityPkg "github.com/andremelinski/pos-goexpert/8-API/pkg/entity"
	"github.com/go-chi/chi/v5"
)

// handler com interface na struct obriga a criar todos os metodos
type ProductHandler struct{
	ProductDB db.ProductInterface
}
// inicia o "controller" colocando a interface que eh utilizada nos produtos
func ProductHandlerInit(db db.ProductInterface)*ProductHandler{
	return &ProductHandler{
		db,
	}
}

func(productHandler *ProductHandler) CreateProduct( w http.ResponseWriter, r *http.Request){
	payload := dto.CreateProductInput{}
	
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	defer r.Body.Close()
	productNormalized, err := entity.NewProduct(payload.Name, payload.Price)
	
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return 
	}
	err = productHandler.ProductDB.Create(productNormalized)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	w.WriteHeader(http.StatusCreated)
}

func(productHandler *ProductHandler) GetProductById( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	productId := chi.URLParam(r,"id")

	if productId=="" {
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	_, err := entityPkg.ParseID(productId)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	product, err := productHandler.ProductDB.FindByID(productId)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
func(productHandler *ProductHandler) GetProducts( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	fmt.Println(limitInt)


	product, err := productHandler.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func(productHandler *ProductHandler) ProductUpdate( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	productToUpdate := entity.Product{}
	
	productId := chi.URLParam(r,"id")

	if productId=="" {
		w.WriteHeader(http.StatusBadRequest)
	}

	err := json.NewDecoder(r.Body).Decode(&productToUpdate)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	productToUpdate.ID, err = entityPkg.ParseID(productId)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	_ ,err = productHandler.ProductDB.FindByID(productId)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}
	err = productHandler.ProductDB.Update(&productToUpdate)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	w.WriteHeader(http.StatusOK)
}

func(productHandler *ProductHandler) ProductDelete( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	productId := chi.URLParam(r,"id")

	if productId=="" {
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err := entityPkg.ParseID(productId)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	err = productHandler.ProductDB.Delete(productId)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("deleted")
}