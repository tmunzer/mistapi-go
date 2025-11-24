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
// Export IDTokens from Mist to import into Conductor to securely allow SSR devices during onboarding
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
// SSR devices can be managed/adopted by Mist.
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
