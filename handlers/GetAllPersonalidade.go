package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jailcomfranssa/REST_api-gorm/models"
	"net/http"
)

func (h handler) GetAllPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade

	if result := h.DB.Find(&personalidades); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(personalidades)
}
