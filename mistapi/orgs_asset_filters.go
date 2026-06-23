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

// OrgsAssetFilters represents a controller struct.
type OrgsAssetFilters struct {
	baseController
}

// NewOrgsAssetFilters creates a new instance of OrgsAssetFilters.
// It takes a baseController as a parameter and returns a pointer to the OrgsAssetFilters.
func NewOrgsAssetFilters(baseController baseController) *OrgsAssetFilters {
	orgsAssetFilters := OrgsAssetFilters{baseController: baseController}
	return &orgsAssetFilters
}

// ListOrgAssetFilters takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.AssetFilter data and
// an error if there was an issue with the request or response.
// List organization-level BLE asset filters. Each filter operates independently, and an asset must match all specified filter properties for that filter to apply.
func (o *OrgsAssetFilters) ListOrgAssetFilters(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.AssetFilter],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/assetfilters")
	req.AppendTemplateParams(orgId)
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

	var result []models.AssetFilter
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.AssetFilter](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgAssetFilter takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.AssetFilter data and
// an error if there was an issue with the request or response.
// Create an organization-level BLE asset filter. Any subset of filter properties can be included, and a matching asset must meet all specified conditions.
func (o *OrgsAssetFilters) CreateOrgAssetFilter(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.AssetFilter) (
	models.ApiResponse[models.AssetFilter],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/assetfilters")
	req.AppendTemplateParams(orgId)
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

	var result models.AssetFilter
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AssetFilter](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgAssetFilter takes context, orgId, assetfilterId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete an organization-level BLE asset filter by filter ID.
func (o *OrgsAssetFilters) DeleteOrgAssetFilter(
	ctx context.Context,
	orgId uuid.UUID,
	assetfilterId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/assetfilters/%v")
	req.AppendTemplateParams(orgId, assetfilterId)
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// GetOrgAssetFilter takes context, orgId, assetfilterId as parameters and
// returns an models.ApiResponse with models.AssetFilter data and
// an error if there was an issue with the request or response.
// Return one organization-level BLE asset filter, including its name, disabled state, and matching criteria.
func (o *OrgsAssetFilters) GetOrgAssetFilter(
	ctx context.Context,
	orgId uuid.UUID,
	assetfilterId uuid.UUID) (
	models.ApiResponse[models.AssetFilter],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/assetfilters/%v")
	req.AppendTemplateParams(orgId, assetfilterId)
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

	var result models.AssetFilter
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AssetFilter](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgAssetFilter takes context, orgId, assetfilterId, body as parameters and
// returns an models.ApiResponse with models.AssetFilter data and
// an error if there was an issue with the request or response.
// Update an organization-level BLE asset filter's name, disabled state, or matching criteria.
func (o *OrgsAssetFilters) UpdateOrgAssetFilter(
	ctx context.Context,
	orgId uuid.UUID,
	assetfilterId uuid.UUID,
	body *models.AssetFilter) (
	models.ApiResponse[models.AssetFilter],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/assetfilters/%v")
	req.AppendTemplateParams(orgId, assetfilterId)
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

	var result models.AssetFilter
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AssetFilter](decoder)
	return models.NewApiResponse(result, resp), err
}
