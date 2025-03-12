package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// OrgsStatsVPNPeers represents a controller struct.
type OrgsStatsVPNPeers struct {
    baseController
}

// NewOrgsStatsVPNPeers creates a new instance of OrgsStatsVPNPeers.
// It takes a baseController as a parameter and returns a pointer to the OrgsStatsVPNPeers.
func NewOrgsStatsVPNPeers(baseController baseController) *OrgsStatsVPNPeers {
    orgsStatsVPNPeers := OrgsStatsVPNPeers{baseController: baseController}
    return &orgsStatsVPNPeers
}

// CountOrgPeerPathStats takes context, orgId, distinct, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count Org Peer Path Stats
func (o *OrgsStatsVPNPeers) CountOrgPeerPathStats(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseCount],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/vpn_peers/count")
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
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
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
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result models.ResponseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchOrgPeerPathStats takes context, orgId, mac, siteId, mType, start, duration, limit as parameters and
// returns an models.ApiResponse with models.VpnPeerStatSearch data and
// an error if there was an issue with the request or response.
// Search Org Peer Path Stats
func (o *OrgsStatsVPNPeers) SearchOrgPeerPathStats(
    ctx context.Context,
    orgId uuid.UUID,
    mac *string,
    siteId *string,
    mType *models.VpnTypeEnum,
    start *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.VpnPeerStatSearch],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/vpn_peers/search")
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
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if start != nil {
        req.QueryParam("start", *start)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    
    var result models.VpnPeerStatSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.VpnPeerStatSearch](decoder)
    return models.NewApiResponse(result, resp), err
}
