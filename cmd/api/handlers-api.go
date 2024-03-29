package main

import (
	"encoding/json"
	"net/http"
	"pickle_ricks_back/internal/cards"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
	Content string `json:"content,omitempty"`
	ID      int    `json:"id,omitempty"`
}

func (app *application) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {

	var payload stripePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	amount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	card := cards.Card{
		Secret:   app.config.stripe.secret,
		Key:      app.config.stripe.key,
		Currency: payload.Currency,
	}

	okay := true

	pi, msg, err := card.Charge(payload.Currency, amount)
	if err != nil {
		okay = false
	}

	if okay {
		out, err := json.MarshalIndent(pi, "", "   ")
		if err != nil {
			app.errorlog.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	} else {
		jsn := jsonResponse{

			OK:      false,
			Message: msg,
			Content: "",
		}
		out, err := json.MarshalIndent(jsn, "", "   ")
		if err != nil {
			app.errorlog.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)

	}

}

func (app *application) GetMeseeksByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	meseeksID, _ := strconv.Atoi(id)
	meseeks, err := app.DB.GetMeseeks(meseeksID)
	if err != nil {
		app.errorlog.Println(err)
		return
	}
	out, err := json.MarshalIndent(meseeks, "", "   ")
	if err != nil {
		app.errorlog.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return
}
