package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type PageData struct {
	PageTitle string
}

func newRouter() *mux.Router {
	//create new router
	router := mux.NewRouter()
	//setup route handlers
	router.HandleFunc("/", homeHandler).Methods("Get")
	router.HandleFunc("/cards", getCardHandler).Methods("Get")
	router.HandleFunc("/cards", createCardHandler).Methods("Post")

	return router
}

func main() {
	//caall method to create new router
	router := newRouter()

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	//start server and log any fatal errors to console
	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("assets/template.html") //parse the html file homepage.html
	if err != nil {                                       // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	//set title of page
	data := PageData{
		PageTitle: "Home",
	}

	err = t.Execute(w, data) //execute the template and pass it the PageData struct to fill in the gaps
	if err != nil {          // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
