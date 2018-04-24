package handlers

import (
	"html/template"
	"model"
	"net/http"
	"resources"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

var app App
var templates *resources.Templates

func Init() {
	app.Router = mux.NewRouter()
	app.setRouters()
	app.setTemplates()
	http.ListenAndServe(":8080", app.Router)
}

func (app *App) setRouters() {
	app.setAuthRoutes()
	app.setMainRoutes()
	app.setUsersRoutes()
	app.setProductRoutes()
	app.setRoleRoutes()
	app.setNotFoundHanlder()
	app.setProfileRoutes()
	app.setAdminDashRoutes()
}

func (app *App) setAdminDashRoutes() {
	app.Router.HandleFunc("/users/update/{id}", app.UpdateUserForm).Methods("GET")
	app.Router.HandleFunc("/users", app.GetAllUsers).Methods("GET")
}

func (app *App) setAuthRoutes() {
	app.Router.HandleFunc("/signin", app.Login).Methods("POST")
	app.Router.HandleFunc("/sign", app.LoginPage).Methods("GET")
	app.Router.HandleFunc("/outs", app.Logout).Methods("GET")
	app.Router.HandleFunc("/signup", app.SignUp).Methods("POST")
	app.Router.HandleFunc("/register", app.SignUpPage).Methods("GET")
}

func (app *App) setUsersRoutes() {
	app.Router.HandleFunc("/users/{id}", app.GetUser).Methods("GET")
	app.Router.HandleFunc("/users", app.CreateUser).Methods("POST")
	app.Router.HandleFunc("/users/{id}", app.UpdateUser).Methods("PUT")
	app.Router.HandleFunc("/users/update/{id}", app.UpdateUser).Methods("POST")
	app.Router.HandleFunc("/users/delete/{id}", app.DeleteUser).Methods("POST")
}

func (app *App) setProductRoutes() {
	app.Router.HandleFunc("/products/{id}", app.GetProduct).Methods("GET")
	app.Router.HandleFunc("/products/{id}", app.CreateProduct).Methods("POST")
	app.Router.HandleFunc("/products/{id}", app.UpdateProduct).Methods("PUT")
	app.Router.HandleFunc("/products/{id}", app.DeleteProduct).Methods("DELETE")
	app.Router.HandleFunc("/products", app.GetAllProducts).Methods("GET")
	app.Router.HandleFunc("/category/clothes", app.GetClothesProducts).Methods("GET")
	app.Router.HandleFunc("/category/techs", app.GetTechsProducts).Methods("GET")
	app.Router.HandleFunc("/products/buy/{id}", app.BuyProduct).Methods("GET")
}

func (app *App) setRoleRoutes() {
	app.Router.HandleFunc("/roles/{id}", app.GetUserRole).Methods("GET")
	app.Router.HandleFunc("/roles/{id}", app.CreateRole).Methods("POST")
	app.Router.HandleFunc("/roles/{id}", app.UpdateUserRole).Methods("PUT")
	app.Router.HandleFunc("/roles/{id}", app.DeleteUserRole).Methods("DELETE")
	app.Router.HandleFunc("/roles", app.GetAllRoles).Methods("GET")
}

func (app *App) setNotFoundHanlder() {
	app.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := resources.GetTemplatesContainer().GetTemplate("error").Execute(w, model.GetNotFoundError()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func (app *App) setProfileRoutes() {
	app.Router.HandleFunc("/profile", app.GetProfile).Methods("GET")
}

func (app *App) setTemplates() {
	templates = resources.GetTemplatesContainer()
	templates.AddTemplate("signin", template.Must(template.ParseFiles("templates/auth/signin.html")))
	templates.AddTemplate("signup", template.Must(template.ParseFiles("templates/auth/signup.html")))

	templates.AddTemplate("error", template.Must(template.ParseFiles("templates/errors/error.html")))
	templates.AddTemplate("main", template.Must(template.ParseFiles("templates/main.html")))
	templates.AddTemplate("shop", template.Must(template.ParseFiles("templates/shop.html")))

	templates.AddTemplate("profile", template.Must(template.ParseFiles("templates/profile.html")))
	templates.AddTemplate("getAllUsers", template.Must(template.ParseFiles("templates/adminDashboard/getAllUsers.html")))
	templates.AddTemplate("message", template.Must(template.ParseFiles("templates/adminDashboard/message.html")))
	templates.AddTemplate("updateUser", template.Must(template.ParseFiles("templates/adminDashboard/updateUser.html")))
	templates.AddTemplate("createUser", template.Must(template.ParseFiles("templates/adminDashboard/createUser.html")))

	templates.AddTemplate("allProducts", template.Must(template.ParseFiles("templates/products/allProducts.html")))
	templates.AddTemplate("buyProduct", template.Must(template.ParseFiles("templates/products/buyProduct.html")))
	templates.AddTemplate("productCategory", template.Must(template.ParseFiles("templates/products/productCategory.html")))

	app.Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
}

func (app *App) setMainRoutes() {
	app.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := resources.GetTemplatesContainer().GetTemplate("main").Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	app.Router.HandleFunc("/shop", func(w http.ResponseWriter, r *http.Request) {
		if err := resources.GetTemplatesContainer().GetTemplate("shop").Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

}
