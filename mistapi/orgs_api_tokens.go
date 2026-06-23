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

// OrgsAPITokens represents a controller struct.
type OrgsAPITokens struct {
	baseController
}

// NewOrgsAPITokens creates a new instance of OrgsAPITokens.
// It takes a baseController as a parameter and returns a pointer to the OrgsAPITokens.
func NewOrgsAPITokens(baseController baseController) *OrgsAPITokens {
	orgsAPITokens := OrgsAPITokens{baseController: baseController}
	return &orgsAPITokens
}

// ListOrgApiTokens takes context, orgId as parameters and
// returns an models.ApiResponse with []models.OrgApitoken data and
// an error if there was an issue with the request or response.
// List organization API tokens, including display names, scoped privileges, allowed source IPs, creator, and last-use metadata.
func (o *OrgsAPITokens) ListOrgApiTokens(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[[]models.OrgApitoken],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/apitokens")
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

	var result []models.OrgApitoken
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.OrgApitoken](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgApiToken takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.OrgApitoken data and
// an error if there was an issue with the request or response.
// Create an organization API token with a display name, scoped privileges, and optional source IP restrictions.
// Note that the full token key is only available at creation time.
func (o *OrgsAPITokens) CreateOrgApiToken(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.OrgApitoken) (
	models.ApiResponse[models.OrgApitoken],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/apitokens")
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

	var result models.OrgApitoken
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.OrgApitoken](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgApiToken takes context, orgId, apitokenId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete an organization API token so it can no longer authenticate API requests.
func (o *OrgsAPITokens) DeleteOrgApiToken(
	ctx context.Context,
	orgId uuid.UUID,
	apitokenId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/apitokens/%v")
	req.AppendTemplateParams(orgId, apitokenId)
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

// GetOrgApiToken takes context, orgId, apitokenId as parameters and
// returns an models.ApiResponse with models.OrgApitoken data and
// an error if there was an issue with the request or response.
// Return metadata for one organization API token. The full token key is only available at creation time and may only be partially shown afterward.
func (o *OrgsAPITokens) GetOrgApiToken(
	ctx context.Context,
	orgId uuid.UUID,
	apitokenId uuid.UUID) (
	models.ApiResponse[models.OrgApitoken],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/apitokens/%v")
	req.AppendTemplateParams(orgId, apitokenId)
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

	var result models.OrgApitoken
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.OrgApitoken](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgApiToken takes context, orgId, apitokenId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Update an organization API token's display name or scoped privileges. Source IP restrictions are defined when the token is created.
func (o *OrgsAPITokens) UpdateOrgApiToken(
	ctx context.Context,
	orgId uuid.UUID,
	apitokenId uuid.UUID,
	body *models.OrgApitoken) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/apitokens/%v")
	req.AppendTemplateParams(orgId, apitokenId)
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}
