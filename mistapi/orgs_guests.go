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

// OrgsGuests represents a controller struct.
type OrgsGuests struct {
	baseController
}

// NewOrgsGuests creates a new instance of OrgsGuests.
// It takes a baseController as a parameter and returns a pointer to the OrgsGuests.
func NewOrgsGuests(baseController baseController) *OrgsGuests {
	orgsGuests := OrgsGuests{baseController: baseController}
	return &orgsGuests
}

// ListOrgGuestAuthorizations takes context, orgId as parameters and
// returns an models.ApiResponse with []models.Guest data and
// an error if there was an issue with the request or response.
// List guest authorization records across the organization, including WLAN, SSID, authentication method, expiration, and guest identity details.
func (o *OrgsGuests) ListOrgGuestAuthorizations(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[[]models.Guest],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/guests")
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

	var result []models.Guest
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Guest](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountOrgGuestAuthorizations takes context, orgId, distinct, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count organization guest authorization records, optionally grouped by `distinct` and filtered by time range.
func (o *OrgsGuests) CountOrgGuestAuthorizations(
	ctx context.Context,
	orgId uuid.UUID,
	distinct *models.OrgGuestsCountDistinctEnum,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/guests/count")
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
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
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

// SearchOrgGuestAuthorization takes context, orgId, wlanId, authMethod, ssid, limit, start, end, duration, sort as parameters and
// returns an models.ApiResponse with models.ResponseGuestSearch data and
// an error if there was an issue with the request or response.
// Search organization guest authorization records with filters for WLAN, SSID, authentication method, and time range.
func (o *OrgsGuests) SearchOrgGuestAuthorization(
	ctx context.Context,
	orgId uuid.UUID,
	wlanId *string,
	authMethod *string,
	ssid *string,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string) (
	models.ApiResponse[models.ResponseGuestSearch],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/guests/search")
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
	if wlanId != nil {
		req.QueryParam("wlan_id", *wlanId)
	}
	if authMethod != nil {
		req.QueryParam("auth_method", *authMethod)
	}
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
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

	var result models.ResponseGuestSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseGuestSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgGuestAuthorization takes context, orgId, guestMac as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete a guest authorization record by guest MAC address.
func (o *OrgsGuests) DeleteOrgGuestAuthorization(
	ctx context.Context,
	orgId uuid.UUID,
	guestMac string) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/guests/%v")
	req.AppendTemplateParams(orgId, guestMac)
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

// GetOrgGuestAuthorization takes context, orgId, guestMac as parameters and
// returns an models.ApiResponse with models.Guest data and
// an error if there was an issue with the request or response.
// Retrieve the guest authorization record associated with a guest MAC address.
func (o *OrgsGuests) GetOrgGuestAuthorization(
	ctx context.Context,
	orgId uuid.UUID,
	guestMac string) (
	models.ApiResponse[models.Guest],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/guests/%v")
	req.AppendTemplateParams(orgId, guestMac)
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

	var result models.Guest
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Guest](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateOrgGuestAuthorization takes context, orgId, guestMac, body as parameters and
// returns an models.ApiResponse with models.Guest data and
// an error if there was an issue with the request or response.
// Update the organization guest authorization record for a guest MAC address, including authorization state, duration, WLAN, and guest identity fields.
func (o *OrgsGuests) UpdateOrgGuestAuthorization(
	ctx context.Context,
	orgId uuid.UUID,
	guestMac string,
	body *models.GuestOrg) (
	models.ApiResponse[models.Guest],
	error) {
	req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/guests/%v")
	req.AppendTemplateParams(orgId, guestMac)
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

	var result models.Guest
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.Guest](decoder)
	return models.NewApiResponse(result, resp), err
}
