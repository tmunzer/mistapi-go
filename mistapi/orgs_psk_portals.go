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

// OrgsPskPortals represents a controller struct.
type OrgsPskPortals struct {
    baseController
}

// NewOrgsPskPortals creates a new instance of OrgsPskPortals.
// It takes a baseController as a parameter and returns a pointer to the OrgsPskPortals.
func NewOrgsPskPortals(baseController baseController) *OrgsPskPortals {
    orgsPskPortals := OrgsPskPortals{baseController: baseController}
    return &orgsPskPortals
}

// ListOrgPskPortals takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.PskPortal data and
// an error if there was an issue with the request or response.
// Get List of Org Psk Portals
func (o *OrgsPskPortals) ListOrgPskPortals(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.PskPortal],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/pskportals")
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
    
    var result []models.PskPortal
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.PskPortal](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgPskPortal takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.PskPortal data and
// an error if there was an issue with the request or response.
// Create Org Psk Portal
func (o *OrgsPskPortals) CreateOrgPskPortal(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.PskPortal) (
    models.ApiResponse[models.PskPortal],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/pskportals")
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
    
    var result models.PskPortal
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.PskPortal](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListOrgPskPortalLogs takes context, orgId, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.ResponsePskPortalLogsSearch data and
// an error if there was an issue with the request or response.
// Get the list of PSK Portals Logs
func (o *OrgsPskPortals) ListOrgPskPortalLogs(
    ctx context.Context,
    orgId uuid.UUID,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponsePskPortalLogsSearch],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/pskportals/logs")
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
    
    var result models.ResponsePskPortalLogsSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponsePskPortalLogsSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountOrgPskPortalLogs takes context, orgId, distinct, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of PskPortal Logs
func (o *OrgsPskPortals) CountOrgPskPortalLogs(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgPskPortalLogsCountDistinctEnum,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/pskportals/logs/count")
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
    
    var result models.ResponseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchOrgPskPortalLogs takes context, orgId, start, end, duration, limit, page, pskName, pskId, pskportalId, id, adminName, adminId, nameId as parameters and
// returns an models.ApiResponse with models.ResponsePskPortalLogsSearch data and
// an error if there was an issue with the request or response.
// Search Org PSK Portal Logs
func (o *OrgsPskPortals) SearchOrgPskPortalLogs(
    ctx context.Context,
    orgId uuid.UUID,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int,
    pskName *string,
    pskId *string,
    pskportalId *string,
    id *uuid.UUID,
    adminName *string,
    adminId *string,
    nameId *uuid.UUID) (
    models.ApiResponse[models.ResponsePskPortalLogsSearch],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/pskportals/logs/search")
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
    if pskName != nil {
        req.QueryParam("psk_name", *pskName)
    }
    if pskId != nil {
        req.QueryParam("psk_id", *pskId)
    }
    if pskportalId != nil {
        req.QueryParam("pskportal_id", *pskportalId)
    }
    if id != nil {
        req.QueryParam("id", *id)
    }
    if adminName != nil {
        req.QueryParam("admin_name", *adminName)
    }
    if adminId != nil {
        req.QueryParam("admin_id", *adminId)
    }
    if nameId != nil {
        req.QueryParam("name_id", *nameId)
    }
    
    var result models.ResponsePskPortalLogsSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponsePskPortalLogsSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgPskPortal takes context, orgId, pskportalId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org Psk Portal
func (o *OrgsPskPortals) DeleteOrgPskPortal(
    ctx context.Context,
    orgId uuid.UUID,
    pskportalId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/pskportals/%v")
    req.AppendTemplateParams(orgId, pskportalId)
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

// GetOrgPskPortal takes context, orgId, pskportalId as parameters and
// returns an models.ApiResponse with models.PskPortal data and
// an error if there was an issue with the request or response.
// Get Org Psk Portal Details
func (o *OrgsPskPortals) GetOrgPskPortal(
    ctx context.Context,
    orgId uuid.UUID,
    pskportalId uuid.UUID) (
    models.ApiResponse[models.PskPortal],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/pskportals/%v")
    req.AppendTemplateParams(orgId, pskportalId)
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
    
    var result models.PskPortal
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.PskPortal](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgPskPortal takes context, orgId, pskportalId, body as parameters and
// returns an models.ApiResponse with models.PskPortal data and
// an error if there was an issue with the request or response.
// Update Org Psk Portal
func (o *OrgsPskPortals) UpdateOrgPskPortal(
    ctx context.Context,
    orgId uuid.UUID,
    pskportalId uuid.UUID,
    body *models.PskPortal) (
    models.ApiResponse[models.PskPortal],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/pskportals/%v")
    req.AppendTemplateParams(orgId, pskportalId)
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
    
    var result models.PskPortal
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.PskPortal](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgPskPortalImage takes context, orgId, pskportalId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete background image for PskPortal
// If image is not uploaded or is deleted, PskPortal will use default image.
func (o *OrgsPskPortals) DeleteOrgPskPortalImage(
    ctx context.Context,
    orgId uuid.UUID,
    pskportalId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "DELETE",
      "/api/v1/orgs/%v/pskportals/%v/portal_image",
    )
    req.AppendTemplateParams(orgId, pskportalId)
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

// UploadOrgPskPortalImage takes context, orgId, pskportalId, file, json as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Upload background image for PskPortal
func (o *OrgsPskPortals) UploadOrgPskPortalImage(
    ctx context.Context,
    orgId uuid.UUID,
    pskportalId uuid.UUID,
    file *models.FileWrapper,
    json *string) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      "/api/v1/orgs/%v/pskportals/%v/portal_image",
    )
    req.AppendTemplateParams(orgId, pskportalId)
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
    if file != nil {
        fileParam := https.FormParam{Key: "file", Value: *file, Headers: http.Header{}}
        formFields = append(formFields, fileParam)
    }
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

// UpdateOrgPskPortalTemplate takes context, orgId, pskportalId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Update Org Psk Portal Template
func (o *OrgsPskPortals) UpdateOrgPskPortalTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    pskportalId uuid.UUID,
    body *models.PskPortalTemplate) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "PUT",
      "/api/v1/orgs/%v/pskportals/%v/portal_template",
    )
    req.AppendTemplateParams(orgId, pskportalId)
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}
