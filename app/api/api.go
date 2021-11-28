package api

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

type JSONError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

type JSONResponse struct {
	JSONData interface{} `json:"json_response"`
}

var apiValidPath = regexp.MustCompile("^/api/sections$")

func ApiError(w http.ResponseWriter, errMessage string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonError, err := json.Marshal(JSONError{Error: errMessage, Code: code})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)
}

func ApiMakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := apiValidPath.FindStringSubmatch(r.URL.Path)
		if len(m) == 0 {
			ApiError(w, "Not found", http.StatusNotFound)
		}
		fn(w, r)
	}
}

func ResponseJSON(anyData interface{}, w http.ResponseWriter, statusCode int) {
	jsonResponse, err := json.Marshal(anyData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
