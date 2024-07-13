package mistapi

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
)

// ConstantsDefinitions represents a controller struct.
type ConstantsDefinitions struct {
	baseController
}

// NewConstantsDefinitions creates a new instance of ConstantsDefinitions.
// It takes a baseController as a parameter and returns a pointer to the ConstantsDefinitions.
func NewConstantsDefinitions(baseController baseController) *ConstantsDefinitions {
	constantsDefinitions := ConstantsDefinitions{baseController: baseController}
	return &constantsDefinitions
}

// ListAlarmDefinitions takes context as parameters and
// returns an models.ApiResponse with []models.ConstAlarmDefinition data and
// an error if there was an issue with the request or response.
// Get List of brief definitions of all the supported alarm types. The
// example field contains an example payload as you would recieve in the
// alarm webhook output.
// HA cluster node names will be specified in the `node` field, if applicable.
func (c *ConstantsDefinitions) ListAlarmDefinitions(ctx context.Context) (
	models.ApiResponse[[]models.ConstAlarmDefinition],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/alarm_defs")
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
	var result []models.ConstAlarmDefinition
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstAlarmDefinition](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListApLedDefinition takes context as parameters and
// returns an models.ApiResponse with []models.ConstApLed data and
// an error if there was an issue with the request or response.
// Get List of AP LED definition
func (c *ConstantsDefinitions) ListApLedDefinition(ctx context.Context) (
	models.ApiResponse[[]models.ConstApLed],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/ap_led_status")
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
	var result []models.ConstApLed
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstApLed](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListAppCategoryDefinitions takes context as parameters and
// returns an models.ApiResponse with []models.ConstAppCategoryDefinition data and
// an error if there was an issue with the request or response.
// Get List of definitions of all the supported Application Categories. The example field contains an example payload as you would recieve in the alarm webhook output.
func (c *ConstantsDefinitions) ListAppCategoryDefinitions(ctx context.Context) (
	models.ApiResponse[[]models.ConstAppCategoryDefinition],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/app_categories")
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
	var result []models.ConstAppCategoryDefinition
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstAppCategoryDefinition](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListAppSubCategoryDefinitions takes context as parameters and
// returns an models.ApiResponse with []models.ConstAppSubcategoryDefinition data and
// an error if there was an issue with the request or response.
// Get List of definitions of all the supported Application sub-categories. The example field contains an example payload as you would recieve in the alarm webhook output.
func (c *ConstantsDefinitions) ListAppSubCategoryDefinitions(ctx context.Context) (
	models.ApiResponse[[]models.ConstAppSubcategoryDefinition],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/app_subcategories")
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
	var result []models.ConstAppSubcategoryDefinition
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstAppSubcategoryDefinition](decoder)
	return models.NewApiResponse(result, resp), err
}
