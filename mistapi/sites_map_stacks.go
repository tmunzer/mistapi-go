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
)

// SitesMapStacks represents a controller struct.
type SitesMapStacks struct {
	baseController
}

// NewSitesMapStacks creates a new instance of SitesMapStacks.
// It takes a baseController as a parameter and returns a pointer to the SitesMapStacks.
func NewSitesMapStacks(baseController baseController) *SitesMapStacks {
	sitesMapStacks := SitesMapStacks{baseController: baseController}
	return &sitesMapStacks
}

// ListSiteMapStacks takes context, siteId, limit, page, name as parameters and
// returns an models.ApiResponse with []models.MapstackResponse data and
// an error if there was an issue with the request or response.
// Get List of Site Map Stacks
func (s *SitesMapStacks) ListSiteMapStacks(
	ctx context.Context,
	siteId uuid.UUID,
	limit *int,
	page *int,
	name *string) (
	models.ApiResponse[[]models.MapstackResponse],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/mapstacks")
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}
	if name != nil {
		req.QueryParam("name", *name)
	}

	var result []models.MapstackResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.MapstackResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateSiteMapStack takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.MapstackResponse data and
// an error if there was an issue with the request or response.
// Create Site Map Stack
func (s *SitesMapStacks) CreateSiteMapStack(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.Mapstack) (
	models.ApiResponse[models.MapstackResponse],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/mapstacks")
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.MapstackResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MapstackResponse](decoder)
	return models.NewApiResponse(result, resp), err
}
