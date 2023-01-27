package main

import (
	"bt/project/connect"
	"bt/project/router"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	//connect database
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Could not load .env file")
	}
	database := connect.Connect()
	defer database.Close()
	routers := router.ConfigRouterProduct(database)

	//cấu hình file public
	routers.PathPrefix("/static/image/").Handler(http.StripPrefix("/static/image/", http.FileServer(http.Dir("./static/image/"))))

	//cấu hình CORS
	handleCross := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Origin", "Accept"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"http://localhost:8080", "http://hoc-tran-test-s3-vue.s3-website-ap-southeast-1.amazonaws.com"}),
		handlers.AllowCredentials(),
	)

	//conn, err := amqp.Dial("amqp://guest:guest@174.138.40.239:5672/")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer conn.Close()
	//
	//fmt.Println("Successfuly Connected To our RMQ Instance")
	//
	//ch, err := conn.Channel()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer ch.Close()
	//
	//fmt.Println("Successfuly created channel")
	//
	//msgs, err := ch.Consume(
	//	"EmailQueue", // queue
	//	"",           // consumer
	//	true,         // auto-ack
	//	false,        // exclusive
	//	false,        // no-local
	//	false,        // no-wait
	//	nil,          // args
	//)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//go func() {
	//	for data := range msgs {
	//		controller.SendEmailBySendGrid(data.Body)
	//	}
	//}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	log.Fatal(http.ListenAndServe(":3000", handleCross(routers)))

}
