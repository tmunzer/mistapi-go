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

// SitesStatsWxRules represents a controller struct.
type SitesStatsWxRules struct {
    baseController
}

// NewSitesStatsWxRules creates a new instance of SitesStatsWxRules.
// It takes a baseController as a parameter and returns a pointer to the SitesStatsWxRules.
func NewSitesStatsWxRules(baseController baseController) *SitesStatsWxRules {
    sitesStatsWxRules := SitesStatsWxRules{baseController: baseController}
    return &sitesStatsWxRules
}

// GetSiteWxRulesUsage takes context, siteId as parameters and
// returns an models.ApiResponse with []models.WxruleStat data and
// an error if there was an issue with the request or response.
// Get Wxlan Rule usage
func (s *SitesStatsWxRules) GetSiteWxRulesUsage(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[[]models.WxruleStat],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/stats/wxrules", siteId),
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
    
    var result []models.WxruleStat
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.WxruleStat](decoder)
    return models.NewApiResponse(result, resp), err
}
