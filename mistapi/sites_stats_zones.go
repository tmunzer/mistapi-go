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

// SitesStatsZones represents a controller struct.
type SitesStatsZones struct {
    baseController
}

// NewSitesStatsZones creates a new instance of SitesStatsZones.
// It takes a baseController as a parameter and returns a pointer to the SitesStatsZones.
func NewSitesStatsZones(baseController baseController) *SitesStatsZones {
    sitesStatsZones := SitesStatsZones{baseController: baseController}
    return &sitesStatsZones
}

// ListSiteZonesStats takes context, siteId, mapId as parameters and
// returns an models.ApiResponse with []models.ZoneStats data and
// an error if there was an issue with the request or response.
// Get List of Site Zones Stats
func (s *SitesStatsZones) ListSiteZonesStats(
    ctx context.Context,
    siteId uuid.UUID,
    mapId *string) (
    models.ApiResponse[[]models.ZoneStats],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/zones", siteId),
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
    if mapId != nil {
        req.QueryParam("map_id", *mapId)
    }
    
    var result []models.ZoneStats
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ZoneStats](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteZoneStats takes context, siteId, zoneType, zoneId as parameters and
// returns an models.ApiResponse with models.ZoneStatsDetails data and
// an error if there was an issue with the request or response.
// Get Detail Zone Stats
func (s *SitesStatsZones) GetSiteZoneStats(
    ctx context.Context,
    siteId uuid.UUID,
    zoneType models.ZoneTypeEnum,
    zoneId uuid.UUID) (
    models.ApiResponse[models.ZoneStatsDetails],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/%v/%v", siteId, zoneType, zoneId),
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
    
    var result models.ZoneStatsDetails
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ZoneStatsDetails](decoder)
    return models.NewApiResponse(result, resp), err
}
