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

// SitesInsights represents a controller struct.
type SitesInsights struct {
	baseController
}

// NewSitesInsights creates a new instance of SitesInsights.
// It takes a baseController as a parameter and returns a pointer to the SitesInsights.
func NewSitesInsights(baseController baseController) *SitesInsights {
	sitesInsights := SitesInsights{baseController: baseController}
	return &sitesInsights
}

// GetSiteInsightMetrics takes context, siteId, metrics, start, end, duration, interval, limit, page as parameters and
// returns an models.ApiResponse with models.InsightMetrics data and
// an error if there was an issue with the request or response.
// Get Site Insight Metrics
func (s *SitesInsights) GetSiteInsightMetrics(
	ctx context.Context,
	siteId uuid.UUID,
	metrics string,
	start *string,
	end *string,
	duration *string,
	interval *string,
	limit *int,
	page *int) (
	models.ApiResponse[models.InsightMetrics],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/insights")
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
	req.QueryParam("metrics", metrics)
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}
	if interval != nil {
		req.QueryParam("interval", *interval)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result models.InsightMetrics
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.InsightMetrics](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteInsightMetricsForAP takes context, siteId, deviceId, metrics, start, end, duration, interval, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseDeviceMetrics data and
// an error if there was an issue with the request or response.
// Get AP Insight Metrics
func (s *SitesInsights) GetSiteInsightMetricsForAP(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID,
	metrics string,
	start *string,
	end *string,
	duration *string,
	interval *string,
	limit *int,
	page *int) (
	models.ApiResponse[models.ResponseDeviceMetrics],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/insights/ap/%v/stats")
	req.AppendTemplateParams(siteId, deviceId)
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
	req.QueryParam("metrics", metrics)
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}
	if interval != nil {
		req.QueryParam("interval", *interval)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result models.ResponseDeviceMetrics
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseDeviceMetrics](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteInsightMetricsForClient takes context, siteId, clientMac, metrics, start, end, duration, interval, limit, page as parameters and
// returns an models.ApiResponse with models.InsightMetrics data and
// an error if there was an issue with the request or response.
// Get Client Insight Metrics
func (s *SitesInsights) GetSiteInsightMetricsForClient(
	ctx context.Context,
	siteId uuid.UUID,
	clientMac string,
	metrics string,
	start *string,
	end *string,
	duration *string,
	interval *string,
	limit *int,
	page *int) (
	models.ApiResponse[models.InsightMetrics],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/insights/client/%v")
	req.AppendTemplateParams(siteId, clientMac)
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
	req.QueryParam("metrics", metrics)
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}
	if interval != nil {
		req.QueryParam("interval", *interval)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result models.InsightMetrics
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.InsightMetrics](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteInsightMetricsForDevice takes context, siteId, metric, deviceMac, portId, start, end, duration, interval, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseDeviceMetrics data and
// an error if there was an issue with the request or response.
// Get AP Insight Metrics
// See metrics possibilities at [List Insight Metrics]($e/Constants%20Definitions/listInsightMetrics)
func (s *SitesInsights) GetSiteInsightMetricsForDevice(
	ctx context.Context,
	siteId uuid.UUID,
	metric string,
	deviceMac string,
	portId *string,
	start *string,
	end *string,
	duration *string,
	interval *string,
	limit *int,
	page *int) (
	models.ApiResponse[models.ResponseDeviceMetrics],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/insights/device/%v/%v")
	req.AppendTemplateParams(siteId, deviceMac, metric)
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
	if portId != nil {
		req.QueryParam("port_id", *portId)
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
	if interval != nil {
		req.QueryParam("interval", *interval)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result models.ResponseDeviceMetrics
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseDeviceMetrics](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteInsightMetricsForGateway takes context, siteId, deviceId, metrics, portId, start, end, duration, interval, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseDeviceMetrics data and
// an error if there was an issue with the request or response.
// Get Gateway Insight Metrics
func (s *SitesInsights) GetSiteInsightMetricsForGateway(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID,
	metrics string,
	portId *string,
	start *string,
	end *string,
	duration *string,
	interval *string,
	limit *int,
	page *int) (
	models.ApiResponse[models.ResponseDeviceMetrics],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/insights/gateway/%v/stats")
	req.AppendTemplateParams(siteId, deviceId)
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
	req.QueryParam("metrics", metrics)
	if portId != nil {
		req.QueryParam("port_id", *portId)
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
	if interval != nil {
		req.QueryParam("interval", *interval)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result models.ResponseDeviceMetrics
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseDeviceMetrics](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteInsightMetricsForMxEdge takes context, siteId, metric, deviceMac, portId, start, end, duration, interval, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseDeviceMetrics data and
// an error if there was an issue with the request or response.
// Get MxEdge Insight Metrics
// See metrics possibilities at [List Insight Metrics](/#operations/listInsightMetrics)
func (s *SitesInsights) GetSiteInsightMetricsForMxEdge(
	ctx context.Context,
	siteId uuid.UUID,
	metric string,
	deviceMac string,
	portId *string,
	start *string,
	end *string,
	duration *string,
	interval *string,
	limit *int,
	page *int) (
	models.ApiResponse[models.ResponseDeviceMetrics],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/insights/mxedge/%v/%v")
	req.AppendTemplateParams(siteId, deviceMac, metric)
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
	if portId != nil {
		req.QueryParam("port_id", *portId)
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
	if interval != nil {
		req.QueryParam("interval", *interval)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result models.ResponseDeviceMetrics
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseDeviceMetrics](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteInsightMetricsForSwitch takes context, siteId, metric, deviceMac, portId, start, end, duration, interval, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseDeviceMetrics data and
// an error if there was an issue with the request or response.
// Get Switch Insight Metrics
// See metrics possibilities at [List Insight Metrics](/#operations/listInsightMetrics)
func (s *SitesInsights) GetSiteInsightMetricsForSwitch(
	ctx context.Context,
	siteId uuid.UUID,
	metric string,
	deviceMac string,
	portId *string,
	start *string,
	end *string,
	duration *string,
	interval *string,
	limit *int,
	page *int) (
	models.ApiResponse[models.ResponseDeviceMetrics],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/insights/switch/%v/%v")
	req.AppendTemplateParams(siteId, deviceMac, metric)
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
	if portId != nil {
		req.QueryParam("port_id", *portId)
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
	if interval != nil {
		req.QueryParam("interval", *interval)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result models.ResponseDeviceMetrics
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseDeviceMetrics](decoder)
	return models.NewApiResponse(result, resp), err
}
