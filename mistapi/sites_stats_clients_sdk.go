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

// SitesStatsClientsSDK represents a controller struct.
type SitesStatsClientsSDK struct {
    baseController
}

// NewSitesStatsClientsSDK creates a new instance of SitesStatsClientsSDK.
// It takes a baseController as a parameter and returns a pointer to the SitesStatsClientsSDK.
func NewSitesStatsClientsSDK(baseController baseController) *SitesStatsClientsSDK {
    sitesStatsClientsSDK := SitesStatsClientsSDK{baseController: baseController}
    return &sitesStatsClientsSDK
}

// GetSiteSdkStatsByMap takes context, siteId, mapId as parameters and
// returns an models.ApiResponse with []models.StatsSdkclient data and
// an error if there was an issue with the request or response.
// Get SdkClient Stats By Map
func (s *SitesStatsClientsSDK) GetSiteSdkStatsByMap(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID) (
    models.ApiResponse[[]models.StatsSdkclient],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/maps/%v/sdkclients", siteId, mapId),
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
    
    var result []models.StatsSdkclient
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.StatsSdkclient](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSdkStats takes context, siteId, sdkclientId as parameters and
// returns an models.ApiResponse with models.SdkstatsWirelessClient data and
// an error if there was an issue with the request or response.
// Get Detail Stats of a SdkClient
func (s *SitesStatsClientsSDK) GetSiteSdkStats(
    ctx context.Context,
    siteId uuid.UUID,
    sdkclientId uuid.UUID) (
    models.ApiResponse[models.SdkstatsWirelessClient],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/sdkclients/%v", siteId, sdkclientId),
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
    
    var result models.SdkstatsWirelessClient
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SdkstatsWirelessClient](decoder)
    return models.NewApiResponse(result, resp), err
}
