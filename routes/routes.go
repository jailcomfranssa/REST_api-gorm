package routes

import (
	"github.com/gorilla/mux"
	"github.com/jailcomfranssa/REST_api-gorm/database"

	"github.com/jailcomfranssa/REST_api-gorm/handlers"
	"log"
	"net/http"
)

func HandleResquest() {

	DB := database.ConectaComBancoDeDados()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/personalidade", h.GetAllPersonalidade).Methods(http.MethodGet)
	router.HandleFunc("/personalidade/{id}", h.GetPersonalidade).Methods(http.MethodGet)
	router.HandleFunc("/personalidade", h.AddPersonalidade).Methods(http.MethodPost)
	router.HandleFunc("/personalidade/{id}", h.UpdatePersonalidade).Methods(http.MethodPut)
	router.HandleFunc("/personalidade/{id}", h.DeletePersonalidade).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
