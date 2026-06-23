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

// SitesStatsOspf represents a controller struct.
type SitesStatsOspf struct {
	baseController
}

// NewSitesStatsOspf creates a new instance of SitesStatsOspf.
// It takes a baseController as a parameter and returns a pointer to the SitesStatsOspf.
func NewSitesStatsOspf(baseController baseController) *SitesStatsOspf {
	sitesStatsOspf := SitesStatsOspf{baseController: baseController}
	return &sitesStatsOspf
}

// CountSiteOspfStats takes context, siteId, distinct, start, end, limit, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count OSPF peer statistics for a site, optionally grouped by the `distinct` field and filtered by time range. Use [Count Org OSPF Stats]($e/Orgs%20Stats%20-%20Ospf/countOrgOspfStats) to count OSPF peer statistics across the organization.
func (s *SitesStatsOspf) CountSiteOspfStats(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.OspfPeerStatsCountDistinctEnum,
	start *string,
	end *string,
	limit *int,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/ospf_peers/count")
	req.AppendTemplateParams(siteId)
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

// SearchSiteOspfStats takes context, siteId, mac, vrfName, peerIp, start, end, limit, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.OspfPeerStatsSearchResult data and
// an error if there was an issue with the request or response.
// Search OSPF peer statistics for a site with filters for device, VRF, peer IP, and time range. Use [Search Org OSPF Stats]($e/Orgs%20Stats%20-%20Ospf/searchOrgOspfStats) to search OSPF peer statistics across the organization.
func (s *SitesStatsOspf) SearchSiteOspfStats(
	ctx context.Context,
	siteId uuid.UUID,
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
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/ospf_peers/search")
	req.AppendTemplateParams(siteId)
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
