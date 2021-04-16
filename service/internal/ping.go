package openapi

import (
	"encoding/json"
	"log"
	"net/http"
)

type pingController struct {
}

func (p *pingController) Routes() Routes {
	return Routes{
		{
			"Ping",
			http.MethodGet,
			"/ping",
			pingHandlerFunc,
		},
	}
}

func (p *pingController) Middleware() func(http.Handler) http.Handler {
	return nil
}

func NewPingController() Router {
	return &pingController{}
}

type statusResponse struct {
	Ok bool
}

func pingHandlerFunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("ping request received")
	var response statusResponse
	response.Ok = true
	responseBody, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "no good", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)

}
