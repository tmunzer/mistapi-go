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

// OrgsIntegrationCradlepoint represents a controller struct.
type OrgsIntegrationCradlepoint struct {
    baseController
}

// NewOrgsIntegrationCradlepoint creates a new instance of OrgsIntegrationCradlepoint.
// It takes a baseController as a parameter and returns a pointer to the OrgsIntegrationCradlepoint.
func NewOrgsIntegrationCradlepoint(baseController baseController) *OrgsIntegrationCradlepoint {
    orgsIntegrationCradlepoint := OrgsIntegrationCradlepoint{baseController: baseController}
    return &orgsIntegrationCradlepoint
}

// DeleteOrgCradlepointConnection takes context, orgId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This deletes the Cradlepoint integration in Mist
func (o *OrgsIntegrationCradlepoint) DeleteOrgCradlepointConnection(
    ctx context.Context,
    orgId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "DELETE",
      "/api/v1/orgs/%v/setting/cradlepoint/setup",
    )
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

// TestOrgCradlepointConnection takes context, orgId as parameters and
// returns an models.ApiResponse with models.TestCradlepoint data and
// an error if there was an issue with the request or response.
// This tests the Cradlepoint integration in Mist
func (o *OrgsIntegrationCradlepoint) TestOrgCradlepointConnection(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.TestCradlepoint],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/setting/cradlepoint/setup")
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
    
    var result models.TestCradlepoint
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.TestCradlepoint](decoder)
    return models.NewApiResponse(result, resp), err
}

// SetupOrgCradlepointConnectionToMist takes context, orgId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This sets up cradlepoint webhooks to send events to Mist
func (o *OrgsIntegrationCradlepoint) SetupOrgCradlepointConnectionToMist(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountCradlepointConfig) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/setting/cradlepoint/setup")
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// UpdateOrgCradlepointConnectionToMist takes context, orgId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This updates the Cradlepoint integration settings in Mist
func (o *OrgsIntegrationCradlepoint) UpdateOrgCradlepointConnectionToMist(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountCradlepointConfig) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/setting/cradlepoint/setup")
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// SyncOrgCradlepointRouters takes context, orgId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This syncs cradlepoint devices with Mist. We’ll also attempt to use the LLDP data from cradlepoint to identify the linkage against Mist Site / Device
func (o *OrgsIntegrationCradlepoint) SyncOrgCradlepointRouters(
    ctx context.Context,
    orgId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/setting/cradlepoint/sync")
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
