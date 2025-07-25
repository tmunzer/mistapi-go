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

// SitesWxRules represents a controller struct.
type SitesWxRules struct {
    baseController
}

// NewSitesWxRules creates a new instance of SitesWxRules.
// It takes a baseController as a parameter and returns a pointer to the SitesWxRules.
func NewSitesWxRules(baseController baseController) *SitesWxRules {
    sitesWxRules := SitesWxRules{baseController: baseController}
    return &sitesWxRules
}

// ListSiteWxRules takes context, siteId, limit, page as parameters and
// returns an models.ApiResponse with []models.WxlanRule data and
// an error if there was an issue with the request or response.
// Get List of Site WxLan Rules
func (s *SitesWxRules) ListSiteWxRules(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.WxlanRule],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/wxrules")
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
    
    var result []models.WxlanRule
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.WxlanRule](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateSiteWxRule takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.WxlanRule data and
// an error if there was an issue with the request or response.
// Create Site WxLan Rule
func (s *SitesWxRules) CreateSiteWxRule(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.WxlanRule) (
    models.ApiResponse[models.WxlanRule],
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/wxrules")
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
    
    var result models.WxlanRule
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WxlanRule](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListSiteWxRulesDerived takes context, siteId as parameters and
// returns an models.ApiResponse with []models.WxlanRule data and
// an error if there was an issue with the request or response.
// Get the list of derived WxLan Rule for a site
func (s *SitesWxRules) ListSiteWxRulesDerived(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[[]models.WxlanRule],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/wxrules/derived")
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
    
    var result []models.WxlanRule
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.WxlanRule](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSiteWxRule takes context, siteId, wxruleId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site WxLan Rule
func (s *SitesWxRules) DeleteSiteWxRule(
    ctx context.Context,
    siteId uuid.UUID,
    wxruleId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/wxrules/%v")
    req.AppendTemplateParams(siteId, wxruleId)
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

// GetSiteWxRule takes context, siteId, wxruleId as parameters and
// returns an models.ApiResponse with models.WxlanRule data and
// an error if there was an issue with the request or response.
// Get Site WxLan Rule Details
func (s *SitesWxRules) GetSiteWxRule(
    ctx context.Context,
    siteId uuid.UUID,
    wxruleId uuid.UUID) (
    models.ApiResponse[models.WxlanRule],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/wxrules/%v")
    req.AppendTemplateParams(siteId, wxruleId)
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
    
    var result models.WxlanRule
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WxlanRule](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSiteWxRule takes context, siteId, wxruleId, body as parameters and
// returns an models.ApiResponse with models.WxlanRule data and
// an error if there was an issue with the request or response.
// Update Site WxLan Rule
func (s *SitesWxRules) UpdateSiteWxRule(
    ctx context.Context,
    siteId uuid.UUID,
    wxruleId uuid.UUID,
    body *models.WxlanRule) (
    models.ApiResponse[models.WxlanRule],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/wxrules/%v")
    req.AppendTemplateParams(siteId, wxruleId)
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
    
    var result models.WxlanRule
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WxlanRule](decoder)
    return models.NewApiResponse(result, resp), err
}
