package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jailcomfranssa/REST_api-gorm/models"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) AddPersonalidade(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var personalidades models.Personalidade
	json.Unmarshal(body, &personalidades)
	if result := h.DB.Create(&personalidades); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(personalidades)

}
