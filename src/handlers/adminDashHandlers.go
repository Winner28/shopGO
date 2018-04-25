package handlers

import "net/http"

func (app *App) UpdateUserForm(w http.ResponseWriter, r *http.Request) {
	userService.UpdateForm(w, r)
}

func (app *App) CreateUserForm(w http.ResponseWriter, r *http.Request) {
	userService.CreateForm(w, r)
}

func (app *App) UpdateProductForm(w http.ResponseWriter, r *http.Request) {
	productService.UpdateForm(w, r)
}

func (app *App) CreateProductForm(w http.ResponseWriter, r *http.Request) {
	productService.CreateForm(w, r)
}

func (app *App) AdminProductsBoard(w http.ResponseWriter, r *http.Request) {
	productService.ProductsBoard(w, r)
}
