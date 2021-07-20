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
	routeApi := routers.PathPrefix("/auth").Subrouter()
	routeApi.Use(middleware.AuthMiddleware)
	RouteProduct(routeApi)

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

	//USERS
	//Register
	routers.Methods(http.MethodPost).Path("/api/register").HandlerFunc(controller.Register)
	//Login
	routers.Methods(http.MethodPost).Path("/api/login").HandlerFunc(controller.Login)
	//Logout
	routers.Methods(http.MethodPost).Path("/api/logout").HandlerFunc(controller.Logout)
	routers.Methods(http.MethodGet).Path("/api/user").HandlerFunc(controller.GetUserById)
	routers.Methods(http.MethodPut).Path("/api/user").HandlerFunc(controller.UpdateUser)
	routers.Methods(http.MethodPost).Path("/api/change-password").HandlerFunc(controller.ChangeUserPassword)

	//API ORDER
	//Get Order By Id
	routers.Methods(http.MethodGet).Path("/api/orders/{user_id}").HandlerFunc(controller.GetOrdersDetailsByUserId)

	//CATEGORY
	//Get All
	routers.HandleFunc("/api/categories", controller.GetAllCategories).Methods(http.MethodGet)

	return routers
}

func RouteProduct(routers *mux.Router) {
	//API ORDER
	//save order to database active = 1
	routers.Methods(http.MethodPut).Path("/api/orders").HandlerFunc(controller.SaveOrderByUserActive)
	//save order to database active = 0
	routers.Methods(http.MethodPost).Path("/api/orders").HandlerFunc(controller.SaveOrderByUserNotActive)
	//infor order
	routers.Methods(http.MethodGet).Path("/api/orders/infororder/{user_id}").HandlerFunc(controller.GetInformationOrder)
}
