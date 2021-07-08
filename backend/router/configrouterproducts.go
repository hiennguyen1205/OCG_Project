package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"bt/project/controller"
)

func ConfigRouterProduct() *mux.Router {
	routers := mux.NewRouter()
	// routers.HandleFunc("/", controller.Hello).Methods(http.MethodGet)

	//API PRODUCTS
	//Get All Products
	routers.HandleFunc("/api/products", controller.GetAllProducts).Methods(http.MethodGet)
	//Get Products By Id
	routers.Methods(http.MethodGet).Path("/api/products/{id_product}").HandlerFunc(controller.GetProductById)
	//Create Product
	routers.Methods(http.MethodPost).Path("/api/products").HandlerFunc(controller.CreateProduct)
	//Delete Product By Id
	routers.Methods(http.MethodDelete).Path("/api/products/{id_product}").HandlerFunc(controller.DeleteProductById)
	//Update Product
	routers.Methods(http.MethodPut).Path("/api/products").HandlerFunc(controller.UpdateProduct)

	//API ORDER
	//Get Order By Id
	routers.Methods(http.MethodGet).Path("/api/orders").HandlerFunc(controller.GetAllOrdersDetails)

	//USERS
	routers.Methods(http.MethodPost).Path("/api/register").HandlerFunc(controller.Register)
	routers.Methods(http.MethodPost).Path("/api/login").HandlerFunc(controller.Login)
	return routers
}
