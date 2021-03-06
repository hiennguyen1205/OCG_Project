package main

import (
	"fmt"
	"log"
	"net/http"

	"bt/project/connect"
	"bt/project/controller"
	"bt/project/router"

	"github.com/gorilla/handlers"
	"github.com/streadway/amqp"
)

func main() {
	routers := router.ConfigRouterProduct()
	//connect database
	connect.Connect()

	//cấu hình file public
	routers.PathPrefix("/static/image/").Handler(http.StripPrefix("/static/image/", http.FileServer(http.Dir("./static/image/"))))

	//cấu hình CORS
	handleCross := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Origin", "Accept"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"http://localhost:8080"}),
		handlers.AllowCredentials(),
	)

	conn, err := amqp.Dial("amqp://guest:guest@174.138.40.239:5672/")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	fmt.Println("Successfuly Connected To our RMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	fmt.Println("Successfuly created channel")

	msgs, err := ch.Consume(
		"EmailQueue", // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for data := range msgs {
			controller.SendEmailBySendGrid(data.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	log.Fatal(http.ListenAndServe(":3000", handleCross(routers)))

}
