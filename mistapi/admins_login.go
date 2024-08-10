package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// AdminsLogin represents a controller struct.
type AdminsLogin struct {
    baseController
}

// NewAdminsLogin creates a new instance of AdminsLogin.
// It takes a baseController as a parameter and returns a pointer to the AdminsLogin.
func NewAdminsLogin(baseController baseController) *AdminsLogin {
    adminsLogin := AdminsLogin{baseController: baseController}
    return &adminsLogin
}

// Login takes context, body as parameters and
// returns an models.ApiResponse with models.ResponseLoginSuccess data and
// an error if there was an issue with the request or response.
// Log in with email/password.
// When 2FA is enabled, there are two ways to login:
// 1. login with two_factor token (with Google Authenticator, etc) 
// 2. login with email/password, generate the token, and use /login/two_factor with the token
func (a *AdminsLogin) Login(
    ctx context.Context,
    body *models.Login) (
    models.ApiResponse[models.ResponseLoginSuccess],
    error) {
    req := a.prepareRequest(ctx, "POST", "/api/v1/login")
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
        "400": {Message: "Login Failed", Unmarshaller: errors.NewResponseLoginFailure},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    var result models.ResponseLoginSuccess
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseLoginSuccess](decoder)
    return models.NewApiResponse(result, resp), err
}

// TwoFactor takes context, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Send 2FA Code
func (a *AdminsLogin) TwoFactor(
    ctx context.Context,
    body *models.TwoFactorString) (
    *http.Response,
    error) {
    req := a.prepareRequest(ctx, "POST", "/api/v1/login/two_factor")
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
        "401": {Message: "two_factor code is incorrect or the user hasn’t login yet"},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "the user doesn’t have 2FA enabled"},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
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
