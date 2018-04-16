package handlers

import (
	"model"
	"net/http"
	"service"
)

func (app *App) GetProduct(w http.ResponseWriter, r *http.Request) {
	service.Get(w, r, model.Product{})
}

func (app *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	service.Create(w, r, model.Product{})
}

func (app *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	service.Delete(w, r, model.Product{})
}

func (app *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	service.Update(w, r, model.Product{})
}

func (app *App) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	service.GetAll(w, r, model.Product{})
}
