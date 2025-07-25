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

// OrgsNACTags represents a controller struct.
type OrgsNACTags struct {
    baseController
}

// NewOrgsNACTags creates a new instance of OrgsNACTags.
// It takes a baseController as a parameter and returns a pointer to the OrgsNACTags.
func NewOrgsNACTags(baseController baseController) *OrgsNACTags {
    orgsNACTags := OrgsNACTags{baseController: baseController}
    return &orgsNACTags
}

// ListOrgNacTags takes context, orgId, mType, name, match, limit, page as parameters and
// returns an models.ApiResponse with []models.NacTag data and
// an error if there was an issue with the request or response.
// Get List of Org NAC Tags
func (o *OrgsNACTags) ListOrgNacTags(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.NacTagTypeEnum,
    name *string,
    match *models.NacTagMatchEnum,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.NacTag],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/nactags")
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
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if name != nil {
        req.QueryParam("name", *name)
    }
    if match != nil {
        req.QueryParam("match", *match)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result []models.NacTag
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.NacTag](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgNacTag takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.NacTag data and
// an error if there was an issue with the request or response.
// Create Org NAC Tag
func (o *OrgsNACTags) CreateOrgNacTag(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.NacTag) (
    models.ApiResponse[models.NacTag],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/nactags")
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
    
    var result models.NacTag
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.NacTag](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgNacTag takes context, orgId, nactagId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org NAC Tag
func (o *OrgsNACTags) DeleteOrgNacTag(
    ctx context.Context,
    orgId uuid.UUID,
    nactagId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/nactags/%v")
    req.AppendTemplateParams(orgId, nactagId)
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

// GetOrgNacTag takes context, orgId, nactagId as parameters and
// returns an models.ApiResponse with models.NacTag data and
// an error if there was an issue with the request or response.
// Get Org NAC Tag
func (o *OrgsNACTags) GetOrgNacTag(
    ctx context.Context,
    orgId uuid.UUID,
    nactagId uuid.UUID) (
    models.ApiResponse[models.NacTag],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/nactags/%v")
    req.AppendTemplateParams(orgId, nactagId)
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
    
    var result models.NacTag
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.NacTag](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgNacTag takes context, orgId, nactagId, body as parameters and
// returns an models.ApiResponse with models.NacTag data and
// an error if there was an issue with the request or response.
// Update Org NAC Tag
func (o *OrgsNACTags) UpdateOrgNacTag(
    ctx context.Context,
    orgId uuid.UUID,
    nactagId uuid.UUID,
    body *models.NacTag) (
    models.ApiResponse[models.NacTag],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/nactags/%v")
    req.AppendTemplateParams(orgId, nactagId)
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
    
    var result models.NacTag
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.NacTag](decoder)
    return models.NewApiResponse(result, resp), err
}
