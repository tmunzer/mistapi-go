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
)

// OrgsStatsOspf represents a controller struct.
type OrgsStatsOspf struct {
	baseController
}

// NewOrgsStatsOspf creates a new instance of OrgsStatsOspf.
// It takes a baseController as a parameter and returns a pointer to the OrgsStatsOspf.
func NewOrgsStatsOspf(baseController baseController) *OrgsStatsOspf {
	orgsStatsOspf := OrgsStatsOspf{baseController: baseController}
	return &orgsStatsOspf
}

// CountOrgOspfStats takes context, orgId, distinct, start, end, limit, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count OSPF peer stats by distinct attribute name
func (o *OrgsStatsOspf) CountOrgOspfStats(
	ctx context.Context,
	orgId uuid.UUID,
	distinct *models.OspfPeerStatsCountDistinctEnum,
	start *string,
	end *string,
	limit *int,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/ospf_peers/count")
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
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.ResponseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchOrgOspfStats takes context, orgId, siteId, mac, vrfName, peerIp, start, end, limit, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.OspfPeerStatsSearchResult data and
// an error if there was an issue with the request or response.
// Search OSPF Neighbor Stats
func (o *OrgsStatsOspf) SearchOrgOspfStats(
	ctx context.Context,
	orgId uuid.UUID,
	siteId *string,
	mac *string,
	vrfName *string,
	peerIp *string,
	start *string,
	end *string,
	limit *int,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.OspfPeerStatsSearchResult],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/ospf_peers/search")
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
	if siteId != nil {
		req.QueryParam("site_id", *siteId)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if vrfName != nil {
		req.QueryParam("vrf_name", *vrfName)
	}
	if peerIp != nil {
		req.QueryParam("peer_ip", *peerIp)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.OspfPeerStatsSearchResult
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.OspfPeerStatsSearchResult](decoder)
	return models.NewApiResponse(result, resp), err
}
