// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
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

// CountOrgBgpStats takes context, orgId, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Org BGP Stats
func (o *OrgsStatsBGPPeers) CountOrgBgpStats(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
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
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    
    var result models.ResponseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchOrgBgpStats takes context, orgId, mac, neighborMac, siteId, vrfName, start, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseSearchBgps data and
// an error if there was an issue with the request or response.
// Search Org BGP Stats
func (o *OrgsStatsBGPPeers) SearchOrgBgpStats(
    ctx context.Context,
    orgId uuid.UUID,
    mac *string,
    neighborMac *string,
    siteId *string,
    vrfName *string,
    start *int,
    duration *string,
    limit *int) (
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
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if neighborMac != nil {
        req.QueryParam("neighbor_mac", *neighborMac)
    }
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if vrfName != nil {
        req.QueryParam("vrf_name", *vrfName)
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
    
    var result models.ResponseSearchBgps
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseSearchBgps](decoder)
    return models.NewApiResponse(result, resp), err
}
