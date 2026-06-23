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

// OrgsPsks represents a controller struct.
type OrgsPsks struct {
	baseController
}

// NewOrgsPsks creates a new instance of OrgsPsks.
// It takes a baseController as a parameter and returns a pointer to the OrgsPsks.
func NewOrgsPsks(baseController baseController) *OrgsPsks {
	orgsPsks := OrgsPsks{baseController: baseController}
	return &orgsPsks
}

// ListOrgPsks takes context, orgId, name, ssid, role, limit, page as parameters and
// returns an models.ApiResponse with []models.Psk data and
// an error if there was an issue with the request or response.
// List organization personal PSKs for WLAN access, optionally filtering by name, SSID, or role.
func (o *OrgsPsks) ListOrgPsks(
	ctx context.Context,
	orgId uuid.UUID,
	name *string,
	ssid *string,
	role *string,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.Psk],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/psks")
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
	if name != nil {
		req.QueryParam("name", *name)
	}
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if role != nil {
		req.QueryParam("role", *role)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.Psk
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Psk](decoder)
	return models.NewApiResponse(result, resp), err
}

// CreateOrgPsk takes context, orgId, upsert, body as parameters and
// returns an models.ApiResponse with models.Psk data and
// an error if there was an issue with the request or response.
// Create an organization personal PSK for WLAN access, including SSID, passphrase, usage mode, optional client MAC binding, role, VLAN, and expiration settings.
// When `usage`==`macs`, corresponding "macs" field will hold a list consisting of client MAC addresses (["xx:xx:xx:xx:xx",...]) or mac patterns(["xx:xx:*","xx*",...]) or both (["xx:xx:xx:xx:xx:xx", "xx:*", ...]). This list is capped at 5000
func (o *OrgsPsks) CreateOrgPsk(
	ctx context.Context,
	orgId uuid.UUID,
	upsert *bool,
	body *models.Psk) (
	models.ApiResponse[models.Psk],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/psks")
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
	if upsert != nil {
		req.QueryParam("upsert", *upsert)
	}
	if body != nil {
		req.Json(body)
	}

	var result models.Psk
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Psk](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgMultiplePsks takes context, orgId, body as parameters and
// returns an models.ApiResponse with []models.Psk data and
// an error if there was an issue with the request or response.
// Update multiple organization personal PSKs in one request, including passphrase, usage mode, role, VLAN, expiration, and notification settings.
func (o *OrgsPsks) UpdateOrgMultiplePsks(
	ctx context.Context,
	orgId uuid.UUID,
	body []models.Psk) (
	models.ApiResponse[[]models.Psk],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/psks")
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

	var result []models.Psk
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Psk](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgPskList takes context, orgId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete one or more organization PSKs by ID.
// The request accepts a single PSK ID string or a list of PSK ID strings.
// **Warning**: If no PSK IDs are provided in the request, all organization PSKs will be deleted and clients will no longer be able to authenticate to any SSID with a PSK.
func (o *OrgsPsks) DeleteOrgPskList(
	ctx context.Context,
	orgId uuid.UUID,
	body *models.PskIdList) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/psks/delete")
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

// ImportOrgPsks takes context, orgId, file as parameters and
// returns an models.ApiResponse with []models.Psk data and
// an error if there was an issue with the request or response.
// Import organization PSKs from a CSV file or JSON payload.
// ## CSV File Format
// ```
// PSK Import CSV File Format:
// name,ssid,passphrase,usage,vlan_id,mac,max_usage,role,expire_time,notify_expiry,expiry_notification_time,notify_on_create_or_edit,email
// Common,warehouse,foryoureyesonly,single,35,a31425f31278,0,student,1618594236
// Justin,reception,visible,multi,1002,200,teacher,1618594236
// Common2,ssid,1245678-xx,single,35,a31425f31278,0,student,1618594236,true,7,true,admin@test.com
// ```
func (o *OrgsPsks) ImportOrgPsks(
	ctx context.Context,
	orgId uuid.UUID,
	file *models.FileWrapper) (
	models.ApiResponse[[]models.Psk],
	error) {
	req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/psks/import")
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
	formFields := []https.FormParam{}
	if file != nil {
		fileParam := https.FormParam{Key: "file", Value: *file, Headers: http.Header{}}
		formFields = append(formFields, fileParam)
	}
	req.FormData(formFields)

	var result []models.Psk
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Psk](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgPsk takes context, orgId, pskId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete an organization personal PSK by PSK ID so clients can no longer authenticate with that key.
func (o *OrgsPsks) DeleteOrgPsk(
	ctx context.Context,
	orgId uuid.UUID,
	pskId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/psks/%v")
	req.AppendTemplateParams(orgId, pskId)
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

// GetOrgPsk takes context, orgId, pskId as parameters and
// returns an models.ApiResponse with models.Psk data and
// an error if there was an issue with the request or response.
// Retrieve details for an organization personal PSK, including SSID, usage mode, MAC binding, role, VLAN, expiration, and notification settings.
func (o *OrgsPsks) GetOrgPsk(
	ctx context.Context,
	orgId uuid.UUID,
	pskId uuid.UUID) (
	models.ApiResponse[models.Psk],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/psks/%v")
	req.AppendTemplateParams(orgId, pskId)
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

	var result models.Psk
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Psk](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgPsk takes context, orgId, pskId, body as parameters and
// returns an models.ApiResponse with models.Psk data and
// an error if there was an issue with the request or response.
// Update an organization personal PSK, including passphrase, usage mode, MAC binding, role, VLAN, expiration, and notification settings.
func (o *OrgsPsks) UpdateOrgPsk(
	ctx context.Context,
	orgId uuid.UUID,
	pskId uuid.UUID,
	body *models.Psk) (
	models.ApiResponse[models.Psk],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/psks/%v")
	req.AppendTemplateParams(orgId, pskId)
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

	var result models.Psk
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Psk](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgPskOldPassphrase takes context, orgId, pskId as parameters and
// returns an models.ApiResponse with models.Psk data and
// an error if there was an issue with the request or response.
// Remove the stored `old_passphrase` from a PSK after rotation.
// If successful, the response returns the PSK with `old_passphrase` removed.
func (o *OrgsPsks) DeleteOrgPskOldPassphrase(
	ctx context.Context,
	orgId uuid.UUID,
	pskId uuid.UUID) (
	models.ApiResponse[models.Psk],
	error) {
	req := o.prepareRequest(
		ctx,
		"POST",
		"/api/v1/orgs/%v/psks/%v/delete_old_passphrase",
	)
	req.AppendTemplateParams(orgId, pskId)
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

	var result models.Psk
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Psk](decoder)
	return models.NewApiResponse(result, resp), err
}
