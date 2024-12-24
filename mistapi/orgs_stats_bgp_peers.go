package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// OrgsStatsBGPPeers represents a controller struct.
type OrgsStatsBGPPeers struct {
    baseController
}

// NewOrgsStatsBGPPeers creates a new instance of OrgsStatsBGPPeers.
// It takes a baseController as a parameter and returns a pointer to the OrgsStatsBGPPeers.
func NewOrgsStatsBGPPeers(baseController baseController) *OrgsStatsBGPPeers {
    orgsStatsBGPPeers := OrgsStatsBGPPeers{baseController: baseController}
    return &orgsStatsBGPPeers
}

// CountOrgBgpStats takes context, orgId as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count Org BGP Stats
func (o *OrgsStatsBGPPeers) CountOrgBgpStats(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/bgp_peers/count")
    req.AppendTemplateParams(orgId)
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
    
    var result models.RepsonseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchOrgBgpStats takes context, orgId as parameters and
// returns an models.ApiResponse with models.ResponseSearchBgps data and
// an error if there was an issue with the request or response.
// Search Org BGP Stats
func (o *OrgsStatsBGPPeers) SearchOrgBgpStats(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponseSearchBgps],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/bgp_peers/search")
    req.AppendTemplateParams(orgId)
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
