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

// OrgsSSORoles represents a controller struct.
type OrgsSSORoles struct {
    baseController
}

// NewOrgsSSORoles creates a new instance of OrgsSSORoles.
// It takes a baseController as a parameter and returns a pointer to the OrgsSSORoles.
func NewOrgsSSORoles(baseController baseController) *OrgsSSORoles {
    orgsSSORoles := OrgsSSORoles{baseController: baseController}
    return &orgsSSORoles
}

// ListOrgSsoRoles takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.SsoRoleMsp data and
// an error if there was an issue with the request or response.
// Get List of Org SSO Roles
func (o *OrgsSSORoles) ListOrgSsoRoles(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.SsoRoleMsp],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/ssoroles")
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
    
    var result []models.SsoRoleMsp
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.SsoRoleMsp](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgSsoRole takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.SsoRoleOrg data and
// an error if there was an issue with the request or response.
// Create Org SSO Role
func (o *OrgsSSORoles) CreateOrgSsoRole(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.SsoRoleOrg) (
    models.ApiResponse[models.SsoRoleOrg],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/ssoroles")
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
    
    var result models.SsoRoleOrg
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SsoRoleOrg](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgSsoRole takes context, orgId, ssoroleId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org SSO Role
func (o *OrgsSSORoles) DeleteOrgSsoRole(
    ctx context.Context,
    orgId uuid.UUID,
    ssoroleId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/ssoroles/%v")
    req.AppendTemplateParams(orgId, ssoroleId)
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

// GetOrgSsoRole takes context, orgId, ssoroleId as parameters and
// returns an models.ApiResponse with models.SsoRoleOrg data and
// an error if there was an issue with the request or response.
// Get Org SSO Role Details
func (o *OrgsSSORoles) GetOrgSsoRole(
    ctx context.Context,
    orgId uuid.UUID,
    ssoroleId uuid.UUID) (
    models.ApiResponse[models.SsoRoleOrg],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/ssoroles/%v")
    req.AppendTemplateParams(orgId, ssoroleId)
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
    
    var result models.SsoRoleOrg
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SsoRoleOrg](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgSsoRole takes context, orgId, ssoroleId, body as parameters and
// returns an models.ApiResponse with models.SsoRoleOrg data and
// an error if there was an issue with the request or response.
// Update Org SSO Role
func (o *OrgsSSORoles) UpdateOrgSsoRole(
    ctx context.Context,
    orgId uuid.UUID,
    ssoroleId uuid.UUID,
    body *models.SsoRoleOrg) (
    models.ApiResponse[models.SsoRoleOrg],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/ssoroles/%v")
    req.AppendTemplateParams(orgId, ssoroleId)
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
    
    var result models.SsoRoleOrg
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SsoRoleOrg](decoder)
    return models.NewApiResponse(result, resp), err
}
