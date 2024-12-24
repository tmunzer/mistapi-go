package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// SitesDevicesWired represents a controller struct.
type SitesDevicesWired struct {
    baseController
}

// NewSitesDevicesWired creates a new instance of SitesDevicesWired.
// It takes a baseController as a parameter and returns a pointer to the SitesDevicesWired.
func NewSitesDevicesWired(baseController baseController) *SitesDevicesWired {
    sitesDevicesWired := SitesDevicesWired{baseController: baseController}
    return &sitesDevicesWired
}

// DeleteSiteLocalSwitchPortConfig takes context, siteId, deviceId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// API Calls delete all the existing port config local overrides, and reapply the configured planed at the device level 
// (with site / template heritance).
func (s *SitesDevicesWired) DeleteSiteLocalSwitchPortConfig(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      "/api/v1/sites/%v/devices/%v/local_port_config",
    )
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// UpdateSiteLocalSwitchPortConfig takes context, siteId, deviceId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// API Calls to add port config local overrides. This can be used by Switch Port Operators or Helpdesk administrators
// to change a Switch Port configuration without having to change the switch configuration.
// The local overrides configured for the switchports with `no_local_overwrite`==`true` won't be applied to the switch configuration. 
// > NOTE:
// >
// > When using the API Call, it is required to put send all overrides in the PUT request Payload, even the existing once. 
// >
// > The current overrides can be retrieved with the API Call [Get Site Device]($e/Sites%20Devices/getSiteDevice). The local overrides will show up separately from the `port_config` in the `local_port_config` so it can be easily identified (and cleared)
func (s *SitesDevicesWired) UpdateSiteLocalSwitchPortConfig(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body map[string]models.JunosLocalPortConfig) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "PUT",
      "/api/v1/sites/%v/devices/%v/local_port_config",
    )
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}
