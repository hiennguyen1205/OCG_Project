package router

import (
	"bt/project/repository"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"

	"bt/project/controller"
	"bt/project/middleware"
)

type server struct {
	AuthController     controller.AuthController
	ProductController  controller.ProductController
	PaymentController  controller.PaymentController
	OrderController    controller.OrderController
	CategoryController controller.CategoryController
}

var controllerServer = &server{}

func ConfigRouterProduct(database *sql.DB) *mux.Router {
	initServer(database)
	routers := mux.NewRouter()
	// routers.HandleFunc("/", controller.Hello).Methods(http.MethodGet)
	routeApi := routers.PathPrefix("/auth").Subrouter()
	routeApi.Use(middleware.AuthMiddleware)
	RouteProduct(routeApi)

	//API PRODUCTS
	//Get Products By Id
	routers.HandleFunc("/api/products/{id_product}", controllerServer.ProductController.GetProductById).Methods(http.MethodGet)
	//Get All Products
	routers.HandleFunc("/api/products", controllerServer.ProductController.GetAllProducts).Methods(http.MethodGet)
	//Create Product
	routers.Methods(http.MethodPost).Path("/api/products").HandlerFunc(controllerServer.ProductController.CreateProduct)
	//Delete Product By Id
	routers.Methods(http.MethodDelete).Path("/api/products/{id_product}").HandlerFunc(controllerServer.ProductController.DeleteProductById)
	//Update Product
	routers.Methods(http.MethodPut).Path("/api/products").HandlerFunc(controllerServer.ProductController.UpdateProduct)

	//USERS
	//Register
	routers.Methods(http.MethodPost).Path("/api/register").HandlerFunc(controllerServer.AuthController.Register)
	//Login
	routers.Methods(http.MethodPost).Path("/api/login").HandlerFunc(controllerServer.AuthController.Login)
	//Logout
	routers.Methods(http.MethodPost).Path("/api/logout").HandlerFunc(controllerServer.AuthController.Logout)
	routers.Methods(http.MethodGet).Path("/api/user").HandlerFunc(controllerServer.AuthController.GetUserById)
	routers.Methods(http.MethodPut).Path("/api/user").HandlerFunc(controllerServer.AuthController.UpdateUser)
	routers.Methods(http.MethodPost).Path("/api/change-password").HandlerFunc(controllerServer.AuthController.ChangeUserPassword)

	//API ORDER
	//Get Order By Id
	routers.Methods(http.MethodGet).Path("/api/orders/{user_id}").HandlerFunc(controllerServer.OrderController.GetOrdersDetailsByUserId)

	//CATEGORY
	//Get All
	routers.HandleFunc("/api/categories", controllerServer.CategoryController.GetAllCategories).Methods(http.MethodGet)

	//PAYMENT
	routers.HandleFunc("/api/payment", controllerServer.PaymentController.HandleCreatePaymentIntent).Methods(http.MethodPost)

	//SIGNAL
	routers.HandleFunc("/api/email", controllerServer.OrderController.ChangeOrderState).Methods(http.MethodPost)

	return routers
}

func RouteProduct(routers *mux.Router) {
	//API ORDER
	//save order to database active = 1
	routers.Methods(http.MethodPut).Path("/api/orders").HandlerFunc(controllerServer.OrderController.SaveOrderByUserActive)
	//save order to database active = 0
	routers.Methods(http.MethodPost).Path("/api/orders").HandlerFunc(controllerServer.OrderController.SaveOrderByUserNotActiveController)
	//infor order
	routers.Methods(http.MethodGet).Path("/api/orders/infororder/{user_id}").HandlerFunc(controllerServer.OrderController.GetInformationOrder)
	//update order detail
	routers.Methods(http.MethodPut).Path("/api/orders/order_details").HandlerFunc(controllerServer.OrderController.UpdateOrderDetails)
	//get all orrder details for admin
	routers.Methods(http.MethodGet).Path("/api/orders/admin").HandlerFunc(controllerServer.OrderController.GetAllOrrderDetailsForAdmin)
	//delete order for admin
	routers.Methods(http.MethodPut).Path("/api/orders/{order_id}").HandlerFunc(controllerServer.OrderController.DeleteOrderDetails)
	//get details order admin
	routers.Methods(http.MethodGet).Path("/api/orders/details/{order_id}").HandlerFunc(controllerServer.OrderController.DetailsOrderAdmin)
}

func initServer(database *sql.DB) {
	//init repository struct
	UserRepository := &repository.UserRepository{
		Db: database,
	}
	ProductRepository := &repository.ProductRepository{
		Db: database,
	}
	OrderRepository := &repository.OrderRepository{
		Db: database,
	}
	CategoryRepository := &repository.CategoryRepository{
		Db: database,
	}
	//set controller struct
	controllerServer.AuthController = controller.AuthController{
		UserRepository: UserRepository,
	}
	controllerServer.ProductController = controller.ProductController{
		ProductRepository: ProductRepository,
	}
	controllerServer.PaymentController = controller.PaymentController{
		OrderRepository: OrderRepository,
	}
	controllerServer.OrderController = controller.OrderController{
		OrderRepository: OrderRepository,
		UserRepository:  UserRepository,
	}
	controllerServer.CategoryController = controller.CategoryController{
		CategoryRepository: CategoryRepository,
	}

}
