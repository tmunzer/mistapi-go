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

// OrgsIntegrationJSE represents a controller struct.
type OrgsIntegrationJSE struct {
    baseController
}

// NewOrgsIntegrationJSE creates a new instance of OrgsIntegrationJSE.
// It takes a baseController as a parameter and returns a pointer to the OrgsIntegrationJSE.
func NewOrgsIntegrationJSE(baseController baseController) *OrgsIntegrationJSE {
    orgsIntegrationJSE := OrgsIntegrationJSE{baseController: baseController}
    return &orgsIntegrationJSE
}

// GetOrgJseInfo takes context, orgId as parameters and
// returns an models.ApiResponse with models.AccountJseInfo data and
// an error if there was an issue with the request or response.
// Retrieves the list of JSE orgs associated with the account.
func (o *OrgsIntegrationJSE) GetOrgJseInfo(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.AccountJseInfo],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/setting/jse/info")
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
    
    var result models.AccountJseInfo
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AccountJseInfo](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgJseIntegration takes context, orgId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete JSE Integration
func (o *OrgsIntegrationJSE) DeleteOrgJseIntegration(
    ctx context.Context,
    orgId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/setting/jse/setup")
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

// GetOrgJseIntegration takes context, orgId as parameters and
// returns an models.ApiResponse with models.AccountJseInfo data and
// an error if there was an issue with the request or response.
// Get Org JSE Integration
func (o *OrgsIntegrationJSE) GetOrgJseIntegration(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.AccountJseInfo],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/setting/jse/setup")
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
    
    var result models.AccountJseInfo
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AccountJseInfo](decoder)
    return models.NewApiResponse(result, resp), err
}

// SetupOrgJseIntegration takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.AccountJseInfo data and
// an error if there was an issue with the request or response.
// In JSE UI: 
// 1. Create custom role with Read access to service_location and RW access to site and IPSec profile APIs. 
// 2. Create a user with the above custom role. - email: john@abc.com 
// 3. Activate the user in the JSE account. 
// 4. Create the service locations on the JSE account.
func (o *OrgsIntegrationJSE) SetupOrgJseIntegration(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountJseConfig) (
    models.ApiResponse[models.AccountJseInfo],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/setting/jse/setup")
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
    
    var result models.AccountJseInfo
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AccountJseInfo](decoder)
    return models.NewApiResponse(result, resp), err
}
