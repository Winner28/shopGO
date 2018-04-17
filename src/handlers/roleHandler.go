package handlers

import (
	"dao"
	"net/http"
	"service"
)

var roleDAO *dao.RoleDAO
var productService *service.ProductService

func init() {
	productDAO = dao.GetProductDAO()
	productService = service.GetProductService(productDAO)
}

func (app *App) GetProduct(w http.ResponseWriter, r *http.Request) {
	productService.Get(w, r)
}

func (app *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	productService.Create(w, r)
}

func (app *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productService.Delete(w, r)
}

func (app *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productService.Update(w, r)
}

func (app *App) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	productService.GetAll(w, r)
}