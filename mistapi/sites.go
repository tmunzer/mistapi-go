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

// Sites represents a controller struct.
type Sites struct {
    baseController
}

// NewSites creates a new instance of Sites.
// It takes a baseController as a parameter and returns a pointer to the Sites.
func NewSites(baseController baseController) *Sites {
    sites := Sites{baseController: baseController}
    return &sites
}

// DeleteSite takes context, siteId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site
func (s *Sites) DeleteSite(
    ctx context.Context,
    siteId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v")
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

// GetSiteInfo takes context, siteId as parameters and
// returns an models.ApiResponse with models.Site data and
// an error if there was an issue with the request or response.
// Provides information about the site, including its name, address,
// timezone, and associated templates. This endpoint is useful for retrieving
// the current configuration and details of a specific site.
func (s *Sites) GetSiteInfo(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.Site],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v")
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
    
    var result models.Site
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Site](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSiteInfo takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.Site data and
// an error if there was an issue with the request or response.
// Updates the configuration and metadata for an existing site. 
// This endpoint allows modification of site properties including location details (address, coordinates, timezone), template associations (alarm, network, RF, security policy templates), site group memberships, and general information (name, notes).
// All fields are optional and only provided fields will be updated.
func (s *Sites) UpdateSiteInfo(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Site) (
    models.ApiResponse[models.Site],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v")
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
    
    var result models.Site
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Site](decoder)
    return models.NewApiResponse(result, resp), err
}
