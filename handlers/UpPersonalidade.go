package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jailcomfranssa/REST_api-gorm/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (h handler) UpdatePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var updatePersonalidade models.Personalidade
	json.Unmarshal(body, &updatePersonalidade)

	var personalidade models.Personalidade

	if result := h.DB.First(&personalidade, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	if updatePersonalidade.Nome == "" {
		personalidade.Nome = personalidade.Nome
		personalidade.Historia = updatePersonalidade.Historia
	} else if updatePersonalidade.Historia == "" {
		personalidade.Nome = updatePersonalidade.Nome
		personalidade.Historia = personalidade.Historia
	} else {
		personalidade.Nome = updatePersonalidade.Nome
		personalidade.Historia = updatePersonalidade.Historia
	}

	h.DB.Save(&personalidade)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(personalidade)

}
