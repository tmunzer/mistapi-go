package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// OrgsStatsTunnels represents a controller struct.
type OrgsStatsTunnels struct {
    baseController
}

// NewOrgsStatsTunnels creates a new instance of OrgsStatsTunnels.
// It takes a baseController as a parameter and returns a pointer to the OrgsStatsTunnels.
func NewOrgsStatsTunnels(baseController baseController) *OrgsStatsTunnels {
    orgsStatsTunnels := OrgsStatsTunnels{baseController: baseController}
    return &orgsStatsTunnels
}

// CountOrgTunnelsStats takes context, orgId, distinct, mType as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count Mist Tunnels Stats
func (o *OrgsStatsTunnels) CountOrgTunnelsStats(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgTunnelCountDistinctEnum,
    mType *models.OrgTunnelTypeCountEnum) (
    models.ApiResponse[models.ResponseCount],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/tunnels/count")
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
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    
    var result models.ResponseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchOrgTunnelsStats takes context, orgId, mxclusterId, siteId, wxtunnelId, ap, mac, node, peerIp, peerHost, ip, tunnelName, protocol, authAlgo, encryptAlgo, ikeVersion, up, mType, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseTunnelSearch data and
// an error if there was an issue with the request or response.
// Search Org Tunnels Stats
func (o *OrgsStatsTunnels) SearchOrgTunnelsStats(
    ctx context.Context,
    orgId uuid.UUID,
    mxclusterId *string,
    siteId *string,
    wxtunnelId *string,
    ap *string,
    mac *string,
    node *string,
    peerIp *string,
    peerHost *string,
    ip *string,
    tunnelName *string,
    protocol *string,
    authAlgo *string,
    encryptAlgo *string,
    ikeVersion *string,
    up *string,
    mType *models.TunnelTypeEnum,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseTunnelSearch],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/tunnels/search")
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
    if mxclusterId != nil {
        req.QueryParam("mxcluster_id", *mxclusterId)
    }
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if wxtunnelId != nil {
        req.QueryParam("wxtunnel_id", *wxtunnelId)
    }
    if ap != nil {
        req.QueryParam("ap", *ap)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if node != nil {
        req.QueryParam("node", *node)
    }
    if peerIp != nil {
        req.QueryParam("peer_ip", *peerIp)
    }
    if peerHost != nil {
        req.QueryParam("peer_host", *peerHost)
    }
    if ip != nil {
        req.QueryParam("ip", *ip)
    }
    if tunnelName != nil {
        req.QueryParam("tunnel_name", *tunnelName)
    }
    if protocol != nil {
        req.QueryParam("protocol", *protocol)
    }
    if authAlgo != nil {
        req.QueryParam("auth_algo", *authAlgo)
    }
    if encryptAlgo != nil {
        req.QueryParam("encrypt_algo", *encryptAlgo)
    }
    if ikeVersion != nil {
        req.QueryParam("ike_version", *ikeVersion)
    }
    if up != nil {
        req.QueryParam("up", *up)
    }
    if mType != nil {
        req.QueryParam("type", *mType)
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
    
    var result models.ResponseTunnelSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseTunnelSearch](decoder)
    return models.NewApiResponse(result, resp), err
}
