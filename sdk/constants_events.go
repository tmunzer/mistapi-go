package mistapi

import (
	"context"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/tmunzer/mistapi/sdk/errors"
	"github.com/tmunzer/mistapi/sdk/models"
)

// ConstantsEvents represents a controller struct.
type ConstantsEvents struct {
	baseController
}

// NewConstantsEvents creates a new instance of ConstantsEvents.
// It takes a baseController as a parameter and returns a pointer to the ConstantsEvents.
func NewConstantsEvents(baseController baseController) *ConstantsEvents {
	constantsEvents := ConstantsEvents{baseController: baseController}
	return &constantsEvents
}

// ListClientEventsDefinitions takes context as parameters and
// returns an models.ApiResponse with []models.ConstEvent data and
// an error if there was an issue with the request or response.
// Get List of List of available Client Events
func (c *ConstantsEvents) ListClientEventsDefinitions(ctx context.Context) (
	models.ApiResponse[[]models.ConstEvent],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/client_events")
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
	var result []models.ConstEvent
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstEvent](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListDeviceEventsDefinitions takes context as parameters and
// returns an models.ApiResponse with []models.ConstEvent data and
// an error if there was an issue with the request or response.
// Get list of available Device Events
func (c *ConstantsEvents) ListDeviceEventsDefinitions(ctx context.Context) (
	models.ApiResponse[[]models.ConstEvent],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/device_events")
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
	var result []models.ConstEvent
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstEvent](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListMxEdgeEventsDefinitions takes context as parameters and
// returns an models.ApiResponse with []models.ConstEvent data and
// an error if there was an issue with the request or response.
// Get List of available MX Edge Events
func (c *ConstantsEvents) ListMxEdgeEventsDefinitions(ctx context.Context) (
	models.ApiResponse[[]models.ConstEvent],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/mxedge_events")
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
	var result []models.ConstEvent
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstEvent](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListNacEventsDefinitions takes context as parameters and
// returns an models.ApiResponse with []models.ConstNacEvent data and
// an error if there was an issue with the request or response.
// Get List of List of available NAC Client Events
func (c *ConstantsEvents) ListNacEventsDefinitions(ctx context.Context) (
	models.ApiResponse[[]models.ConstNacEvent],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/nac_events")
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
	var result []models.ConstNacEvent
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstNacEvent](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListOtherDeviceEventsDefinitions takes context as parameters and
// returns an models.ApiResponse with []models.ConstEvent data and
// an error if there was an issue with the request or response.
// Supported Events Type
func (c *ConstantsEvents) ListOtherDeviceEventsDefinitions(ctx context.Context) (
	models.ApiResponse[[]models.ConstEvent],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/otherdevice_events")
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
	var result []models.ConstEvent
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstEvent](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSystemEventsDefinitions takes context as parameters and
// returns an models.ApiResponse with []models.ConstEvent data and
// an error if there was an issue with the request or response.
// Get List of List of available System Events
func (c *ConstantsEvents) ListSystemEventsDefinitions(ctx context.Context) (
	models.ApiResponse[[]models.ConstEvent],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/system_events")
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
	var result []models.ConstEvent
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstEvent](decoder)
	return models.NewApiResponse(result, resp), err
}
