package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func AuthMiddleware(authHost, authPath string, authPort int) func(http.Handler) http.Handler {
	url := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", authHost, authPort),
		Path:   authPath,
	}
	authUrl := url.String()
	log.Printf("initializing auth middleware using url: %s", authUrl)
	client := getHttpClientWithTimeout()
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("authorizing request")
			token := getBearerToken(r)
			if token == "" {
				http.Error(w, "Token not found", http.StatusForbidden)
				return
			}
			if resp, err := sendIntrospectRequest(r.Context(), client, authUrl, token); err != nil || !resp.Valid {
				http.Error(w, "Token invalid", http.StatusForbidden)
				return
			} else {
				ctxWithScope := withScopes(r.Context(), resp.Scopes)
				rWithScopes := r.WithContext(ctxWithScope)
				next.ServeHTTP(w, rWithScopes)
			}
		})
	}
}

type introspectRequest struct {
	AccessToken string `json:"access_token"`
}

type introspectResponse struct {
	Valid  bool
	Scopes []string
}

func sendIntrospectRequest(ctx context.Context, client *http.Client, authUrl string, token string) (introspectResponse, error) {
	var responseStruct introspectResponse
	requestBodyStruct := introspectRequest{
		AccessToken: token,
	}
	requestBodyBytes, err := json.Marshal(requestBodyStruct)
	if err != nil {
		log.Printf("error marshalling request for introspection: %s", err)
		return responseStruct, err
	}
	requestBodyReader := bytes.NewReader(requestBodyBytes)
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, authUrl, requestBodyReader)
	if err != nil {
		log.Printf("error creating request for introspection: %s", err)
		return responseStruct, err
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Printf("error for introspection: %s", err)
		return responseStruct, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("unable to verify token")
		return responseStruct, fmt.Errorf("unable to verify token. status: %d", resp.StatusCode)
	}
	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading introspection response: %s", err)
		return responseStruct, err
	}
	if err := json.Unmarshal(responseBytes, &responseStruct); err != nil {
		log.Printf("error unmarshalling introspection response: %s", err)
		return responseStruct, err
	}
	return responseStruct, nil
}

func getHttpClientWithTimeout() *http.Client {
	return &http.Client{
		Timeout: time.Second * 3,
	}
}

func getBearerToken(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	splitToken := strings.Split(bearer, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}
	return splitToken[1]
}
