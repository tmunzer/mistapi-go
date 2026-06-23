// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
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
// Authenticate an administrator with email and password. A successful login creates the browser session cookies, including the `csrftoken` value used with the `X-CSRFToken` header on later API requests.
// When 2FA is enabled, either include the `two_factor` code in this request or submit the first factor here and complete the login with [Two Factor]($e/Admins%20Login/twoFactor).
func (a *AdminsLogin) Login(
	ctx context.Context,
	body *models.Login) (
	models.ApiResponse[models.ResponseLoginSuccess],
	error) {
	req := a.prepareRequest(ctx, "POST", "/api/v1/login")

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
// returns an *Response and
// an error if there was an issue with the request or response.
// Complete a two-factor login by submitting the 2FA code after the initial email/password step has created a pending login session.
func (a *AdminsLogin) TwoFactor(
	ctx context.Context,
	body *models.TwoFactorString) (
	*http.Response,
	error) {
	req := a.prepareRequest(ctx, "POST", "/api/v1/login/two_factor")

	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "two_factor code is incorrect or the user hasn't login yet"},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "The user doesn't have 2FA enabled"},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
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
