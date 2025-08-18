// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

// SelfOAuth2 represents a controller struct.
type SelfOAuth2 struct {
	baseController
}

// NewSelfOAuth2 creates a new instance of SelfOAuth2.
// It takes a baseController as a parameter and returns a pointer to the SelfOAuth2.
func NewSelfOAuth2(baseController baseController) *SelfOAuth2 {
	selfOAuth2 := SelfOAuth2{baseController: baseController}
	return &selfOAuth2
}

// GetOauth2UrlForLinking takes context, provider, forward as parameters and
// returns an models.ApiResponse with models.ResponseSelfOauthUrl data and
// an error if there was an issue with the request or response.
// Obtain Authorization URL for Linking
func (s *SelfOAuth2) GetOauth2UrlForLinking(
	ctx context.Context,
	provider string,
	forward *string) (
	models.ApiResponse[models.ResponseSelfOauthUrl],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/self/oauth/%v")
	req.AppendTemplateParams(provider)
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
	if forward != nil {
		req.QueryParam("forward", *forward)
	}

	var result models.ResponseSelfOauthUrl
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSelfOauthUrl](decoder)
	return models.NewApiResponse(result, resp), err
}

// LinkOauth2MistAccount takes context, provider, body as parameters and
// returns an models.ApiResponse with models.ResponseSelfOauthLinkSuccess data and
// an error if there was an issue with the request or response.
// Link Mist account with an OAuth2 Provider
func (s *SelfOAuth2) LinkOauth2MistAccount(
	ctx context.Context,
	provider string,
	body *models.CodeString) (
	models.ApiResponse[models.ResponseSelfOauthLinkSuccess],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/self/oauth/%v")
	req.AppendTemplateParams(provider)
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
		"400": {Message: "Authorization Error", Unmarshaller: errors.NewResponseSelfOauthLinkFailure},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.ResponseSelfOauthLinkSuccess
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSelfOauthLinkSuccess](decoder)
	return models.NewApiResponse(result, resp), err
}
