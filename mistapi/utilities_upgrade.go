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

// UtilitiesUpgrade represents a controller struct.
type UtilitiesUpgrade struct {
	baseController
}

// NewUtilitiesUpgrade creates a new instance of UtilitiesUpgrade.
// It takes a baseController as a parameter and returns a pointer to the UtilitiesUpgrade.
func NewUtilitiesUpgrade(baseController baseController) *UtilitiesUpgrade {
	utilitiesUpgrade := UtilitiesUpgrade{baseController: baseController}
	return &utilitiesUpgrade
}

// ListOrgDeviceUpgrades takes context, orgId as parameters and
// returns an models.ApiResponse with []models.UpgradeOrgDevicesItem data and
// an error if there was an issue with the request or response.
// List organization-level device upgrade jobs, including the site-level upgrade jobs created under each organization upgrade.
func (u *UtilitiesUpgrade) ListOrgDeviceUpgrades(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[[]models.UpgradeOrgDevicesItem],
	error) {
	req := u.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/devices/upgrade")
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

	var result []models.UpgradeOrgDevicesItem
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.UpgradeOrgDevicesItem](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpgradeOrgDevices takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.ResponseUpgradeOrgDevices data and
// an error if there was an issue with the request or response.
// Start an organization-level device upgrade job across selected sites. The request selects device type, sites, models, firmware versions, and upgrade strategy; AP-specific and Junos-specific options apply only where supported.
func (u *UtilitiesUpgrade) UpgradeOrgDevices(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.UpgradeOrgDevices) (
	models.ApiResponse[models.ResponseUpgradeOrgDevices],
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/devices/upgrade")
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

	var result models.ResponseUpgradeOrgDevices
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseUpgradeOrgDevices](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetOrgDeviceUpgrade takes context, orgId, upgradeId as parameters and
// returns an models.ApiResponse with models.ResponseUpgradeOrgDevices data and
// an error if there was an issue with the request or response.
// Retrieve details for an organization-level device upgrade job, including per-site upgrade status and device targets.
func (u *UtilitiesUpgrade) GetOrgDeviceUpgrade(
	ctx context.Context,
	orgId uuid.UUID,
	upgradeId uuid.UUID) (
	models.ApiResponse[models.ResponseUpgradeOrgDevices],
	error) {
	req := u.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/devices/upgrade/%v")
	req.AppendTemplateParams(orgId, upgradeId)
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

	var result models.ResponseUpgradeOrgDevices
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseUpgradeOrgDevices](decoder)
	return models.NewApiResponse(result, resp), err
}

// CancelOrgDeviceUpgrade takes context, orgId, upgradeId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Cancel an organization-level device upgrade job on a best-effort basis. Devices that have already completed the upgrade are not changed.
func (u *UtilitiesUpgrade) CancelOrgDeviceUpgrade(
	ctx context.Context,
	orgId uuid.UUID,
	upgradeId uuid.UUID) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/devices/upgrade/%v/cancel")
	req.AppendTemplateParams(orgId, upgradeId)
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

// ListOrgAvailableDeviceVersions takes context, orgId, mType, model as parameters and
// returns an models.ApiResponse with []models.DeviceVersionItem data and
// an error if there was an issue with the request or response.
// List available firmware versions for organization devices, optionally filtered by device type and model.
func (u *UtilitiesUpgrade) ListOrgAvailableDeviceVersions(
	ctx context.Context,
	orgId uuid.UUID,
	mType *models.DeviceTypeDefaultApEnum,
	model *string) (
	models.ApiResponse[[]models.DeviceVersionItem],
	error) {
	req := u.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/devices/versions")
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
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}

	var result []models.DeviceVersionItem
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.DeviceVersionItem](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpgradeOrgJsiDevice takes context, orgId, deviceMac, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Start a software upgrade for a JSI-connected device identified by MAC address using the requested target version.
func (u *UtilitiesUpgrade) UpgradeOrgJsiDevice(
	ctx context.Context,
	orgId uuid.UUID,
	deviceMac string,
	body *models.VersionString) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/jsi/devices/%v/upgrade")
	req.AppendTemplateParams(orgId, deviceMac)
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// ListOrgMxEdgeUpgrades takes context, orgId as parameters and
// returns an models.ApiResponse with []models.ResponseMxedgeUpgrade data and
// an error if there was an issue with the request or response.
// List Mist Edge upgrade requests for the organization, including status, rollout strategy, target versions, and per-status target counts.
func (u *UtilitiesUpgrade) ListOrgMxEdgeUpgrades(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[[]models.ResponseMxedgeUpgrade],
	error) {
	req := u.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/mxedges/upgrade")
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

	var result []models.ResponseMxedgeUpgrade
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ResponseMxedgeUpgrade](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpgradeOrgMxEdges takes context, orgId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Schedule a Mist Edge upgrade for selected Mist Edges, using service target versions or an optional Linux distro upgrade with rollout strategy and canary settings.
func (u *UtilitiesUpgrade) UpgradeOrgMxEdges(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.MxedgeUpgradeMulti) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/mxedges/upgrade")
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// GetOrgMxEdgeUpgrade takes context, orgId, upgradeId as parameters and
// returns an models.ApiResponse with models.ResponseMxedgeUpgrade data and
// an error if there was an issue with the request or response.
// Retrieve status, rollout strategy, target versions, and target counts for a specific Mist Edge upgrade request.
func (u *UtilitiesUpgrade) GetOrgMxEdgeUpgrade(
	ctx context.Context,
	orgId uuid.UUID,
	upgradeId uuid.UUID) (
	models.ApiResponse[models.ResponseMxedgeUpgrade],
	error) {
	req := u.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/mxedges/upgrade/%v")
	req.AppendTemplateParams(orgId, upgradeId)
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

	var result models.ResponseMxedgeUpgrade
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseMxedgeUpgrade](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgMxEdgeUpgrade takes context, orgId, upgradeId, body as parameters and
// returns an models.ApiResponse with models.ResponseMxedgeUpgrade data and
// an error if there was an issue with the request or response.
// Update a queued Mist Edge upgrade request, such as target versions, rollout strategy, start time, or target Mist Edge IDs. Only upgrades in `queued` state can be updated.
func (u *UtilitiesUpgrade) UpdateOrgMxEdgeUpgrade(
	ctx context.Context,
	orgId uuid.UUID,
	upgradeId uuid.UUID,
	body *models.MxedgeUpgradeMulti) (
	models.ApiResponse[models.ResponseMxedgeUpgrade],
	error) {
	req := u.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/mxedges/upgrade/%v")
	req.AppendTemplateParams(orgId, upgradeId)
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

	var result models.ResponseMxedgeUpgrade
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseMxedgeUpgrade](decoder)
	return models.NewApiResponse(result, resp), err
}

// CancelOrgMxEdgeUpgrade takes context, orgId, upgradeId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Cancel a Mist Edge upgrade request on a best-effort basis. Mist Edges that have already been upgraded are not changed.
func (u *UtilitiesUpgrade) CancelOrgMxEdgeUpgrade(
	ctx context.Context,
	orgId uuid.UUID,
	upgradeId uuid.UUID) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/mxedges/upgrade/%v/cancel")
	req.AppendTemplateParams(orgId, upgradeId)
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

// ListOrgSsrUpgrades takes context, orgId as parameters and
// returns an models.ApiResponse with []models.ResponseSsrUpgrade data and
// an error if there was an issue with the request or response.
// List SSR firmware upgrade jobs for the organization, including status, rollout strategy, target versions, release channel, and device counts.
func (u *UtilitiesUpgrade) ListOrgSsrUpgrades(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[[]models.ResponseSsrUpgrade],
	error) {
	req := u.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/ssr/upgrade")
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

	var result []models.ResponseSsrUpgrade
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ResponseSsrUpgrade](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpgradeOrgSsrs takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.ResponseSsrUpgrade data and
// an error if there was an issue with the request or response.
// Create an SSR firmware upgrade job for selected devices, with firmware version or channel, rollout strategy, and optional download or reboot timing.
func (u *UtilitiesUpgrade) UpgradeOrgSsrs(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.SsrUpgradeMulti) (
	models.ApiResponse[models.ResponseSsrUpgrade],
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/ssr/upgrade")
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

	var result models.ResponseSsrUpgrade
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSsrUpgrade](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetOrgSsrUpgrade takes context, orgId, upgradeId as parameters and
// returns an models.ApiResponse with models.ResponseSsrUpgradeStatus data and
// an error if there was an issue with the request or response.
// Return detailed status for an SSR firmware upgrade job, including target device IDs grouped by upgrade status.
func (u *UtilitiesUpgrade) GetOrgSsrUpgrade(
	ctx context.Context,
	orgId uuid.UUID,
	upgradeId uuid.UUID) (
	models.ApiResponse[models.ResponseSsrUpgradeStatus],
	error) {
	req := u.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/ssr/upgrade/%v/cancel")
	req.AppendTemplateParams(orgId, upgradeId)
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

	var result models.ResponseSsrUpgradeStatus
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSsrUpgradeStatus](decoder)
	return models.NewApiResponse(result, resp), err
}

// CancelOrgSsrUpgrade takes context, orgId, upgradeId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Cancel an SSR firmware upgrade job on a best-effort basis. Devices that have already upgraded are not changed.
func (u *UtilitiesUpgrade) CancelOrgSsrUpgrade(
	ctx context.Context,
	orgId uuid.UUID,
	upgradeId uuid.UUID) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/ssr/upgrade/%v/cancel")
	req.AppendTemplateParams(orgId, upgradeId)
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

// ListOrgAvailableSsrVersions takes context, orgId, channel, mac as parameters and
// returns an models.ApiResponse with []models.SsrVersion data and
// an error if there was an issue with the request or response.
// List SSR firmware versions available for upgrade, optionally filtered by release channel and one or more SSR MAC addresses.
func (u *UtilitiesUpgrade) ListOrgAvailableSsrVersions(
	ctx context.Context,
	orgId uuid.UUID,
	channel *models.SsrVersionChannelEnum,
	mac *string) (
	models.ApiResponse[[]models.SsrVersion],
	error) {
	req := u.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/ssr/versions")
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
	if channel != nil {
		req.QueryParam("channel", *channel)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}

	var result []models.SsrVersion
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.SsrVersion](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSiteDeviceUpgrades takes context, siteId, status as parameters and
// returns an models.ApiResponse with []models.ResponseSiteDeviceUpgradesItem data and
// an error if there was an issue with the request or response.
// List device upgrade operations for a site, optionally filtered by upgrade status. Use [List Org Device Upgrades]($e/Utilities%20Upgrade/listOrgDeviceUpgrades) to retrieve device upgrade operations across the organization.
func (u *UtilitiesUpgrade) ListSiteDeviceUpgrades(
	ctx context.Context,
	siteId uuid.UUID,
	status *models.UpgradeDeviceStatusEnum) (
	models.ApiResponse[[]models.ResponseSiteDeviceUpgradesItem],
	error) {
	req := u.prepareRequest(ctx, "GET", "/api/v1/sites/%v/devices/upgrade")
	req.AppendTemplateParams(siteId)
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
	if status != nil {
		req.QueryParam("status", *status)
	}

	var result []models.ResponseSiteDeviceUpgradesItem
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ResponseSiteDeviceUpgradesItem](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpgradeSiteDevices takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.ResponseUpgradeId data and
// an error if there was an issue with the request or response.
// Upgrade Site Device
// **Note**: this call doesn’t guarantee the devices to be upgraded right away (they may be offline)
func (u *UtilitiesUpgrade) UpgradeSiteDevices(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.UpgradeSiteDevices) (
	models.ApiResponse[models.ResponseUpgradeId],
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/upgrade")
	req.AppendTemplateParams(siteId)
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

	var result models.ResponseUpgradeId
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseUpgradeId](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteDeviceUpgrade takes context, siteId, upgradeId as parameters and
// returns an models.ApiResponse with models.ResponseSiteDeviceUpgrade data and
// an error if there was an issue with the request or response.
// Get Site Device Upgrade
func (u *UtilitiesUpgrade) GetSiteDeviceUpgrade(
	ctx context.Context,
	siteId uuid.UUID,
	upgradeId uuid.UUID) (
	models.ApiResponse[models.ResponseSiteDeviceUpgrade],
	error) {
	req := u.prepareRequest(ctx, "GET", "/api/v1/sites/%v/devices/upgrade/%v")
	req.AppendTemplateParams(siteId, upgradeId)
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

	var result models.ResponseSiteDeviceUpgrade
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSiteDeviceUpgrade](decoder)
	return models.NewApiResponse(result, resp), err
}

// CancelSiteDeviceUpgrade takes context, siteId, upgradeId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Best effort to cancel an upgrade. Devices which are already upgraded wont be touched
func (u *UtilitiesUpgrade) CancelSiteDeviceUpgrade(
	ctx context.Context,
	siteId uuid.UUID,
	upgradeId uuid.UUID) (
	*http.Response,
	error) {
	req := u.prepareRequest(
		ctx,
		"POST",
		"/api/v1/sites/%v/devices/upgrade/%v/cancel",
	)
	req.AppendTemplateParams(siteId, upgradeId)
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

// ListSiteAvailableDeviceVersions takes context, siteId, mType, model as parameters and
// returns an models.ApiResponse with []models.DeviceVersionItem data and
// an error if there was an issue with the request or response.
// List firmware versions available for devices in a site, optionally filtered by device type and model. Use [List Org Available Device Versions]($e/Utilities%20Upgrade/listOrgAvailableDeviceVersions) to retrieve available device versions across the organization.
func (u *UtilitiesUpgrade) ListSiteAvailableDeviceVersions(
	ctx context.Context,
	siteId uuid.UUID,
	mType *models.DeviceTypeDefaultApEnum,
	model *string) (
	models.ApiResponse[[]models.DeviceVersionItem],
	error) {
	req := u.prepareRequest(ctx, "GET", "/api/v1/sites/%v/devices/versions")
	req.AppendTemplateParams(siteId)
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
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}

	var result []models.DeviceVersionItem
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.DeviceVersionItem](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpgradeDevice takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.ResponseDeviceUpgrade data and
// an error if there was an issue with the request or response.
// Device Upgrade
func (u *UtilitiesUpgrade) UpgradeDevice(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID,
	body *models.DeviceUpgrade) (
	models.ApiResponse[models.ResponseDeviceUpgrade],
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/%v/upgrade")
	req.AppendTemplateParams(siteId, deviceId)
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

	var result models.ResponseDeviceUpgrade
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseDeviceUpgrade](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSiteMxEdgeUpgrades takes context, siteId as parameters and
// returns an models.ApiResponse with []models.ResponseMxedgeUpgrade data and
// an error if there was an issue with the request or response.
// List Mist Edge upgrade operations for a site. Use [List Org Mist Edge Upgrades]($e/Utilities%20Upgrade/listOrgMxEdgeUpgrades) to retrieve Mist Edge upgrade operations across the organization.
func (u *UtilitiesUpgrade) ListSiteMxEdgeUpgrades(
	ctx context.Context,
	siteId uuid.UUID) (
	models.ApiResponse[[]models.ResponseMxedgeUpgrade],
	error) {
	req := u.prepareRequest(ctx, "GET", "/api/v1/sites/%v/mxedges/upgrade")
	req.AppendTemplateParams(siteId)
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

	var result []models.ResponseMxedgeUpgrade
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ResponseMxedgeUpgrade](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpgradeSiteMxEdges takes context, siteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Upgrade Mist Edges in a Site.
// See [Org Mist Edges](/#tag/Utilities-Upgrade/operation/upgradeOrgMxEdges) for package upgrades
// See [Org Mist Edges Distro](/#tag/Utilities-Upgrade/operation/upgradeOrgMxEdges) for distro upgrades
func (u *UtilitiesUpgrade) UpgradeSiteMxEdges(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.MxedgeUpgradeMulti) (
	*http.Response,
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/mxedges/upgrade")
	req.AppendTemplateParams(siteId)
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}

// GetSiteMxEdgeUpgrade takes context, siteId, upgradeId as parameters and
// returns an models.ApiResponse with models.ResponseMxedgeUpgrade data and
// an error if there was an issue with the request or response.
// Get Mist Edge Upgrade
func (u *UtilitiesUpgrade) GetSiteMxEdgeUpgrade(
	ctx context.Context,
	siteId uuid.UUID,
	upgradeId uuid.UUID) (
	models.ApiResponse[models.ResponseMxedgeUpgrade],
	error) {
	req := u.prepareRequest(ctx, "GET", "/api/v1/sites/%v/mxedges/upgrade/%v")
	req.AppendTemplateParams(siteId, upgradeId)
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

	var result models.ResponseMxedgeUpgrade
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseMxedgeUpgrade](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateSiteMxEdgeUpgrade takes context, siteId, upgradeId, body as parameters and
// returns an models.ApiResponse with models.ResponseMxedgeUpgrade data and
// an error if there was an issue with the request or response.
// Update Mist Edge Upgrade. Only upgrades in `queued` state can be updated.
func (u *UtilitiesUpgrade) UpdateSiteMxEdgeUpgrade(
	ctx context.Context,
	siteId uuid.UUID,
	upgradeId uuid.UUID,
	body *models.MxedgeUpgradeMulti) (
	models.ApiResponse[models.ResponseMxedgeUpgrade],
	error) {
	req := u.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/mxedges/upgrade/%v")
	req.AppendTemplateParams(siteId, upgradeId)
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

	var result models.ResponseMxedgeUpgrade
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseMxedgeUpgrade](decoder)
	return models.NewApiResponse(result, resp), err
}

// CancelSiteMxEdgeUpgrade takes context, siteId, upgradeId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Cancel Mist Edge Upgrade. Best effort to cancel an upgrade. MxEdges which are already upgraded won't be touched.
func (u *UtilitiesUpgrade) CancelSiteMxEdgeUpgrade(
	ctx context.Context,
	siteId uuid.UUID,
	upgradeId uuid.UUID) (
	*http.Response,
	error) {
	req := u.prepareRequest(
		ctx,
		"POST",
		"/api/v1/sites/%v/mxedges/upgrade/%v/cancel",
	)
	req.AppendTemplateParams(siteId, upgradeId)
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

// GetSiteSsrUpgrade takes context, siteId, upgradeId as parameters and
// returns an models.ApiResponse with models.ResponseSsrUpgradeStatus data and
// an error if there was an issue with the request or response.
// Get Specific Site SSR Upgrade
func (u *UtilitiesUpgrade) GetSiteSsrUpgrade(
	ctx context.Context,
	siteId uuid.UUID,
	upgradeId uuid.UUID) (
	models.ApiResponse[models.ResponseSsrUpgradeStatus],
	error) {
	req := u.prepareRequest(ctx, "GET", "/api/v1/sites/%v/ssr/upgrade/%v")
	req.AppendTemplateParams(siteId, upgradeId)
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

	var result models.ResponseSsrUpgradeStatus
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSsrUpgradeStatus](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpgradeSsr takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.ResponseSsrUpgrade data and
// an error if there was an issue with the request or response.
// Upgrade Site SSR device
func (u *UtilitiesUpgrade) UpgradeSsr(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID,
	body *models.SsrUpgrade) (
	models.ApiResponse[models.ResponseSsrUpgrade],
	error) {
	req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/ssr/%v/upgrade")
	req.AppendTemplateParams(siteId, deviceId)
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

	var result models.ResponseSsrUpgrade
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSsrUpgrade](decoder)
	return models.NewApiResponse(result, resp), err
}
