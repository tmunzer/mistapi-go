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

// SitesStatsIoTEndpoints represents a controller struct.
type SitesStatsIoTEndpoints struct {
	baseController
}

// NewSitesStatsIoTEndpoints creates a new instance of SitesStatsIoTEndpoints.
// It takes a baseController as a parameter and returns a pointer to the SitesStatsIoTEndpoints.
func NewSitesStatsIoTEndpoints(baseController baseController) *SitesStatsIoTEndpoints {
	sitesStatsIoTEndpoints := SitesStatsIoTEndpoints{baseController: baseController}
	return &sitesStatsIoTEndpoints
}

// CountSiteIotEndpoints takes context, siteId, distinct, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count IoT Endpoints
func (s *SitesStatsIoTEndpoints) CountSiteIotEndpoints(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteIotendpointsCountDistinctEnum,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/iotendpoints/count")
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

// SearchSiteIotEndpoints takes context, siteId, apMac, mac, mType, mfg, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseIotEndpointsSearch data and
// an error if there was an issue with the request or response.
// Search IoT Endpoints
func (s *SitesStatsIoTEndpoints) SearchSiteIotEndpoints(
	ctx context.Context,
	siteId uuid.UUID,
	apMac *string,
	mac *string,
	mType *string,
	mfg *string,
	limit *int,
	start *string,
	end *string,
	duration *string) (
	models.ApiResponse[models.ResponseIotEndpointsSearch],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/iotendpoints/search")
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
	if apMac != nil {
		req.QueryParam("ap_mac", *apMac)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if mfg != nil {
		req.QueryParam("mfg", *mfg)
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

	var result models.ResponseIotEndpointsSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseIotEndpointsSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// RejoinSiteIotEndpointZigbee takes context, siteId, id as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Trigger a Zigbee endpoint to rejoin the network
func (s *SitesStatsIoTEndpoints) RejoinSiteIotEndpointZigbee(
	ctx context.Context,
	siteId uuid.UUID,
	id uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		"/api/v1/sites/%v/iotendpoints/%v/zigbee_rejoin",
	)
	req.AppendTemplateParams(siteId, id)
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
