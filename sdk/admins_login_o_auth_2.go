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

// AdminsLoginOAuth2 represents a controller struct.
type AdminsLoginOAuth2 struct {
	baseController
}

// NewAdminsLoginOAuth2 creates a new instance of AdminsLoginOAuth2.
// It takes a baseController as a parameter and returns a pointer to the AdminsLoginOAuth2.
func NewAdminsLoginOAuth2(baseController baseController) *AdminsLoginOAuth2 {
	adminsLoginOAuth2 := AdminsLoginOAuth2{baseController: baseController}
	return &adminsLoginOAuth2
}

// UnlinkOauth2Provider takes context, provider as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Unlink OAuth2 Provider
func (a *AdminsLoginOAuth2) UnlinkOauth2Provider(
	ctx context.Context,
	provider string) (
	*http.Response,
	error) {
	req := a.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/login/oauth/%v", provider),
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

// GetOauth2AuthorizationUrlForLogin takes context, provider, forward as parameters and
// returns an models.ApiResponse with models.ResponseLoginOauthUrl data and
// an error if there was an issue with the request or response.
// Obtain Authorization URL for Login
func (a *AdminsLoginOAuth2) GetOauth2AuthorizationUrlForLogin(
	ctx context.Context,
	provider string,
	forward *string) (
	models.ApiResponse[models.ResponseLoginOauthUrl],
	error) {
	req := a.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/login/oauth/%v", provider),
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
	if forward != nil {
		req.QueryParam("forward", *forward)
	}

	var result models.ResponseLoginOauthUrl
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseLoginOauthUrl](decoder)
	return models.NewApiResponse(result, resp), err
}

// LoginOauth2 takes context, provider, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Login via OAuth2
func (a *AdminsLoginOAuth2) LoginOauth2(
	ctx context.Context,
	provider string,
	body *models.CodeString) (
	*http.Response,
	error) {
	req := a.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/login/oauth/%v", provider),
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
