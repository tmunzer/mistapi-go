package mistapigo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/tmunzer/mistapigo/sdk/errors"
	"github.com/tmunzer/mistapigo/sdk/models"
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
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// **Note**: another call to ```GET /api/v1/self``` is required to see the new set of privileges
func (a *Admins) VerifyAdminInvite(
	ctx context.Context,
	token string) (
	*http.Response,
	error) {
	req := a.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/invite/verify/%v", token),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not Found", Unmarshaller: errors.NewResponseDetailString},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// RegisterNewAdmin takes context, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Register a new admin and his/her org
// An email will also be sent to the user with a link to `/verify/register?token={token}`
// ### reCAPTCHA
// Google reCAPTCHA is the choice to prevent bot registration
// It needs this
// &lt;script src='https://www.google.com/recaptcha/api.js' &gt;&lt;/script&gt;
// and this &lt;div&gt; in the desired place
// ```html
// <div class="g-recaptcha" data_sitekey="6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd"></div>
// ```
// Use GET /api/v1/register/recaptcha to read the current setting.
// Response example:
// ```json
// {
// "flavor": "google",
// "required": true,
// "sitekey": "6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd"
// }
// ```
// ### hCaptcha
// Alternative to reCAPTCHA is hCaptcha to prevent bot registration
// It needs this script
// &lt;script src='https://js.hcaptcha.com/1/api.js' async defer &gt;&lt;/script&gt;
// and this &lt;div&gt; in the desired place
// ```html
// <div class="h-recaptcha" data_sitekey="6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd"></div>
// ```
// Use GET /api/v1/register/recaptcha?recaptcha_flavor=hcaptcha to read the current setting for hcaptcha with reply.
// Response example:
// ```json
// {
// "flavor": "hcaptcha",
// "required": true,
// "sitekey": "6LdAewsTAAAAAE25XKQhPEQ2FiMTft-WrZXQ5NUd"
// }"
// ```
func (a *Admins) RegisterNewAdmin(
	ctx context.Context,
	body *models.AdminInvite) (
	*http.Response,
	error) {
	req := a.prepareRequest(ctx, "POST", "/api/v1/register")
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
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

// GetAdminRegistrationInfo takes context, recaptchaFlavor as parameters and
// returns an models.ApiResponse with models.Recaptcha data and
// an error if there was an issue with the request or response.
// Get Registration Information
func (a *Admins) GetAdminRegistrationInfo(
	ctx context.Context,
	recaptchaFlavor *models.RecaptchaFlavorEnum) (
	models.ApiResponse[models.Recaptcha],
	error) {
	req := a.prepareRequest(ctx, "GET", "/api/v1/register/recaptcha")
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
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
// Verify registration
func (a *Admins) VerifyRegistration(
	ctx context.Context,
	token string) (
	models.ApiResponse[models.ResponseVerifyTokenSuccess],
	error) {
	req := a.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/register/verify/%v", token),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Response if verification expired or already registered", Unmarshaller: errors.NewResponseDetailString},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Response if secret is invalid", Unmarshaller: errors.NewResponseDetailString},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	var result models.ResponseVerifyTokenSuccess
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseVerifyTokenSuccess](decoder)
	return models.NewApiResponse(result, resp), err
}
