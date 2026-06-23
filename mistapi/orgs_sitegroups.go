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

// OrgsSitegroups represents a controller struct.
type OrgsSitegroups struct {
	baseController
}

// NewOrgsSitegroups creates a new instance of OrgsSitegroups.
// It takes a baseController as a parameter and returns a pointer to the OrgsSitegroups.
func NewOrgsSitegroups(baseController baseController) *OrgsSitegroups {
	orgsSitegroups := OrgsSitegroups{baseController: baseController}
	return &orgsSitegroups
}

// ListOrgSiteGroups takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.Sitegroup data and
// an error if there was an issue with the request or response.
// List site groups configured in the organization. A site group collects site IDs so sites can be managed or referenced as a group.
func (o *OrgsSitegroups) ListOrgSiteGroups(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.Sitegroup],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/sitegroups")
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

	var result []models.Sitegroup
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Sitegroup](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgSiteGroup takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Sitegroup data and
// an error if there was an issue with the request or response.
// Create a site group in the organization with a display name and optional site membership.
func (o *OrgsSitegroups) CreateOrgSiteGroup(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Sitegroup) (
	models.ApiResponse[models.Sitegroup],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/sitegroups")
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

	var result models.Sitegroup
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sitegroup](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgSiteGroup takes context, orgId, sitegroupId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Remove a site group from the organization.
func (o *OrgsSitegroups) DeleteOrgSiteGroup(
	ctx context.Context,
	orgId uuid.UUID,
	sitegroupId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/sitegroups/%v")
	req.AppendTemplateParams(orgId, sitegroupId)
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

// GetOrgSiteGroup takes context, orgId, sitegroupId as parameters and
// returns an models.ApiResponse with models.Sitegroup data and
// an error if there was an issue with the request or response.
// Return a site group, including its display name and the site IDs included in the group.
func (o *OrgsSitegroups) GetOrgSiteGroup(
	ctx context.Context,
	orgId uuid.UUID,
	sitegroupId uuid.UUID) (
	models.ApiResponse[models.Sitegroup],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/sitegroups/%v")
	req.AppendTemplateParams(orgId, sitegroupId)
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

	var result models.Sitegroup
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sitegroup](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgSiteGroup takes context, orgId, sitegroupId, body as parameters and
// returns an models.ApiResponse with models.Sitegroup data and
// an error if there was an issue with the request or response.
// Update the display name used to identify a site group in the organization.
func (o *OrgsSitegroups) UpdateOrgSiteGroup(
	ctx context.Context,
	orgId uuid.UUID,
	sitegroupId uuid.UUID,
	body *models.NameString) (
	models.ApiResponse[models.Sitegroup],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/sitegroups/%v")
	req.AppendTemplateParams(orgId, sitegroupId)
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

	var result models.Sitegroup
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sitegroup](decoder)
	return models.NewApiResponse(result, resp), err
}
