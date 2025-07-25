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

// SitesMxEdges represents a controller struct.
type SitesMxEdges struct {
    baseController
}

// NewSitesMxEdges creates a new instance of SitesMxEdges.
// It takes a baseController as a parameter and returns a pointer to the SitesMxEdges.
func NewSitesMxEdges(baseController baseController) *SitesMxEdges {
    sitesMxEdges := SitesMxEdges{baseController: baseController}
    return &sitesMxEdges
}

// ListSiteMxEdges takes context, siteId, limit, page as parameters and
// returns an models.ApiResponse with []models.Mxedge data and
// an error if there was an issue with the request or response.
// Get List of Site Mist Edges
func (s *SitesMxEdges) ListSiteMxEdges(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Mxedge],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/mxedges")
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
    
    var result []models.Mxedge
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Mxedge](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountSiteMxEdgeEvents takes context, siteId, distinct, mxedgeId, mxclusterId, mType, service, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Mist Edge Events
func (s *SitesMxEdges) CountSiteMxEdgeEvents(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteMxedgeEventsCountDistinctEnum,
    mxedgeId *string,
    mxclusterId *string,
    mType *string,
    service *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/mxedges/events/count")
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
    if mxedgeId != nil {
        req.QueryParam("mxedge_id", *mxedgeId)
    }
    if mxclusterId != nil {
        req.QueryParam("mxcluster_id", *mxclusterId)
    }
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if service != nil {
        req.QueryParam("service", *service)
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

// SearchSiteMistEdgeEvents takes context, siteId, mxedgeId, mxclusterId, mType, service, component, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseMxedgeEventsSearch data and
// an error if there was an issue with the request or response.
// Search Site Mist Edge Events
func (s *SitesMxEdges) SearchSiteMistEdgeEvents(
    ctx context.Context,
    siteId uuid.UUID,
    mxedgeId *string,
    mxclusterId *string,
    mType *string,
    service *string,
    component *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseMxedgeEventsSearch],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/mxedges/events/search")
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
    if mxedgeId != nil {
        req.QueryParam("mxedge_id", *mxedgeId)
    }
    if mxclusterId != nil {
        req.QueryParam("mxcluster_id", *mxclusterId)
    }
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if service != nil {
        req.QueryParam("service", *service)
    }
    if component != nil {
        req.QueryParam("component", *component)
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
    
    var result models.ResponseMxedgeEventsSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseMxedgeEventsSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSiteMxEdge takes context, siteId, mxedgeId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site Mist Edge
func (s *SitesMxEdges) DeleteSiteMxEdge(
    ctx context.Context,
    siteId uuid.UUID,
    mxedgeId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/mxedges/%v")
    req.AppendTemplateParams(siteId, mxedgeId)
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

// GetSiteMxEdge takes context, siteId, mxedgeId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Get Site Mist Edge
func (s *SitesMxEdges) GetSiteMxEdge(
    ctx context.Context,
    siteId uuid.UUID,
    mxedgeId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/mxedges/%v")
    req.AppendTemplateParams(siteId, mxedgeId)
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

// UpdateSiteMxEdge takes context, siteId, mxedgeId, body as parameters and
// returns an models.ApiResponse with models.Mxedge data and
// an error if there was an issue with the request or response.
// Update Site Mist Edge settings
func (s *SitesMxEdges) UpdateSiteMxEdge(
    ctx context.Context,
    siteId uuid.UUID,
    mxedgeId uuid.UUID,
    body *models.Mxedge) (
    models.ApiResponse[models.Mxedge],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/mxedges/%v")
    req.AppendTemplateParams(siteId, mxedgeId)
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
    
    var result models.Mxedge
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Mxedge](decoder)
    return models.NewApiResponse(result, resp), err
}

// UploadSiteMxEdgeSupportFiles takes context, siteId, mxedgeId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Support / Upload Mist Edge support files
func (s *SitesMxEdges) UploadSiteMxEdgeSupportFiles(
    ctx context.Context,
    siteId uuid.UUID,
    mxedgeId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/mxedges/%v/support")
    req.AppendTemplateParams(siteId, mxedgeId)
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
