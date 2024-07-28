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

// SitesMapsAutoPlacement represents a controller struct.
type SitesMapsAutoPlacement struct {
    baseController
}

// NewSitesMapsAutoPlacement creates a new instance of SitesMapsAutoPlacement.
// It takes a baseController as a parameter and returns a pointer to the SitesMapsAutoPlacement.
func NewSitesMapsAutoPlacement(baseController baseController) *SitesMapsAutoPlacement {
    sitesMapsAutoPlacement := SitesMapsAutoPlacement{baseController: baseController}
    return &sitesMapsAutoPlacement
}

// DeleteSiteApAutoOrientation takes context, mapId, siteId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This API is called to force stop auto placement for a given map
func (s *SitesMapsAutoPlacement) DeleteSiteApAutoOrientation(
    ctx context.Context,
    mapId uuid.UUID,
    siteId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/sites/%v/maps/%v/auto_orient", siteId, mapId),
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
        "400": {Message: "Autoplacement was not triggered"},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// StartSiteApAutoOrientation takes context, mapId, siteId, body as parameters and
// returns an models.ApiResponse with models.ResponseAutoOrientation data and
// an error if there was an issue with the request or response.
// This API is called to trigger a map for auto orientation. For auto orient feature to work, BLE data needs to be collected from the APs on the map. This precess is not disruptive unlike FTM collection. Repeated POST to this endpoint while a map is still running will be rejected.
// List of devices to provide suggestions for is an optional parameter that can be given to this API. This will provide auto orient suggestions only for the devices specified. If no list of devices is provided, all APs asociated with that map are considered by default
func (s *SitesMapsAutoPlacement) StartSiteApAutoOrientation(
    ctx context.Context,
    mapId uuid.UUID,
    siteId uuid.UUID,
    body *models.AutoOrient) (
    models.ApiResponse[models.ResponseAutoOrientation],
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/maps/%v/auto_orient", siteId, mapId),
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
        "400": {Message: "Bad Request", Unmarshaller: errors.NewResponseDetailString},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ResponseAutoOrientation
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseAutoOrientation](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSiteApAutoplacement takes context, siteId, mapId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This API is called to force stop auto placement for a given map
func (s *SitesMapsAutoPlacement) DeleteSiteApAutoplacement(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/sites/%v/maps/%v/auto_placement", siteId, mapId),
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
        "400": {Message: "Autoplacement was not triggered"},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// GetSiteApAutoPlacement takes context, siteId, mapId as parameters and
// returns an models.ApiResponse with models.ResponseAutoPlacementInfo data and
// an error if there was an issue with the request or response.
// This API is called to view the current status of auto placement for a given map.
// #### Status Descriptions
// | Status | Description |
// | --- | --- |
// | `pending` | Autoplacement has not been requested for this map |
// | `inprogress` | Autoplacement is currently processing |
// | `done` | The autoplacement process has completed |
// | `data_needed` | Additional position data is required for autoplacement. Users should verify the requested anchor APs have a position on the map |
// | `invalid_model` | Autoplacement is not supported on the model of the APs on the map |
// | `invalid_version` | Autoplacement is not supported with the APs current firmware version |
// | `error` | There was an error in the autoplacement process |
func (s *SitesMapsAutoPlacement) GetSiteApAutoPlacement(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID) (
    models.ApiResponse[models.ResponseAutoPlacementInfo],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/maps/%v/auto_placement", siteId, mapId),
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    
    var result models.ResponseAutoPlacementInfo
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseAutoPlacementInfo](decoder)
    return models.NewApiResponse(result, resp), err
}

// RunSiteApAutoplacement takes context, siteId, mapId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This API is called to trigger a map for auto placement. For auto placement feature to work, RTT-FTM data need to be collected from the APs on the map. This scan is disruptive and therefore the user must be notified of service disrution during the functioning of auto placement Repeated POST to this endpoint while a map is still running will be rejected.
// List of devices to provide suggestions for is an optional parameter that can be given to this API. This will provide autoplacement suggestions only for the devices specified. If no list of devices is provided, all APs asociated with that map are considered by default
func (s *SitesMapsAutoPlacement) RunSiteApAutoplacement(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID,
    body *models.AutoPlacement) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/maps/%v/auto_placement", siteId, mapId),
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
        "400": {Message: "* Map has less than 3 APs associated with it to perform auto placement \n* Auto AP Placement is already in progress for this Map\n* Autoplacement data does not exist or has gone stale"},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
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

// ClearSiteApAutoOrient takes context, siteId, mapId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This API is used to destroy the autoorientations of a map or subset of APs on a map.
func (s *SitesMapsAutoPlacement) ClearSiteApAutoOrient(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID,
    body *models.MacAddresses) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/maps/%v/clear_auto_orient", siteId, mapId),
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
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

// ClearSiteApAutoplacement takes context, siteId, mapId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This API is used to destroy the cached autoplacement locations of a map or subset of APs on a map.
func (s *SitesMapsAutoPlacement) ClearSiteApAutoplacement(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID,
    body *models.MacAddresses) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/maps/%v/clear_autoplacement", siteId, mapId),
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
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

// ConfirmSiteApLocalizationData takes context, siteId, mapId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// This API is used to accept or reject the cached autoplacement and auto orientation values of a map or subset of APs on a map. A rejected AP will retain its current X,Y and orientation until accpeted.
func (s *SitesMapsAutoPlacement) ConfirmSiteApLocalizationData(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID,
    body *models.UseAutoApValues) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/maps/%v/use_auto_ap_values", siteId, mapId),
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
        "400": {Message: "Invalid localization service expected: placement or orientation"},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
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
