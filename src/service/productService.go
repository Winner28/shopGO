package service

import (
	"errors"
	"fmt"
	"log"
	"managers"
	"model"
	"net/http"
	"resources"
	"response"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type productDAO interface {
	Get(ID int) (model.Product, error)
	Create(product model.Product, category model.Category) (model.Product, error)
	Update(product model.Product, category model.Category) (model.Product, error)
	Delete(ID int) error
	FindAll() ([]model.Product, error)
	GetClothesProducts() ([]model.Product, error)
	GetTechsProducts() ([]model.Product, error)
	CreateOrder(order model.Order, productID int) (model.Order, error)
}

type ProductService struct {
	DAO productDAO
}

func GetProductService(dao productDAO) *ProductService {
	return &ProductService{dao}
}

func validateProductInput(product model.Product) error {
	if product.Description == "" || product.Name == "" {
		return errors.New("Fields cant be empty")
	}
	return nil
}

func (service *ProductService) Create(w http.ResponseWriter, r *http.Request) {
	if managers.GetSessionManager().UserLoggedIn(r) && secureService.checkIfAdmin(r) {
		price, err := strconv.ParseFloat(r.FormValue("price"), 32)
		if err != nil {
			if err := resources.GetTemplatesContainer().GetTemplate("message").Execute(w, model.ErrorWhileCreatingProduct(err.Error())); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		product, category := service.getProductAndCategoryFromRequest(r)
		product.Price = price
		if err = validateProductInput(product); err != nil {
			if err := resources.GetTemplatesContainer().GetTemplate("message").Execute(w, model.ErrorWhileCreatingProduct(err.Error())); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		product, err = service.DAO.Create(product, category)
		if err != nil {
			if err := resources.GetTemplatesContainer().GetTemplate("message").Execute(w, model.ErrorWhileCreatingProduct(err.Error())); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		if err := resources.GetTemplatesContainer().GetTemplate("message").Execute(w, model.ProductSuccessfullyCreated(product)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetAccessDeniedError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (service *ProductService) Update(w http.ResponseWriter, r *http.Request) {
	if managers.GetSessionManager().UserLoggedIn(r) && secureService.checkIfAdmin(r) {
		var price float64
		var err error
		ID, _ := strconv.Atoi(mux.Vars(r)["id"])
		if r.FormValue("price") == "" {
			price = 0
		} else {
			price, err = strconv.ParseFloat(r.FormValue("price"), 32)
			if err != nil {
				if err := resources.GetTemplatesContainer().GetTemplate("message").Execute(w, model.ErrorWhileCreatingProduct(err.Error())); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}
		}
		product, category := service.getProductAndCategoryFromRequest(r)
		product.Price = price
		product.ID = ID
		log.Println("Updating product:", product)
		product, err = service.DAO.Update(product, category)

		if err != nil {
			if err := resources.GetTemplatesContainer().GetTemplate("message").Execute(w, model.ErrorWhileUpdatingProduct(err.Error())); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		if err := resources.GetTemplatesContainer().GetTemplate("message").Execute(w, model.ProductSuccessfullyUpdated(product)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
			if err := resources.GetTemplatesContainer().GetTemplate("message").Execute(w, model.ErrorSystemProblems(err.Error(), "products")); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		if err := resources.GetTemplatesContainer().GetTemplate("message").Execute(w, model.ProductSuccessfullyDeleted()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetAccessDeniedError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
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

func (service *ProductService) BuyProductPage(w http.ResponseWriter, r *http.Request) {
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

func (service *ProductService) BuyProduct(w http.ResponseWriter, r *http.Request) {
	if managers.GetSessionManager().UserLoggedIn(r) {
		params := mux.Vars(r)
		order, error := service.getOrderFromRequest(r)
		if error != nil {
			// logic
			return
		} else {
			productID, _ := strconv.Atoi(params["id"])
			order, err := service.DAO.CreateOrder(order, productID)
			if err != nil {
				resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetError(http.StatusInternalServerError, "Server error"))
				log.Println("Failed to create an order!", err.Error())
			} else {
				fmt.Fprintln(w, "U successfully buy product")
				log.Println("Order created:", order)
			}
		}
	} else {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetError(http.StatusForbidden, "If you want buy a product you need to login first!")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func (service *ProductService) GetClothesProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting clothes products")
	if products, err := service.DAO.GetClothesProducts(); err == nil {
		products[0].Category = "Clothes"
		if err := resources.GetTemplatesContainer().GetTemplate("productCategory").Execute(w, products); err != nil {
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
		products[0].Category = "Technologies"
		if err := resources.GetTemplatesContainer().GetTemplate("productCategory").Execute(w, products); err != nil {
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

func (service *ProductService) getProductAndCategoryFromRequest(r *http.Request) (model.Product, model.Category) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	category := r.FormValue("category")
	return model.Product{
			Name:        name,
			Description: description,
		},
		model.Category{
			Name: category,
		}

}

func (service *ProductService) UpdateForm(w http.ResponseWriter, r *http.Request) {
	if access := getSecureService().checkIfAdmin(r); !access {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetAccessDeniedError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			return
		}
	}

	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["id"])

	if product, err := service.DAO.Get(ID); err == nil {
		if err := resources.GetTemplatesContainer().GetTemplate("updateProduct").Execute(w, product); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err := resources.GetTemplatesContainer().GetTemplate("message").Execute(w, model.WeDontHaveSuchUser()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (service *ProductService) CreateForm(w http.ResponseWriter, r *http.Request) {
	if access := getSecureService().checkIfAdmin(r); !access {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetAccessDeniedError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			return
		}
	}

	if err := resources.GetTemplatesContainer().GetTemplate("createProduct").Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (service *ProductService) ProductsBoard(w http.ResponseWriter, r *http.Request) {
	if access := getSecureService().checkIfAdmin(r); !access {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetAccessDeniedError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			return
		}
	}
	products, err := service.DAO.FindAll()
	if err != nil {
		if err := resources.GetTemplatesContainer().GetTemplate("message").Execute(w, model.ErrorSystemProblems(err.Error(), "products")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if err := resources.GetTemplatesContainer().GetTemplate("productsBoard").Execute(w, products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (service *ProductService) getOrderFromRequest(r *http.Request) (model.Order, error) {
	var order model.Order

	ammount := r.FormValue("ammount")
	if valid := validateAmmountInput(ammount); !valid {
		return model.Order{}, errors.New("Bad ammount input! PLease enter a number beetween 1 and 100")
	}

	userID := managers.GetSessionManager().GetUserID(r)
	amm, _ := strconv.Atoi(ammount)

	order.UserID = userID
	order.Amount = amm
	order.Date = time.Now()
	return order, nil
}

func validateAmmountInput(ammount string) bool {
	if amm, error := strconv.Atoi(ammount); error != nil || amm < 1 {
		return false
	}

	return true

}
