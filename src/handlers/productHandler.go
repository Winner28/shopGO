package handlers

import "net/http"
import "service"

func (app *App) GetProduct(w http.ResponseWriter, r *http.Request) {
	service.GetProduct(w, r)
}

func (app *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	service.CreateProduct(w, r)
}

func (app *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	service.DeleteProduct(w, r)
}

func (app *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	service.UpdateProduct(w, r)
}

func (app *App) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	service.GetAllProducts(w, r)
}
