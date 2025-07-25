// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// AdminsRecoverPassword represents a controller struct.
type AdminsRecoverPassword struct {
    baseController
}

// NewAdminsRecoverPassword creates a new instance of AdminsRecoverPassword.
// It takes a baseController as a parameter and returns a pointer to the AdminsRecoverPassword.
func NewAdminsRecoverPassword(baseController baseController) *AdminsRecoverPassword {
    adminsRecoverPassword := AdminsRecoverPassword{baseController: baseController}
    return &adminsRecoverPassword
}

// RecoverPassword takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Recover Password
// An email will also be sent to the user with a link to https://manage.mist.com/verify/recover?token=:token
func (a *AdminsRecoverPassword) RecoverPassword(
    ctx context.Context,
    body *models.Recover) (
    *http.Response,
    error) {
    req := a.prepareRequest(ctx, "POST", "/api/v1/recover")
    
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

// VerifyRecoverPassword takes context, token as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Verify Recover Password
// With correct verification, the user will be authenticated. UI can then prompt for new password
func (a *AdminsRecoverPassword) VerifyRecoverPassword(
    ctx context.Context,
    token string) (
    *http.Response,
    error) {
    req := a.prepareRequest(ctx, "POST", "/api/v1/recover/verify/%v")
    req.AppendTemplateParams(token)
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
