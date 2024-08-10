package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
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
func (c *ConstantsMisc) ListSiteLanguages(ctx context.Context) (
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
func (c *ConstantsMisc) GetLicenseTypes(ctx context.Context) (
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
