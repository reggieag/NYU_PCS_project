/*
 * Toy API
 *
 * Toy API for testing RESTler
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer) Router {
	return &DefaultApiController{ service: s }
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"DataDataIdDelete",
			strings.ToUpper("Delete"),
			"/data/{dataId}",
			c.DataDataIdDelete,
		},
		{
			"DataDataIdGet",
			strings.ToUpper("Get"),
			"/data/{dataId}",
			c.DataDataIdGet,
		},
		{
			"DataDataIdPost",
			strings.ToUpper("Post"),
			"/data/{dataId}",
			c.DataDataIdPost,
		},
		{
			"DataGet",
			strings.ToUpper("Get"),
			"/data",
			c.DataGet,
		},
		{
			"DataPost",
			strings.ToUpper("Post"),
			"/data",
			c.DataPost,
		},
	}
}

// DataDataIdDelete - Delete data by id
func (c *DefaultApiController) DataDataIdDelete(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	dataId, err := parseInt32Parameter(params["dataId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := c.service.DataDataIdDelete(r.Context(), dataId)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// DataDataIdGet - Get data by id
func (c *DefaultApiController) DataDataIdGet(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	dataId, err := parseInt32Parameter(params["dataId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := c.service.DataDataIdGet(r.Context(), dataId)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// DataDataIdPost - Update data by id
func (c *DefaultApiController) DataDataIdPost(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	dataId, err := parseInt32Parameter(params["dataId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	inlineObject1 := &InlineObject1{}
	if err := json.NewDecoder(r.Body).Decode(&inlineObject1); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := c.service.DataDataIdPost(r.Context(), dataId, *inlineObject1)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// DataGet - Returns a list of data.
func (c *DefaultApiController) DataGet(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.DataGet(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// DataPost - Add a new entry
func (c *DefaultApiController) DataPost(w http.ResponseWriter, r *http.Request) { 
	inlineObject := &InlineObject{}
	if err := json.NewDecoder(r.Body).Decode(&inlineObject); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := c.service.DataPost(r.Context(), *inlineObject)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}