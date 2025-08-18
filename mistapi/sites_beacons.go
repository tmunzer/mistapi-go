// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"net/http"
)

// SitesBeacons represents a controller struct.
type SitesBeacons struct {
	baseController
}

// NewSitesBeacons creates a new instance of SitesBeacons.
// It takes a baseController as a parameter and returns a pointer to the SitesBeacons.
func NewSitesBeacons(baseController baseController) *SitesBeacons {
	sitesBeacons := SitesBeacons{baseController: baseController}
	return &sitesBeacons
}

// ListSiteBeacons takes context, siteId, limit, page as parameters and
// returns an models.ApiResponse with []models.Beacon data and
// an error if there was an issue with the request or response.
// Get List of Site Beacons
func (s *SitesBeacons) ListSiteBeacons(
	ctx context.Context,
	siteId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.Beacon],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/beacons")
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.Beacon
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Beacon](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateSiteBeacon takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.Beacon data and
// an error if there was an issue with the request or response.
// Create Site Beacon
func (s *SitesBeacons) CreateSiteBeacon(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.Beacon) (
	models.ApiResponse[models.Beacon],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/beacons")
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.Beacon
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Beacon](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteSiteBeacon takes context, siteId, beaconId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site Beacon
func (s *SitesBeacons) DeleteSiteBeacon(
	ctx context.Context,
	siteId uuid.UUID,
	beaconId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/beacons/%v")
	req.AppendTemplateParams(siteId, beaconId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// GetSiteBeacon takes context, siteId, beaconId as parameters and
// returns an models.ApiResponse with models.Beacon data and
// an error if there was an issue with the request or response.
// Get Site Beacon Details
func (s *SitesBeacons) GetSiteBeacon(
	ctx context.Context,
	siteId uuid.UUID,
	beaconId uuid.UUID) (
	models.ApiResponse[models.Beacon],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/beacons/%v")
	req.AppendTemplateParams(siteId, beaconId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})

	var result models.Beacon
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Beacon](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateSiteBeacon takes context, siteId, beaconId, body as parameters and
// returns an models.ApiResponse with models.Beacon data and
// an error if there was an issue with the request or response.
// Update Site Beacon
func (s *SitesBeacons) UpdateSiteBeacon(
	ctx context.Context,
	siteId uuid.UUID,
	beaconId uuid.UUID,
	body *models.Beacon) (
	models.ApiResponse[models.Beacon],
	error) {
	req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/beacons/%v")
	req.AppendTemplateParams(siteId, beaconId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
			NewAndAuth(
				NewAuth("basicAuth"),
				NewAuth("csrfToken"),
			),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.Beacon
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Beacon](decoder)
	return models.NewApiResponse(result, resp), err
}
