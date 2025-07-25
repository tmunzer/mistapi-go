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

// OrgsNACRules represents a controller struct.
type OrgsNACRules struct {
    baseController
}

// NewOrgsNACRules creates a new instance of OrgsNACRules.
// It takes a baseController as a parameter and returns a pointer to the OrgsNACRules.
func NewOrgsNACRules(baseController baseController) *OrgsNACRules {
    orgsNACRules := OrgsNACRules{baseController: baseController}
    return &orgsNACRules
}

// ListOrgNacRules takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.NacRule data and
// an error if there was an issue with the request or response.
// Get List of Org NAC Rules
func (o *OrgsNACRules) ListOrgNacRules(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.NacRule],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/nacrules")
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
    
    var result []models.NacRule
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.NacRule](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgNacRule takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.NacRule data and
// an error if there was an issue with the request or response.
// Create Org NAC Rule
func (o *OrgsNACRules) CreateOrgNacRule(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.NacRule) (
    models.ApiResponse[models.NacRule],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/nacrules")
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
    
    var result models.NacRule
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.NacRule](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgNacRule takes context, orgId, nacruleId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org NAC Rule
func (o *OrgsNACRules) DeleteOrgNacRule(
    ctx context.Context,
    orgId uuid.UUID,
    nacruleId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/nacrules/%v")
    req.AppendTemplateParams(orgId, nacruleId)
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

// GetOrgNacRule takes context, orgId, nacruleId as parameters and
// returns an models.ApiResponse with models.NacRule data and
// an error if there was an issue with the request or response.
// Get Org NAC Rule
func (o *OrgsNACRules) GetOrgNacRule(
    ctx context.Context,
    orgId uuid.UUID,
    nacruleId uuid.UUID) (
    models.ApiResponse[models.NacRule],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/nacrules/%v")
    req.AppendTemplateParams(orgId, nacruleId)
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
    
    var result models.NacRule
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.NacRule](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgNacRule takes context, orgId, nacruleId, body as parameters and
// returns an models.ApiResponse with models.NacRule data and
// an error if there was an issue with the request or response.
// Update Org NAC Rule
func (o *OrgsNACRules) UpdateOrgNacRule(
    ctx context.Context,
    orgId uuid.UUID,
    nacruleId uuid.UUID,
    body *models.NacRule) (
    models.ApiResponse[models.NacRule],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/nacrules/%v")
    req.AppendTemplateParams(orgId, nacruleId)
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
    
    var result models.NacRule
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.NacRule](decoder)
    return models.NewApiResponse(result, resp), err
}
