package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// OrgsJSE represents a controller struct.
type OrgsJSE struct {
    baseController
}

// NewOrgsJSE creates a new instance of OrgsJSE.
// It takes a baseController as a parameter and returns a pointer to the OrgsJSE.
func NewOrgsJSE(baseController baseController) *OrgsJSE {
    orgsJSE := OrgsJSE{baseController: baseController}
    return &orgsJSE
}

// GetOrgJseInfo takes context, orgId as parameters and
// returns an models.ApiResponse with models.AccountJseInfo data and
// an error if there was an issue with the request or response.
// Retrieves the list of JSE orgs associated with the account.
func (o *OrgsJSE) GetOrgJseInfo(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.AccountJseInfo],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/setting/jse/info", orgId),
    )
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    
    var result models.AccountJseInfo
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AccountJseInfo](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgJsecCredential takes context, orgId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete JSE credential
func (o *OrgsJSE) DeleteOrgJsecCredential(
    ctx context.Context,
    orgId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/orgs/%v/setting/jse/setup", orgId),
    )
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// GetOrgJsecCredential takes context, orgId as parameters and
// returns an models.ApiResponse with models.AccountJseInfo data and
// an error if there was an issue with the request or response.
// Get Org JSE Credential
func (o *OrgsJSE) GetOrgJsecCredential(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.AccountJseInfo],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/setting/jse/setup", orgId),
    )
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    
    var result models.AccountJseInfo
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AccountJseInfo](decoder)
    return models.NewApiResponse(result, resp), err
}

// SetupOrgJsecCredential takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.AccountJseInfo data and
// an error if there was an issue with the request or response.
// in JSE UI: 
// 1. Create custom role with Read access to service_location and RW access to site and IPSec profile APIs. 
// 2. Create a user with the above custom role. - email: john@abc.com 
// 3. Activate the user in the JSE account. 
// 4. Create the service locations on the JSE account.
func (o *OrgsJSE) SetupOrgJsecCredential(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountJseConfig) (
    models.ApiResponse[models.AccountJseInfo],
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/setting/jse/setup", orgId),
    )
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
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
