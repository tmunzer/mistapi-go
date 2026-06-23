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

// OrgsMarvisInvites represents a controller struct.
type OrgsMarvisInvites struct {
	baseController
}

// NewOrgsMarvisInvites creates a new instance of OrgsMarvisInvites.
// It takes a baseController as a parameter and returns a pointer to the OrgsMarvisInvites.
func NewOrgsMarvisInvites(baseController baseController) *OrgsMarvisInvites {
	orgsMarvisInvites := OrgsMarvisInvites{baseController: baseController}
	return &orgsMarvisInvites
}

// ListOrgMarvisClientInvites takes context, orgId as parameters and
// returns an models.ApiResponse with []models.MarvisClient data and
// an error if there was an issue with the request or response.
// List Marvis Client onboarding invite profiles for the organization, including enrollment URLs and enabled telemetry, location, and synthetic-test capabilities.
func (o *OrgsMarvisInvites) ListOrgMarvisClientInvites(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[[]models.MarvisClient],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/marvisinvites")
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

	var result []models.MarvisClient
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.MarvisClient](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgMarvisClientInvite takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.MarvisClient data and
// an error if there was an issue with the request or response.
// Create a Marvis Client onboarding invite profile for the organization, defining the enabled telemetry, location, and synthetic-test capabilities.
func (o *OrgsMarvisInvites) CreateOrgMarvisClientInvite(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.MarvisClient) (
	models.ApiResponse[models.MarvisClient],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/marvisinvites")
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

	var result models.MarvisClient
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MarvisClient](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgMarvisClientInvite takes context, orgId, marvisinviteId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete a Marvis Client onboarding invite profile by ID.
func (o *OrgsMarvisInvites) DeleteOrgMarvisClientInvite(
	ctx context.Context,
	orgId uuid.UUID,
	marvisinviteId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/marvisinvites/%v")
	req.AppendTemplateParams(orgId, marvisinviteId)
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

// GetOrgMarvisClientInvite takes context, orgId, marvisinviteId as parameters and
// returns an models.ApiResponse with models.MarvisClient data and
// an error if there was an issue with the request or response.
// Retrieve a Marvis Client onboarding invite profile, including enrollment URL and enabled client capabilities.
func (o *OrgsMarvisInvites) GetOrgMarvisClientInvite(
	ctx context.Context,
	orgId uuid.UUID,
	marvisinviteId uuid.UUID) (
	models.ApiResponse[models.MarvisClient],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/marvisinvites/%v")
	req.AppendTemplateParams(orgId, marvisinviteId)
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

	var result models.MarvisClient
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MarvisClient](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgMarvisClientInvite takes context, orgId, marvisinviteId, body as parameters and
// returns an models.ApiResponse with models.MarvisClient data and
// an error if there was an issue with the request or response.
// Update a Marvis Client onboarding invite profile, including enabled telemetry, location, and synthetic-test capabilities.
func (o *OrgsMarvisInvites) UpdateOrgMarvisClientInvite(
	ctx context.Context,
	orgId uuid.UUID,
	marvisinviteId uuid.UUID,
	body *models.MarvisClient) (
	models.ApiResponse[models.MarvisClient],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/marvisinvites/%v")
	req.AppendTemplateParams(orgId, marvisinviteId)
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

	var result models.MarvisClient
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MarvisClient](decoder)
	return models.NewApiResponse(result, resp), err
}
