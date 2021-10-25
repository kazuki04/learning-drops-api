package router

import (
	"net/http"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/learning-drops-api/app/api"
	"github.com/learning-drops-api/handler"
	"github.com/learning-drops-api/middleware"
)

func New() *mux.Router {
	jwtMiddleware := midleware.JwtMiddleware()
	r := mux.NewRouter()

	r.HandleFunc("/api/section", api.ApiMakeHandler(handler.SaveSectionHandler))

	// This route is always accessible
	r.Handle("/api/public", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := "Hello from a public endpoint! You don't need to be authenticated to see this."
		midleware.ResponseJSON(message, w, http.StatusOK)
	}))

	// This route is only accessible if the user has a valid Access Token
	// We are chaining the jwtmiddleware middleware into the negroni handler function which will check
	// for a valid token.
	r.Handle("/api/private", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			message := "Hello from a private endpoint! You need to be authenticated to see this."
			midleware.ResponseJSON(message, w, http.StatusOK)
		}))))

	r.Handle("/api/private-scoped", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
			token := authHeaderParts[1]

			hasScope := midleware.CheckScope("read:messages", token)

			if !hasScope {
				message := "Insufficient scope."
				midleware.ResponseJSON(message, w, http.StatusForbidden)
				return
			}
			message := "Hello from a private endpoint! You need to be authenticated to see this."
			midleware.ResponseJSON(message, w, http.StatusOK)
		}))))

	return r
}
