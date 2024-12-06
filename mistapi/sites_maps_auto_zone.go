package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// SitesMapsAutoZone represents a controller struct.
type SitesMapsAutoZone struct {
    baseController
}

// NewSitesMapsAutoZone creates a new instance of SitesMapsAutoZone.
// It takes a baseController as a parameter and returns a pointer to the SitesMapsAutoZone.
func NewSitesMapsAutoZone(baseController baseController) *SitesMapsAutoZone {
    sitesMapsAutoZone := SitesMapsAutoZone{baseController: baseController}
    return &sitesMapsAutoZone
}

// DeleteSiteMapAutoZone takes context, mapId, siteId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This API starts the auto zones service for a specified map. This map must have an image to parse for the auto zones service. Repeated POST requests to this endpoint while the auto zones service is proccessing the map or awaiting review will be rejected.
func (s *SitesMapsAutoZone) DeleteSiteMapAutoZone(
    ctx context.Context,
    mapId uuid.UUID,
    siteId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/sites/%v/maps/%v/auto_zones", siteId, mapId),
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

// GetSiteMapAutoZoneStatus takes context, mapId, siteId as parameters and
// returns an models.ApiResponse with models.ResponseAutoZone data and
// an error if there was an issue with the request or response.
// This API provides the current status of the auto zones service for a given map
func (s *SitesMapsAutoZone) GetSiteMapAutoZoneStatus(
    ctx context.Context,
    mapId uuid.UUID,
    siteId uuid.UUID) (
    models.ApiResponse[models.ResponseAutoZone],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/maps/%v/auto_zones", siteId, mapId),
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    
    var result models.ResponseAutoZone
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseAutoZone](decoder)
    return models.NewApiResponse(result, resp), err
}

// StartSiteMapAutoZone takes context, mapId, siteId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// This API starts the auto zones service for a specified map. This map must have an image to parse for the auto zones service. Reppeated POST requests to this endpoint while the auto zones service is proccessing the map will be rejected.
func (s *SitesMapsAutoZone) StartSiteMapAutoZone(
    ctx context.Context,
    mapId uuid.UUID,
    siteId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/maps/%v/auto_zones", siteId, mapId),
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
