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

// SitesZones represents a controller struct.
type SitesZones struct {
    baseController
}

// NewSitesZones creates a new instance of SitesZones.
// It takes a baseController as a parameter and returns a pointer to the SitesZones.
func NewSitesZones(baseController baseController) *SitesZones {
    sitesZones := SitesZones{baseController: baseController}
    return &sitesZones
}

// ListSiteZones takes context, siteId, limit, page as parameters and
// returns an models.ApiResponse with []models.Zone data and
// an error if there was an issue with the request or response.
// Get List of Site Zones
func (s *SitesZones) ListSiteZones(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Zone],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/zones")
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
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result []models.Zone
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Zone](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateSiteZone takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.Zone data and
// an error if there was an issue with the request or response.
// Create Site Zone
func (s *SitesZones) CreateSiteZone(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Zone) (
    models.ApiResponse[models.Zone],
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/zones")
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
    
    var result models.Zone
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Zone](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSiteZone takes context, siteId, zoneId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site Zone
func (s *SitesZones) DeleteSiteZone(
    ctx context.Context,
    siteId uuid.UUID,
    zoneId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/zones/%v")
    req.AppendTemplateParams(siteId, zoneId)
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

// GetSiteZone takes context, siteId, zoneId as parameters and
// returns an models.ApiResponse with models.Zone data and
// an error if there was an issue with the request or response.
// Get Site Zone Details
func (s *SitesZones) GetSiteZone(
    ctx context.Context,
    siteId uuid.UUID,
    zoneId uuid.UUID) (
    models.ApiResponse[models.Zone],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/zones/%v")
    req.AppendTemplateParams(siteId, zoneId)
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
    
    var result models.Zone
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Zone](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSiteZone takes context, siteId, zoneId, body as parameters and
// returns an models.ApiResponse with models.Zone data and
// an error if there was an issue with the request or response.
// Update Site Zone
func (s *SitesZones) UpdateSiteZone(
    ctx context.Context,
    siteId uuid.UUID,
    zoneId uuid.UUID,
    body *models.Zone) (
    models.ApiResponse[models.Zone],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/zones/%v")
    req.AppendTemplateParams(siteId, zoneId)
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
    
    var result models.Zone
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Zone](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountSiteZoneSessions takes context, siteId, zoneType, distinct, userType, user, scopeId, scope, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Site Zone Sessions
func (s *SitesZones) CountSiteZoneSessions(
    ctx context.Context,
    siteId uuid.UUID,
    zoneType models.ZoneTypeEnum,
    distinct *models.SiteZoneCountDistinctEnum,
    userType *models.RfClientTypeEnum,
    user *string,
    scopeId *string,
    scope *models.ZoneScopeEnum,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/%v/count")
    req.AppendTemplateParams(siteId, zoneType)
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
    if userType != nil {
        req.QueryParam("user_type", *userType)
    }
    if user != nil {
        req.QueryParam("user", *user)
    }
    if scopeId != nil {
        req.QueryParam("scope_id", *scopeId)
    }
    if scope != nil {
        req.QueryParam("scope", *scope)
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

// SearchSiteZoneSessions takes context, siteId, zoneType, userType, user, scopeId, scope, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseZoneSearch data and
// an error if there was an issue with the request or response.
// Search Zone Sessions
func (s *SitesZones) SearchSiteZoneSessions(
    ctx context.Context,
    siteId uuid.UUID,
    zoneType models.ZoneTypeEnum,
    userType *models.RfClientTypeEnum,
    user *string,
    scopeId *string,
    scope *models.VisitsScopeEnum,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseZoneSearch],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/%v/visits/search")
    req.AppendTemplateParams(siteId, zoneType)
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
    if userType != nil {
        req.QueryParam("user_type", *userType)
    }
    if user != nil {
        req.QueryParam("user", *user)
    }
    if scopeId != nil {
        req.QueryParam("scope_id", *scopeId)
    }
    if scope != nil {
        req.QueryParam("scope", *scope)
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
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result models.ResponseZoneSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseZoneSearch](decoder)
    return models.NewApiResponse(result, resp), err
}
