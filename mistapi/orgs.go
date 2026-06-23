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

// Orgs represents a controller struct.
type Orgs struct {
	baseController
}

// NewOrgs creates a new instance of Orgs.
// It takes a baseController as a parameter and returns a pointer to the Orgs.
func NewOrgs(baseController baseController) *Orgs {
	orgs := Orgs{baseController: baseController}
	return &orgs
}

// CreateOrg takes context, body as parameters and
// returns an models.ApiResponse with models.Org data and
// an error if there was an issue with the request or response.
// Create a Mist organization with organization-level defaults such as display name, support-access setting, default alarm template, and admin session lifetime.
func (o *Orgs) CreateOrg(
	ctx context.Context,
	body *models.Org) (
	models.ApiResponse[models.Org],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs")

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
	var result models.Org
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Org](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrg takes context, orgId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete an organization and its organization-level resources. Use MSP organization assignment endpoints when only changing MSP ownership or association. This action is only allowed when the Organization Inventory is empty.
func (o *Orgs) DeleteOrg(
	ctx context.Context,
	orgId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v")
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// GetOrg takes context, orgId as parameters and
// returns an models.ApiResponse with models.Org data and
// an error if there was an issue with the request or response.
// Return organization details, including name, MSP ownership, organization group membership, support-access setting, and session settings.
func (o *Orgs) GetOrg(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[models.Org],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v")
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

	var result models.Org
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Org](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrg takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Org data and
// an error if there was an issue with the request or response.
// Update organization-level settings such as display name, support-access setting, organization groups, default alarm template, or admin session lifetime.
func (o *Orgs) UpdateOrg(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Org) (
	models.ApiResponse[models.Org],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v")
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

	var result models.Org
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Org](decoder)
	return models.NewApiResponse(result, resp), err
}

// CloneOrg takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Org data and
// an error if there was an issue with the request or response.
// Create an Org by cloning from another one. Org Settings, Templates, Wxlan Tags, Wxlan Tunnels, Wxlan Rules, Org Wlans will be copied. Sites and Site Groups will not be copied, and therefore, the copied template will not be applied to any site/sitegroups.
func (o *Orgs) CloneOrg(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.NameString) (
	models.ApiResponse[models.Org],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/clone")
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

	var result models.Org
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Org](decoder)
	return models.NewApiResponse(result, resp), err
}
