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

// SitesWxTags represents a controller struct.
type SitesWxTags struct {
    baseController
}

// NewSitesWxTags creates a new instance of SitesWxTags.
// It takes a baseController as a parameter and returns a pointer to the SitesWxTags.
func NewSitesWxTags(baseController baseController) *SitesWxTags {
    sitesWxTags := SitesWxTags{baseController: baseController}
    return &sitesWxTags
}

// ListSiteWxTags takes context, siteId, limit, page as parameters and
// returns an models.ApiResponse with []models.WxlanTag data and
// an error if there was an issue with the request or response.
// Get List of Site WxTags
func (s *SitesWxTags) ListSiteWxTags(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.WxlanTag],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/wxtags")
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
    
    var result []models.WxlanTag
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.WxlanTag](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateSiteWxTag takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.WxlanTag data and
// an error if there was an issue with the request or response.
// Create Site WxTag
func (s *SitesWxTags) CreateSiteWxTag(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.WxlanTag) (
    models.ApiResponse[models.WxlanTag],
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/wxtags")
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
    
    var result models.WxlanTag
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WxlanTag](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteApplicationList takes context, siteId as parameters and
// returns an models.ApiResponse with []models.SearchWxtagAppsItem data and
// an error if there was an issue with the request or response.
// Get Application List
func (s *SitesWxTags) GetSiteApplicationList(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[[]models.SearchWxtagAppsItem],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/wxtags/apps")
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
    
    var result []models.SearchWxtagAppsItem
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.SearchWxtagAppsItem](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSiteWxTag takes context, siteId, wxtagId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site WxTag
func (s *SitesWxTags) DeleteSiteWxTag(
    ctx context.Context,
    siteId uuid.UUID,
    wxtagId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/wxtags/%v")
    req.AppendTemplateParams(siteId, wxtagId)
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

// GetSiteWxTag takes context, siteId, wxtagId as parameters and
// returns an models.ApiResponse with models.WxlanTag data and
// an error if there was an issue with the request or response.
// Get Site WxTag Details
func (s *SitesWxTags) GetSiteWxTag(
    ctx context.Context,
    siteId uuid.UUID,
    wxtagId uuid.UUID) (
    models.ApiResponse[models.WxlanTag],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/wxtags/%v")
    req.AppendTemplateParams(siteId, wxtagId)
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
    
    var result models.WxlanTag
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WxlanTag](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSiteWxTag takes context, siteId, wxtagId, body as parameters and
// returns an models.ApiResponse with models.WxlanTag data and
// an error if there was an issue with the request or response.
// Update Site WxTag
func (s *SitesWxTags) UpdateSiteWxTag(
    ctx context.Context,
    siteId uuid.UUID,
    wxtagId uuid.UUID,
    body *models.WxlanTag) (
    models.ApiResponse[models.WxlanTag],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/wxtags/%v")
    req.AppendTemplateParams(siteId, wxtagId)
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
    
    var result models.WxlanTag
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WxlanTag](decoder)
    return models.NewApiResponse(result, resp), err
}
