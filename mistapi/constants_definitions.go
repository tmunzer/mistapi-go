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

// ListApChannels takes context, countryCode as parameters and
// returns an models.ApiResponse with models.ConstApChannel data and
// an error if there was an issue with the request or response.
// Return supported AP radio channels for the requested country code, based on the regulatory domain used by Mist AP configuration.
func (c *ConstantsDefinitions) ListApChannels(
	ctx context.Context,
	countryCode *string) (
	models.ApiResponse[models.ConstApChannel],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/ap_channels")

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

// ListApLEslVersions takes context as parameters and
// returns an models.ApiResponse with []models.ConstApEslVersion data and
// an error if there was an issue with the request or response.
// Return Electronic Shelf Label (ESL) firmware versions available per AP model.
func (c *ConstantsDefinitions) ListApLEslVersions(ctx context.Context) (
	models.ApiResponse[[]models.ConstApEslVersion],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/ap_esl_versions")

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
	var result []models.ConstApEslVersion
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstApEslVersion](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListApLedDefinition takes context as parameters and
// returns an models.ApiResponse with []models.ConstApLed data and
// an error if there was an issue with the request or response.
// Return AP LED status definitions, including the values used to describe LED behavior and the corresponding `error_code` value.
func (c *ConstantsDefinitions) ListApLedDefinition(ctx context.Context) (
	models.ApiResponse[[]models.ConstApLed],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/ap_led_status")

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
// Return supported application categories used for application identification, traffic classification, and policy matching.
func (c *ConstantsDefinitions) ListAppCategoryDefinitions(ctx context.Context) (
	models.ApiResponse[[]models.ConstAppCategoryDefinition],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/app_categories")

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
// Return supported application subcategories used for application identification, traffic classification, and policy matching.
func (c *ConstantsDefinitions) ListAppSubCategoryDefinitions(ctx context.Context) (
	models.ApiResponse[[]models.ConstAppSubcategoryDefinition],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/app_subcategories")

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
	var result []models.ConstAppSubcategoryDefinition
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstAppSubcategoryDefinition](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListApplications takes context as parameters and
// returns an models.ApiResponse with []models.ConstApplicationDefinition data and
// an error if there was an issue with the request or response.
// Return applications recognized by Juniper Mist devices for traffic classification and application analytics.
func (c *ConstantsDefinitions) ListApplications(ctx context.Context) (
	models.ApiResponse[[]models.ConstApplicationDefinition],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/applications")

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
	var result []models.ConstApplicationDefinition
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstApplicationDefinition](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListCountryCodes takes context, extend as parameters and
// returns an models.ApiResponse with []models.ConstCountry data and
// an error if there was an issue with the request or response.
// Return supported country codes for Mist configuration. Set `extend=true` to include additional country codes when available.
func (c *ConstantsDefinitions) ListCountryCodes(
	ctx context.Context,
	extend *bool) (
	models.ApiResponse[[]models.ConstCountry],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/countries")

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
	if extend != nil {
		req.QueryParam("extend", *extend)
	}
	var result []models.ConstCountry
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstCountry](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListFingerprintTypes takes context as parameters and
// returns an models.ApiResponse with models.ConstFingerprintTypes data and
// an error if there was an issue with the request or response.
// Return supported client fingerprint attribute values (`family`, `model`, `mfg`, and `os_type`) that can be used in [Mist NAC Rules]($h/Orgs%20NAC%20Rules/_overview) `matching` conditions.
func (c *ConstantsDefinitions) ListFingerprintTypes(ctx context.Context) (
	models.ApiResponse[models.ConstFingerprintTypes],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/fingerprint_types")

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
	var result models.ConstFingerprintTypes
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ConstFingerprintTypes](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListGatewayApplications takes context as parameters and
// returns an models.ApiResponse with []models.ConstGatewayApplicationsDefinition data and
// an error if there was an issue with the request or response.
// Return applications recognized by Mist gateways for traffic classification.
func (c *ConstantsDefinitions) ListGatewayApplications(ctx context.Context) (
	models.ApiResponse[[]models.ConstGatewayApplicationsDefinition],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/gateway_applications")

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
// Return supported insight metric names and metadata used by insight and SLE APIs.
func (c *ConstantsDefinitions) ListInsightMetrics(ctx context.Context) (
	models.ApiResponse[map[string]models.ConstInsightMetricsProperty],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/insight_metrics")

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
// Return supported language codes for localized Mist configuration and user-facing portals.
func (c *ConstantsDefinitions) ListSiteLanguages(ctx context.Context) (
	models.ApiResponse[[]models.ConstLanguage],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/languages")

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
	var result []models.ConstLanguage
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstLanguage](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListLicenseTypes takes context as parameters and
// returns an models.ApiResponse with []models.ConstLicenseType data and
// an error if there was an issue with the request or response.
// Return Mist license type definitions used by inventory and subscription APIs.
func (c *ConstantsDefinitions) ListLicenseTypes(ctx context.Context) (
	models.ApiResponse[[]models.ConstLicenseType],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/license_types")

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
	var result []models.ConstLicenseType
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstLicenseType](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListMarvisClientVersions takes context as parameters and
// returns an models.ApiResponse with []models.ConstMarvisClientVersion data and
// an error if there was an issue with the request or response.
// Return available Marvis Client application versions for tracking or managing Marvis Client deployments.
func (c *ConstantsDefinitions) ListMarvisClientVersions(ctx context.Context) (
	models.ApiResponse[[]models.ConstMarvisClientVersion],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/marvisclient_versions")

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
	var result []models.ConstMarvisClientVersion
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstMarvisClientVersion](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListStates takes context, countryCode as parameters and
// returns an models.ApiResponse with []models.ConstState data and
// an error if there was an issue with the request or response.
// Return ISO state or province codes for the supplied country code.
func (c *ConstantsDefinitions) ListStates(
	ctx context.Context,
	countryCode string) (
	models.ApiResponse[[]models.ConstState],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/states")

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
	req.QueryParam("country_code", countryCode)
	var result []models.ConstState
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstState](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListTrafficTypes takes context as parameters and
// returns an models.ApiResponse with []models.ConstTrafficType data and
// an error if there was an issue with the request or response.
// Return traffic type definitions used to classify recognized network traffic.
func (c *ConstantsDefinitions) ListTrafficTypes(ctx context.Context) (
	models.ApiResponse[[]models.ConstTrafficType],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/traffic_types")

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
	var result []models.ConstTrafficType
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstTrafficType](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListWebhookTopics takes context as parameters and
// returns an models.ApiResponse with []models.ConstWebhookTopic data and
// an error if there was an issue with the request or response.
// Return webhook topic definitions that can be subscribed to in webhook configuration.
func (c *ConstantsDefinitions) ListWebhookTopics(ctx context.Context) (
	models.ApiResponse[[]models.ConstWebhookTopic],
	error) {
	req := c.prepareRequest(ctx, "GET", "/api/v1/const/webhook_topics")

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
	var result []models.ConstWebhookTopic
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConstWebhookTopic](decoder)
	return models.NewApiResponse(result, resp), err
}
