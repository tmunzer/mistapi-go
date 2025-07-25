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

// SitesRfdiags represents a controller struct.
type SitesRfdiags struct {
    baseController
}

// NewSitesRfdiags creates a new instance of SitesRfdiags.
// It takes a baseController as a parameter and returns a pointer to the SitesRfdiags.
func NewSitesRfdiags(baseController baseController) *SitesRfdiags {
    sitesRfdiags := SitesRfdiags{baseController: baseController}
    return &sitesRfdiags
}

// GetSiteSiteRfdiagRecording takes context, siteId, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with [][]models.RfDiagInfoItem data and
// an error if there was an issue with the request or response.
// List RF Glass Recording
func (s *SitesRfdiags) GetSiteSiteRfdiagRecording(
    ctx context.Context,
    siteId uuid.UUID,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[[][]models.RfDiagInfoItem],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/rfdiags")
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
    
    var result [][]models.RfDiagInfoItem
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[][]models.RfDiagInfoItem](decoder)
    return models.NewApiResponse(result, resp), err
}

// StartSiteRecording takes context, siteId, body as parameters and
// returns an models.ApiResponse with []models.RfDiagInfoItem data and
// an error if there was an issue with the request or response.
// Start RF Glass Recording
func (s *SitesRfdiags) StartSiteRecording(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.RfDiag) (
    models.ApiResponse[[]models.RfDiagInfoItem],
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/rfdiags")
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
    
    var result []models.RfDiagInfoItem
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.RfDiagInfoItem](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSiteRfdiagRecording takes context, siteId, rfdiagId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Recording
func (s *SitesRfdiags) DeleteSiteRfdiagRecording(
    ctx context.Context,
    siteId uuid.UUID,
    rfdiagId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/rfdiags/%v")
    req.AppendTemplateParams(siteId, rfdiagId)
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

// GetSiteRfdiagRecording takes context, siteId, rfdiagId as parameters and
// returns an models.ApiResponse with []models.RfDiagInfoItem data and
// an error if there was an issue with the request or response.
// Get RF Diag Recording Details
func (s *SitesRfdiags) GetSiteRfdiagRecording(
    ctx context.Context,
    siteId uuid.UUID,
    rfdiagId uuid.UUID) (
    models.ApiResponse[[]models.RfDiagInfoItem],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/rfdiags/%v")
    req.AppendTemplateParams(siteId, rfdiagId)
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
    
    var result []models.RfDiagInfoItem
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.RfDiagInfoItem](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSiteRfdiagRecording takes context, siteId, rfdiagId, body as parameters and
// returns an models.ApiResponse with []models.RfDiagInfoItem data and
// an error if there was an issue with the request or response.
// Update Recording
func (s *SitesRfdiags) UpdateSiteRfdiagRecording(
    ctx context.Context,
    siteId uuid.UUID,
    rfdiagId uuid.UUID,
    body *models.RfDiag) (
    models.ApiResponse[[]models.RfDiagInfoItem],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/rfdiags/%v")
    req.AppendTemplateParams(siteId, rfdiagId)
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
    
    var result []models.RfDiagInfoItem
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.RfDiagInfoItem](decoder)
    return models.NewApiResponse(result, resp), err
}

// DownloadSiteRfdiagRecording takes context, siteId, rfdiagId as parameters and
// returns an models.ApiResponse with []byte data and
// an error if there was an issue with the request or response.
// Download Recording
// Download raw_events blob
func (s *SitesRfdiags) DownloadSiteRfdiagRecording(
    ctx context.Context,
    siteId uuid.UUID,
    rfdiagId uuid.UUID) (
    models.ApiResponse[[]byte],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/rfdiags/%v/download")
    req.AppendTemplateParams(siteId, rfdiagId)
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
    
    stream, resp, err := req.CallAsStream()
    if err != nil {
        return models.NewApiResponse(stream, resp), err
    }
    return models.NewApiResponse(stream, resp), err
}

// StopSiteRfdiagRecording takes context, siteId, rfdiagId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// If the recording session is active for the given rfdiag_id, it will finish the recording. duration and end_time will be updated to reflect the correct values.
func (s *SitesRfdiags) StopSiteRfdiagRecording(
    ctx context.Context,
    siteId uuid.UUID,
    rfdiagId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/rfdiags/%v/stop")
    req.AppendTemplateParams(siteId, rfdiagId)
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
