package controllers

import (
	"fmt"
	"github.com/learning-drops-api/config"
	"net/http"

	"github.com/gorilla/mux"
)

func StartWebServer() error {
	r := mux.NewRouter()
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), r)
}
