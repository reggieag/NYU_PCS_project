package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/go-oauth2/oauth2/v4"
)

type RequestStruct struct {
	AccessToken string `json:"access_token"`
}

type ResponseStruct struct {
	Valid  bool
	Scopes []string
}

func createIntrospectHandler(tokenStore oauth2.TokenStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("introspect called")
		ctx := r.Context()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var requestBody RequestStruct
		if err := json.Unmarshal(body, &requestBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := tokenStore.GetByAccess(ctx, requestBody.AccessToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if token == nil {
			http.Error(w, "Token not found", http.StatusNotFound)
			return
		}
		var response ResponseStruct
		if expireTime := token.GetAccessExpiresIn(); expireTime < 0 {
			log.Printf("token expire time: %d", expireTime)
			response.Valid = false
		} else {
			response.Valid = true
		}
		response.Scopes = strings.Split(token.GetScope(), ",")
		responseBody, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBody)
	}
}
