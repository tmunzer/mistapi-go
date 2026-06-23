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

// OrgsIntegrationSkyATP represents a controller struct.
type OrgsIntegrationSkyATP struct {
	baseController
}

// NewOrgsIntegrationSkyATP creates a new instance of OrgsIntegrationSkyATP.
// It takes a baseController as a parameter and returns a pointer to the OrgsIntegrationSkyATP.
func NewOrgsIntegrationSkyATP(baseController baseController) *OrgsIntegrationSkyATP {
	orgsIntegrationSkyATP := OrgsIntegrationSkyATP{baseController: baseController}
	return &orgsIntegrationSkyATP
}

// UdpateOrgAtpAllowedList takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.SkyatpList data and
// an error if there was an issue with the request or response.
// Update the Sky ATP SecIntel allowlist with domain and IP address entries for the organization.
func (o *OrgsIntegrationSkyATP) UdpateOrgAtpAllowedList(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.SkyatpList) (
	models.ApiResponse[models.SkyatpList],
	error) {
	req := o.prepareRequest(
		ctx,
		"PUT",
		"/api/v1/orgs/%v/setting/skyatp/secintel_allowlist",
	)
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

	var result models.SkyatpList
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SkyatpList](decoder)
	return models.NewApiResponse(result, resp), err
}

// UdpateOrgAtpBlockedList takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.SkyatpList data and
// an error if there was an issue with the request or response.
// Update the Sky ATP SecIntel blocklist with domain and IP address entries for the organization.
func (o *OrgsIntegrationSkyATP) UdpateOrgAtpBlockedList(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.SkyatpList) (
	models.ApiResponse[models.SkyatpList],
	error) {
	req := o.prepareRequest(
		ctx,
		"PUT",
		"/api/v1/orgs/%v/setting/skyatp/secintel_blocklist",
	)
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

	var result models.SkyatpList
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SkyatpList](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgSkyAtpIntegration takes context, orgId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Remove the Sky ATP integration configuration from the organization.
func (o *OrgsIntegrationSkyATP) DeleteOrgSkyAtpIntegration(
	ctx context.Context,
	orgId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/setting/skyatp/setup")
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// GetOrgSkyAtpIntegration takes context, orgId as parameters and
// returns an models.ApiResponse with models.AccountSkyatpData data and
// an error if there was an issue with the request or response.
// Return the Sky ATP integration configuration, including linked realm information and generated SecIntel allowlist and blocklist URLs.
func (o *OrgsIntegrationSkyATP) GetOrgSkyAtpIntegration(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[models.AccountSkyatpData],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/setting/skyatp/setup")
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

	var result models.AccountSkyatpData
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AccountSkyatpData](decoder)
	return models.NewApiResponse(result, resp), err
}

// SetupOrgAtpIntegration takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.AccountSkyatpData data and
// an error if there was an issue with the request or response.
// Configure the Sky ATP integration by linking or creating the Sky ATP realm with the supplied cloud, realm, username, and password. The integration enables Security Intelligence and Advanced Anti-Malware features, with SecIntel configuration for command-and-control, DNS feeds, infected hosts, blocklists, and allowlists.
func (o *OrgsIntegrationSkyATP) SetupOrgAtpIntegration(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.AccountSkyatpConfig) (
	models.ApiResponse[models.AccountSkyatpData],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/setting/skyatp/setup")
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

	var result models.AccountSkyatpData
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AccountSkyatpData](decoder)
	return models.NewApiResponse(result, resp), err
}

// UdpateOrgAtpIntegration takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.AccountSkyatpInfo data and
// an error if there was an issue with the request or response.
// Update Sky ATP SecIntel feed configuration, including the third-party threat feeds enabled for the organization.
func (o *OrgsIntegrationSkyATP) UdpateOrgAtpIntegration(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.AccountSkyatpData) (
	models.ApiResponse[models.AccountSkyatpInfo],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/setting/skyatp/setup")
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

	var result models.AccountSkyatpInfo
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.AccountSkyatpInfo](decoder)
	return models.NewApiResponse(result, resp), err
}
