package mistapi

import (
    "context"
    "fmt"
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

// GetSiteInsightMetricsForClient takes context, siteId, clientMac, metric, page, limit, start, end, duration, interval as parameters and
// returns an models.ApiResponse with models.InsightMetrics data and
// an error if there was an issue with the request or response.
// Get Client Insight Metrics
// See metrics possibilities at /api/v1/const/insight_metrics
func (s *SitesInsights) GetSiteInsightMetricsForClient(
    ctx context.Context,
    siteId uuid.UUID,
    clientMac string,
    metric string,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string,
    interval *string) (
    models.ApiResponse[models.InsightMetrics],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/insights/client/%v/%v", siteId, clientMac, metric),
    )
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    if page != nil {
        req.QueryParam("page", *page)
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
    if interval != nil {
        req.QueryParam("interval", *interval)
    }
    
    var result models.InsightMetrics
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.InsightMetrics](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteInsightMetricsForDevice takes context, siteId, metric, deviceMac, page, limit, start, end, duration, interval as parameters and
// returns an models.ApiResponse with models.ResponseDeviceMetrics data and
// an error if there was an issue with the request or response.
// Get AP Insight Metrics
// See metrics possibilities at /api/v1/const/insight_metrics
func (s *SitesInsights) GetSiteInsightMetricsForDevice(
    ctx context.Context,
    siteId uuid.UUID,
    metric string,
    deviceMac string,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string,
    interval *string) (
    models.ApiResponse[models.ResponseDeviceMetrics],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/insights/device/%v/%v", siteId, deviceMac, metric),
    )
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    if page != nil {
        req.QueryParam("page", *page)
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
    if interval != nil {
        req.QueryParam("interval", *interval)
    }
    
    var result models.ResponseDeviceMetrics
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseDeviceMetrics](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteInsightMetrics takes context, siteId, metric, page, limit, start, end, duration, interval as parameters and
// returns an models.ApiResponse with models.InsightMetrics data and
// an error if there was an issue with the request or response.
// Get Site Insight Metrics
// See metrics possibilities at /api/v1/const/insight_metrics
func (s *SitesInsights) GetSiteInsightMetrics(
    ctx context.Context,
    siteId uuid.UUID,
    metric string,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string,
    interval *string) (
    models.ApiResponse[models.InsightMetrics],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/insights/%v", siteId, metric),
    )
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    if page != nil {
        req.QueryParam("page", *page)
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
    if interval != nil {
        req.QueryParam("interval", *interval)
    }
    
    var result models.InsightMetrics
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.InsightMetrics](decoder)
    return models.NewApiResponse(result, resp), err
}
