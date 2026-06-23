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
// List organization gateway templates, which provide reusable WAN gateway configuration that can be applied to gateways at sites.
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
// Create an organization gateway template with reusable WAN gateway networks, ports, routing, and service-policy configuration.
// Gateway templates can be applied to multiple sites within the organization to provide consistent gateway configuration across sites.
// To assign a gateway template to a site, use the [Update Site]($e/Sites/updateSiteInfo) endpoint and specify the gateway template ID in the `gatewaytemplate_id` field of the request body.
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
// Delete an organization gateway template and remove that reusable gateway configuration object from the organization.
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

// GetOrgGatewayTemplate takes context, orgId, gatewaytemplateId as parameters and
// returns an models.ApiResponse with models.GatewayTemplate data and
// an error if there was an issue with the request or response.
// Retrieve the configuration stored in a specific organization gateway template.
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
// Update the configuration stored in an organization gateway template.
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

	var result models.GatewayTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GatewayTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}
