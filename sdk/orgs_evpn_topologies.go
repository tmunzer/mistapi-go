package mistapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi/sdk/errors"
	"github.com/tmunzer/mistapi/sdk/models"
)

// OrgsEVPNTopologies represents a controller struct.
type OrgsEVPNTopologies struct {
	baseController
}

// NewOrgsEVPNTopologies creates a new instance of OrgsEVPNTopologies.
// It takes a baseController as a parameter and returns a pointer to the OrgsEVPNTopologies.
func NewOrgsEVPNTopologies(baseController baseController) *OrgsEVPNTopologies {
	orgsEVPNTopologies := OrgsEVPNTopologies{baseController: baseController}
	return &orgsEVPNTopologies
}

// ListOrgEvpnTopologies takes context, orgId, page, limit as parameters and
// returns an models.ApiResponse with []models.EvpnTopology data and
// an error if there was an issue with the request or response.
// Get List of the existing Org EVPN topologies
func (o *OrgsEVPNTopologies) ListOrgEvpnTopologies(
	ctx context.Context,
	orgId uuid.UUID,
	page *int,
	limit *int) (
	models.ApiResponse[[]models.EvpnTopology],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/evpn_topologies", orgId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
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

	var result []models.EvpnTopology
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.EvpnTopology](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgEvpnTopology takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.EvpnTopology data and
// an error if there was an issue with the request or response.
// While all the `evpn_id` / `downlink_ips` can be specifidd by hand, the easiest way is to call the `build_vpn_topology` API, allowing you to examine the diff, and update it yourself. You can also simply call it with `overwrite=true` which will apply the updates for you.
// **Notes:**
// 1. You can use `core` / `distribution` / `access` to create a CLOS topology
// 2. You can also use `core` / `distribution` to form a 2-tier EVPN topology where ESI-Lag is configured distribution to connect to access switches
// 3. In a small/medium campus, `collapsed-core` can be used where core switches are the inter-connected to do EVPN
// 4. The API uses a few pre-defined parameters and best-practices to generate the configs. It can be customized by using `evpn_options` in Site Setting / Network Template. (e.g. a different subnet for the underlay)
// #### Collapsed Core
// In a small-medium campus, EVPN can also be enabled only at the core switches (up to 4) by assigning all participating switches with `collapsed-core role`. When there are more than 2 switches, a ring-like topology will be formed.
// #### ESI-Lag
// If the access switchess does not have EVPN support, you can take advantage of EVPN by setting up ESI-Lag on distribution switches
// #### Leaf / Access / Collapsed-Core
// For leaf nodes in a EVPN topology, you’d have to configure the IPs for networks that would participate in EVPN. Optionally, VRFs to isolate traffic from one tenant verus another
func (o *OrgsEVPNTopologies) CreateOrgEvpnTopology(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.EvpnTopology) (
	models.ApiResponse[models.EvpnTopology],
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/orgs/%v/evpn_topologies", orgId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.EvpnTopology
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.EvpnTopology](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgEvpnTopology takes context, orgId, evpnTopologyId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete the Org EVPN Topology
func (o *OrgsEVPNTopologies) DeleteOrgEvpnTopology(
	ctx context.Context,
	orgId uuid.UUID,
	evpnTopologyId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/orgs/%v/evpn_topologies/%v", orgId, evpnTopologyId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// GetOrgEvpnTolopogy takes context, orgId, evpnTopologyId as parameters and
// returns an models.ApiResponse with models.EvpnTopology data and
// an error if there was an issue with the request or response.
// Get One EVPN Topology Detail
func (o *OrgsEVPNTopologies) GetOrgEvpnTolopogy(
	ctx context.Context,
	orgId uuid.UUID,
	evpnTopologyId uuid.UUID) (
	models.ApiResponse[models.EvpnTopology],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/orgs/%v/evpn_topologies/%v", orgId, evpnTopologyId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	var result models.EvpnTopology
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.EvpnTopology](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgEvpnTopology takes context, orgId, evpnTopologyId, body as parameters and
// returns an models.ApiResponse with models.EvpnTopology data and
// an error if there was an issue with the request or response.
// Update the EVPN Topolgy
func (o *OrgsEVPNTopologies) UpdateOrgEvpnTopology(
	ctx context.Context,
	orgId uuid.UUID,
	evpnTopologyId uuid.UUID,
	body *models.EvpnTopology) (
	models.ApiResponse[models.EvpnTopology],
	error) {
	req := o.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/orgs/%v/evpn_topologies/%v", orgId, evpnTopologyId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.EvpnTopology
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.EvpnTopology](decoder)
	return models.NewApiResponse(result, resp), err
}
