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

// OrgsDevicesSSR represents a controller struct.
type OrgsDevicesSSR struct {
	baseController
}

// NewOrgsDevicesSSR creates a new instance of OrgsDevicesSSR.
// It takes a baseController as a parameter and returns a pointer to the OrgsDevicesSSR.
func NewOrgsDevicesSSR(baseController baseController) *OrgsDevicesSSR {
	orgsDevicesSSR := OrgsDevicesSSR{baseController: baseController}
	return &orgsDevicesSSR
}

// GetOrg128TRegistrationCommands takes context, orgId, ttl, assetIds as parameters and
// returns an models.ApiResponse with models.ResponseRouterSsrRegisterCmd data and
// an error if there was an issue with the request or response.
// Deprecated: getOrg128TRegistrationCommands is deprecated
// 128T devices can be managed/adopted by Mist.
func (o *OrgsDevicesSSR) GetOrg128TRegistrationCommands(
	ctx context.Context,
	orgId uuid.UUID,
	ttl *int,
	assetIds []string) (
	models.ApiResponse[models.ResponseRouterSsrRegisterCmd],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/128routers/register_cmd")
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
	if ttl != nil {
		req.QueryParam("ttl", *ttl)
	}
	if assetIds != nil {
		req.QueryParam("asset_ids", assetIds)
	}

	var result models.ResponseRouterSsrRegisterCmd
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseRouterSsrRegisterCmd](decoder)
	return models.NewApiResponse(result, resp), err
}

// ExportOrgSsrIdTokens takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.ResponseSsrExportIdTokens data and
// an error if there was an issue with the request or response.
// Export SSR ID tokens for the requested device MAC addresses so they can be imported into Conductor during SSR onboarding.
func (o *OrgsDevicesSSR) ExportOrgSsrIdTokens(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.MacAddresses) (
	models.ApiResponse[models.ResponseSsrExportIdTokens],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/ssr/export_idtokens")
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

	var result models.ResponseSsrExportIdTokens
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSsrExportIdTokens](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetOrgSsrRegistrationCommands takes context, orgId, ttl, assetIds as parameters and
// returns an models.ApiResponse with models.ResponseRouterSsrRegisterCmd data and
// an error if there was an issue with the request or response.
// Return the registration token and conductor or router commands used to register SSR routers with Mist. The optional TTL controls token validity, and asset IDs can restrict registration to specific assets.
func (o *OrgsDevicesSSR) GetOrgSsrRegistrationCommands(
	ctx context.Context,
	orgId uuid.UUID,
	ttl *int,
	assetIds []string) (
	models.ApiResponse[models.ResponseRouterSsrRegisterCmd],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/ssr/register_cmd")
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
	if ttl != nil {
		req.QueryParam("ttl", *ttl)
	}
	if assetIds != nil {
		req.QueryParam("asset_ids", assetIds)
	}

	var result models.ResponseRouterSsrRegisterCmd
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseRouterSsrRegisterCmd](decoder)
	return models.NewApiResponse(result, resp), err
}
