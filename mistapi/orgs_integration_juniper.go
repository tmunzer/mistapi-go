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

// OrgsIntegrationJuniper represents a controller struct.
type OrgsIntegrationJuniper struct {
    baseController
}

// NewOrgsIntegrationJuniper creates a new instance of OrgsIntegrationJuniper.
// It takes a baseController as a parameter and returns a pointer to the OrgsIntegrationJuniper.
func NewOrgsIntegrationJuniper(baseController baseController) *OrgsIntegrationJuniper {
    orgsIntegrationJuniper := OrgsIntegrationJuniper{baseController: baseController}
    return &orgsIntegrationJuniper
}

// LinkOrgToJuniperJuniperAccount takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.AccountJuniperInfo data and
// an error if there was an issue with the request or response.
// Link Juniper Accounts
func (o *OrgsIntegrationJuniper) LinkOrgToJuniperJuniperAccount(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountJuniperConfig) (
    models.ApiResponse[models.AccountJuniperInfo],
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      "/api/v1/orgs/%v/setting/juniper/link_accounts",
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
        "400": {Message: "Account already linked", Unmarshaller: errors.NewResponseDetailString},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.AccountJuniperInfo
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AccountJuniperInfo](decoder)
    return models.NewApiResponse(result, resp), err
}

// UnlinkOrgFromJuniperCustomerId takes context, orgId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Unlink Juniper Customer ID
// `linked_by` field is only required if there are duplicate account_names.
func (o *OrgsIntegrationJuniper) UnlinkOrgFromJuniperCustomerId(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountJuniperInfo) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "DELETE",
      "/api/v1/orgs/%v/setting/juniper/unlink_account",
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
