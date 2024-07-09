package mistapigo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapigo/sdk/errors"
	"github.com/tmunzer/mistapigo/sdk/models"
)

// SitesMapsAutoOrientation represents a controller struct.
type SitesMapsAutoOrientation struct {
	baseController
}

// NewSitesMapsAutoOrientation creates a new instance of SitesMapsAutoOrientation.
// It takes a baseController as a parameter and returns a pointer to the SitesMapsAutoOrientation.
func NewSitesMapsAutoOrientation(baseController baseController) *SitesMapsAutoOrientation {
	sitesMapsAutoOrientation := SitesMapsAutoOrientation{baseController: baseController}
	return &sitesMapsAutoOrientation
}

// DeleteSiteApAutoOrientation takes context, mapId, siteId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This API is called to force stop auto placement for a given map
func (s *SitesMapsAutoOrientation) DeleteSiteApAutoOrientation(
	ctx context.Context,
	mapId uuid.UUID,
	siteId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/sites/%v/maps/%v/auto_orient", siteId, mapId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Autoplacement was not triggered"},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// StartSiteApAutoOrientation takes context, mapId, siteId, body as parameters and
// returns an models.ApiResponse with models.ResponseAutoOrientation data and
// an error if there was an issue with the request or response.
// This API is called to trigger a map for auto orientation. For auto orient feature to work, BLE data needs to be collected from the APs on the map. This precess is not disruptive unlike FTM collection. Repeated POST to this endpoint while a map is still running will be rejected.
// List of devices to provide suggestions for is an optional parameter that can be given to this API. This will provide auto orient suggestions only for the devices specified. If no list of devices is provided, all APs asociated with that map are considered by default
func (s *SitesMapsAutoOrientation) StartSiteApAutoOrientation(
	ctx context.Context,
	mapId uuid.UUID,
	siteId uuid.UUID,
	body *models.AutoOrient) (
	models.ApiResponse[models.ResponseAutoOrientation],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/maps/%v/auto_orient", siteId, mapId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Request", Unmarshaller: errors.NewResponseDetailString},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.ResponseAutoOrientation
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseAutoOrientation](decoder)
	return models.NewApiResponse(result, resp), err
}

// ClearSiteApAutoOrient takes context, siteId, mapId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This API is used to destroy the autoorientations of a map or subset of APs on a map.
func (s *SitesMapsAutoOrientation) ClearSiteApAutoOrient(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID,
	body *models.MacAddresses) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/maps/%v/clear_auto_orient", siteId, mapId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}
