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

// GetSiteInsightMetricsForClient takes context, siteId, clientMac, metric, start, end, duration, interval, limit, page as parameters and
// returns an models.ApiResponse with models.InsightMetrics data and
// an error if there was an issue with the request or response.
// Get Client Insight Metrics
// See metrics possibilities at [List Insight Metrics]($e/Constants%20Definitions/listInsightMetrics)
func (s *SitesInsights) GetSiteInsightMetricsForClient(
    ctx context.Context,
    siteId uuid.UUID,
    clientMac string,
    metric string,
    start *int,
    end *int,
    duration *string,
    interval *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.InsightMetrics],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/insights/client/%v/%v")
    req.AppendTemplateParams(siteId, clientMac, metric)
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

// GetSiteInsightMetricsForDevice takes context, siteId, metric, deviceMac, start, end, duration, interval, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseDeviceMetrics data and
// an error if there was an issue with the request or response.
// Get AP Insight Metrics
// See metrics possibilities at [List Insight Metrics]($e/Constants%20Definitions/listInsightMetrics)
func (s *SitesInsights) GetSiteInsightMetricsForDevice(
    ctx context.Context,
    siteId uuid.UUID,
    metric string,
    deviceMac string,
    start *int,
    end *int,
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

// GetSiteInsightMetrics takes context, siteId, metric, start, end, duration, interval, limit, page as parameters and
// returns an models.ApiResponse with models.InsightMetrics data and
// an error if there was an issue with the request or response.
// Get Site Insight Metrics
// See metrics possibilities at [List Insight Metrics]($e/Constants%20Definitions/listInsightMetrics)
func (s *SitesInsights) GetSiteInsightMetrics(
    ctx context.Context,
    siteId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    interval *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.InsightMetrics],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/insights/%v")
    req.AppendTemplateParams(siteId, metric)
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
