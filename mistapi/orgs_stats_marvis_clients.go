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

// OrgsStatsMarvisClients represents a controller struct.
type OrgsStatsMarvisClients struct {
	baseController
}

// NewOrgsStatsMarvisClients creates a new instance of OrgsStatsMarvisClients.
// It takes a baseController as a parameter and returns a pointer to the OrgsStatsMarvisClients.
func NewOrgsStatsMarvisClients(baseController baseController) *OrgsStatsMarvisClients {
	orgsStatsMarvisClients := OrgsStatsMarvisClients{baseController: baseController}
	return &orgsStatsMarvisClients
}

// CountOrgMarvisClientsStats takes context, orgId, distinct, deviceId, wifiMac, wifiIp, hostname, model, mfg, serial, osType, osVersion, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count Marvis Client stats records by a distinct field.
func (o *OrgsStatsMarvisClients) CountOrgMarvisClientsStats(
	ctx context.Context,
	orgId uuid.UUID,
	distinct *string,
	deviceId *string,
	wifiMac *string,
	wifiIp *string,
	hostname *string,
	model *string,
	mfg *string,
	serial *string,
	osType *string,
	osVersion *string,
	limit *int,
	start *string,
	end *string,
	duration *string) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/marvisclients/count")
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
	if deviceId != nil {
		req.QueryParam("device_id", *deviceId)
	}
	if wifiMac != nil {
		req.QueryParam("wifi_mac", *wifiMac)
	}
	if wifiIp != nil {
		req.QueryParam("wifi_ip", *wifiIp)
	}
	if hostname != nil {
		req.QueryParam("hostname", *hostname)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if mfg != nil {
		req.QueryParam("mfg", *mfg)
	}
	if serial != nil {
		req.QueryParam("serial", *serial)
	}
	if osType != nil {
		req.QueryParam("os_type", *osType)
	}
	if osVersion != nil {
		req.QueryParam("os_version", *osVersion)
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

	var result models.ResponseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchOrgMarvisClientsStats takes context, orgId, deviceId, wifiMac, wifiIp, hostname, model, mfg, serial, osType, osVersion, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.StatsMarvisClientsSearch data and
// an error if there was an issue with the request or response.
// Search Marvis Client stats records across the organization.
func (o *OrgsStatsMarvisClients) SearchOrgMarvisClientsStats(
	ctx context.Context,
	orgId uuid.UUID,
	deviceId *string,
	wifiMac *string,
	wifiIp *string,
	hostname *string,
	model *string,
	mfg *string,
	serial *string,
	osType *string,
	osVersion *string,
	limit *int,
	start *string,
	end *string,
	duration *string) (
	models.ApiResponse[models.StatsMarvisClientsSearch],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/marvisclients/search")
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
	if deviceId != nil {
		req.QueryParam("device_id", *deviceId)
	}
	if wifiMac != nil {
		req.QueryParam("wifi_mac", *wifiMac)
	}
	if wifiIp != nil {
		req.QueryParam("wifi_ip", *wifiIp)
	}
	if hostname != nil {
		req.QueryParam("hostname", *hostname)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if mfg != nil {
		req.QueryParam("mfg", *mfg)
	}
	if serial != nil {
		req.QueryParam("serial", *serial)
	}
	if osType != nil {
		req.QueryParam("os_type", *osType)
	}
	if osVersion != nil {
		req.QueryParam("os_version", *osVersion)
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

	var result models.StatsMarvisClientsSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.StatsMarvisClientsSearch](decoder)
	return models.NewApiResponse(result, resp), err
}
