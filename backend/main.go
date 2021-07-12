package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"bt/project/connect"
	"bt/project/router"
)

func main() {
	routers := router.ConfigRouterProduct()
	//connect database
	connect.Connect()

	//cấu hình file public
	routers.PathPrefix("/static/image/").Handler(http.StripPrefix("/static/image/", http.FileServer(http.Dir("./static/image/"))))

	//cấu hình CORS
	handleCross :=handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), 
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), 
		handlers.AllowedOrigins([]string{"*"}),
	)
	log.Fatal(http.ListenAndServe(":3000", handleCross(routers)))
}
