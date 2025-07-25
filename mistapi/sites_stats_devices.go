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

// SitesStatsDevices represents a controller struct.
type SitesStatsDevices struct {
    baseController
}

// NewSitesStatsDevices creates a new instance of SitesStatsDevices.
// It takes a baseController as a parameter and returns a pointer to the SitesStatsDevices.
func NewSitesStatsDevices(baseController baseController) *SitesStatsDevices {
    sitesStatsDevices := SitesStatsDevices{baseController: baseController}
    return &sitesStatsDevices
}

// ListSiteDevicesStats takes context, siteId, mType, status, limit, page as parameters and
// returns an models.ApiResponse with []models.StatsDevice data and
// an error if there was an issue with the request or response.
// Get List of Site Devices Stats
func (s *SitesStatsDevices) ListSiteDevicesStats(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.DeviceTypeWithAllEnum,
    status *models.StatDeviceStatusFilterEnum,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.StatsDevice],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/devices")
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
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if status != nil {
        req.QueryParam("status", *status)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result []models.StatsDevice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.StatsDevice](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteDeviceStats takes context, siteId, deviceId, fields as parameters and
// returns an models.ApiResponse with models.StatsDevice data and
// an error if there was an issue with the request or response.
// Get Site Device Stats Details
func (s *SitesStatsDevices) GetSiteDeviceStats(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    fields *string) (
    models.ApiResponse[models.StatsDevice],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/devices/%v")
    req.AppendTemplateParams(siteId, deviceId)
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
    if fields != nil {
        req.QueryParam("fields", *fields)
    }
    
    var result models.StatsDevice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.StatsDevice](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteAllClientsStatsByDevice takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with []models.StatsWirelessClient data and
// an error if there was an issue with the request or response.
// Get wireless client stat by Device
func (s *SitesStatsDevices) GetSiteAllClientsStatsByDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[[]models.StatsWirelessClient],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/devices/%v/clients")
    req.AppendTemplateParams(siteId, deviceId)
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
    
    var result []models.StatsWirelessClient
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.StatsWirelessClient](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteGatewayMetrics takes context, siteId as parameters and
// returns an models.ApiResponse with models.GatewayMetrics data and
// an error if there was an issue with the request or response.
// Get Site Gateway Metrics
func (s *SitesStatsDevices) GetSiteGatewayMetrics(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.GatewayMetrics],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/gateways/metrics")
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
    
    var result models.GatewayMetrics
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.GatewayMetrics](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSwitchesMetrics takes context, siteId, mType, scope, switchMac as parameters and
// returns an models.ApiResponse with models.ResponseSwitchMetrics data and
// an error if there was an issue with the request or response.
// Get version compliance metrics for managed or monitored switches
func (s *SitesStatsDevices) GetSiteSwitchesMetrics(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.SwitchMetricTypeEnum,
    scope *models.SwitchMetricScopeEnum,
    switchMac *string) (
    models.ApiResponse[models.ResponseSwitchMetrics],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/switches/metrics")
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
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if scope != nil {
        req.QueryParam("scope", *scope)
    }
    if switchMac != nil {
        req.QueryParam("switch_mac", *switchMac)
    }
    
    var result models.ResponseSwitchMetrics
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseSwitchMetrics](decoder)
    return models.NewApiResponse(result, resp), err
}
