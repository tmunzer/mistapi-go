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

// OrgsRFTemplates represents a controller struct.
type OrgsRFTemplates struct {
	baseController
}

// NewOrgsRFTemplates creates a new instance of OrgsRFTemplates.
// It takes a baseController as a parameter and returns a pointer to the OrgsRFTemplates.
func NewOrgsRFTemplates(baseController baseController) *OrgsRFTemplates {
	orgsRFTemplates := OrgsRFTemplates{baseController: baseController}
	return &orgsRFTemplates
}

// ListOrgRfTemplates takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.RfTemplate data and
// an error if there was an issue with the request or response.
// List organization RF templates used by RRM to apply radio settings across sites.
func (o *OrgsRFTemplates) ListOrgRfTemplates(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.RfTemplate],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/rftemplates")
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

	var result []models.RfTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.RfTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgRfTemplate takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.RfTemplate data and
// an error if there was an issue with the request or response.
// Create an organization RF template with 2.4, 5, and 6 GHz radio
// settings, country code, scanning behavior, antenna gain, and model-specific
// overrides.
// To assign a RF template to a site, use the [Update Site]($e/Sites/updateSiteInfo) endpoint and specify the RF template ID in the `rftemplate_id` field of the request body.
func (o *OrgsRFTemplates) CreateOrgRfTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.RfTemplate) (
	models.ApiResponse[models.RfTemplate],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/rftemplates")
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

	var result models.RfTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RfTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgRfTemplate takes context, orgId, rftemplateId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete an organization RF template by template ID so it can no longer be applied to sites.
func (o *OrgsRFTemplates) DeleteOrgRfTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	rftemplateId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/rftemplates/%v")
	req.AppendTemplateParams(orgId, rftemplateId)
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

// GetOrgRfTemplate takes context, orgId, rftemplateId as parameters and
// returns an models.ApiResponse with models.RfTemplate data and
// an error if there was an issue with the request or response.
// Retrieve details for a specific organization RF template, including radio settings, country code, scanning behavior, antenna gain, and model-specific overrides.
func (o *OrgsRFTemplates) GetOrgRfTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	rftemplateId uuid.UUID) (
	models.ApiResponse[models.RfTemplate],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/rftemplates/%v")
	req.AppendTemplateParams(orgId, rftemplateId)
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

	var result models.RfTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RfTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgRfTemplate takes context, orgId, rftemplateId, body as parameters and
// returns an models.ApiResponse with models.RfTemplate data and
// an error if there was an issue with the request or response.
// Update an organization RF template, including radio settings, country code, scanning behavior, antenna gain, and model-specific overrides.
func (o *OrgsRFTemplates) UpdateOrgRfTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	rftemplateId uuid.UUID,
	body *models.RfTemplate) (
	models.ApiResponse[models.RfTemplate],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/rftemplates/%v")
	req.AppendTemplateParams(orgId, rftemplateId)
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

	var result models.RfTemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RfTemplate](decoder)
	return models.NewApiResponse(result, resp), err
}
