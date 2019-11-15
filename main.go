package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lukaszozimek/stock-exchange-service/controller"
	"github.com/lukaszozimek/stock-exchange-service/controller/financial_instruments_controller"
	"github.com/lukaszozimek/stock-exchange-service/security"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()
	router.Use(security.JwtAuthentication) //attach JWT auth middleware

	router.HandleFunc("/api/v1/user/register", controller.CreateAccount).Methods("POST")
	router.HandleFunc("/api/v1/user/login", controller.Authenticate).Methods("POST")
	router.HandleFunc("/api/v1/user/{id}/contact", controller.CreateContact).Methods("POST")
	router.HandleFunc("/api/v1/user/{id}/contact", controller.GetContactsFor).Methods("GET")

	router.HandleFunc("/api/v1/stock", financial_instruments.FindAllStock).Methods("GET")
	router.HandleFunc("/api/v1/stock/{id}", financial_instruments.FindOneStock).Methods("GET")
	router.HandleFunc("/api/v1/stock", financial_instruments.CreateStock).Methods("POST")
	router.HandleFunc("/api/v1/stock", financial_instruments.UpdateStock).Methods("PUT")
	router.HandleFunc("/api/v1/stock/{id}", financial_instruments.DeleteStock).Methods("DELETE")

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
