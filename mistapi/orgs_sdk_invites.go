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

// OrgsSDKInvites represents a controller struct.
type OrgsSDKInvites struct {
	baseController
}

// NewOrgsSDKInvites creates a new instance of OrgsSDKInvites.
// It takes a baseController as a parameter and returns a pointer to the OrgsSDKInvites.
func NewOrgsSDKInvites(baseController baseController) *OrgsSDKInvites {
	orgsSDKInvites := OrgsSDKInvites{baseController: baseController}
	return &orgsSDKInvites
}

// ActivateSdkInvite takes context, secret, body as parameters and
// returns an models.ApiResponse with models.ResponseMobileVerifySecret data and
// an error if there was an issue with the request or response.
// Activate a mobile SDK invite by verifying the invite secret and binding it to the supplied device identifier. The response returns the device-specific secret used by the mobile SDK client.
func (o *OrgsSDKInvites) ActivateSdkInvite(
	ctx context.Context,
	secret string,
	body *models.DeviceIdString) (
	models.ApiResponse[models.ResponseMobileVerifySecret],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/mobile/verify/%v")
	req.AppendTemplateParams(secret)
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

	var result models.ResponseMobileVerifySecret
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseMobileVerifySecret](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSdkInvites takes context, orgId as parameters and
// returns an models.ApiResponse with []models.Sdkinvite data and
// an error if there was an issue with the request or response.
// List SDK invites configured for the organization. SDK invites are used to onboard mobile SDK clients and can define whether an invite is enabled, limited by usage quota, or scoped to a site.
func (o *OrgsSDKInvites) ListSdkInvites(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[[]models.Sdkinvite],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/sdkinvites")
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

	var result []models.Sdkinvite
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Sdkinvite](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateSdkInvite takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Sdkinvite data and
// an error if there was an issue with the request or response.
// Create an SDK invite that mobile SDK clients can use to onboard into the organization. The invite can be enabled or disabled, limited by usage quota, and associated with a site.
func (o *OrgsSDKInvites) CreateSdkInvite(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.Sdkinvite) (
	models.ApiResponse[models.Sdkinvite],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/sdkinvites")
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

	var result models.Sdkinvite
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sdkinvite](decoder)
	return models.NewApiResponse(result, resp), err
}

// RevokeSdkInvite takes context, orgId, sdkinviteId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Revoke an SDK invite so it can no longer be used for mobile SDK client onboarding.
func (o *OrgsSDKInvites) RevokeSdkInvite(
	ctx context.Context,
	orgId uuid.UUID,
	sdkinviteId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/sdkinvites/%v")
	req.AppendTemplateParams(orgId, sdkinviteId)
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

// GetSdkInvite takes context, orgId, sdkinviteId as parameters and
// returns an models.ApiResponse with models.Sdkinvite data and
// an error if there was an issue with the request or response.
// Return the configuration and status of an SDK invite, including enablement, expiration time, usage quota, and site scope.
func (o *OrgsSDKInvites) GetSdkInvite(
	ctx context.Context,
	orgId uuid.UUID,
	sdkinviteId uuid.UUID) (
	models.ApiResponse[models.Sdkinvite],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/sdkinvites/%v")
	req.AppendTemplateParams(orgId, sdkinviteId)
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

	var result models.Sdkinvite
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sdkinvite](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateSdkInvite takes context, orgId, sdkinviteId, body as parameters and
// returns an models.ApiResponse with models.Sdkinvite data and
// an error if there was an issue with the request or response.
// Update an SDK invite's onboarding settings, such as its display name, enabled state, expiration time, quota, or site association.
func (o *OrgsSDKInvites) UpdateSdkInvite(
	ctx context.Context,
	orgId uuid.UUID,
	sdkinviteId uuid.UUID,
	body *models.Sdkinvite) (
	models.ApiResponse[models.Sdkinvite],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/sdkinvites/%v")
	req.AppendTemplateParams(orgId, sdkinviteId)
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

	var result models.Sdkinvite
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Sdkinvite](decoder)
	return models.NewApiResponse(result, resp), err
}

// SendSdkInviteEmail takes context, orgId, sdkinviteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Send the SDK invite to a recipient email address so the recipient can onboard a mobile SDK client.
func (o *OrgsSDKInvites) SendSdkInviteEmail(
	ctx context.Context,
	orgId uuid.UUID,
	sdkinviteId uuid.UUID,
	body *models.EmailString) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/sdkinvites/%v/email")
	req.AppendTemplateParams(orgId, sdkinviteId)
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

// GetSdkInviteQrCode takes context, orgId, sdkinviteId as parameters and
// returns an models.ApiResponse with []byte data and
// an error if there was an issue with the request or response.
// Download a QR code image for the SDK invite so it can be scanned by a mobile SDK client during onboarding.
func (o *OrgsSDKInvites) GetSdkInviteQrCode(
	ctx context.Context,
	orgId uuid.UUID,
	sdkinviteId uuid.UUID) (
	models.ApiResponse[[]byte],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/sdkinvites/%v/qrcode")
	req.AppendTemplateParams(orgId, sdkinviteId)
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

	stream, resp, err := req.CallAsStream()
	if err != nil {
		return models.NewApiResponse(stream, resp), err
	}
	return models.NewApiResponse(stream, resp), err
}

// SendSdkInviteSms takes context, orgId, sdkinviteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Send the SDK invite to a phone number by SMS so the recipient can onboard a mobile SDK client.
func (o *OrgsSDKInvites) SendSdkInviteSms(
	ctx context.Context,
	orgId uuid.UUID,
	sdkinviteId uuid.UUID,
	body *models.SdkInviteSms) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/sdkinvites/%v/sms")
	req.AppendTemplateParams(orgId, sdkinviteId)
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
