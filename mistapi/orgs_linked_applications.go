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

// OrgsLinkedApplications represents a controller struct.
type OrgsLinkedApplications struct {
    baseController
}

// NewOrgsLinkedApplications creates a new instance of OrgsLinkedApplications.
// It takes a baseController as a parameter and returns a pointer to the OrgsLinkedApplications.
func NewOrgsLinkedApplications(baseController baseController) *OrgsLinkedApplications {
    orgsLinkedApplications := OrgsLinkedApplications{baseController: baseController}
    return &orgsLinkedApplications
}

// LinkOrgToJuniperJuniperAccount takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.AccountJuniperInfo data and
// an error if there was an issue with the request or response.
// Link Juniper Accounts
func (o *OrgsLinkedApplications) LinkOrgToJuniperJuniperAccount(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountJuniperConfig) (
    models.ApiResponse[models.AccountJuniperInfo],
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/setting/juniper/link_accounts", orgId),
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
        "400": {Message: "account already linked", Unmarshaller: errors.NewResponseDetailString},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
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
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Unlink Juniper Customer ID
// `linked_by` field is only required if there are duplicate account_names.
func (o *OrgsLinkedApplications) UnlinkOrgFromJuniperCustomerId(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountJuniperInfo) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/orgs/%v/setting/juniper/unlink_account", orgId),
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// GetOrgOauthAppLinkedStatus takes context, orgId, appName, forward as parameters and
// returns an models.ApiResponse with models.ResponseOauthAppLink data and
// an error if there was an issue with the request or response.
// Get Org Level OAuth Application Linked Status
func (o *OrgsLinkedApplications) GetOrgOauthAppLinkedStatus(
    ctx context.Context,
    orgId uuid.UUID,
    appName models.OauthAppNameEnum,
    forward string) (
    models.ApiResponse[models.ResponseOauthAppLink],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/setting/%v/link_accounts", orgId, appName),
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
    req.QueryParam("forward", forward)
    
    var result models.ResponseOauthAppLink
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseOauthAppLink](decoder)
    return models.NewApiResponse(result, resp), err
}

// AddOrgOauthAppAccounts takes context, orgId, appName, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Add Jamf, VMware Authorization With Mist Portal
func (o *OrgsLinkedApplications) AddOrgOauthAppAccounts(
    ctx context.Context,
    orgId uuid.UUID,
    appName models.OauthAppNameEnum,
    body *models.AccountOauthAdd) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/setting/%v/link_accounts", orgId, appName),
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
        "400": {Message: "Unsuccessful"},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// UpdateOrgOauthAppAccounts takes context, orgId, appName, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Update Zoom, Teams, Intune Authorization.
// Request Payload, These Field And Values Will Be Specific To Each Of The Third Party Apps Accounts.
func (o *OrgsLinkedApplications) UpdateOrgOauthAppAccounts(
    ctx context.Context,
    orgId uuid.UUID,
    appName models.OauthAppNameEnum,
    body *models.AccountOauthConfig) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/orgs/%v/setting/%v/link_accounts", orgId, appName),
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// DeleteOrgOauthAppAuthorization takes context, orgId, appName, accountId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete Org Level OAuth Application Authorization With Mist Portal
func (o *OrgsLinkedApplications) DeleteOrgOauthAppAuthorization(
    ctx context.Context,
    orgId uuid.UUID,
    appName models.OauthAppNameEnum,
    accountId string) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/orgs/%v/setting/%v/link_accounts/%v", orgId, appName, accountId),
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
        "400": {Message: "Unsuccessful"},
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
