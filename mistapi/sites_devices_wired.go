package mistapi

import (
    "context"
    "fmt"
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
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Sometimes HelpDesk Admin needs to change port configs
func (s *SitesDevicesWired) DeleteSiteLocalSwitchPortConfig(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/local_port_config", siteId, deviceId),
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// UpdateSiteLocalSwitchPortConfig takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Sometimes HelpDesk Admin needs to change port configs
func (s *SitesDevicesWired) UpdateSiteLocalSwitchPortConfig(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.JunosPortConfig) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/local_port_config", siteId, deviceId),
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}
