package service

import (
	"encoding/json"
	"log"
	"managers"
	"model"
	"net/http"
	"resources"
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
	GetClothesProducts() ([]model.Product, error)
	GetTechsProducts() ([]model.Product, error)
}

type ProductService struct {
	DAO productDAO
}

func GetProductService(dao productDAO) *ProductService {
	return &ProductService{dao}
}

func (service *ProductService) Create(w http.ResponseWriter, r *http.Request) {
	if managers.GetSessionManager().UserLoggedIn(r) && secureService.checkIfAdmin(r) {
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
	} else {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetAccessDeniedError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
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
	if managers.GetSessionManager().UserLoggedIn(r) && secureService.checkIfAdmin(r) {

		params := mux.Vars(r)
		ID, _ := strconv.Atoi(params["id"])
		if err := service.DAO.Delete(ID); err != nil {
			response.RespondError(w, http.StatusNotFound, err.Error())
		}
		response.RespondJSON(w, http.StatusOK, "Product Successfully Deleted")
	} else {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetAccessDeniedError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (service *ProductService) Update(w http.ResponseWriter, r *http.Request) {
	response.RespondJSON(w, http.StatusOK, "Value Updated")
}

func (service *ProductService) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := service.DAO.FindAll()
	if err != nil {
		response.RespondError(w, http.StatusInternalServerError, "Server error")
	}
	if err := resources.GetTemplatesContainer().GetTemplate("allProducts").Execute(w, products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (service *ProductService) BuyProduct(w http.ResponseWriter, r *http.Request) {
	if managers.GetSessionManager().UserLoggedIn(r) {
		params := mux.Vars(r)
		ID, _ := strconv.Atoi(params["id"])
		product, err := service.DAO.Get(ID)
		if err != nil {
			if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetInternalServerErrorError()); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			if err := resources.GetTemplatesContainer().GetTemplate("buyProduct").Execute(w, product); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	} else {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.NotAuthorizedError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func (service *ProductService) GetClothesProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting clothes products")
	if products, err := service.DAO.GetClothesProducts(); err == nil {
		if err := resources.GetTemplatesContainer().GetTemplate("clothesProducts").Execute(w, products); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetInternalServerErrorError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (service *ProductService) GetTechsProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting techs clothes")
	if products, err := service.DAO.GetTechsProducts(); err == nil {
		if err := resources.GetTemplatesContainer().GetTemplate("techsProducts").Execute(w, products); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetInternalServerErrorError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
