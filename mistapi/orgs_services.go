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

// OrgsServices represents a controller struct.
type OrgsServices struct {
	baseController
}

// NewOrgsServices creates a new instance of OrgsServices.
// It takes a baseController as a parameter and returns a pointer to the OrgsServices.
func NewOrgsServices(baseController baseController) *OrgsServices {
	orgsServices := OrgsServices{baseController: baseController}
	return &orgsServices
}

// ListOrgServices takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.Service data and
// an error if there was an issue with the request or response.
// List organization service definitions. Services describe applications, application categories, URLs, hostnames, subnets, or custom protocol and port match criteria used by gateway and SSR policies.
func (o *OrgsServices) ListOrgServices(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.Service],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/services")
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

	var result []models.Service
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Service](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgService takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Service data and
// an error if there was an issue with the request or response.
// Create an organization service definition with the match criteria
// used by gateway and SSR policies, such as applications, URLs, hostnames,
// subnets, or custom protocol and port rules.
// Services can be user in the service policies to allow or deny traffic matching the service or to apply specific inspection settings or steering rules.
func (o *OrgsServices) CreateOrgService(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Service) (
	models.ApiResponse[models.Service],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/services")
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

	var result models.Service
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Service](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgService takes context, orgId, serviceId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Remove an organization service definition from the available service catalog.
func (o *OrgsServices) DeleteOrgService(
	ctx context.Context,
	orgId uuid.UUID,
	serviceId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/services/%v")
	req.AppendTemplateParams(orgId, serviceId)
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

// GetOrgService takes context, orgId, serviceId as parameters and
// returns an models.ApiResponse with models.Service data and
// an error if there was an issue with the request or response.
// Return an organization service definition, including its matching mode, match values, traffic classification, and optional SSR path-selection thresholds.
func (o *OrgsServices) GetOrgService(
	ctx context.Context,
	orgId uuid.UUID,
	serviceId uuid.UUID) (
	models.ApiResponse[models.Service],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/services/%v")
	req.AppendTemplateParams(orgId, serviceId)
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

	var result models.Service
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Service](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgService takes context, orgId, serviceId, body as parameters and
// returns an models.ApiResponse with models.Service data and
// an error if there was an issue with the request or response.
// Update an organization service definition, including its matching mode, match values, traffic classification, or optional SSR path-selection thresholds.
func (o *OrgsServices) UpdateOrgService(
	ctx context.Context,
	orgId uuid.UUID,
	serviceId uuid.UUID,
	body *models.Service) (
	models.ApiResponse[models.Service],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/services/%v")
	req.AppendTemplateParams(orgId, serviceId)
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

	var result models.Service
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Service](decoder)
	return models.NewApiResponse(result, resp), err
}
