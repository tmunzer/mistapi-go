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

// OrgsSiteTemplates represents a controller struct.
type OrgsSiteTemplates struct {
    baseController
}

// NewOrgsSiteTemplates creates a new instance of OrgsSiteTemplates.
// It takes a baseController as a parameter and returns a pointer to the OrgsSiteTemplates.
func NewOrgsSiteTemplates(baseController baseController) *OrgsSiteTemplates {
    orgsSiteTemplates := OrgsSiteTemplates{baseController: baseController}
    return &orgsSiteTemplates
}

// ListOrgSiteTemplates takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.SiteTemplate data and
// an error if there was an issue with the request or response.
// Get List of Org Site Templates
func (o *OrgsSiteTemplates) ListOrgSiteTemplates(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.SiteTemplate],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/sitetemplates")
    req.AppendTemplateParams(orgId)
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
    
    var result []models.SiteTemplate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.SiteTemplate](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgSiteTemplate takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.SiteTemplate data and
// an error if there was an issue with the request or response.
// Create Org Site Template
func (o *OrgsSiteTemplates) CreateOrgSiteTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.SiteTemplate) (
    models.ApiResponse[models.SiteTemplate],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/sitetemplates")
    req.AppendTemplateParams(orgId)
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
    
    var result models.SiteTemplate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SiteTemplate](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgSiteTemplate takes context, orgId, sitetemplateId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org Site Template
func (o *OrgsSiteTemplates) DeleteOrgSiteTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    sitetemplateId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/sitetemplates/%v")
    req.AppendTemplateParams(orgId, sitetemplateId)
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

// GetOrgSiteTemplate takes context, orgId, sitetemplateId as parameters and
// returns an models.ApiResponse with models.SiteTemplate data and
// an error if there was an issue with the request or response.
// Get Org Site Template
func (o *OrgsSiteTemplates) GetOrgSiteTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    sitetemplateId uuid.UUID) (
    models.ApiResponse[models.SiteTemplate],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/sitetemplates/%v")
    req.AppendTemplateParams(orgId, sitetemplateId)
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
    
    var result models.SiteTemplate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SiteTemplate](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgSiteTemplate takes context, orgId, sitetemplateId, body as parameters and
// returns an models.ApiResponse with models.SiteTemplate data and
// an error if there was an issue with the request or response.
// Update Org Site Template
func (o *OrgsSiteTemplates) UpdateOrgSiteTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    sitetemplateId uuid.UUID,
    body *models.SiteTemplate) (
    models.ApiResponse[models.SiteTemplate],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/sitetemplates/%v")
    req.AppendTemplateParams(orgId, sitetemplateId)
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
    
    var result models.SiteTemplate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SiteTemplate](decoder)
    return models.NewApiResponse(result, resp), err
}
