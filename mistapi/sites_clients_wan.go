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

// SitesClientsWan represents a controller struct.
type SitesClientsWan struct {
	baseController
}

// NewSitesClientsWan creates a new instance of SitesClientsWan.
// It takes a baseController as a parameter and returns a pointer to the SitesClientsWan.
func NewSitesClientsWan(baseController baseController) *SitesClientsWan {
	sitesClientsWan := SitesClientsWan{baseController: baseController}
	return &sitesClientsWan
}

// CountSiteWanClientEvents takes context, siteId, distinct, mType, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count WAN client events for a site, optionally grouped by the `distinct` field and filtered by event type and time range. Use [Count Org WAN Client Events]($e/Orgs%20Clients%20-%20Wan/countOrgWanClientEvents) to count WAN client events across the organization.
func (s *SitesClientsWan) CountSiteWanClientEvents(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteWanClientEventsDistinctEnum,
	mType *string,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/wan_client/events/count")
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
	if mType != nil {
		req.QueryParam("type", *mType)
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

// CountSiteWanClients takes context, siteId, distinct, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count WAN clients for a site, optionally grouped by the `distinct` field and filtered by time range. Use [Count Org WAN Clients]($e/Orgs%20Clients%20-%20Wan/countOrgWanClients) to count WAN clients across the organization.
func (s *SitesClientsWan) CountSiteWanClients(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteWanClientsCountDistinctEnum,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/wan_clients/count")
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

// SearchSiteWanClientEvents takes context, siteId, mType, mac, hostname, ip, mfg, nacruleId, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.SearchEventsWanClient data and
// an error if there was an issue with the request or response.
// Search WAN client events for a site with filters for client identity, manufacturer, NAC rule, event type, and time range. Use [Search Org WAN Client Events]($e/Orgs%20Clients%20-%20Wan/searchOrgWanClientEvents) to search WAN client events across the organization.
func (s *SitesClientsWan) SearchSiteWanClientEvents(
	ctx context.Context,
	siteId uuid.UUID,
	mType *string,
	mac *string,
	hostname *string,
	ip *string,
	mfg *string,
	nacruleId *string,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.SearchEventsWanClient],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/wan_clients/events/search")
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
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if hostname != nil {
		req.QueryParam("hostname", *hostname)
	}
	if ip != nil {
		req.QueryParam("ip", *ip)
	}
	if mfg != nil {
		req.QueryParam("mfg", *mfg)
	}
	if nacruleId != nil {
		req.QueryParam("nacrule_id", *nacruleId)
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

	var result models.SearchEventsWanClient
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SearchEventsWanClient](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteWanClients takes context, siteId, hostname, ip, ipSrc, mac, mfg, network, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.SearchWanClient data and
// an error if there was an issue with the request or response.
// Search WAN clients for a site with filters for hostname, IP address, source IP, MAC address, manufacturer, network, and time range. Use [Search Org WAN Clients]($e/Orgs%20Clients%20-%20Wan/searchOrgWanClients) to search WAN clients across the organization.
func (s *SitesClientsWan) SearchSiteWanClients(
	ctx context.Context,
	siteId uuid.UUID,
	hostname *string,
	ip *string,
	ipSrc *string,
	mac *string,
	mfg *string,
	network *string,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.SearchWanClient],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/wan_clients/search")
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
	if hostname != nil {
		req.QueryParam("hostname", *hostname)
	}
	if ip != nil {
		req.QueryParam("ip", *ip)
	}
	if ipSrc != nil {
		req.QueryParam("ip_src", *ipSrc)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if mfg != nil {
		req.QueryParam("mfg", *mfg)
	}
	if network != nil {
		req.QueryParam("network", *network)
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

	var result models.SearchWanClient
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SearchWanClient](decoder)
	return models.NewApiResponse(result, resp), err
}
