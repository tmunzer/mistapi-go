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

// OrgsWxTags represents a controller struct.
type OrgsWxTags struct {
	baseController
}

// NewOrgsWxTags creates a new instance of OrgsWxTags.
// It takes a baseController as a parameter and returns a pointer to the OrgsWxTags.
func NewOrgsWxTags(baseController baseController) *OrgsWxTags {
	orgsWxTags := OrgsWxTags{baseController: baseController}
	return &orgsWxTags
}

// ListOrgWxTags takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.WxlanTag data and
// an error if there was an issue with the request or response.
// Get List of Org WxLAN Tags
func (o *OrgsWxTags) ListOrgWxTags(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.WxlanTag],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/wxtags")
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

	var result []models.WxlanTag
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.WxlanTag](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgWxTag takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.WxlanTag data and
// an error if there was an issue with the request or response.
// Create WxLAN Tag
func (o *OrgsWxTags) CreateOrgWxTag(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.WxlanTag) (
	models.ApiResponse[models.WxlanTag],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/wxtags")
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

	var result models.WxlanTag
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.WxlanTag](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetOrgApplicationList takes context, orgId as parameters and
// returns an models.ApiResponse with []models.SearchWxtagAppsItem data and
// an error if there was an issue with the request or response.
// Get Application List
func (o *OrgsWxTags) GetOrgApplicationList(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[[]models.SearchWxtagAppsItem],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/wxtags/apps")
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

	var result []models.SearchWxtagAppsItem
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.SearchWxtagAppsItem](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgWxTag takes context, orgId, wxtagId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete WxLAN Tag
func (o *OrgsWxTags) DeleteOrgWxTag(
	ctx context.Context,
	orgId uuid.UUID,
	wxtagId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/wxtags/%v")
	req.AppendTemplateParams(orgId, wxtagId)
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

// GetOrgWxTag takes context, orgId, wxtagId as parameters and
// returns an models.ApiResponse with models.WxlanTag data and
// an error if there was an issue with the request or response.
// Get WxLAN Tag Details
func (o *OrgsWxTags) GetOrgWxTag(
	ctx context.Context,
	orgId uuid.UUID,
	wxtagId uuid.UUID) (
	models.ApiResponse[models.WxlanTag],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/wxtags/%v")
	req.AppendTemplateParams(orgId, wxtagId)
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

	var result models.WxlanTag
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.WxlanTag](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgWxTag takes context, orgId, wxtagId, body as parameters and
// returns an models.ApiResponse with models.WxlanTag data and
// an error if there was an issue with the request or response.
// Update WxLAN Tag
func (o *OrgsWxTags) UpdateOrgWxTag(
	ctx context.Context,
	orgId uuid.UUID,
	wxtagId uuid.UUID,
	body *models.WxlanTag) (
	models.ApiResponse[models.WxlanTag],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/wxtags/%v")
	req.AppendTemplateParams(orgId, wxtagId)
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

	var result models.WxlanTag
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.WxlanTag](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetOrgCurrentMatchingClientsOfAWxTag takes context, orgId, wxtagId as parameters and
// returns an models.ApiResponse with []models.WxtagClient data and
// an error if there was an issue with the request or response.
// Get Current Matching Clients of a WXLAN Tag
func (o *OrgsWxTags) GetOrgCurrentMatchingClientsOfAWxTag(
	ctx context.Context,
	orgId uuid.UUID,
	wxtagId uuid.UUID) (
	models.ApiResponse[[]models.WxtagClient],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/wxtags/%v/clients")
	req.AppendTemplateParams(orgId, wxtagId)
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

	var result []models.WxtagClient
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.WxtagClient](decoder)
	return models.NewApiResponse(result, resp), err
}
