package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"bt/project/controller"
	"bt/project/middleware"
)

func ConfigRouterProduct() *mux.Router {
	routers := mux.NewRouter()
	// routers.HandleFunc("/", controller.Hello).Methods(http.MethodGet)
	routeApi := routers.PathPrefix("/c").Subrouter()
	routeApi.Use(middleware.AuthMiddleware)

	//API PRODUCTS
	//Get All Products
	routers.HandleFunc("/api/products", controller.GetAllProducts).Methods(http.MethodGet)
	//Get Products By Id

	http.Handle("/api/products/{id_product}", middleware.AuthMiddleware(routers))
	//Create Product
	routers.Methods(http.MethodPost).Path("/api/products").HandlerFunc(controller.CreateProduct)
	//Delete Product By Id
	routers.Methods(http.MethodDelete).Path("/api/products/{id_product}").HandlerFunc(controller.DeleteProductById)
	//Update Product
	routers.Methods(http.MethodPut).Path("/api/products").HandlerFunc(controller.UpdateProduct)

	//API ORDER
	//Get Order By Id
	routers.Methods(http.MethodGet).Path("/api/orders").HandlerFunc(controller.GetOrdersDetailsByUserId)
	//tecs check valid orders
	routers.Methods(http.MethodGet).Path("/api/orders/check").HandlerFunc(controller.CheckValid)
	//save order to database
	routers.Methods(http.MethodPut).Path("/api/orders").HandlerFunc(controller.SaveOrderByUserActive)

	//USERS
	routers.Methods(http.MethodPost).Path("/api/register").HandlerFunc(controller.Register)
	routers.Methods(http.MethodPost).Path("/api/login").HandlerFunc(controller.Login)
	routers.Methods(http.MethodGet).Path("/api/logout").HandlerFunc(controller.Logout)

	//CATEGORY
	routers.HandleFunc("/api/categories", controller.GetAllCategories).Methods(http.MethodGet)

	return routers
}
