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
	"net/http"
)

// OrgsNetworks represents a controller struct.
type OrgsNetworks struct {
	baseController
}

// NewOrgsNetworks creates a new instance of OrgsNetworks.
// It takes a baseController as a parameter and returns a pointer to the OrgsNetworks.
func NewOrgsNetworks(baseController baseController) *OrgsNetworks {
	orgsNetworks := OrgsNetworks{baseController: baseController}
	return &orgsNetworks
}

// ListOrgNetworks takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.Network data and
// an error if there was an issue with the request or response.
// List organization-level Layer 3 network definitions used for service routes and gateway or network template configuration.
func (o *OrgsNetworks) ListOrgNetworks(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.Network],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/networks")
	req.AppendTemplateParams(orgId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.Network
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Network](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgNetwork takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Network data and
// an error if there was an issue with the request or response.
// Create an organization network definition with subnet, gateway,
// VLAN, access, NAT, multicast, tenant, and VPN access settings used for service
// routes and gateway or network template configuration.
// Networks can be used
// - in the gateway configuration to define Layer 3 network settings and policies for traffic entering or leaving the gateway through a specific interface, such as a corporate LAN, guest Wi-Fi, or DMZ network.
// - in the service policies to allow or deny traffic matching the network or to apply specific inspection settings or steering rules.
func (o *OrgsNetworks) CreateOrgNetwork(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Network) (
	models.ApiResponse[models.Network],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/networks")
	req.AppendTemplateParams(orgId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.Network
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Network](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgNetwork takes context, orgId, networkId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete an organization network definition by network ID so it can no longer be referenced by gateway or network template configuration.
func (o *OrgsNetworks) DeleteOrgNetwork(
	ctx context.Context,
	orgId uuid.UUID,
	networkId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/networks/%v")
	req.AppendTemplateParams(orgId, networkId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// GetOrgNetwork takes context, orgId, networkId as parameters and
// returns an models.ApiResponse with models.Network data and
// an error if there was an issue with the request or response.
// Retrieve details for a specific organization network, including subnet, gateway, VLAN, access, NAT, multicast, tenant, and VPN access settings.
func (o *OrgsNetworks) GetOrgNetwork(
	ctx context.Context,
	orgId uuid.UUID,
	networkId uuid.UUID) (
	models.ApiResponse[models.Network],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/networks/%v")
	req.AppendTemplateParams(orgId, networkId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})

	var result models.Network
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Network](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgNetwork takes context, orgId, networkId, body as parameters and
// returns an models.ApiResponse with models.Network data and
// an error if there was an issue with the request or response.
// Update an organization network definition, including subnet, gateway, VLAN, access, NAT, multicast, tenant, and VPN access settings.
func (o *OrgsNetworks) UpdateOrgNetwork(
	ctx context.Context,
	orgId uuid.UUID,
	networkId uuid.UUID,
	body *models.Network) (
	models.ApiResponse[models.Network],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/networks/%v")
	req.AppendTemplateParams(orgId, networkId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.Network
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Network](decoder)
	return models.NewApiResponse(result, resp), err
}
