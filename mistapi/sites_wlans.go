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

// SitesWlans represents a controller struct.
type SitesWlans struct {
    baseController
}

// NewSitesWlans creates a new instance of SitesWlans.
// It takes a baseController as a parameter and returns a pointer to the SitesWlans.
func NewSitesWlans(baseController baseController) *SitesWlans {
    sitesWlans := SitesWlans{baseController: baseController}
    return &sitesWlans
}

// ListSiteWlans takes context, siteId, limit, page as parameters and
// returns an models.ApiResponse with []models.Wlan data and
// an error if there was an issue with the request or response.
// Get List of Site WLANs
func (s *SitesWlans) ListSiteWlans(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Wlan],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/wlans")
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
    
    var result []models.Wlan
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Wlan](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateSiteWlan takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.Wlan data and
// an error if there was an issue with the request or response.
// Create Site WLAN
func (s *SitesWlans) CreateSiteWlan(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Wlan) (
    models.ApiResponse[models.Wlan],
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/wlans")
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
    
    var result models.Wlan
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Wlan](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListSiteWlansDerived takes context, siteId, resolve, wlanId as parameters and
// returns an models.ApiResponse with []models.Wlan data and
// an error if there was an issue with the request or response.
// Get the list of derived Wlans for a Site
func (s *SitesWlans) ListSiteWlansDerived(
    ctx context.Context,
    siteId uuid.UUID,
    resolve *bool,
    wlanId *string) (
    models.ApiResponse[[]models.Wlan],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/wlans/derived")
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
    if resolve != nil {
        req.QueryParam("resolve", *resolve)
    }
    if wlanId != nil {
        req.QueryParam("wlan_id", *wlanId)
    }
    
    var result []models.Wlan
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Wlan](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSiteWlan takes context, siteId, wlanId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site WLAN
func (s *SitesWlans) DeleteSiteWlan(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/wlans/%v")
    req.AppendTemplateParams(siteId, wlanId)
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

// GetSiteWlan takes context, siteId, wlanId as parameters and
// returns an models.ApiResponse with models.Wlan data and
// an error if there was an issue with the request or response.
// Get Site WLAN
func (s *SitesWlans) GetSiteWlan(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId uuid.UUID) (
    models.ApiResponse[models.Wlan],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/wlans/%v")
    req.AppendTemplateParams(siteId, wlanId)
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
    
    var result models.Wlan
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Wlan](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSiteWlan takes context, siteId, wlanId, body as parameters and
// returns an models.ApiResponse with models.Wlan data and
// an error if there was an issue with the request or response.
// Update Site WLAN
func (s *SitesWlans) UpdateSiteWlan(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId uuid.UUID,
    body *models.Wlan) (
    models.ApiResponse[models.Wlan],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/wlans/%v")
    req.AppendTemplateParams(siteId, wlanId)
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
    
    var result models.Wlan
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Wlan](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSiteWlanPortalImage takes context, siteId, wlanId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site WLAN Portal Image
func (s *SitesWlans) DeleteSiteWlanPortalImage(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/wlans/%v/portal_image")
    req.AppendTemplateParams(siteId, wlanId)
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

// UploadSiteWlanPortalImage takes context, siteId, wlanId, file, json as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// WLAN Portal Image Upload
func (s *SitesWlans) UploadSiteWlanPortalImage(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId uuid.UUID,
    file models.FileWrapper,
    json *string) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/wlans/%v/portal_image")
    req.AppendTemplateParams(siteId, wlanId)
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
    formFields := []https.FormParam{}
    fileParam := https.FormParam{Key: "file", Value: file, Headers: http.Header{}}
    formFields = append(formFields, fileParam)
    if json != nil {
        jsonParam := https.FormParam{Key: "json", Value: *json, Headers: http.Header{}}
        formFields = append(formFields, jsonParam)
    }
    req.FormData(formFields)
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// UpdateSiteWlanPortalTemplate takes context, siteId, wlanId, body as parameters and
// returns an models.ApiResponse with models.WlanPortalTemplate data and
// an error if there was an issue with the request or response.
// Update a Portal Template
// #### Sponsor Email Template
// Sponsor Email Template supports following template variables:
// | **Name** | **Description** |
// | --- | --- |
// | approve_url | Renders URL to approve the request; optionally &minutes=N query param can be appended to change the Authorization period of the guest, where N is a valid integer denoting number of minutes a guest remains authorized |
// | deny_url | Renders URL to reject the request |
// | guest_email | Renders Email ID of the guest |
// | guest_name | Renders Name of the guest |
// | field1 | Renders value of the Custom Field 1 |
// | field2 | Renders value of the Custom Field 2 |
// | company | Renders value of the Company field |
// | sponsor_link_validity_duration | Renders validity time of the request (i.e. Approve/Deny URL) |
// | auth_expire_minutes | Renders Wlan-level configured Guest Authorization Expiration time period (in minutes), If not configured then default (1 day in minutes) |
func (s *SitesWlans) UpdateSiteWlanPortalTemplate(
    ctx context.Context,
    siteId uuid.UUID,
    wlanId uuid.UUID,
    body *models.WlanPortalTemplate) (
    models.ApiResponse[models.WlanPortalTemplate],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/wlans/%v/portal_template")
    req.AppendTemplateParams(siteId, wlanId)
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
    
    var result models.WlanPortalTemplate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WlanPortalTemplate](decoder)
    return models.NewApiResponse(result, resp), err
}
