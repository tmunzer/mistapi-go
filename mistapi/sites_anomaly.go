// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
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
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/anomaly/client/%v/%v")
    req.AppendTemplateParams(siteId, clientMac, metric)
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

// GetSiteAnomalyEventsForDevice takes context, siteId, metric, deviceMac as parameters and
// returns an models.ApiResponse with models.ResponseAnomalySearch data and
// an error if there was an issue with the request or response.
// Get Device Anomaly Events
func (s *SitesAnomaly) GetSiteAnomalyEventsForDevice(
    ctx context.Context,
    siteId uuid.UUID,
    metric string,
    deviceMac string) (
    models.ApiResponse[models.ResponseAnomalySearch],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/anomaly/device/%v/%v")
    req.AppendTemplateParams(siteId, deviceMac, metric)
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

// ListSiteAnomalyEvents takes context, siteId, metric as parameters and
// returns an models.ApiResponse with models.ResponseAnomalySearch data and
// an error if there was an issue with the request or response.
// List Site Anomaly Events
func (s *SitesAnomaly) ListSiteAnomalyEvents(
    ctx context.Context,
    siteId uuid.UUID,
    metric string) (
    models.ApiResponse[models.ResponseAnomalySearch],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/anomaly/%v")
    req.AppendTemplateParams(siteId, metric)
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
