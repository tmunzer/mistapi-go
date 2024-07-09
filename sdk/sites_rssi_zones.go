package mistapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi/sdk/errors"
	"github.com/tmunzer/mistapi/sdk/models"
)

// SitesRSSIZones represents a controller struct.
type SitesRSSIZones struct {
	baseController
}

// NewSitesRSSIZones creates a new instance of SitesRSSIZones.
// It takes a baseController as a parameter and returns a pointer to the SitesRSSIZones.
func NewSitesRSSIZones(baseController baseController) *SitesRSSIZones {
	sitesRSSIZones := SitesRSSIZones{baseController: baseController}
	return &sitesRSSIZones
}

// ListSiteRssiZones takes context, siteId, page, limit as parameters and
// returns an models.ApiResponse with []models.RssiZone data and
// an error if there was an issue with the request or response.
// Get List of Site RSSI Zone (RSSI-based)
func (s *SitesRSSIZones) ListSiteRssiZones(
	ctx context.Context,
	siteId uuid.UUID,
	page *int,
	limit *int) (
	models.ApiResponse[[]models.RssiZone],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/rssizones", siteId),
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
	if page != nil {
		req.QueryParam("page", *page)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}

	var result []models.RssiZone
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.RssiZone](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateSiteRssiZone takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.RssiZone data and
// an error if there was an issue with the request or response.
// Create RSSI Zone
func (s *SitesRSSIZones) CreateSiteRssiZone(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.RssiZone) (
	models.ApiResponse[models.RssiZone],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/rssizones", siteId),
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

	var result models.RssiZone
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RssiZone](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteSiteRssiZone takes context, siteId, rssizoneId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete Site RSSI Zone
func (s *SitesRSSIZones) DeleteSiteRssiZone(
	ctx context.Context,
	siteId uuid.UUID,
	rssizoneId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/sites/%v/rssizones/%v", siteId, rssizoneId),
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

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// GetSiteRssiZone takes context, siteId, rssizoneId as parameters and
// returns an models.ApiResponse with []models.RssiZone data and
// an error if there was an issue with the request or response.
// Get Site RSSI Zone details
func (s *SitesRSSIZones) GetSiteRssiZone(
	ctx context.Context,
	siteId uuid.UUID,
	rssizoneId uuid.UUID) (
	models.ApiResponse[[]models.RssiZone],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/rssizones/%v", siteId, rssizoneId),
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

	var result []models.RssiZone
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.RssiZone](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateSiteRssiZone takes context, siteId, rssizoneId, body as parameters and
// returns an models.ApiResponse with models.RssiZone data and
// an error if there was an issue with the request or response.
// Update Site RSSI Zone
func (s *SitesRSSIZones) UpdateSiteRssiZone(
	ctx context.Context,
	siteId uuid.UUID,
	rssizoneId uuid.UUID,
	body *models.RssiZone) (
	models.ApiResponse[models.RssiZone],
	error) {
	req := s.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/sites/%v/rssizones/%v", siteId, rssizoneId),
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

	var result models.RssiZone
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RssiZone](decoder)
	return models.NewApiResponse(result, resp), err
}
