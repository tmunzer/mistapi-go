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

// OrgsSecurityZones represents a controller struct.
type OrgsSecurityZones struct {
	baseController
}

// NewOrgsSecurityZones creates a new instance of OrgsSecurityZones.
// It takes a baseController as a parameter and returns a pointer to the OrgsSecurityZones.
func NewOrgsSecurityZones(baseController baseController) *OrgsSecurityZones {
	orgsSecurityZones := OrgsSecurityZones{baseController: baseController}
	return &orgsSecurityZones
}

// ListOrgSecurityZones takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.Securityzone data and
// an error if there was an issue with the request or response.
// List organization-level security zones. Security zones are used by SRX gateways, and the zone `name` is used as the security zone name on the device.
func (o *OrgsSecurityZones) ListOrgSecurityZones(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.Securityzone],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/securityzones")
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

	var result []models.Securityzone
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Securityzone](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgSecurityZone takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Securityzone data and
// an error if there was an issue with the request or response.
// Create an organization-level security zone for SRX gateways. Networks reference a zone through their `zone_id`, and several networks can share the same zone.
func (o *OrgsSecurityZones) CreateOrgSecurityZone(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Securityzone) (
	models.ApiResponse[models.Securityzone],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/securityzones")
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

	var result models.Securityzone
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Securityzone](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgSecurityZone takes context, orgId, securityzoneId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete an organization security zone by ID so it can no longer be referenced by a Network `zone_id`.
func (o *OrgsSecurityZones) DeleteOrgSecurityZone(
	ctx context.Context,
	orgId uuid.UUID,
	securityzoneId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/securityzones/%v")
	req.AppendTemplateParams(orgId, securityzoneId)
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

// GetOrgSecurityZone takes context, orgId, securityzoneId as parameters and
// returns an models.ApiResponse with models.Securityzone data and
// an error if there was an issue with the request or response.
// Retrieve details for a specific organization security zone.
func (o *OrgsSecurityZones) GetOrgSecurityZone(
	ctx context.Context,
	orgId uuid.UUID,
	securityzoneId uuid.UUID) (
	models.ApiResponse[models.Securityzone],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/securityzones/%v")
	req.AppendTemplateParams(orgId, securityzoneId)
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

	var result models.Securityzone
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Securityzone](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgSecurityZone takes context, orgId, securityzoneId, body as parameters and
// returns an models.ApiResponse with models.Securityzone data and
// an error if there was an issue with the request or response.
// Update an organization security zone. Renaming a zone changes the security zone name used on the device for every Network referencing it.
func (o *OrgsSecurityZones) UpdateOrgSecurityZone(
	ctx context.Context,
	orgId uuid.UUID,
	securityzoneId uuid.UUID,
	body *models.Securityzone) (
	models.ApiResponse[models.Securityzone],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/securityzones/%v")
	req.AppendTemplateParams(orgId, securityzoneId)
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

	var result models.Securityzone
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Securityzone](decoder)
	return models.NewApiResponse(result, resp), err
}
