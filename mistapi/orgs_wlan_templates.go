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

// OrgsWLANTemplates represents a controller struct.
type OrgsWLANTemplates struct {
    baseController
}

// NewOrgsWLANTemplates creates a new instance of OrgsWLANTemplates.
// It takes a baseController as a parameter and returns a pointer to the OrgsWLANTemplates.
func NewOrgsWLANTemplates(baseController baseController) *OrgsWLANTemplates {
    orgsWLANTemplates := OrgsWLANTemplates{baseController: baseController}
    return &orgsWLANTemplates
}

// ListOrgTemplates takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.Template data and
// an error if there was an issue with the request or response.
// Get List of Org WLAN Templates
func (o *OrgsWLANTemplates) ListOrgTemplates(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Template],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/templates")
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
    
    var result []models.Template
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Template](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgTemplate takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Template data and
// an error if there was an issue with the request or response.
// Create Org Template
func (o *OrgsWLANTemplates) CreateOrgTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Template) (
    models.ApiResponse[models.Template],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/templates")
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
    
    var result models.Template
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Template](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgTemplate takes context, orgId, templateId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org Template
func (o *OrgsWLANTemplates) DeleteOrgTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    templateId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/templates/%v")
    req.AppendTemplateParams(orgId, templateId)
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

// GetOrgTemplate takes context, orgId, templateId as parameters and
// returns an models.ApiResponse with models.Template data and
// an error if there was an issue with the request or response.
// Get Org Template Details
func (o *OrgsWLANTemplates) GetOrgTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    templateId uuid.UUID) (
    models.ApiResponse[models.Template],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/templates/%v")
    req.AppendTemplateParams(orgId, templateId)
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
    
    var result models.Template
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Template](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgTemplate takes context, orgId, templateId, body as parameters and
// returns an models.ApiResponse with models.Template data and
// an error if there was an issue with the request or response.
// Update Org Template
func (o *OrgsWLANTemplates) UpdateOrgTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    templateId uuid.UUID,
    body *models.Template) (
    models.ApiResponse[models.Template],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/templates/%v")
    req.AppendTemplateParams(orgId, templateId)
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
    
    var result models.Template
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Template](decoder)
    return models.NewApiResponse(result, resp), err
}

// CloneOrgTemplate takes context, orgId, templateId, body as parameters and
// returns an models.ApiResponse with models.Template data and
// an error if there was an issue with the request or response.
// Clone Org Template
func (o *OrgsWLANTemplates) CloneOrgTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    templateId uuid.UUID,
    body *models.NameString) (
    models.ApiResponse[models.Template],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/templates/%v/clone")
    req.AppendTemplateParams(orgId, templateId)
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
    
    var result models.Template
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Template](decoder)
    return models.NewApiResponse(result, resp), err
}
