package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// MSPsLogo represents a controller struct.
type MSPsLogo struct {
    baseController
}

// NewMSPsLogo creates a new instance of MSPsLogo.
// It takes a baseController as a parameter and returns a pointer to the MSPsLogo.
func NewMSPsLogo(baseController baseController) *MSPsLogo {
    mSPsLogo := MSPsLogo{baseController: baseController}
    return &mSPsLogo
}

// DeleteMspLogo takes context, mspId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete MSP Logo
func (m *MSPsLogo) DeleteMspLogo(
    ctx context.Context,
    mspId uuid.UUID) (
    *http.Response,
    error) {
    req := m.prepareRequest(ctx, "DELETE", "/api/v1/msps/%v/logo")
    req.AppendTemplateParams(mspId)
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

// PostMspLogo takes context, mspId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Upload Logo (only for advanced msp tier)
func (m *MSPsLogo) PostMspLogo(
    ctx context.Context,
    mspId uuid.UUID,
    body *models.MspLogo) (
    *http.Response,
    error) {
    req := m.prepareRequest(ctx, "POST", "/api/v1/msps/%v/logo")
    req.AppendTemplateParams(mspId)
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
