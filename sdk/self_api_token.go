package mistapigo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapigo/sdk/errors"
	"github.com/tmunzer/mistapigo/sdk/models"
)

// SelfAPIToken represents a controller struct.
type SelfAPIToken struct {
	baseController
}

// NewSelfAPIToken creates a new instance of SelfAPIToken.
// It takes a baseController as a parameter and returns a pointer to the SelfAPIToken.
func NewSelfAPIToken(baseController baseController) *SelfAPIToken {
	selfAPIToken := SelfAPIToken{baseController: baseController}
	return &selfAPIToken
}

// ListApiTokens takes context as parameters and
// returns an models.ApiResponse with []models.UserApitoken data and
// an error if there was an issue with the request or response.
// Get List of Current User API Tokens
func (s *SelfAPIToken) ListApiTokens(ctx context.Context) (
	models.ApiResponse[[]models.UserApitoken],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/self/apitokens")
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
	var result []models.UserApitoken
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.UserApitoken](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateApiToken takes context, body as parameters and
// returns an models.ApiResponse with []models.UserApitoken data and
// an error if there was an issue with the request or response.
// Create API Token
// Note that the key is only available during creation time.
func (s *SelfAPIToken) CreateApiToken(
	ctx context.Context,
	body *models.UserApitoken) (
	models.ApiResponse[[]models.UserApitoken],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/self/apitokens")
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
	var result []models.UserApitoken
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.UserApitoken](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteApiToken takes context, apitokenId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete an API Token
func (s *SelfAPIToken) DeleteApiToken(
	ctx context.Context,
	apitokenId uuid.UUID) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/self/apitokens/%v", apitokenId),
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

// GetApiToken takes context, apitokenId as parameters and
// returns an models.ApiResponse with models.UserApitoken data and
// an error if there was an issue with the request or response.
// Get User API Token
func (s *SelfAPIToken) GetApiToken(
	ctx context.Context,
	apitokenId uuid.UUID) (
	models.ApiResponse[models.UserApitoken],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/self/apitokens/%v", apitokenId),
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

	var result models.UserApitoken
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.UserApitoken](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateApiToken takes context, apitokenId, body as parameters and
// returns an models.ApiResponse with models.UserApitoken data and
// an error if there was an issue with the request or response.
// Update User API Token
func (s *SelfAPIToken) UpdateApiToken(
	ctx context.Context,
	apitokenId uuid.UUID,
	body *models.UserApitoken) (
	models.ApiResponse[models.UserApitoken],
	error) {
	req := s.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/self/apitokens/%v", apitokenId),
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

	var result models.UserApitoken
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.UserApitoken](decoder)
	return models.NewApiResponse(result, resp), err
}
