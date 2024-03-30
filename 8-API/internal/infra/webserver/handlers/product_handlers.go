package handlers

import (
	"encoding/json"
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

// Create Product godoc
// @Summary      Create product
// @Description  Create products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateProductInput  true  "product request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /products [post]
// @Security ApiKeyAuth
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

// GetProduct godoc
// @Summary      Get a product
// @Description  Get a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "product ID" Format(uuid)
// @Success      200  {object}  entity.Product
// @Failure      404
// @Failure      500  {object}  Error
// @Router       /products/{id} [get]
// @Security ApiKeyAuth
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

// ListAccounts godoc
// @Summary      List products
// @Description  get all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page      query     string  false  "page number"
// @Param        limit     query     string  false  "limit"
// @Success      200       {array}   entity.Product
// @Failure      404       {object}  Error
// @Failure      500       {object}  Error
// @Router       /products [get]
func(productHandler *ProductHandler) GetProducts( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	product, err := productHandler.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Update a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id        	path      string                  true  "product ID" Format(uuid)
// @Param        request     body      dto.CreateProductInput  true  "product request"
// @Success      200
// @Failure      404
// @Failure      500       {object}  Error
// @Router       /products/{id} [put]
// @Security ApiKeyAuth
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

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id        path      string                  true  "product ID" Format(uuid)
// @Success      200
// @Failure      404
// @Failure      500       {object}  Error
// @Router       /products/{id} [delete]
// @Security ApiKeyAuth
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