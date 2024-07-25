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

// SitesStatsBeacons represents a controller struct.
type SitesStatsBeacons struct {
    baseController
}

// NewSitesStatsBeacons creates a new instance of SitesStatsBeacons.
// It takes a baseController as a parameter and returns a pointer to the SitesStatsBeacons.
func NewSitesStatsBeacons(baseController baseController) *SitesStatsBeacons {
    sitesStatsBeacons := SitesStatsBeacons{baseController: baseController}
    return &sitesStatsBeacons
}

// ListSiteBeaconsStats takes context, siteId, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with []models.BeaconStatsItems data and
// an error if there was an issue with the request or response.
// Get List of Site Beacons Stats
func (s *SitesStatsBeacons) ListSiteBeaconsStats(
    ctx context.Context,
    siteId uuid.UUID,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[[]models.BeaconStatsItems],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/beacons", siteId),
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
    if page != nil {
        req.QueryParam("page", *page)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    
    var result []models.BeaconStatsItems
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.BeaconStatsItems](decoder)
    return models.NewApiResponse(result, resp), err
}
