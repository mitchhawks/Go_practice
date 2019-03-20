package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Card struct {
	Name    string `json:"name"`
	Barcode string `json:"barcode"`
	Type    string `json:"type"`
}

var cards []Card

func getCardHandler(w http.ResponseWriter, r *http.Request) {
	cardListBytes, err := json.Marshal(cards)
	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of cards to the response
	w.Write(cardListBytes)
}
func createCardHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Card
	card := Card{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the card from the form info
	card.Name = r.Form.Get("name")
	card.Barcode = r.Form.Get("barcode")
	card.Type = r.Form.Get("type")

	// Append our existing list of cards with a new entry
	cards = append(cards, card)

	//Finally, we redirect the user to the original HTMl page
	http.Redirect(w, r, "/", http.StatusFound)
}
