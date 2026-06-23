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

// MSPsOrgs represents a controller struct.
type MSPsOrgs struct {
	baseController
}

// NewMSPsOrgs creates a new instance of MSPsOrgs.
// It takes a baseController as a parameter and returns a pointer to the MSPsOrgs.
func NewMSPsOrgs(baseController baseController) *MSPsOrgs {
	mSPsOrgs := MSPsOrgs{baseController: baseController}
	return &mSPsOrgs
}

// ListMspOrgs takes context, mspId as parameters and
// returns an models.ApiResponse with []models.Org data and
// an error if there was an issue with the request or response.
// List organizations managed by this MSP account, including MSP ownership and organization-group membership metadata.
func (m *MSPsOrgs) ListMspOrgs(
	ctx context.Context,
	mspId uuid.UUID) (
	models.ApiResponse[[]models.Org],
	error) {
	req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v/orgs")
	req.AppendTemplateParams(mspId)
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

	var result []models.Org
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Org](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateMspOrg takes context, mspId, body as parameters and
// returns an models.ApiResponse with models.Org data and
// an error if there was an issue with the request or response.
// Create an organization under this MSP, optionally assigning it to organization groups and setting organization-level defaults.
func (m *MSPsOrgs) CreateMspOrg(
	ctx context.Context,
	mspId uuid.UUID,
	body *models.Org) (
	models.ApiResponse[models.Org],
	error) {
	req := m.prepareRequest(ctx, "POST", "/api/v1/msps/%v/orgs")
	req.AppendTemplateParams(mspId)
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

// ManageMspOrgs takes context, mspId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Assign existing organizations to this MSP account or unassign organizations from it by providing the desired operation and org IDs.
func (m *MSPsOrgs) ManageMspOrgs(
	ctx context.Context,
	mspId uuid.UUID,
	body *models.MspOrgChange) (
	*http.Response,
	error) {
	req := m.prepareRequest(ctx, "PUT", "/api/v1/msps/%v/orgs")
	req.AppendTemplateParams(mspId)
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// SearchMspOrgs takes context, mspId, name, orgId, subInsufficient, trialEnabled, usageTypes, limit, sort, start, end, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseOrgSearch data and
// an error if there was an issue with the request or response.
// Search organizations under this MSP using organization identifiers, names, subscription state, trial state, usage types, and time-based filters.
func (m *MSPsOrgs) SearchMspOrgs(
	ctx context.Context,
	mspId uuid.UUID,
	name *string,
	orgId *uuid.UUID,
	subInsufficient *bool,
	trialEnabled *bool,
	usageTypes []string,
	limit *int,
	sort *string,
	start *string,
	end *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseOrgSearch],
	error) {
	req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v/orgs/search")
	req.AppendTemplateParams(mspId)
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
	if name != nil {
		req.QueryParam("name", *name)
	}
	if orgId != nil {
		req.QueryParam("org_id", *orgId)
	}
	if subInsufficient != nil {
		req.QueryParam("sub_insufficient", *subInsufficient)
	}
	if trialEnabled != nil {
		req.QueryParam("trial_enabled", *trialEnabled)
	}
	if usageTypes != nil {
		req.QueryParam("usage_types", usageTypes)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.ResponseOrgSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseOrgSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteMspOrg takes context, mspId, orgId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete an organization managed by this MSP. Use the MSP organization assignment endpoint when only assigning or unassigning organizations from the MSP account.
func (m *MSPsOrgs) DeleteMspOrg(
	ctx context.Context,
	mspId uuid.UUID,
	orgId uuid.UUID) (
	*http.Response,
	error) {
	req := m.prepareRequest(ctx, "DELETE", "/api/v1/msps/%v/orgs/%v")
	req.AppendTemplateParams(mspId, orgId)
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

// GetMspOrg takes context, mspId, orgId as parameters and
// returns an models.ApiResponse with models.Org data and
// an error if there was an issue with the request or response.
// Return details for one organization managed by this MSP, including MSP ownership, organization groups, support access, and session settings.
func (m *MSPsOrgs) GetMspOrg(
	ctx context.Context,
	mspId uuid.UUID,
	orgId uuid.UUID) (
	models.ApiResponse[models.Org],
	error) {
	req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v/orgs/%v")
	req.AppendTemplateParams(mspId, orgId)
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

// UpdateMspOrg takes context, mspId, orgId, body as parameters and
// returns an models.ApiResponse with models.Org data and
// an error if there was an issue with the request or response.
// Update organization settings for an organization managed by this MSP.
func (m *MSPsOrgs) UpdateMspOrg(
	ctx context.Context,
	mspId uuid.UUID,
	orgId uuid.UUID,
	body *models.Org) (
	models.ApiResponse[models.Org],
	error) {
	req := m.prepareRequest(ctx, "PUT", "/api/v1/msps/%v/orgs/%v")
	req.AppendTemplateParams(mspId, orgId)
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

// ListMspOrgStats takes context, mspId, limit, page as parameters and
// returns an models.ApiResponse with []models.StatsOrg data and
// an error if there was an issue with the request or response.
// List organization statistics for organizations managed by this MSP, including device, inventory, site, and SLE summary fields.
func (m *MSPsOrgs) ListMspOrgStats(
	ctx context.Context,
	mspId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.StatsOrg],
	error) {
	req := m.prepareRequest(ctx, "GET", "/api/v1/msps/%v/stats/orgs")
	req.AppendTemplateParams(mspId)
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

	var result []models.StatsOrg
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.StatsOrg](decoder)
	return models.NewApiResponse(result, resp), err
}
