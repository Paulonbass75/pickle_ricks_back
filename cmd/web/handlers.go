package main

import (
	"net/http"
	"pickle_ricks_back/models"
)

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {

	if err := app.renderTemplate(w, r, "terminal", &templateData{
		
	}, "stripe-js"); err != nil {
		app.errorlog.Println(err)
	}
}

func (app *application) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {
	// Assume that the payment was successful
	err := r.ParseForm()
	if err != nil {
		app.errorlog.Println(err)
	}
	// read posted data
	cardHolder := r.Form.Get("cardholder_name")
	email := r.Form.Get("email")
	paymentIntent := r.Form.Get("payment_intent")
	paymentMethod := r.Form.Get("payment_method")
	paymentAmount := r.Form.Get("payment_amount")
	paymentCurrency := r.Form.Get("payment_Currency")

	data := make(map[string]interface{})
	data["cardholder"] = cardHolder
	data["email"] = email
	data["payment_intent"] = paymentIntent
	data["payment_method"] = paymentMethod
	data["payment_Amount"] = paymentAmount
	data["payment_Currency"] = paymentCurrency
	if err := app.renderTemplate(w, r, "succeeded", &templateData{
		Data: data,
	}); err != nil {
		app.errorlog.Println(err)
	}
}

func (app *application) ChargeOnce(w http.ResponseWriter, r *http.Request) {

	meseeks := models.Meseeks{
		ID:             1,
		Name:           "Mr. Meseeks",
		Description:    "I'm Mr. Meeseeks, look at me!",
		InventoryLevel: 10,
		Price:          1000,
	}

	data := make(map[string]interface{})
	data["meseeks"] = meseeks

	if err := app.renderTemplate(w, r, "buy-once", &templateData{
		Data: data,
	}, "stripe-js"); err != nil {
		app.errorlog.Println(err)
	}

}
