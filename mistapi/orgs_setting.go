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

// OrgsSetting represents a controller struct.
type OrgsSetting struct {
	baseController
}

// NewOrgsSetting creates a new instance of OrgsSetting.
// It takes a baseController as a parameter and returns a pointer to the OrgsSetting.
func NewOrgsSetting(baseController baseController) *OrgsSetting {
	orgsSetting := OrgsSetting{baseController: baseController}
	return &orgsSetting
}

// GetOrgSettings takes context, orgId as parameters and
// returns an models.ApiResponse with models.OrgSetting data and
// an error if there was an issue with the request or response.
// Return organization-wide settings, including feature flags, automatic device assignment rules, management connectivity, packet capture, security controls, and integration configuration.
func (o *OrgsSetting) GetOrgSettings(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[models.OrgSetting],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/setting")
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

	var result models.OrgSetting
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.OrgSetting](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgSettings takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.OrgSetting data and
// an error if there was an issue with the request or response.
// Update organization-wide settings such as automatic device assignment rules, management connectivity, packet capture, password policy, security controls, tags, and integration options.
func (o *OrgsSetting) UpdateOrgSettings(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.OrgSetting) (
	models.ApiResponse[models.OrgSetting],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/setting")
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

	var result models.OrgSetting
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.OrgSetting](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgWirelessClientsBlocklist takes context, orgId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Clear the organization wireless client blocklist by removing all blocked client MAC addresses.
func (o *OrgsSetting) DeleteOrgWirelessClientsBlocklist(
	ctx context.Context,
	orgId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/setting/blacklist")
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

// CreateOrgWirelessClientsBlocklist takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.MacAddresses data and
// an error if there was an issue with the request or response.
// Replace the organization wireless client blocklist with the supplied client MAC addresses. The list can contain up to 1000 MAC addresses; retrieve the current list from the `blacklist_url` field in organization settings.
func (o *OrgsSetting) CreateOrgWirelessClientsBlocklist(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.MacAddresses) (
	models.ApiResponse[models.MacAddresses],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/setting/blacklist")
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

	var result models.MacAddresses
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MacAddresses](decoder)
	return models.NewApiResponse(result, resp), err
}

// SetOrgCustomBucket takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.ResponsePcapBucketConfig data and
// an error if there was an issue with the request or response.
// Start custom packet capture bucket setup by saving the bucket name and having Mist write a `MIST_TOKEN` file to the bucket. Complete ownership verification with the verify endpoint by submitting the token contents.
func (o *OrgsSetting) SetOrgCustomBucket(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.PcapBucket) (
	models.ApiResponse[models.ResponsePcapBucketConfig],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/setting/pcap_bucket/setup")
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

	var result models.ResponsePcapBucketConfig
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponsePcapBucketConfig](decoder)
	return models.NewApiResponse(result, resp), err
}

// VerifyOrgCustomBucket takes context, orgId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Verify ownership of a custom packet capture bucket by submitting the token read from the `MIST_TOKEN` file. If verification succeeds, Mist creates a `VERIFIED` file in the bucket.
func (o *OrgsSetting) VerifyOrgCustomBucket(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.PcapBucketVerify) (
	*http.Response,
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		"/api/v1/orgs/%v/setting/pcap_bucket/verify",
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

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}
