package mistapi

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"net/http"
)

// OrgsSDKTemplates represents a controller struct.
type OrgsSDKTemplates struct {
	baseController
}

// NewOrgsSDKTemplates creates a new instance of OrgsSDKTemplates.
// It takes a baseController as a parameter and returns a pointer to the OrgsSDKTemplates.
func NewOrgsSDKTemplates(baseController baseController) *OrgsSDKTemplates {
	orgsSDKTemplates := OrgsSDKTemplates{baseController: baseController}
	return &orgsSDKTemplates
}

// ListSdkTemplates takes context, orgId as parameters and
// returns an models.ApiResponse with []models.Sdktemplate data and
// an error if there was an issue with the request or response.
// Get List of Org SDK Templates
func (o *OrgsSDKTemplates) ListSdkTemplates(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[[]models.Sdktemplate],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/sdktemplates", orgId),
	)
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

	var result []models.Sdktemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Sdktemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateSdkTemplate takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Sdktemplate data and
// an error if there was an issue with the request or response.
// Create SDK Template
func (o *OrgsSDKTemplates) CreateSdkTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Sdktemplate) (
	models.ApiResponse[models.Sdktemplate],
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/orgs/%v/sdktemplates", orgId),
	)
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
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.Sdktemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sdktemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteSdkTemplate takes context, orgId, sdktemplateId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete SDK Template
func (o *OrgsSDKTemplates) DeleteSdkTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	sdktemplateId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/orgs/%v/sdktemplates/%v", orgId, sdktemplateId),
	)
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// GetSdkTemplate takes context, orgId, sdktemplateId as parameters and
// returns an models.ApiResponse with models.Sdktemplate data and
// an error if there was an issue with the request or response.
// Get SDK Template Details
func (o *OrgsSDKTemplates) GetSdkTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	sdktemplateId uuid.UUID) (
	models.ApiResponse[models.Sdktemplate],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/sdktemplates/%v", orgId, sdktemplateId),
	)
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

	var result models.Sdktemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sdktemplate](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateSdkTemplate takes context, orgId, sdktemplateId, body as parameters and
// returns an models.ApiResponse with models.Sdktemplate data and
// an error if there was an issue with the request or response.
// Update SDK Template
func (o *OrgsSDKTemplates) UpdateSdkTemplate(
	ctx context.Context,
	orgId uuid.UUID,
	sdktemplateId uuid.UUID,
	body *models.Sdktemplate) (
	models.ApiResponse[models.Sdktemplate],
	error) {
	req := o.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/orgs/%v/sdktemplates/%v", orgId, sdktemplateId),
	)
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
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.Sdktemplate
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sdktemplate](decoder)
	return models.NewApiResponse(result, resp), err
}
