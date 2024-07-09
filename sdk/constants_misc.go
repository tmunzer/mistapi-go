package mistapigo

import (
	"context"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/tmunzer/mistapigo/sdk/errors"
	"github.com/tmunzer/mistapigo/sdk/models"
)

// ConstantsMisc represents a controller struct.
type ConstantsMisc struct {
	baseController
}

// NewConstantsMisc creates a new instance of ConstantsMisc.
// It takes a baseController as a parameter and returns a pointer to the ConstantsMisc.
func NewConstantsMisc(baseController baseController) *ConstantsMisc {
	constantsMisc := ConstantsMisc{baseController: baseController}
	return &constantsMisc
}

// ListApChannels takes context, countryCode as parameters and
// returns an models.ApiResponse with models.ConstApChannel data and
// an error if there was an issue with the request or response.
// Get List of List of Available channels per country code
func (c *ConstantsMisc) ListApChannels(
	ctx context.Context,
	countryCode *string) (
	models.ApiResponse[models.ConstApChannel],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/ap_channels")
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
	if countryCode != nil {
		req.QueryParam("country_code", *countryCode)
	}
	var result models.ConstApChannel
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ConstApChannel](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListApplications takes context as parameters and
// returns an models.ApiResponse with []models.ConstApplicationDefinition data and
// an error if there was an issue with the request or response.
// Get List of a list of applications that Juniper-Mist APs recognize
func (c *ConstantsMisc) ListApplications(ctx context.Context) (
	models.ApiResponse[[]models.ConstApplicationDefinition],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/applications")
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
	var result []models.ConstApplicationDefinition
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstApplicationDefinition](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListCountryCodes takes context as parameters and
// returns an models.ApiResponse with []models.ConstCountry data and
// an error if there was an issue with the request or response.
// Get List of List of available Country Codes
func (c *ConstantsMisc) ListCountryCodes(ctx context.Context) (
	models.ApiResponse[[]models.ConstCountry],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/countries")
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
	var result []models.ConstCountry
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstCountry](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetGatewayDefaultConfig takes context, model, ha as parameters and
// returns an models.ApiResponse with interface{} data and
// an error if there was an issue with the request or response.
// Generate Default Gateway Config
func (c *ConstantsMisc) GetGatewayDefaultConfig(
	ctx context.Context,
	model string,
	ha *string) (
	models.ApiResponse[interface{}],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/default_gateway_config")
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
	req.QueryParam("model", model)
	if ha != nil {
		req.QueryParam("ha", *ha)
	}
	var result interface{}
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[interface{}](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListGatewayApplications takes context as parameters and
// returns an models.ApiResponse with []models.ConstGatewayApplicationsDefinition data and
// an error if there was an issue with the request or response.
// Get the full list of applications that we recognize
func (c *ConstantsMisc) ListGatewayApplications(ctx context.Context) (
	models.ApiResponse[[]models.ConstGatewayApplicationsDefinition],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/gateway_applications")
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
	var result []models.ConstGatewayApplicationsDefinition
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstGatewayApplicationsDefinition](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListInsightMetrics takes context as parameters and
// returns an models.ApiResponse with map[string]models.ConstInsightMetricsProperty data and
// an error if there was an issue with the request or response.
// List Insight Metrics
func (c *ConstantsMisc) ListInsightMetrics(ctx context.Context) (
	models.ApiResponse[map[string]models.ConstInsightMetricsProperty],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/insight_metrics")
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
	var result map[string]models.ConstInsightMetricsProperty
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[map[string]models.ConstInsightMetricsProperty](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSiteLanguages takes context as parameters and
// returns an models.ApiResponse with []models.ConstLanguage data and
// an error if there was an issue with the request or response.
// Get List of Languages
func (c *ConstantsMisc) ListSiteLanguages(ctx context.Context) (
	models.ApiResponse[[]models.ConstLanguage],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/languages")
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
	var result []models.ConstLanguage
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstLanguage](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetLicenseTypes takes context as parameters and
// returns an models.ApiResponse with []models.ConstLicenseType data and
// an error if there was an issue with the request or response.
// Get License Types
func (c *ConstantsMisc) GetLicenseTypes(ctx context.Context) (
	models.ApiResponse[[]models.ConstLicenseType],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/license_types")
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
	var result []models.ConstLicenseType
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstLicenseType](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListTrafficTypes takes context as parameters and
// returns an models.ApiResponse with []models.ConstTrafficType data and
// an error if there was an issue with the request or response.
// Get List of identified traffic
func (c *ConstantsMisc) ListTrafficTypes(ctx context.Context) (
	models.ApiResponse[[]models.ConstTrafficType],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/traffic_types")
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
	var result []models.ConstTrafficType
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstTrafficType](decoder)
	return models.NewApiResponse(result, resp), err
}
