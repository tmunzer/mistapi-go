package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// SitesStatsApps represents a controller struct.
type SitesStatsApps struct {
    baseController
}

// NewSitesStatsApps creates a new instance of SitesStatsApps.
// It takes a baseController as a parameter and returns a pointer to the SitesStatsApps.
func NewSitesStatsApps(baseController baseController) *SitesStatsApps {
    sitesStatsApps := SitesStatsApps{baseController: baseController}
    return &sitesStatsApps
}

// CountSiteApps takes context, siteId, distinct, deviceMac, app, wired as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Applications
func (s *SitesStatsApps) CountSiteApps(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteAppsCountDistinctEnum,
    deviceMac *string,
    app *string,
    wired *string) (
    models.ApiResponse[models.ResponseCount],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/apps/count")
    req.AppendTemplateParams(siteId)
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
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
    }
    if deviceMac != nil {
        req.QueryParam("device_mac", *deviceMac)
    }
    if app != nil {
        req.QueryParam("app", *app)
    }
    if wired != nil {
        req.QueryParam("wired", *wired)
    }
    
    var result models.ResponseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCount](decoder)
    return models.NewApiResponse(result, resp), err
}
