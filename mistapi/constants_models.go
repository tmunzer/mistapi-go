package mistapi

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
)

// ConstantsModels represents a controller struct.
type ConstantsModels struct {
	baseController
}

// NewConstantsModels creates a new instance of ConstantsModels.
// It takes a baseController as a parameter and returns a pointer to the ConstantsModels.
func NewConstantsModels(baseController baseController) *ConstantsModels {
	constantsModels := ConstantsModels{baseController: baseController}
	return &constantsModels
}

// ListDeviceModels takes context as parameters and
// returns an models.ApiResponse with []models.ListDeviceModelsResponse data and
// an error if there was an issue with the request or response.
// Get list of AP device models for the Mist Site
func (c *ConstantsModels) ListDeviceModels(ctx context.Context) (
	models.ApiResponse[[]models.ListDeviceModelsResponse],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/device_models")
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	var result []models.ListDeviceModelsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ListDeviceModelsResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListMxEdgeModels takes context as parameters and
// returns an models.ApiResponse with []models.ConstMxedgeModel data and
// an error if there was an issue with the request or response.
// Get List of available Mx Edge models
func (c *ConstantsModels) ListMxEdgeModels(ctx context.Context) (
	models.ApiResponse[[]models.ConstMxedgeModel],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/mxedge_models")
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	var result []models.ConstMxedgeModel
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstMxedgeModel](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSupportedOtherDeviceModels takes context as parameters and
// returns an models.ApiResponse with []models.ConstOtherDeviceModel data and
// an error if there was an issue with the request or response.
// Supported OtherDevice Models
func (c *ConstantsModels) ListSupportedOtherDeviceModels(ctx context.Context) (
	models.ApiResponse[[]models.ConstOtherDeviceModel],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/otherdevice_models")
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
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	var result []models.ConstOtherDeviceModel
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstOtherDeviceModel](decoder)
	return models.NewApiResponse(result, resp), err
}
