package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// SitesStatsBGPPeers represents a controller struct.
type SitesStatsBGPPeers struct {
    baseController
}

// NewSitesStatsBGPPeers creates a new instance of SitesStatsBGPPeers.
// It takes a baseController as a parameter and returns a pointer to the SitesStatsBGPPeers.
func NewSitesStatsBGPPeers(baseController baseController) *SitesStatsBGPPeers {
    sitesStatsBGPPeers := SitesStatsBGPPeers{baseController: baseController}
    return &sitesStatsBGPPeers
}

// CountSiteBgpStats takes context, siteId, state, distinct as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count BGP Stats
func (s *SitesStatsBGPPeers) CountSiteBgpStats(
    ctx context.Context,
    siteId uuid.UUID,
    state *string,
    distinct *string) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/bgp_peers/count")
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
    if state != nil {
        req.QueryParam("state", *state)
    }
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
    }
    
    var result models.RepsonseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchSiteBgpStats takes context, siteId as parameters and
// returns an models.ApiResponse with models.ResponseSearchBgps data and
// an error if there was an issue with the request or response.
// Search BGP Stats
func (s *SitesStatsBGPPeers) SearchSiteBgpStats(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.ResponseSearchBgps],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/bgp_peers/search")
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
    
    var result models.ResponseSearchBgps
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseSearchBgps](decoder)
    return models.NewApiResponse(result, resp), err
}
