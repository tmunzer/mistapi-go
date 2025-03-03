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

// OrgsIntegrationSkyATP represents a controller struct.
type OrgsIntegrationSkyATP struct {
    baseController
}

// NewOrgsIntegrationSkyATP creates a new instance of OrgsIntegrationSkyATP.
// It takes a baseController as a parameter and returns a pointer to the OrgsIntegrationSkyATP.
func NewOrgsIntegrationSkyATP(baseController baseController) *OrgsIntegrationSkyATP {
    orgsIntegrationSkyATP := OrgsIntegrationSkyATP{baseController: baseController}
    return &orgsIntegrationSkyATP
}

// DeleteOrgSkyAtpIntegration takes context, orgId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete SkyATP Integration
func (o *OrgsIntegrationSkyATP) DeleteOrgSkyAtpIntegration(
    ctx context.Context,
    orgId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/setting/skyatp/setup")
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// GetOrgSkyAtpIntegration takes context, orgId as parameters and
// returns an models.ApiResponse with models.AccountSkyatpInfo data and
// an error if there was an issue with the request or response.
// Get Org SkyATP Integration
func (o *OrgsIntegrationSkyATP) GetOrgSkyAtpIntegration(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.AccountSkyatpInfo],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/setting/skyatp/setup")
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
    
    var result models.AccountSkyatpInfo
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AccountSkyatpInfo](decoder)
    return models.NewApiResponse(result, resp), err
}

// SetupOrgAtpIntegration takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.AccountSkyatpInfo data and
// an error if there was an issue with the request or response.
// 1. Login to the Sky ATP realm through the Mist UI by providing the realm, username and password.
// 2. Sky ATP API is invoked which creates the realm using above details.
// 3. Sky ATP by default will provide functionality for Security-Intelligence and Advanced Anti Malware.
// 4. Security Intelligence will provide configuration for CC, DNS Feeds, Infected Host, Blocklists and Allowlists.
func (o *OrgsIntegrationSkyATP) SetupOrgAtpIntegration(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountSkyatpConfig) (
    models.ApiResponse[models.AccountSkyatpInfo],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/setting/skyatp/setup")
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
    
    var result models.AccountSkyatpInfo
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AccountSkyatpInfo](decoder)
    return models.NewApiResponse(result, resp), err
}
