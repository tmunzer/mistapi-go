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

// OrgsNetworkTemplates represents a controller struct.
type OrgsNetworkTemplates struct {
	baseController
}

// NewOrgsNetworkTemplates creates a new instance of OrgsNetworkTemplates.
// It takes a baseController as a parameter and returns a pointer to the OrgsNetworkTemplates.
func NewOrgsNetworkTemplates(baseController baseController) *OrgsNetworkTemplates {
	orgsNetworkTemplates := OrgsNetworkTemplates{baseController: baseController}
	return &orgsNetworkTemplates
}

// ListOrgNetworkTemplates takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.NetworkTemplate data and
// an error if there was an issue with the request or response.
// List organization network templates that provide switch network, port, management, routing, NAC, and service configuration at the organization level.
func (o *OrgsNetworkTemplates) ListOrgNetworkTemplates(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.NetworkTemplate],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/networktemplates")
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

	var result []models.NetworkTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.NetworkTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgNetworkTemplate takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.NetworkTemplate data and
// an error if there was an issue with the request or response.
// Create an organization network template with network, port usage, switch management, routing, NAC, and service  configuration at the organization level.
// Network templates can be applied to multiple sites within the organization to provide consistent network configuration across sites.
// To assign a network template to a site, use the [Update Site]($e/Sites/updateSiteInfo) endpoint and specify the network template ID in the `networktemplate_id` field of the request body.
func (o *OrgsNetworkTemplates) CreateOrgNetworkTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.NetworkTemplate) (
	models.ApiResponse[models.NetworkTemplate],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/networktemplates")
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

	var result models.NetworkTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.NetworkTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgNetworkTemplate takes context, orgId, networktemplateId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete an organization network template by template ID so it can no longer be applied to sites or site groups.
func (o *OrgsNetworkTemplates) DeleteOrgNetworkTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	networktemplateId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/networktemplates/%v")
	req.AppendTemplateParams(orgId, networktemplateId)
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

// GetOrgNetworkTemplate takes context, orgId, networktemplateId as parameters and
// returns an models.ApiResponse with models.NetworkTemplate data and
// an error if there was an issue with the request or response.
// Retrieve details for a specific organization network template, including network, port usage, switch management, routing, NAC, and service defaults.
func (o *OrgsNetworkTemplates) GetOrgNetworkTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	networktemplateId uuid.UUID) (
	models.ApiResponse[models.NetworkTemplate],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/networktemplates/%v")
	req.AppendTemplateParams(orgId, networktemplateId)
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

	var result models.NetworkTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.NetworkTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgNetworkTemplate takes context, orgId, networktemplateId, body as parameters and
// returns an models.ApiResponse with models.NetworkTemplate data and
// an error if there was an issue with the request or response.
// Update an organization network template, including network, port usage, switch management, routing, NAC, and service defaults.
func (o *OrgsNetworkTemplates) UpdateOrgNetworkTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	networktemplateId uuid.UUID,
	body *models.NetworkTemplate) (
	models.ApiResponse[models.NetworkTemplate],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/networktemplates/%v")
	req.AppendTemplateParams(orgId, networktemplateId)
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

	var result models.NetworkTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.NetworkTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}
