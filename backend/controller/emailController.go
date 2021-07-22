package controller

import (
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmailBySendGrid(e interface{}) {
	log.Println(e)
	from := mail.NewEmail("Công ty Nội thất nhà xinh", "nhandt@reactplus.com")
	subject := "Thư cảm ơn"
	to := mail.NewEmail("Người nhận", "simsonabcxyz@gmail.com")
	plainTextContent := "Game de vai lon"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("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")
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
