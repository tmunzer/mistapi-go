package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// SitesDevicesWireless represents a controller struct.
type SitesDevicesWireless struct {
    baseController
}

// NewSitesDevicesWireless creates a new instance of SitesDevicesWireless.
// It takes a baseController as a parameter and returns a pointer to the SitesDevicesWireless.
func NewSitesDevicesWireless(baseController baseController) *SitesDevicesWireless {
    sitesDevicesWireless := SitesDevicesWireless{baseController: baseController}
    return &sitesDevicesWireless
}

// ListSiteDeviceRadioChannels takes context, siteId, countryCode as parameters and
// returns an models.ApiResponse with models.ResponseDeviceRadioChannels data and
// an error if there was an issue with the request or response.
// Get a list of allowed channels (per channel width)
func (s *SitesDevicesWireless) ListSiteDeviceRadioChannels(
    ctx context.Context,
    siteId uuid.UUID,
    countryCode *string) (
    models.ApiResponse[models.ResponseDeviceRadioChannels],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/devices/ap_channels")
    req.AppendTemplateParams(siteId)
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
    
    var result models.ResponseDeviceRadioChannels
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseDeviceRadioChannels](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteDeviceIotPort takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with map[string]int data and
// an error if there was an issue with the request or response.
// Returns the current state of each enabled IoT pin configured as an output.
func (s *SitesDevicesWireless) GetSiteDeviceIotPort(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[map[string]int],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/devices/%v/iot")
    req.AppendTemplateParams(siteId, deviceId)
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
    
    var result map[string]int
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[map[string]int](decoder)
    return models.NewApiResponse(result, resp), err
}

// SetSiteDeviceIotPort takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with map[string]int data and
// an error if there was an issue with the request or response.
// **Note**: For each IoT pin referenced:
// * The pin must be enabled using the Device `iot_config` API
// * The pin must support the output direction
func (s *SitesDevicesWireless) SetSiteDeviceIotPort(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body map[string]int) (
    models.ApiResponse[map[string]int],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/devices/%v/iot")
    req.AppendTemplateParams(siteId, deviceId)
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
    
    var result map[string]int
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[map[string]int](decoder)
    return models.NewApiResponse(result, resp), err
}
