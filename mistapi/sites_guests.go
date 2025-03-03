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

// SitesGuests represents a controller struct.
type SitesGuests struct {
    baseController
}

// NewSitesGuests creates a new instance of SitesGuests.
// It takes a baseController as a parameter and returns a pointer to the SitesGuests.
func NewSitesGuests(baseController baseController) *SitesGuests {
    sitesGuests := SitesGuests{baseController: baseController}
    return &sitesGuests
}

// ListSiteAllGuestAuthorizations takes context, siteId, wlanId as parameters and
// returns an models.ApiResponse with []models.Guest data and
// an error if there was an issue with the request or response.
// Get List of Site Guest Authorizations
func (s *SitesGuests) ListSiteAllGuestAuthorizations(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId *string) (
    models.ApiResponse[[]models.Guest],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/guests")
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
    if wlanId != nil {
        req.QueryParam("wlan_id", *wlanId)
    }
    
    var result []models.Guest
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Guest](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountSiteGuestAuthorizations takes context, siteId, distinct, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count Authorized Guest
func (s *SitesGuests) CountSiteGuestAuthorizations(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteGuestsCountDistinctEnum,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseCount],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/guests/count")
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
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result models.ResponseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListSiteAllGuestAuthorizationsDerived takes context, siteId, wlanId, crossSite as parameters and
// returns an models.ApiResponse with []models.Guest data and
// an error if there was an issue with the request or response.
// Get List of Site Guest Authorizations
func (s *SitesGuests) ListSiteAllGuestAuthorizationsDerived(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId *string,
    crossSite *bool) (
    models.ApiResponse[[]models.Guest],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/guests/derived")
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
    if wlanId != nil {
        req.QueryParam("wlan_id", *wlanId)
    }
    if crossSite != nil {
        req.QueryParam("cross_site", *crossSite)
    }
    
    var result []models.Guest
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Guest](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchSiteGuestAuthorization takes context, siteId, wlanId, authMethod, ssid, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseGuestSearch data and
// an error if there was an issue with the request or response.
// Search Authorized Guest
func (s *SitesGuests) SearchSiteGuestAuthorization(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId *string,
    authMethod *string,
    ssid *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseGuestSearch],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/guests/search")
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
    if wlanId != nil {
        req.QueryParam("wlan_id", *wlanId)
    }
    if authMethod != nil {
        req.QueryParam("auth_method", *authMethod)
    }
    if ssid != nil {
        req.QueryParam("ssid", *ssid)
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
    
    var result models.ResponseGuestSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseGuestSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSiteGuestAuthorization takes context, siteId, guestMac as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Guest Authorization
func (s *SitesGuests) DeleteSiteGuestAuthorization(
    ctx context.Context,
    siteId uuid.UUID,
    guestMac string) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/guests/%v")
    req.AppendTemplateParams(siteId, guestMac)
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

// GetSiteGuestAuthorization takes context, siteId, guestMac as parameters and
// returns an models.ApiResponse with models.Guest data and
// an error if there was an issue with the request or response.
// Get Guest Authorization
func (s *SitesGuests) GetSiteGuestAuthorization(
    ctx context.Context,
    siteId uuid.UUID,
    guestMac string) (
    models.ApiResponse[models.Guest],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/guests/%v")
    req.AppendTemplateParams(siteId, guestMac)
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
    
    var result models.Guest
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Guest](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSiteGuestAuthorization takes context, siteId, guestMac, body as parameters and
// returns an models.ApiResponse with models.Guest data and
// an error if there was an issue with the request or response.
// Update Guest Authorization
func (s *SitesGuests) UpdateSiteGuestAuthorization(
    ctx context.Context,
    siteId uuid.UUID,
    guestMac string,
    body *models.Guest) (
    models.ApiResponse[models.Guest],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/guests/%v")
    req.AppendTemplateParams(siteId, guestMac)
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
    
    var result models.Guest
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Guest](decoder)
    return models.NewApiResponse(result, resp), err
}
