package service

import (
	"encoding/json"
	"model"
	"net/http"
	"response"
	"strconv"

	"github.com/gorilla/mux"
)

type productDAO interface {
	Get(ID int) (model.Product, error)
	Create(product model.Product) (model.Product, error)
	Update(ID int, product model.Product) (model.Product, error)
	Delete(ID int) error
	FindAll() ([]model.Product, error)
}

type ProductService struct {
	DAO productDAO
}

func GetProductService(dao productDAO) *ProductService {
	return &ProductService{dao}
}

func (service *ProductService) Create(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["id"])
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		response.RespondError(w, http.StatusMethodNotAllowed, err.Error())
		return
	}
	product.ID = ID
	product, err := service.DAO.Create(product)
	if err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	response.RespondJSON(w, http.StatusCreated, product)
}

func (service *ProductService) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["id"])
	product, err := service.DAO.Get(ID)
	if err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	response.RespondJSON(w, http.StatusOK, product)
}

func (service *ProductService) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["id"])
	if err := service.DAO.Delete(ID); err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
	}
	response.RespondJSON(w, http.StatusOK, "Product Successfully Deleted")
}

func (service *ProductService) Update(w http.ResponseWriter, r *http.Request) {
	response.RespondJSON(w, http.StatusOK, "Value Updated")
}

func (service *ProductService) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := service.DAO.FindAll()
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Server error")
	}
	response.RespondJSON(w, http.StatusOK, products)
}
