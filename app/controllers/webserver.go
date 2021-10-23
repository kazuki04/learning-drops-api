package controllers

import (
	"fmt"
	"net/http"

	"github.com/learning-drops-api/config"
	"github.com/learning-drops-api/router"

	"github.com/rs/cors"
)

var corsWrapper = cors.New(cors.Options{
	AllowedOrigins: []string{"http://localhost:3000"},
	AllowedMethods: []string{"GET", "POST"},
	AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
})

func StartWebServer() error {
	r := router.New()
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), corsWrapper.Handler(r))
}
