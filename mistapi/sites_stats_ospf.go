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
// Count by Distinct Attributes of OSPF peers stats
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

// SearchSiteOspfStats takes context, siteIdTemplate, siteId, mac, vrfName, peerIp, start, end, limit, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.OspfPeerStatsSearchResult data and
// an error if there was an issue with the request or response.
// Search OSPF Neighbor Stats
func (s *SitesStatsOspf) SearchSiteOspfStats(
	ctx context.Context,
	siteIdTemplate uuid.UUID,
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
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/ospf_peers/search")
	req.AppendTemplateParams(siteIdTemplate)
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
