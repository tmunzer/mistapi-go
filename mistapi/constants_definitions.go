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

// ListAlarmDefinitions takes context as parameters and
// returns an models.ApiResponse with []models.ConstAlarmDefinition data and
// an error if there was an issue with the request or response.
// Get List of brief definitions of all the supported alarm types.
// The example field contains an example payload as you would recieve in the alarm webhook output.
// HA cluster node names will be specified in the `node` field, if applicable.'
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    var result []models.ConstAlarmDefinition
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ConstAlarmDefinition](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListApChannels takes context, countryCode as parameters and
// returns an models.ApiResponse with models.ConstApChannel data and
// an error if there was an issue with the request or response.
// Get List of List of Available channels per country code
func (c *ConstantsDefinitions) ListApChannels(
    ctx context.Context,
    countryCode *string) (
    models.ApiResponse[models.ConstApChannel],
    error) {
    req := c.prepareRequest(ctx, "GET", "/api/v1/const/ap_channels")
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
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
// Get List of a list of applications that Juniper-Mist APs recognize
func (c *ConstantsDefinitions) ListApplications(ctx context.Context) (
    models.ApiResponse[[]models.ConstApplicationDefinition],
    error) {
    req := c.prepareRequest(ctx, "GET", "/api/v1/const/applications")
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
// Get List of available Country Codes
func (c *ConstantsDefinitions) ListCountryCodes(
    ctx context.Context,
    extend *bool) (
    models.ApiResponse[[]models.ConstCountry],
    error) {
    req := c.prepareRequest(ctx, "GET", "/api/v1/const/countries")
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

// ListGatewayApplications takes context as parameters and
// returns an models.ApiResponse with []models.ConstGatewayApplicationsDefinition data and
// an error if there was an issue with the request or response.
// Get the full list of applications that we recognize
func (c *ConstantsDefinitions) ListGatewayApplications(ctx context.Context) (
    models.ApiResponse[[]models.ConstGatewayApplicationsDefinition],
    error) {
    req := c.prepareRequest(ctx, "GET", "/api/v1/const/gateway_applications")
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
func (c *ConstantsDefinitions) ListInsightMetrics(ctx context.Context) (
    models.ApiResponse[map[string]models.ConstInsightMetricsProperty],
    error) {
    req := c.prepareRequest(ctx, "GET", "/api/v1/const/insight_metrics")
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
func (c *ConstantsDefinitions) ListSiteLanguages(ctx context.Context) (
    models.ApiResponse[[]models.ConstLanguage],
    error) {
    req := c.prepareRequest(ctx, "GET", "/api/v1/const/languages")
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
func (c *ConstantsDefinitions) GetLicenseTypes(ctx context.Context) (
    models.ApiResponse[[]models.ConstLicenseType],
    error) {
    req := c.prepareRequest(ctx, "GET", "/api/v1/const/license_types")
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
    var result []models.ConstLicenseType
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ConstLicenseType](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListStates takes context, countryCode as parameters and
// returns an models.ApiResponse with []models.ConstState data and
// an error if there was an issue with the request or response.
// Get List of ISO States based on country code
func (c *ConstantsDefinitions) ListStates(
    ctx context.Context,
    countryCode string) (
    models.ApiResponse[[]models.ConstState],
    error) {
    req := c.prepareRequest(ctx, "GET", "/api/v1/const/states")
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
// Get List of identified traffic
func (c *ConstantsDefinitions) ListTrafficTypes(ctx context.Context) (
    models.ApiResponse[[]models.ConstTrafficType],
    error) {
    req := c.prepareRequest(ctx, "GET", "/api/v1/const/traffic_types")
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
    var result []models.ConstTrafficType
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ConstTrafficType](decoder)
    return models.NewApiResponse(result, resp), err
}
