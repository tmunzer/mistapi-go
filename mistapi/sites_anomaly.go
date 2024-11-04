package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// SitesAnomaly represents a controller struct.
type SitesAnomaly struct {
    baseController
}

// NewSitesAnomaly creates a new instance of SitesAnomaly.
// It takes a baseController as a parameter and returns a pointer to the SitesAnomaly.
func NewSitesAnomaly(baseController baseController) *SitesAnomaly {
    sitesAnomaly := SitesAnomaly{baseController: baseController}
    return &sitesAnomaly
}

// GetSiteAnomalyEventsForClient takes context, siteId, clientMac, metric as parameters and
// returns an models.ApiResponse with models.ResponseAnomalySearch data and
// an error if there was an issue with the request or response.
// Get Client Anomaly Events
func (s *SitesAnomaly) GetSiteAnomalyEventsForClient(
    ctx context.Context,
    siteId uuid.UUID,
    clientMac string,
    metric string) (
    models.ApiResponse[models.ResponseAnomalySearch],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/anomaly/client/%v/%v", siteId, clientMac, metric),
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
    
    var result models.ResponseAnomalySearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseAnomalySearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteAnomalyEventsforDevice takes context, siteId, metric, deviceMac as parameters and
// returns an models.ApiResponse with models.ResponseAnomalySearch data and
// an error if there was an issue with the request or response.
// Get Device Anomaly Events
func (s *SitesAnomaly) GetSiteAnomalyEventsforDevice(
    ctx context.Context,
    siteId uuid.UUID,
    metric string,
    deviceMac string) (
    models.ApiResponse[models.ResponseAnomalySearch],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/anomaly/device/%v/%v", siteId, deviceMac, metric),
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
    
    var result models.ResponseAnomalySearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseAnomalySearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteAnomalyEvents takes context, siteId, metric as parameters and
// returns an models.ApiResponse with models.ResponseAnomalySearch data and
// an error if there was an issue with the request or response.
// Get Site Anomaly Events
func (s *SitesAnomaly) GetSiteAnomalyEvents(
    ctx context.Context,
    siteId uuid.UUID,
    metric string) (
    models.ApiResponse[models.ResponseAnomalySearch],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/anomaly/%v", siteId, metric),
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
    
    var result models.ResponseAnomalySearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseAnomalySearch](decoder)
    return models.NewApiResponse(result, resp), err
}
