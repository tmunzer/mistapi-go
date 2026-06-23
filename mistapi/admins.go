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

// Admins represents a controller struct.
type Admins struct {
	baseController
}

// NewAdmins creates a new instance of Admins.
// It takes a baseController as a parameter and returns a pointer to the Admins.
func NewAdmins(baseController baseController) *Admins {
	admins := Admins{baseController: baseController}
	return &admins
}

// VerifyAdminInvite takes context, token as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Accept an administrator invite using the invite verification token from the invite email. This public endpoint does not require authentication. After a successful verification, call [Get Self]($e/Self%20Account/getSelf) to refresh the authenticated admin profile and retrieve the newly granted privileges.
func (a *Admins) VerifyAdminInvite(
	ctx context.Context,
	token string) (
	*http.Response,
	error) {
	req := a.prepareRequest(ctx, "POST", "/api/v1/invite/verify/%v")
	req.AppendTemplateParams(token)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not Found", Unmarshaller: errors.NewResponseDetailString},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// RegisterNewAdmin takes context, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Register a new administrator account and initial organization. This public endpoint does not require authentication. Mist sends a verification email containing a link such as `/verify/register?token={token}`; use [Verify Registration]($e/Admins/verifyRegistration) to complete registration with that token.
// Use [Get Registration Information]($e/Admins/getAdminRegistrationInfo) before submitting this request to determine whether CAPTCHA is required, which CAPTCHA provider to render, and which public site key to use. If CAPTCHA is required, include the provider response token in `recaptcha` and the provider name in `recaptcha_flavor`.
func (a *Admins) RegisterNewAdmin(
	ctx context.Context,
	body *models.AdminInvite) (
	*http.Response,
	error) {
	req := a.prepareRequest(ctx, "POST", "/api/v1/register")

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

// GetAdminRegistrationInfo takes context, recaptchaFlavor as parameters and
// returns an models.ApiResponse with models.Recaptcha data and
// an error if there was an issue with the request or response.
// Return the public CAPTCHA settings required for administrator registration. This public endpoint does not require authentication. Use the returned `flavor`, `required`, and `sitekey` values to render the correct CAPTCHA challenge before calling [Register New Admin]($e/Admins/registerNewAdmin).
func (a *Admins) GetAdminRegistrationInfo(
	ctx context.Context,
	recaptchaFlavor *models.RecaptchaFlavorEnum) (
	models.ApiResponse[models.Recaptcha],
	error) {
	req := a.prepareRequest(ctx, "GET", "/api/v1/register/recaptcha")

	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})
	if recaptchaFlavor != nil {
		req.QueryParam("recaptcha_flavor", *recaptchaFlavor)
	}
	var result models.Recaptcha
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Recaptcha](decoder)
	return models.NewApiResponse(result, resp), err
}

// VerifyRegistration takes context, token as parameters and
// returns an models.ApiResponse with models.ResponseVerifyTokenSuccess data and
// an error if there was an issue with the request or response.
// Verify a new administrator registration using the token from the registration email. This public endpoint does not require authentication. A successful verification creates a login session and may also apply a pending invitation; the response indicates whether an invitation could not be applied automatically.
func (a *Admins) VerifyRegistration(
	ctx context.Context,
	token string) (
	models.ApiResponse[models.ResponseVerifyTokenSuccess],
	error) {
	req := a.prepareRequest(ctx, "POST", "/api/v1/register/verify/%v")
	req.AppendTemplateParams(token)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Response if verification expired or already registered", Unmarshaller: errors.NewResponseDetailString},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Response if secret is invalid", Unmarshaller: errors.NewResponseDetailString},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})

	var result models.ResponseVerifyTokenSuccess
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseVerifyTokenSuccess](decoder)
	return models.NewApiResponse(result, resp), err
}
