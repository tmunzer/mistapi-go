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

// OrgsServicePolicies represents a controller struct.
type OrgsServicePolicies struct {
	baseController
}

// NewOrgsServicePolicies creates a new instance of OrgsServicePolicies.
// It takes a baseController as a parameter and returns a pointer to the OrgsServicePolicies.
func NewOrgsServicePolicies(baseController baseController) *OrgsServicePolicies {
	orgsServicePolicies := OrgsServicePolicies{baseController: baseController}
	return &orgsServicePolicies
}

// ListOrgServicePolicies takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.OrgServicePolicy data and
// an error if there was an issue with the request or response.
// List organization-level service policies. Service policies match tenants to services or service groups and define the allow or deny action plus optional inspection controls.
func (o *OrgsServicePolicies) ListOrgServicePolicies(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.OrgServicePolicy],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/servicepolicies")
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

	var result []models.OrgServicePolicy
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.OrgServicePolicy](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgServicePolicy takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.OrgServicePolicy data and
// an error if there was an issue with the request or response.
// Create an organization-level service policy that matches tenants
// to services and applies an allow or deny action with optional security inspection
// settings.
// Organization-level service policies can be imported in the gateway templates and gateway policies.
func (o *OrgsServicePolicies) CreateOrgServicePolicy(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.OrgServicePolicy) (
	models.ApiResponse[models.OrgServicePolicy],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/servicepolicies")
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

	var result models.OrgServicePolicy
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.OrgServicePolicy](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgServicePolicy takes context, orgId, servicepolicyId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Remove an organization-level service policy from the available policy set.
func (o *OrgsServicePolicies) DeleteOrgServicePolicy(
	ctx context.Context,
	orgId uuid.UUID,
	servicepolicyId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/servicepolicies/%v")
	req.AppendTemplateParams(orgId, servicepolicyId)
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

// GetOrgServicePolicy takes context, orgId, servicepolicyId as parameters and
// returns an models.ApiResponse with models.OrgServicePolicy data and
// an error if there was an issue with the request or response.
// Return an organization-level service policy, including the matched tenants and services, action, local routing, path preference, and inspection settings.
func (o *OrgsServicePolicies) GetOrgServicePolicy(
	ctx context.Context,
	orgId uuid.UUID,
	servicepolicyId uuid.UUID) (
	models.ApiResponse[models.OrgServicePolicy],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/servicepolicies/%v")
	req.AppendTemplateParams(orgId, servicepolicyId)
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

	var result models.OrgServicePolicy
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.OrgServicePolicy](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgServicePolicy takes context, orgId, servicepolicyId, body as parameters and
// returns an models.ApiResponse with models.OrgServicePolicy data and
// an error if there was an issue with the request or response.
// Update an organization-level service policy, including its tenant and service matches, action, local routing, path preference, or inspection settings.
func (o *OrgsServicePolicies) UpdateOrgServicePolicy(
	ctx context.Context,
	orgId uuid.UUID,
	servicepolicyId uuid.UUID,
	body *models.OrgServicePolicy) (
	models.ApiResponse[models.OrgServicePolicy],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/servicepolicies/%v")
	req.AppendTemplateParams(orgId, servicepolicyId)
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

	var result models.OrgServicePolicy
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.OrgServicePolicy](decoder)
	return models.NewApiResponse(result, resp), err
}
