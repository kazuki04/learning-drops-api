package handler

import (
	"encoding/json"
	"github.com/learning-drops-api/app/models"
	// "github.com/learning-drops-api/middleware"
	"github.com/learning-drops-api/app/api"
	"net/http"
)

func SaveSectionHandler(w http.ResponseWriter, r *http.Request) {
	var section = models.Section{}
	json.NewDecoder(r.Body).Decode(&section)
	section.Create()
	api.ResponseJSON(section, w, http.StatusOK)
}
