package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"bt/project/controller"
)

func ConfigRouterProduct() *mux.Router {
	routers := mux.NewRouter()
	// routers.HandleFunc("/", controller.Hello).Methods(http.MethodGet)
	routeApi := routers.PathPrefix("/c").Subrouter()
	routeApi.Use(controller.AuthMiddleware)
	routeProduct(routeApi)

	//API PRODUCTS
	//Get Products By Id
	routers.HandleFunc("/api/products/{id_product}", controller.GetProductById).Methods(http.MethodGet)

	//Get All Products
	routers.HandleFunc("/api/products", controller.GetAllProducts).Methods(http.MethodGet)
	//Create Product
	routers.Methods(http.MethodPost).Path("/api/products").HandlerFunc(controller.CreateProduct)
	//Delete Product By Id
	routers.Methods(http.MethodDelete).Path("/api/products/{id_product}").HandlerFunc(controller.DeleteProductById)
	//Update Product
	routers.Methods(http.MethodPut).Path("/api/products").HandlerFunc(controller.UpdateProduct)

	//API ORDER
	//Get Order By Id
	routers.Methods(http.MethodGet).Path("/api/orders/{user_id}").HandlerFunc(controller.GetOrdersDetailsByUserId)
	//save order to database active = 1
	routers.Methods(http.MethodPut).Path("/api/orders").HandlerFunc(controller.SaveOrderByUserActive)
	//save order to database active = 0
	routers.Methods(http.MethodPost).Path("/api/orders").HandlerFunc(controller.SaveOrderByUserNotActive)

	//USERS
	routers.Methods(http.MethodPost).Path("/api/register").HandlerFunc(controller.Register)
	routers.Methods(http.MethodPost).Path("/api/login").HandlerFunc(controller.Login)
	routers.Methods(http.MethodPost).Path("/api/logout").HandlerFunc(controller.Logout)

	//CATEGORY
	routers.HandleFunc("/api/categories", controller.GetAllCategories).Methods(http.MethodGet)

	return routers
}

func routeProduct(router *mux.Router) {
	router.HandleFunc("/api/products", controller.GetAllProducts).Methods(http.MethodGet)
}
