package controller

import (
	"bt/project/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

type PaymentMethodId struct {
	Id string `json:"payment_method_id"`
}

type ConfirmPaymentResponse struct {
	RequiresAction            bool    `json:"requires_action"`
	PaymentIntentClientSecret *string `json:"payment_intent_client_secret"`
}

func HandleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	stripe.Key = "sk_test_51JFBJVAuOJBTYkNRmXP9Bah92NkAdHKKJo7yZI9t78zwdF68DSBBjbFoLDD4LXOLZHYnyviPghKalVZVGiDqIBAy003NcEHr19"
	c := GetCookie(w, r)
	requestBody, _ := ioutil.ReadAll(r.Body)
	var pmKey PaymentMethodId
	json.Unmarshal(requestBody, &pmKey)
	fmt.Println(requestBody)
	fmt.Println(pmKey.Id)
	intIdUser, _ := strconv.Atoi(c.Value)
	amount, _ := repository.CalcAmount(intIdUser)
	params := &stripe.PaymentIntentParams{
		PaymentMethod:      stripe.String(pmKey.Id),
		Amount:             stripe.Int64(amount),
		Currency:           stripe.String(string(stripe.CurrencyVND)),
		Confirm:            stripe.Bool(true),
		ConfirmationMethod: stripe.String(string(stripe.PaymentIntentConfirmationMethodManual)),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		// Try to safely cast a generic error to a stripe.Error so that we can get at
		// some additional Stripe-specific information about what went wrong.
		if stripeErr, ok := err.(*stripe.Error); ok {
			fmt.Printf("Other Stripe error occurred: %v\n", stripeErr.Error())
			http.Error(w, stripeErr.Error(), 400)
		} else {
			fmt.Printf("Other error occurred: %v\n", err.Error())
			http.Error(w, "Unknown server error", 500)
		}

		return
	}

	if pi.Status == stripe.PaymentIntentStatusRequiresAction &&
		pi.NextAction.Type == "use_stripe_sdk" {

		// Tell the client to handle the action
		w.WriteHeader(http.StatusPaymentRequired)
		json.NewEncoder(w).Encode(ConfirmPaymentResponse{
			RequiresAction:            true,
			PaymentIntentClientSecret: &pi.ClientSecret,
		})

	} else if pi.Status == stripe.PaymentIntentStatusSucceeded {

		// The payment didn’t need any additional actions and completed!
		// Handle post-payment fulfillment
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Success")

	} else {

		// Invalid status
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Invalid Payment Intent status: %s", pi.Status)

	}

}

// func generatePaymentResponse(intent *stripe.PaymentIntent, w http.ResponseWriter) {
// 	if intent.Status == stripe.PaymentIntentStatusRequiresAction &&
// 		intent.NextAction.Type == "use_stripe_sdk" {

// 		// Tell the client to handle the action
// 		w.WriteHeader(http.StatusPaymentRequired)
// 		json.NewEncoder(w).Encode(ConfirmPaymentResponse{
// 			RequiresAction:            true,
// 			PaymentIntentClientSecret: &intent.ClientSecret,
// 		})

// 	} else if intent.Status == stripe.PaymentIntentStatusSucceeded {

// 		// The payment didn’t need any additional actions and completed!
// 		// Handle post-payment fulfillment
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprint(w, "Success")

// 	} else {

// 		// Invalid status
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintf(w, "Invalid Payment Intent status: %s", intent.Status)

// 	}
// }

// type PaymentGateway interface {
// 	auth(interface{}) (*Transaction, error)
// 	capture(interface{}) (*Transaction, error)
// }

// type transactionStatus string

// var authorized transactionStatus = "authorized"
// var captured transactionStatus = "captured"

// type Transaction struct {
// 	status transactionStatus
// 	amount int64
// }

// type Stripe struct {
// }

// func calcAmountToCharge(write http.ResponseWriter, request *http.Request) int64 {
// 	c, err := request.Cookie("id")
// 	if err != nil {
// 		log.Println("không lấy được cookie")
// 		statusCode := http.StatusUnauthorized
// 		http.Error(write, "Token doesnt exist", statusCode)
// 	}
// 	intIdUser, _ := strconv.Atoi(c.Value)
// 	result := repository.CalcAmount(intIdUser)

// 	return result

// }
// func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
// func (s Stripe) auth(paymentMethod interface{}) (*Transaction, error) {
// 	stripe.Key = "sk_test_51JFBJVAuOJBTYkNRmXP9Bah92NkAdHKKJo7yZI9t78zwdF68DSBBjbFoLDD4LXOLZHYnyviPghKalVZVGiDqIBAy003NcEHr19"

// 	amount := int64(1099)

// 	params := &stripe.PaymentIntentParams{
// 		Amount:        stripe.Int64(amount),
// 		PaymentMethod: stripe.String("pm_1JDqVPDNFpZPwqTEnFZCSdnR"),
// 		Currency:      stripe.String(string(stripe.CurrencyVND)),
// 		PaymentMethodTypes: stripe.StringSlice([]string{
// 			"card",
// 		}),
// 	}
// 	pi, err := paymentintent.New(params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Transaction{
// 		status: authorized,
// 		amount: pi.Amount,
// 	}, nil
// }

// func (s Stripe) capture(paymentMethod interface{}) (*Transaction, error) {
// 	stripe.Key = "sk_test_51JFBJVAuOJBTYkNRmXP9Bah92NkAdHKKJo7yZI9t78zwdF68DSBBjbFoLDD4LXOLZHYnyviPghKalVZVGiDqIBAy003NcEHr19"

// 	res, err := paymentintent.Confirm(
// 		"pi_1JDqW7DNFpZPwqTE4IcLGZbv",
// 		&stripe.PaymentIntentConfirmParams{
// 			PaymentMethod: stripe.String("pm_1JDqVPDNFpZPwqTEnFZCSdnR"),
// 		},
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Transaction{
// 		status: captured,
// 		amount: res.Amount,
// 	}, nil
// }

// type Factory struct {
// }

// func (f Factory) makePaymentGateway(gateway string) (PaymentGateway, error) {

// 	switch gateway {
// 	case "stripe":
// 		return Stripe{}, nil
// 	}

// 	return nil, errors.New(fmt.Sprintf("payment %s method not support", gateway))
// }

// func authorize(w http.ResponseWriter, req *http.Request) {

// 	gateway, err := Factory{}.makePaymentGateway("")

// 	if err != nil {

// 		return
// 	}

// 	//
// 	gateway.auth(nil)

// }

// func capture(w http.ResponseWriter, req *http.Request) {

// 	for name, headers := range req.Header {
// 		for _, h := range headers {
// 			fmt.Fprintf(w, "%v: %v\n", name, h)
// 		}
// 	}
// }
