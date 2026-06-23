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
// returns an *Response and
// an error if there was an issue with the request or response.
// Unlink the specified OAuth2 provider from the authenticated administrator account so it can no longer be used for that account's OAuth login.
func (a *AdminsLoginOAuth2) UnlinkOauth2Provider(
	ctx context.Context,
	provider string) (
	*http.Response,
	error) {
	req := a.prepareRequest(ctx, "DELETE", "/api/v1/login/oauth/%v")
	req.AppendTemplateParams(provider)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
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

// GetOauth2AuthorizationUrlForLogin takes context, provider, forward as parameters and
// returns an models.ApiResponse with models.ResponseLoginOauthUrl data and
// an error if there was an issue with the request or response.
// Return the provider authorization URL used to start an OAuth2 login or account-linking flow. When `forward` is provided, the provider redirects back to that callback URL after authorization.
func (a *AdminsLoginOAuth2) GetOauth2AuthorizationUrlForLogin(
	ctx context.Context,
	provider string,
	forward *string) (
	models.ApiResponse[models.ResponseLoginOauthUrl],
	error) {
	req := a.prepareRequest(ctx, "GET", "/api/v1/login/oauth/%v")
	req.AppendTemplateParams(provider)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
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
// returns an *Response and
// an error if there was an issue with the request or response.
// Complete an OAuth2 login by exchanging the provider authorization code for a Mist administrator session.
func (a *AdminsLoginOAuth2) LoginOauth2(
	ctx context.Context,
	provider string,
	body *models.CodeString) (
	*http.Response,
	error) {
	req := a.prepareRequest(ctx, "POST", "/api/v1/login/oauth/%v")
	req.AppendTemplateParams(provider)
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
