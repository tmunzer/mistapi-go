package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "mistapi/errors"
    "mistapi/models"
)

// SitesDevicesDiscoveredSwitches represents a controller struct.
type SitesDevicesDiscoveredSwitches struct {
    baseController
}

// NewSitesDevicesDiscoveredSwitches creates a new instance of SitesDevicesDiscoveredSwitches.
// It takes a baseController as a parameter and returns a pointer to the SitesDevicesDiscoveredSwitches.
func NewSitesDevicesDiscoveredSwitches(baseController baseController) *SitesDevicesDiscoveredSwitches {
    sitesDevicesDiscoveredSwitches := SitesDevicesDiscoveredSwitches{baseController: baseController}
    return &sitesDevicesDiscoveredSwitches
}

// SearchSiteDiscoveredSwitchesMetrics takes context, siteId, scope, mType, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseDiscoveredSwitchMetrics data and
// an error if there was an issue with the request or response.
// Search Discovered Switch Metrics
func (s *SitesDevicesDiscoveredSwitches) SearchSiteDiscoveredSwitchesMetrics(
    ctx context.Context,
    siteId uuid.UUID,
    scope *models.DiscoveredSwitchesMetricScopeEnum,
    mType *models.DiscoveredSwitchMetricTypeEnum,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseDiscoveredSwitchMetrics],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/discovered_switch_metrics/search", siteId),
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
    if scope != nil {
        req.QueryParam("scope", *scope)
    }
    if mType != nil {
        req.QueryParam("type", *mType)
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
    
    var result models.ResponseDiscoveredSwitchMetrics
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseDiscoveredSwitchMetrics](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountSiteDiscoveredSwitches takes context, siteId, distinct, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count Discovered Switches
func (s *SitesDevicesDiscoveredSwitches) CountSiteDiscoveredSwitches(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteDiscoveredSwitchesCountDistinctEnum,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/discovered_switches/count", siteId),
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
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
    }
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
    
    var result models.RepsonseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteDiscoveredSwitchesMetrics takes context, siteId, threshold, systemName as parameters and
// returns an models.ApiResponse with models.ResponseDswitchesMetrics data and
// an error if there was an issue with the request or response.
// Discovered switches related metrics, lists related switch system names & details if not compliant
func (s *SitesDevicesDiscoveredSwitches) GetSiteDiscoveredSwitchesMetrics(
    ctx context.Context,
    siteId uuid.UUID,
    threshold *string,
    systemName *string) (
    models.ApiResponse[models.ResponseDswitchesMetrics],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/discovered_switches/metrics", siteId),
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
    if threshold != nil {
        req.QueryParam("threshold", *threshold)
    }
    if systemName != nil {
        req.QueryParam("system_name", *systemName)
    }
    
    var result models.ResponseDswitchesMetrics
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseDswitchesMetrics](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchSiteDiscoveredSwitches takes context, siteId, adopted, systemName, hostname, vendor, model, version, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseDiscoveredSwitches data and
// an error if there was an issue with the request or response.
// Search Discovered Switches
func (s *SitesDevicesDiscoveredSwitches) SearchSiteDiscoveredSwitches(
    ctx context.Context,
    siteId uuid.UUID,
    adopted *bool,
    systemName *string,
    hostname *string,
    vendor *string,
    model *string,
    version *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseDiscoveredSwitches],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/discovered_switches/search", siteId),
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
    if adopted != nil {
        req.QueryParam("adopted", *adopted)
    }
    if systemName != nil {
        req.QueryParam("system_name", *systemName)
    }
    if hostname != nil {
        req.QueryParam("hostname", *hostname)
    }
    if vendor != nil {
        req.QueryParam("vendor", *vendor)
    }
    if model != nil {
        req.QueryParam("model", *model)
    }
    if version != nil {
        req.QueryParam("version", *version)
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
    
    var result models.ResponseDiscoveredSwitches
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseDiscoveredSwitches](decoder)
    return models.NewApiResponse(result, resp), err
}
