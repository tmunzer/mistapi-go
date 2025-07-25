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

// OrgsIDPProfiles represents a controller struct.
type OrgsIDPProfiles struct {
    baseController
}

// NewOrgsIDPProfiles creates a new instance of OrgsIDPProfiles.
// It takes a baseController as a parameter and returns a pointer to the OrgsIDPProfiles.
func NewOrgsIDPProfiles(baseController baseController) *OrgsIDPProfiles {
    orgsIDPProfiles := OrgsIDPProfiles{baseController: baseController}
    return &orgsIDPProfiles
}

// ListOrgIdpProfiles takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.IdpProfile data and
// an error if there was an issue with the request or response.
// Get the list of Org IDP Profiles
func (o *OrgsIDPProfiles) ListOrgIdpProfiles(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.IdpProfile],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/idpprofiles")
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
    
    var result []models.IdpProfile
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.IdpProfile](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgIdpProfile takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.IdpProfile data and
// an error if there was an issue with the request or response.
// Create Org IDP Profile
func (o *OrgsIDPProfiles) CreateOrgIdpProfile(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.IdpProfile) (
    models.ApiResponse[models.IdpProfile],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/idpprofiles")
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
    
    var result models.IdpProfile
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.IdpProfile](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgIdpProfile takes context, orgId, idpprofileId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org IDP Profile
func (o *OrgsIDPProfiles) DeleteOrgIdpProfile(
    ctx context.Context,
    orgId uuid.UUID,
    idpprofileId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/idpprofiles/%v")
    req.AppendTemplateParams(orgId, idpprofileId)
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

// GetOrgIdpProfile takes context, orgId, idpprofileId as parameters and
// returns an models.ApiResponse with models.IdpProfile data and
// an error if there was an issue with the request or response.
// Get Org IDP Profile
func (o *OrgsIDPProfiles) GetOrgIdpProfile(
    ctx context.Context,
    orgId uuid.UUID,
    idpprofileId uuid.UUID) (
    models.ApiResponse[models.IdpProfile],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/idpprofiles/%v")
    req.AppendTemplateParams(orgId, idpprofileId)
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
    
    var result models.IdpProfile
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.IdpProfile](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgIdpProfile takes context, orgId, idpprofileId, body as parameters and
// returns an models.ApiResponse with models.IdpProfile data and
// an error if there was an issue with the request or response.
// Update Org IDP Profile
func (o *OrgsIDPProfiles) UpdateOrgIdpProfile(
    ctx context.Context,
    orgId uuid.UUID,
    idpprofileId uuid.UUID,
    body *models.IdpProfile) (
    models.ApiResponse[models.IdpProfile],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/idpprofiles/%v")
    req.AppendTemplateParams(orgId, idpprofileId)
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
    
    var result models.IdpProfile
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.IdpProfile](decoder)
    return models.NewApiResponse(result, resp), err
}
