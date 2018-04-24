package handlers

import (
	"dao"
	"net/http"
	"service"
)

var productDAO *dao.ProductDAO
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

func (app *App) BuyProduct(w http.ResponseWriter, r *http.Request) {
	productService.BuyProduct(w, r)
}

func (app *App) GetClothesProducts(w http.ResponseWriter, r *http.Request) {
	productService.GetClothesProducts(w, r)
}

func (app *App) GetTechsProducts(w http.ResponseWriter, r *http.Request) {
	productService.GetTechsProducts(w, r)
}
