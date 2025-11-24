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

// CountOrgTunnelsStats takes context, orgId, distinct, mType, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Mist Tunnels Stats
func (o *OrgsStatsTunnels) CountOrgTunnelsStats(
	ctx context.Context,
	orgId uuid.UUID,
	distinct *models.OrgTunnelCountDistinctEnum,
	mType *models.OrgTunnelTypeCountEnum,
	limit *int) (
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

// SearchOrgTunnelsStats takes context, orgId, mxclusterId, siteId, wxtunnelId, ap, mac, node, peerIp, peerHost, ip, tunnelName, protocol, authAlgo, encryptAlgo, ikeVersion, up, mType, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseTunnelSearch data and
// an error if there was an issue with the request or response.
// By default the endpoint returns only `wxtunnel` type stats, to get `wan` type stats
// you need to specify `type=wan` in the query parameters.
// Tunnel types:
// - `wxtunnel` (default) - A WxLan Tunnel (WxTunnel) are used to create a secure connection between Juniper Mist Access Points and third-party VPN concentrators using protocols such as L2TPv3 or dmvpn.
// - `wan` - A WAN Tunnel is a secure connection between two Gateways, typically used for site-to-site or mesh connectivity. It can be configured with various protocols and encryption methods.
// If `type` is not specified or `type`==`wxtunnel`, the following parameters are supported:
// - `mxcluster_id` - the MX cluster ID
// - `site_id` - the site ID
// - `wxtunnel_id` - the WX tunnel ID
// - `ap` - the AP MAC address
// If `type`==`wan`, the following parameters are supported:
// - `mac` - the MAC address of the WAN device
// - `node` - the node ID
// - `peer_ip` - the peer IP address
// - `peer_host` - the peer host name
// - `ip` - the IP address of the WAN device
// - `tunnel_name` - the name of the tunnel
// - `protocol` - the protocol used for the tunnel
// - `auth_algo` - the authentication algorithm used for the tunnel
// - `encrypt_algo` - the encryption algorithm used for the tunnel
// - `ike_version` - the IKE version used for the tunnel
// - `up` - the status of the tunnel (up or down)
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
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
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
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.ResponseTunnelSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseTunnelSearch](decoder)
	return models.NewApiResponse(result, resp), err
}
