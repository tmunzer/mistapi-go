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
// Start password recovery for an administrator account. This public endpoint does not require authentication. When the request is accepted, Mist sends an email containing a recovery link such as `https://manage.mist.com/verify/recover?token=:token`. If CAPTCHA is required, include the CAPTCHA response token and provider flavor in the request body.
func (a *AdminsRecoverPassword) RecoverPassword(
	ctx context.Context,
	body *models.Recover) (
	*http.Response,
	error) {
	req := a.prepareRequest(ctx, "POST", "/api/v1/recover")

	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
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

// VerifyRecoverPassword takes context, token as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Verify a password recovery token from the recovery email. This public endpoint does not require authentication. When the token is valid, the user is authenticated for the recovery flow so the client can prompt for and submit a new password.
func (a *AdminsRecoverPassword) VerifyRecoverPassword(
	ctx context.Context,
	token string) (
	*http.Response,
	error) {
	req := a.prepareRequest(ctx, "POST", "/api/v1/recover/verify/%v")
	req.AppendTemplateParams(token)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}
