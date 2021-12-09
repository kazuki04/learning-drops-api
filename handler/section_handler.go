package handler

import (
	"encoding/json"
	"github.com/learning-drops-api/app/api"
	"github.com/learning-drops-api/app/models"
	"net/http"
)

func SaveSectionHandler(w http.ResponseWriter, r *http.Request) {
	var section = models.Section{}
	json.NewDecoder(r.Body).Decode(&section)
	section.Create()
	api.ResponseJSON(section, w, http.StatusOK)
}

func GetSectionsHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	sections := models.GetSections(userId)
	api.ResponseJSON(sections, w, http.StatusOK)
}
