package main

import (
	"net/http"

	"bt/project/connect"
	"bt/project/router"
)

func main() {
	routers := router.ConfigRouterProduct()
	//connect database
	connect.Connect()

	//đoạn này không biết có cần config CORS không
	// These two lines are important if you're designing a front-end to utilise this API methods
	// allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	// allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	//https://medium.com/@gautamprajapati/writing-a-simple-e-commerce-website-api-in-go-programming-language-9f671324b4dd

	err := http.ListenAndServe(":3000", (routers))
	// log.Fatal(http.ListenAndServe(":3000", routers))
	if err != nil {
		panic(err)
	}
}
