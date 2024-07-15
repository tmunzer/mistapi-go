package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// SelfMFA represents a controller struct.
type SelfMFA struct {
    baseController
}

// NewSelfMFA creates a new instance of SelfMFA.
// It takes a baseController as a parameter and returns a pointer to the SelfMFA.
func NewSelfMFA(baseController baseController) *SelfMFA {
    selfMFA := SelfMFA{baseController: baseController}
    return &selfMFA
}

// GenerateSecretFor2faVerification takes context, by as parameters and
// returns an models.ApiResponse with models.ResponseTwoFactorJson data and
// an error if there was an issue with the request or response.
// Generate Secret Key for 2FA verification
func (s *SelfMFA) GenerateSecretFor2faVerification(
    ctx context.Context,
    by *models.MfaSecretTypeEnum) (
    models.ApiResponse[models.ResponseTwoFactorJson],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/self/two_factor/token")
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    if by != nil {
        req.QueryParam("by", *by)
    }
    var result models.ResponseTwoFactorJson
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseTwoFactorJson](decoder)
    return models.NewApiResponse(result, resp), err
}

// VerifyTwoFactor takes context, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Verify Two-factor (OTP)
// To verify two-factor authentication by using a code generated by app (e.g. Google Authenticator, Authy). Upon successful verification, the `two_factor_passed` will be set to true if it hasn’t already been.
func (s *SelfMFA) VerifyTwoFactor(
    ctx context.Context,
    body *models.TwoFactorCode) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/self/two_factor/verify")
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
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
