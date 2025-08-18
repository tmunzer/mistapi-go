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

// OrgsSecurityPolicies represents a controller struct.
type OrgsSecurityPolicies struct {
	baseController
}

// NewOrgsSecurityPolicies creates a new instance of OrgsSecurityPolicies.
// It takes a baseController as a parameter and returns a pointer to the OrgsSecurityPolicies.
func NewOrgsSecurityPolicies(baseController baseController) *OrgsSecurityPolicies {
	orgsSecurityPolicies := OrgsSecurityPolicies{baseController: baseController}
	return &orgsSecurityPolicies
}

// ListOrgSecPolicies takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.Secpolicy data and
// an error if there was an issue with the request or response.
// Get List of Org Security Policies
func (o *OrgsSecurityPolicies) ListOrgSecPolicies(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.Secpolicy],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/secpolicies")
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

	var result []models.Secpolicy
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Secpolicy](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgSecPolicy takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Secpolicy data and
// an error if there was an issue with the request or response.
// Create Org Security Policy
func (o *OrgsSecurityPolicies) CreateOrgSecPolicy(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Secpolicy) (
	models.ApiResponse[models.Secpolicy],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/secpolicies")
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

	var result models.Secpolicy
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Secpolicy](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgSecPolicy takes context, orgId, secpolicyId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org Security Policy
func (o *OrgsSecurityPolicies) DeleteOrgSecPolicy(
	ctx context.Context,
	orgId uuid.UUID,
	secpolicyId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/secpolicies/%v")
	req.AppendTemplateParams(orgId, secpolicyId)
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

// GetOrgSecPolicy takes context, orgId, secpolicyId as parameters and
// returns an models.ApiResponse with models.Secpolicy data and
// an error if there was an issue with the request or response.
// Get Org Security Policy
func (o *OrgsSecurityPolicies) GetOrgSecPolicy(
	ctx context.Context,
	orgId uuid.UUID,
	secpolicyId uuid.UUID) (
	models.ApiResponse[models.Secpolicy],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/secpolicies/%v")
	req.AppendTemplateParams(orgId, secpolicyId)
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

	var result models.Secpolicy
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Secpolicy](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgSecPolicy takes context, orgId, secpolicyId, body as parameters and
// returns an models.ApiResponse with models.Secpolicy data and
// an error if there was an issue with the request or response.
// Update Org Security Policy
func (o *OrgsSecurityPolicies) UpdateOrgSecPolicy(
	ctx context.Context,
	orgId uuid.UUID,
	secpolicyId uuid.UUID,
	body *models.Secpolicy) (
	models.ApiResponse[models.Secpolicy],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/secpolicies/%v")
	req.AppendTemplateParams(orgId, secpolicyId)
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

	var result models.Secpolicy
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Secpolicy](decoder)
	return models.NewApiResponse(result, resp), err
}
