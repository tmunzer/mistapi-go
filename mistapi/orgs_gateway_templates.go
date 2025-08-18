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

// OrgsGatewayTemplates represents a controller struct.
type OrgsGatewayTemplates struct {
	baseController
}

// NewOrgsGatewayTemplates creates a new instance of OrgsGatewayTemplates.
// It takes a baseController as a parameter and returns a pointer to the OrgsGatewayTemplates.
func NewOrgsGatewayTemplates(baseController baseController) *OrgsGatewayTemplates {
	orgsGatewayTemplates := OrgsGatewayTemplates{baseController: baseController}
	return &orgsGatewayTemplates
}

// ListOrgGatewayTemplates takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.GatewayTemplate data and
// an error if there was an issue with the request or response.
// Get List of Org Gateway Templates
func (o *OrgsGatewayTemplates) ListOrgGatewayTemplates(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.GatewayTemplate],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/gatewaytemplates")
	req.AppendTemplateParams(orgId)
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

	var result []models.GatewayTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.GatewayTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgGatewayTemplate takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.GatewayTemplate data and
// an error if there was an issue with the request or response.
// Create Org Gateway Template
func (o *OrgsGatewayTemplates) CreateOrgGatewayTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.GatewayTemplate) (
	models.ApiResponse[models.GatewayTemplate],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/gatewaytemplates")
	req.AppendTemplateParams(orgId)
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

	var result models.GatewayTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GatewayTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgGatewayTemplate takes context, orgId, gatewaytemplateId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Organization Gateway Template
func (o *OrgsGatewayTemplates) DeleteOrgGatewayTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	gatewaytemplateId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/gatewaytemplates/%v")
	req.AppendTemplateParams(orgId, gatewaytemplateId)
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

// GetOrgGatewayTemplate takes context, orgId, gatewaytemplateId as parameters and
// returns an models.ApiResponse with models.GatewayTemplate data and
// an error if there was an issue with the request or response.
// Get Organization Gateway Template details
func (o *OrgsGatewayTemplates) GetOrgGatewayTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	gatewaytemplateId uuid.UUID) (
	models.ApiResponse[models.GatewayTemplate],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/gatewaytemplates/%v")
	req.AppendTemplateParams(orgId, gatewaytemplateId)
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

	var result models.GatewayTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GatewayTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgGatewayTemplate takes context, orgId, gatewaytemplateId, body as parameters and
// returns an models.ApiResponse with models.GatewayTemplate data and
// an error if there was an issue with the request or response.
// Update Organization Gateway Template
func (o *OrgsGatewayTemplates) UpdateOrgGatewayTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	gatewaytemplateId uuid.UUID,
	body *models.GatewayTemplate) (
	models.ApiResponse[models.GatewayTemplate],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/gatewaytemplates/%v")
	req.AppendTemplateParams(orgId, gatewaytemplateId)
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

	var result models.GatewayTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GatewayTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}
