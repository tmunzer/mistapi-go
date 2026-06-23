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

// OrgsMxEdges represents a controller struct.
type OrgsMxEdges struct {
	baseController
}

// NewOrgsMxEdges creates a new instance of OrgsMxEdges.
// It takes a baseController as a parameter and returns a pointer to the OrgsMxEdges.
func NewOrgsMxEdges(baseController baseController) *OrgsMxEdges {
	orgsMxEdges := OrgsMxEdges{baseController: baseController}
	return &orgsMxEdges
}

// ListOrgMxEdges takes context, orgId, forSite, limit, page as parameters and
// returns an models.ApiResponse with []models.Mxedge data and
// an error if there was an issue with the request or response.
// List Mist Edge appliances in the organization, optionally filtering for org-level, site-level, or all Mist Edges with `for_site`.
func (o *OrgsMxEdges) ListOrgMxEdges(
	ctx context.Context,
	orgId uuid.UUID,
	forSite *models.MxedgeForSiteEnum,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.Mxedge],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/mxedges")
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
	if forSite != nil {
		req.QueryParam("for_site", *forSite)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.Mxedge
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Mxedge](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgMxEdge takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Mxedge data and
// an error if there was an issue with the request or response.
// Create a Mist Edge appliance configuration in the organization, including cluster assignment, management, services, and tunnel termination settings.
func (o *OrgsMxEdges) CreateOrgMxEdge(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Mxedge) (
	models.ApiResponse[models.Mxedge],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/mxedges")
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

	var result models.Mxedge
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Mxedge](decoder)
	return models.NewApiResponse(result, resp), err
}

// AssignOrgMxEdgeToSite takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.ResponseAssignSuccess data and
// an error if there was an issue with the request or response.
// Assign one or more Mist Edge appliances from the organization to a site by Mist Edge ID and site ID.
func (o *OrgsMxEdges) AssignOrgMxEdgeToSite(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.MxedgesAssign) (
	models.ApiResponse[models.ResponseAssignSuccess],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/mxedges/assign")
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

	var result models.ResponseAssignSuccess
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseAssignSuccess](decoder)
	return models.NewApiResponse(result, resp), err
}

// ClaimOrgMxEdge takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.ResponseClaimMxEdge data and
// an error if there was an issue with the request or response.
// Claim one or more Mist Edge appliances into the organization using their claim codes.
func (o *OrgsMxEdges) ClaimOrgMxEdge(
	ctx context.Context,
	orgId uuid.UUID,
	body []string) (
	models.ApiResponse[models.ResponseClaimMxEdge],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/mxedges/claim")
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

	var result models.ResponseClaimMxEdge
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseClaimMxEdge](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountOrgMxEdges takes context, orgId, distinct, mxedgeId, siteId, mxclusterId, model, distro, tuntermVersion, sort, stats, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count organization Mist Edge records, optionally grouped by `distinct` and filtered by Mist Edge, cluster, site, model, distro, tunnel termination version, and time range.
func (o *OrgsMxEdges) CountOrgMxEdges(
	ctx context.Context,
	orgId uuid.UUID,
	distinct *models.OrgMxedgeCountDistinctEnum,
	mxedgeId *string,
	siteId *string,
	mxclusterId *string,
	model *string,
	distro *string,
	tuntermVersion *string,
	sort *string,
	stats *bool,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/mxedges/count")
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
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}
	if mxedgeId != nil {
		req.QueryParam("mxedge_id", *mxedgeId)
	}
	if siteId != nil {
		req.QueryParam("site_id", *siteId)
	}
	if mxclusterId != nil {
		req.QueryParam("mxcluster_id", *mxclusterId)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if distro != nil {
		req.QueryParam("distro", *distro)
	}
	if tuntermVersion != nil {
		req.QueryParam("tunterm_version", *tuntermVersion)
	}
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if stats != nil {
		req.QueryParam("stats", *stats)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}

	var result models.ResponseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountOrgSiteMxEdgeEvents takes context, orgId, distinct, mxedgeId, mxclusterId, mType, service, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count Mist Edge event records across the organization, optionally grouped by `distinct` and filtered by Mist Edge, cluster, event type, service, and time range.
func (o *OrgsMxEdges) CountOrgSiteMxEdgeEvents(
	ctx context.Context,
	orgId uuid.UUID,
	distinct *models.OrgMxedgeEventsCountDistinctEnum,
	mxedgeId *string,
	mxclusterId *string,
	mType *string,
	service *string,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/mxedges/events/count")
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
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}
	if mxedgeId != nil {
		req.QueryParam("mxedge_id", *mxedgeId)
	}
	if mxclusterId != nil {
		req.QueryParam("mxcluster_id", *mxclusterId)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if service != nil {
		req.QueryParam("service", *service)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}

	var result models.ResponseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchOrgMistEdgeEvents takes context, orgId, mxedgeId, mxclusterId, mType, service, component, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseMxedgeEventsSearch data and
// an error if there was an issue with the request or response.
// Search Mist Edge event records across the organization with filters for Mist Edge, cluster, event type, service, component, and time range.
func (o *OrgsMxEdges) SearchOrgMistEdgeEvents(
	ctx context.Context,
	orgId uuid.UUID,
	mxedgeId *string,
	mxclusterId *string,
	mType *string,
	service *string,
	component *string,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseMxedgeEventsSearch],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/mxedges/events/search")
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
	if mxedgeId != nil {
		req.QueryParam("mxedge_id", *mxedgeId)
	}
	if mxclusterId != nil {
		req.QueryParam("mxcluster_id", *mxclusterId)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if service != nil {
		req.QueryParam("service", *service)
	}
	if component != nil {
		req.QueryParam("component", *component)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.ResponseMxedgeEventsSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseMxedgeEventsSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchOrgMxEdges takes context, orgId, hostname, mxedgeId, mxclusterId, model, distro, tuntermVersion, siteId, stats, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseMxedgeSearch data and
// an error if there was an issue with the request or response.
// Search organization Mist Edge records with filters for hostname, Mist Edge, cluster, site, model, distro, tunnel termination version, and time range.
func (o *OrgsMxEdges) SearchOrgMxEdges(
	ctx context.Context,
	orgId uuid.UUID,
	hostname *string,
	mxedgeId *string,
	mxclusterId *string,
	model *string,
	distro *string,
	tuntermVersion *string,
	siteId *string,
	stats *bool,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseMxedgeSearch],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/mxedges/search")
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
	if hostname != nil {
		req.QueryParam("hostname", *hostname)
	}
	if mxedgeId != nil {
		req.QueryParam("mxedge_id", *mxedgeId)
	}
	if mxclusterId != nil {
		req.QueryParam("mxcluster_id", *mxclusterId)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if distro != nil {
		req.QueryParam("distro", *distro)
	}
	if tuntermVersion != nil {
		req.QueryParam("tunterm_version", *tuntermVersion)
	}
	if siteId != nil {
		req.QueryParam("site_id", *siteId)
	}
	if stats != nil {
		req.QueryParam("stats", *stats)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.ResponseMxedgeSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseMxedgeSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// UnassignOrgMxEdgeFromSite takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.ResponseAssignSuccess data and
// an error if there was an issue with the request or response.
// Unassign one or more Mist Edge appliances from their current site while keeping them in the organization.
func (o *OrgsMxEdges) UnassignOrgMxEdgeFromSite(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.MxedgesUnassign) (
	models.ApiResponse[models.ResponseAssignSuccess],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/mxedges/unassign")
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

	var result models.ResponseAssignSuccess
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseAssignSuccess](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetOrgMxEdgeUpgradeInfo takes context, orgId, channel, distro as parameters and
// returns an models.ApiResponse with []models.MxedgeUpgradeInfoItems data and
// an error if there was an issue with the request or response.
// Retrieve available Mist Edge package versions by upgrade channel and distro.
func (o *OrgsMxEdges) GetOrgMxEdgeUpgradeInfo(
	ctx context.Context,
	orgId uuid.UUID,
	channel *models.GetOrgMxedgeUpgradeInfoChannelEnum,
	distro *string) (
	models.ApiResponse[[]models.MxedgeUpgradeInfoItems],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/mxedges/versions")
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
	if channel != nil {
		req.QueryParam("channel", *channel)
	}
	if distro != nil {
		req.QueryParam("distro", *distro)
	}

	var result []models.MxedgeUpgradeInfoItems
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.MxedgeUpgradeInfoItems](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgMxEdge takes context, orgId, mxedgeId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete a Mist Edge appliance record from the organization by Mist Edge ID.
func (o *OrgsMxEdges) DeleteOrgMxEdge(
	ctx context.Context,
	orgId uuid.UUID,
	mxedgeId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/mxedges/%v")
	req.AppendTemplateParams(orgId, mxedgeId)
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

// GetOrgMxEdge takes context, orgId, mxedgeId as parameters and
// returns an models.ApiResponse with models.Mxedge data and
// an error if there was an issue with the request or response.
// Retrieve configuration and registration details for a specific Mist Edge appliance in the organization.
func (o *OrgsMxEdges) GetOrgMxEdge(
	ctx context.Context,
	orgId uuid.UUID,
	mxedgeId uuid.UUID) (
	models.ApiResponse[models.Mxedge],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/mxedges/%v")
	req.AppendTemplateParams(orgId, mxedgeId)
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

	var result models.Mxedge
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Mxedge](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgMxEdge takes context, orgId, mxedgeId, body as parameters and
// returns an models.ApiResponse with models.Mxedge data and
// an error if there was an issue with the request or response.
// Update a Mist Edge appliance configuration, including model, name, management IP, services, and tunnel termination settings.
func (o *OrgsMxEdges) UpdateOrgMxEdge(
	ctx context.Context,
	orgId uuid.UUID,
	mxedgeId uuid.UUID,
	body *models.Mxedge) (
	models.ApiResponse[models.Mxedge],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/mxedges/%v")
	req.AppendTemplateParams(orgId, mxedgeId)
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

	var result models.Mxedge
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Mxedge](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgMxEdgeImage takes context, orgId, mxedgeId, imageNumber as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete a numbered image attachment from a Mist Edge appliance.
func (o *OrgsMxEdges) DeleteOrgMxEdgeImage(
	ctx context.Context,
	orgId uuid.UUID,
	mxedgeId uuid.UUID,
	imageNumber int) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/mxedges/%v/image/%v")
	req.AppendTemplateParams(orgId, mxedgeId, imageNumber)
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

// AddOrgMxEdgeImage takes context, orgId, mxedgeId, imageNumber, file, json as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Upload and attach an image file to a Mist Edge appliance. A Mist Edge can have up to three image attachments.
func (o *OrgsMxEdges) AddOrgMxEdgeImage(
	ctx context.Context,
	orgId uuid.UUID,
	mxedgeId uuid.UUID,
	imageNumber int,
	file models.FileWrapper,
	json *string) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/mxedges/%v/image/%v")
	req.AppendTemplateParams(orgId, mxedgeId, imageNumber)
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
	formFields := []https.FormParam{}
	fileParam := https.FormParam{Key: "file", Value: file, Headers: http.Header{}}
	formFields = append(formFields, fileParam)
	if json != nil {
		jsonParam := https.FormParam{Key: "json", Value: *json, Headers: http.Header{}}
		formFields = append(formFields, jsonParam)
	}
	req.FormData(formFields)

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// RestartOrgMxEdge takes context, orgId, mxedgeId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Restart the registration workflow for a Mist Edge replacement by disconnecting the currently registered appliance so another Mist Edge can register.
func (o *OrgsMxEdges) RestartOrgMxEdge(
	ctx context.Context,
	orgId uuid.UUID,
	mxedgeId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/mxedges/%v/restart")
	req.AppendTemplateParams(orgId, mxedgeId)
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

// BounceOrgMxEdgeDataPorts takes context, orgId, mxedgeId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Bounce one or more TunTerm data ports on a Mist Edge, optionally setting the hold time between port bounces.
func (o *OrgsMxEdges) BounceOrgMxEdgeDataPorts(
	ctx context.Context,
	orgId uuid.UUID,
	mxedgeId uuid.UUID,
	body *models.UtilsTuntermBouncePort) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		"/api/v1/orgs/%v/mxedges/%v/services/tunterm/bounce_port",
	)
	req.AppendTemplateParams(orgId, mxedgeId)
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

// DisconnectOrgMxEdgeTuntermAps takes context, orgId, mxedgeId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Disconnect specific APs from the Mist Edge TunTerm service by AP MAC address.
func (o *OrgsMxEdges) DisconnectOrgMxEdgeTuntermAps(
	ctx context.Context,
	orgId uuid.UUID,
	mxedgeId uuid.UUID,
	body *models.MacAddresses) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		"/api/v1/orgs/%v/mxedges/%v/services/tunterm/disconnect_aps",
	)
	req.AppendTemplateParams(orgId, mxedgeId)
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

// ControlOrgMxEdgeServices takes context, orgId, mxedgeId, name, action as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Start, stop, or restart a named Mist Edge service such as `tunterm`, `mxagent`, or `radsecproxy`.
func (o *OrgsMxEdges) ControlOrgMxEdgeServices(
	ctx context.Context,
	orgId uuid.UUID,
	mxedgeId uuid.UUID,
	name models.MxedgeServiceNameEnum,
	action models.MxedgeServiceActionEnum) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/mxedges/%v/services/%v/%v")
	req.AppendTemplateParams(orgId, mxedgeId, name, action)
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

// UploadOrgMxEdgeSupportFiles takes context, orgId, mxedgeId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Trigger upload of support files from a Mist Edge for troubleshooting.
func (o *OrgsMxEdges) UploadOrgMxEdgeSupportFiles(
	ctx context.Context,
	orgId uuid.UUID,
	mxedgeId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/mxedges/%v/support")
	req.AppendTemplateParams(orgId, mxedgeId)
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

// UnregisterOrgMxEdge takes context, orgId, mxedgeId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Unregister a Mist Edge during a replacement workflow by disconnecting the currently registered appliance so another Mist Edge can register.
func (o *OrgsMxEdges) UnregisterOrgMxEdge(
	ctx context.Context,
	orgId uuid.UUID,
	mxedgeId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/mxedges/%v/unregister")
	req.AppendTemplateParams(orgId, mxedgeId)
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

// GetOrgMxEdgeVmParams takes context, orgId, mxedgeId as parameters and
// returns an models.ApiResponse with models.MxedgeVmParams data and
// an error if there was an issue with the request or response.
// Retrieve VM deployment parameters for a Mist Edge, including model, optional name, and base64 user data.
func (o *OrgsMxEdges) GetOrgMxEdgeVmParams(
	ctx context.Context,
	orgId uuid.UUID,
	mxedgeId uuid.UUID) (
	models.ApiResponse[models.MxedgeVmParams],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/mxedges/%v/vm_params")
	req.AppendTemplateParams(orgId, mxedgeId)
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

	var result models.MxedgeVmParams
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MxedgeVmParams](decoder)
	return models.NewApiResponse(result, resp), err
}
