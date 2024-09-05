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

// SitesWANUsages represents a controller struct.
type SitesWANUsages struct {
    baseController
}

// NewSitesWANUsages creates a new instance of SitesWANUsages.
// It takes a baseController as a parameter and returns a pointer to the SitesWANUsages.
func NewSitesWANUsages(baseController baseController) *SitesWANUsages {
    sitesWANUsages := SitesWANUsages{baseController: baseController}
    return &sitesWANUsages
}

// CountSiteWanUsage takes context, siteId, mac, peerMac, portId, peerPortId, policy, tenant, pathType, distinct, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count Site WAN Uages
func (s *SitesWANUsages) CountSiteWanUsage(
    ctx context.Context,
    siteId uuid.UUID,
    mac *string,
    peerMac *string,
    portId *string,
    peerPortId *string,
    policy *string,
    tenant *string,
    pathType *string,
    distinct *models.WanUsagesCountDisctinctEnum,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/wan_usages/count", siteId),
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
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if peerMac != nil {
        req.QueryParam("peer_mac", *peerMac)
    }
    if portId != nil {
        req.QueryParam("port_id", *portId)
    }
    if peerPortId != nil {
        req.QueryParam("peer_port_id", *peerPortId)
    }
    if policy != nil {
        req.QueryParam("policy", *policy)
    }
    if tenant != nil {
        req.QueryParam("tenant", *tenant)
    }
    if pathType != nil {
        req.QueryParam("path_type", *pathType)
    }
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
    
    var result models.RepsonseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchSiteWanUsage takes context, siteId, mac, peerMac, portId, peerPortId, policy, tenant, pathType, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.SearchWanUsage data and
// an error if there was an issue with the request or response.
// Search Site WAN Uages
func (s *SitesWANUsages) SearchSiteWanUsage(
    ctx context.Context,
    siteId uuid.UUID,
    mac *string,
    peerMac *string,
    portId *string,
    peerPortId *string,
    policy *string,
    tenant *string,
    pathType *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.SearchWanUsage],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/wan_usages/search", siteId),
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
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if peerMac != nil {
        req.QueryParam("peer_mac", *peerMac)
    }
    if portId != nil {
        req.QueryParam("port_id", *portId)
    }
    if peerPortId != nil {
        req.QueryParam("peer_port_id", *peerPortId)
    }
    if policy != nil {
        req.QueryParam("policy", *policy)
    }
    if tenant != nil {
        req.QueryParam("tenant", *tenant)
    }
    if pathType != nil {
        req.QueryParam("path_type", *pathType)
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
    
    var result models.SearchWanUsage
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SearchWanUsage](decoder)
    return models.NewApiResponse(result, resp), err
}
