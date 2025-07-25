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

// SitesVBeacons represents a controller struct.
type SitesVBeacons struct {
    baseController
}

// NewSitesVBeacons creates a new instance of SitesVBeacons.
// It takes a baseController as a parameter and returns a pointer to the SitesVBeacons.
func NewSitesVBeacons(baseController baseController) *SitesVBeacons {
    sitesVBeacons := SitesVBeacons{baseController: baseController}
    return &sitesVBeacons
}

// ListSiteVBeacons takes context, siteId, limit, page as parameters and
// returns an models.ApiResponse with []models.Vbeacon data and
// an error if there was an issue with the request or response.
// Get List of Site Virtual Beacons
func (s *SitesVBeacons) ListSiteVBeacons(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Vbeacon],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/vbeacons")
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
    
    var result []models.Vbeacon
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Vbeacon](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateSiteVBeacon takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.Vbeacon data and
// an error if there was an issue with the request or response.
// Create Virtual Beacon
func (s *SitesVBeacons) CreateSiteVBeacon(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Vbeacon) (
    models.ApiResponse[models.Vbeacon],
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/vbeacons")
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
    
    var result models.Vbeacon
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Vbeacon](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSiteVBeacon takes context, siteId, vbeaconId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site Virtual Beacon
func (s *SitesVBeacons) DeleteSiteVBeacon(
    ctx context.Context,
    siteId uuid.UUID,
    vbeaconId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/vbeacons/%v")
    req.AppendTemplateParams(siteId, vbeaconId)
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

// GetSiteVBeacon takes context, siteId, vbeaconId as parameters and
// returns an models.ApiResponse with models.Vbeacon data and
// an error if there was an issue with the request or response.
// Get Site Virtual Beacon Details
func (s *SitesVBeacons) GetSiteVBeacon(
    ctx context.Context,
    siteId uuid.UUID,
    vbeaconId uuid.UUID) (
    models.ApiResponse[models.Vbeacon],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/vbeacons/%v")
    req.AppendTemplateParams(siteId, vbeaconId)
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
    
    var result models.Vbeacon
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Vbeacon](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSiteVBeacon takes context, siteId, vbeaconId, body as parameters and
// returns an models.ApiResponse with models.Vbeacon data and
// an error if there was an issue with the request or response.
// Update Site Virtual Beacon
func (s *SitesVBeacons) UpdateSiteVBeacon(
    ctx context.Context,
    siteId uuid.UUID,
    vbeaconId uuid.UUID,
    body *models.Vbeacon) (
    models.ApiResponse[models.Vbeacon],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/vbeacons/%v")
    req.AppendTemplateParams(siteId, vbeaconId)
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
    
    var result models.Vbeacon
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Vbeacon](decoder)
    return models.NewApiResponse(result, resp), err
}
