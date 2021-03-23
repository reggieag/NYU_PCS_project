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
	"context"
	"log"
	"net/http"
)

// DefaultApiService is a service that implents the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
	db ToyAPIDB
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService(db ToyAPIDB) DefaultApiServicer {
	return &DefaultApiService{
		db,
	}
}

// DataDataIdDelete - Delete data by id
func (s *DefaultApiService) DataDataIdDelete(ctx context.Context, dataId int32) (ImplResponse, error) {
	log.Printf("deleting row with id %d", dataId)
	err := s.db.DeleteDataByID(ctx, int(dataId))
	if err != nil {
		log.Printf("unable to delete: %s", err)
		return Response(http.StatusNotFound, nil), nil
	}
	log.Printf("deleted row")
	return Response(http.StatusOK, nil), nil
}

// DataDataIdGet - Get data by id
func (s *DefaultApiService) DataDataIdGet(ctx context.Context, dataId int32) (ImplResponse, error) {
	log.Printf("getting row with id %d", dataId)
	resp, err := s.db.GetDataByID(ctx, int(dataId))
	if err != nil {
		log.Printf("error fetching data: %s", err)
		return Response(http.StatusNotFound, nil), nil
	}
	log.Printf("fetched data")
	return Response(http.StatusOK, resp), nil
}

// DataDataIdPost - Update data by id
func (s *DefaultApiService) DataDataIdPost(ctx context.Context, dataId int32, inlineObject1 InlineObject1) (ImplResponse, error) {
	log.Printf("updating row with id %d with name %s, quantity %d", dataId, inlineObject1.Name, inlineObject1.Quantity)
	if inlineObject1.Name == "" && inlineObject1.Quantity == 0 {
		log.Printf("both name and quantity are invalid")
		return Response(http.StatusBadRequest, nil), nil
	}
	log.Printf("updating db")
	resp, err := s.db.UpdateDataByID(ctx, int(dataId), inlineObject1.Name, int(inlineObject1.Quantity))
	if err != nil {
		log.Printf("error updating db: %s", err)
		return Response(http.StatusNotFound, nil), nil
	}
	log.Printf("updated db")
	return Response(http.StatusOK, resp), nil
}

// DataGet - Returns a list of data.
func (s *DefaultApiService) DataGet(ctx context.Context) (ImplResponse, error) {
	log.Printf("receiving get all data request")
	resp, err := s.db.GetData(ctx)
	if err != nil {
		log.Printf("erro getting data: %s", err)
		return Response(http.StatusInternalServerError, nil), err
	}
	log.Printf("got %d rows back", len(resp))
	return Response(http.StatusOK, resp), nil
}

// DataPost - Add a new entry
func (s *DefaultApiService) DataPost(ctx context.Context, inlineObject InlineObject) (ImplResponse, error) {
	log.Printf("receiving add data with name %s quantity %d", inlineObject.Name, inlineObject.Quantity)
	if inlineObject.Name == "" && inlineObject.Quantity == 0 {
		log.Printf("both fields are 'null'. Bad request.")
		return Response(http.StatusBadRequest, nil), nil
	}
	log.Printf("adding data to db")
	resp, err := s.db.AddData(ctx, inlineObject.Name, int(inlineObject.Quantity))
	if err != nil {
		log.Printf("error adding to db: %s", err)
		return Response(http.StatusInternalServerError, nil), err
	}
	log.Printf("data added")

	return Response(http.StatusOK, resp), nil
}
