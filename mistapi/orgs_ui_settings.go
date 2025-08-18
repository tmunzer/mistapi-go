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

// OrgsUISettings represents a controller struct.
type OrgsUISettings struct {
    baseController
}

// NewOrgsUISettings creates a new instance of OrgsUISettings.
// It takes a baseController as a parameter and returns a pointer to the OrgsUISettings.
func NewOrgsUISettings(baseController baseController) *OrgsUISettings {
    orgsUISettings := OrgsUISettings{baseController: baseController}
    return &orgsUISettings
}

// ListOrgUiSettings takes context, orgId as parameters and
// returns an models.ApiResponse with []models.OrgUiSettings data and
// an error if there was an issue with the request or response.
// List the Orgs UI settings/databoard
func (o *OrgsUISettings) ListOrgUiSettings(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.OrgUiSettings],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/uisettings")
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
    
    var result []models.OrgUiSettings
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.OrgUiSettings](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgUiSettings takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.OrgUiSettings data and
// an error if there was an issue with the request or response.
// Create an Org UI settings/databoard
func (o *OrgsUISettings) CreateOrgUiSettings(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.OrgUiSettings) (
    models.ApiResponse[models.OrgUiSettings],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/uisettings")
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
    
    var result models.OrgUiSettings
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.OrgUiSettings](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgUiSetting takes context, orgId, uisettingId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete an Org UI settings
func (o *OrgsUISettings) DeleteOrgUiSetting(
    ctx context.Context,
    orgId uuid.UUID,
    uisettingId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/uisettings/%v")
    req.AppendTemplateParams(orgId, uisettingId)
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

// GetOrgUiSetting takes context, orgId, uisettingId as parameters and
// returns an models.ApiResponse with models.OrgUiSettings data and
// an error if there was an issue with the request or response.
// Get an Org UI settings/databoard
func (o *OrgsUISettings) GetOrgUiSetting(
    ctx context.Context,
    orgId uuid.UUID,
    uisettingId uuid.UUID) (
    models.ApiResponse[models.OrgUiSettings],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/uisettings/%v")
    req.AppendTemplateParams(orgId, uisettingId)
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
    
    var result models.OrgUiSettings
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.OrgUiSettings](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgUiSetting takes context, orgId, uisettingId, body as parameters and
// returns an models.ApiResponse with models.OrgUiSettings data and
// an error if there was an issue with the request or response.
// Org UI settings/databoard
func (o *OrgsUISettings) UpdateOrgUiSetting(
    ctx context.Context,
    orgId uuid.UUID,
    uisettingId uuid.UUID,
    body *models.OrgUiSettings) (
    models.ApiResponse[models.OrgUiSettings],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/uisettings/%v")
    req.AppendTemplateParams(orgId, uisettingId)
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
    
    var result models.OrgUiSettings
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.OrgUiSettings](decoder)
    return models.NewApiResponse(result, resp), err
}
