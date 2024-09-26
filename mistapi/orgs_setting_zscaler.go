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

// OrgsSettingZscaler represents a controller struct.
type OrgsSettingZscaler struct {
    baseController
}

// NewOrgsSettingZscaler creates a new instance of OrgsSettingZscaler.
// It takes a baseController as a parameter and returns a pointer to the OrgsSettingZscaler.
func NewOrgsSettingZscaler(baseController baseController) *OrgsSettingZscaler {
    orgsSettingZscaler := OrgsSettingZscaler{baseController: baseController}
    return &orgsSettingZscaler
}

// DeleteOrgZscalerCredential takes context, orgId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// To delete Zscaler credential
func (o *OrgsSettingZscaler) DeleteOrgZscalerCredential(
    ctx context.Context,
    orgId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/orgs/%v/setting/zscaler/setup", orgId),
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

// GetOrgZscalerCredential takes context, orgId as parameters and
// returns an models.ApiResponse with models.AccountZscalerInfo data and
// an error if there was an issue with the request or response.
// To get Zscaler credential
func (o *OrgsSettingZscaler) GetOrgZscalerCredential(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.AccountZscalerInfo],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/setting/zscaler/setup", orgId),
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    
    var result models.AccountZscalerInfo
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AccountZscalerInfo](decoder)
    return models.NewApiResponse(result, resp), err
}

// SetupOrgZscalerCredential takes context, orgId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// To setup Zscaler credential
func (o *OrgsSettingZscaler) SetupOrgZscalerCredential(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountZscalerConfig) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/setting/zscaler/setup", orgId),
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
