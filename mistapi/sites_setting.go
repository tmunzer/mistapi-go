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

// SitesSetting represents a controller struct.
type SitesSetting struct {
    baseController
}

// NewSitesSetting creates a new instance of SitesSetting.
// It takes a baseController as a parameter and returns a pointer to the SitesSetting.
func NewSitesSetting(baseController baseController) *SitesSetting {
    sitesSetting := SitesSetting{baseController: baseController}
    return &sitesSetting
}

// GetSiteSetting takes context, siteId as parameters and
// returns an models.ApiResponse with models.SiteSetting data and
// an error if there was an issue with the request or response.
// Get the Site Settings
func (s *SitesSetting) GetSiteSetting(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.SiteSetting],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/setting")
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
    
    var result models.SiteSetting
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SiteSetting](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSiteSettings takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.SiteSetting data and
// an error if there was an issue with the request or response.
// Update Site Settings
func (s *SitesSetting) UpdateSiteSettings(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.SiteSetting) (
    models.ApiResponse[models.SiteSetting],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/setting")
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.SiteSetting
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SiteSetting](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSiteWirelessClientsBlocklist takes context, siteId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site Blacklist Station Clients
func (s *SitesSetting) DeleteSiteWirelessClientsBlocklist(
    ctx context.Context,
    siteId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/setting/blacklist")
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// CreateSiteWirelessClientsBlocklist takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.MacAddresses data and
// an error if there was an issue with the request or response.
// This endpoint is to provide list of client macs for annotation blacklist.
// Retrieve the current clients list `blacklist_url` under Site:Setting
func (s *SitesSetting) CreateSiteWirelessClientsBlocklist(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.MacAddresses],
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/setting/blacklist")
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.MacAddresses
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.MacAddresses](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSettingDerived takes context, siteId as parameters and
// returns an models.ApiResponse with models.SiteSettingDerived data and
// an error if there was an issue with the request or response.
// Get the Derived Site Settings, generated by merging the Org level templates (network templates, gateway templates) and the Site level configuration. If the same parameter is defined in both scopes, the Site level one is used. In addition, the Zoom and Teams accounts are also merged into the derived settings.
func (s *SitesSetting) GetSiteSettingDerived(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.SiteSettingDerived],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/setting/derived")
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
    
    var result models.SiteSettingDerived
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SiteSettingDerived](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSiteWatchedStations takes context, siteId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site Watched Station Clients
func (s *SitesSetting) DeleteSiteWatchedStations(
    ctx context.Context,
    siteId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      "/api/v1/sites/%v/setting/watched_station",
    )
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// CreateSiteWatchedStations takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.MacAddresses data and
// an error if there was an issue with the request or response.
// This endpoint is to provide list of client macs for annotation as watched station.
// Retrieve the current clients list from `watched_station_url` under Site:Setting
func (s *SitesSetting) CreateSiteWatchedStations(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.MacAddresses],
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/setting/watched_station")
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.MacAddresses
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.MacAddresses](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSiteWirelessClientsAllowlist takes context, siteId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site Whitelist Station Clients
func (s *SitesSetting) DeleteSiteWirelessClientsAllowlist(
    ctx context.Context,
    siteId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/setting/whitelist")
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// CreateSiteWirelessClientsAllowlist takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.MacAddresses data and
// an error if there was an issue with the request or response.
// This endpoint is to provide list of client macs for annotation as whitelist.
// Retrieve the current clients list from `whitelist_url` under Site:Setting
func (s *SitesSetting) CreateSiteWirelessClientsAllowlist(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.MacAddresses],
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/setting/whitelist")
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.MacAddresses
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.MacAddresses](decoder)
    return models.NewApiResponse(result, resp), err
}
