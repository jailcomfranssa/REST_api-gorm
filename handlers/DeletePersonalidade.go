package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jailcomfranssa/REST_api-gorm/models"
	"net/http"
	"strconv"
)

func (h handler) DeletePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var personalidade models.Personalidade

	if result := h.DB.First(&personalidade, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	h.DB.Delete(&personalidade)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
