package mistapigo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/tmunzer/mistapigo/sdk/errors"
	"github.com/tmunzer/mistapigo/sdk/models"
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
// returns an models.ApiResponse with  data and
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

// VerifyRecoverPasssword takes context, token as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Verify Recover Password
// With correct verification, the user will be authenticated. UI can then prompt for new password
func (a *AdminsRecoverPassword) VerifyRecoverPasssword(
	ctx context.Context,
	token string) (
	*http.Response,
	error) {
	req := a.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/recover/verify/%v", token),
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
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}
