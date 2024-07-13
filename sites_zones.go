package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "mistapi/errors"
    "mistapi/models"
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

// ListSiteZonesStats takes context, siteId, mapId as parameters and
// returns an models.ApiResponse with []models.ZoneStats data and
// an error if there was an issue with the request or response.
// Get List of Site Zones Stats
func (s *SitesZones) ListSiteZonesStats(
    ctx context.Context,
    siteId uuid.UUID,
    mapId *string) (
    models.ApiResponse[[]models.ZoneStats],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/zones", siteId),
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
    if mapId != nil {
        req.QueryParam("map_id", *mapId)
    }
    
    var result []models.ZoneStats
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ZoneStats](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteZoneStats takes context, siteId, zoneType, zoneId as parameters and
// returns an models.ApiResponse with models.ZoneStatsDetails data and
// an error if there was an issue with the request or response.
// Get Detail Zone Stats
func (s *SitesZones) GetSiteZoneStats(
    ctx context.Context,
    siteId uuid.UUID,
    zoneType models.ZoneTypeEnum,
    zoneId uuid.UUID) (
    models.ApiResponse[models.ZoneStatsDetails],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/%v/%v", siteId, zoneType, zoneId),
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
    
    var result models.ZoneStatsDetails
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ZoneStatsDetails](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListSiteZones takes context, siteId, page, limit as parameters and
// returns an models.ApiResponse with []models.Zone data and
// an error if there was an issue with the request or response.
// Get List of Site Zones
func (s *SitesZones) ListSiteZones(
    ctx context.Context,
    siteId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.Zone],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/zones", siteId),
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
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/zones", siteId),
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
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete Site Zone
func (s *SitesZones) DeleteSiteZone(
    ctx context.Context,
    siteId uuid.UUID,
    zoneId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/sites/%v/zones/%v", siteId, zoneId),
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
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
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/zones/%v", siteId, zoneId),
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
    req := s.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/sites/%v/zones/%v", siteId, zoneId),
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

// CountSiteZoneSessions takes context, siteId, zoneType, distinct, userType, user, scopeId, scope, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count Site Zone Sessions
func (s *SitesZones) CountSiteZoneSessions(
    ctx context.Context,
    siteId uuid.UUID,
    zoneType models.ZoneTypeEnum,
    distinct *models.SiteZoneCountDistinctEnum,
    userType *models.RfClientTypeEnum,
    user *string,
    scopeId *string,
    scope *models.ZoneScopeEnum,
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
      fmt.Sprintf("/api/v1/sites/%v/%v/count", siteId, zoneType),
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

// SearchSiteZoneSessions takes context, siteId, zoneType, userType, user, scopeId, scope, page, limit, start, end, duration as parameters and
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
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseZoneSearch],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/%v/visits/search", siteId, zoneType),
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
    
    var result models.ResponseZoneSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseZoneSearch](decoder)
    return models.NewApiResponse(result, resp), err
}
