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

// OrgsClientsMarvis represents a controller struct.
type OrgsClientsMarvis struct {
	baseController
}

// NewOrgsClientsMarvis creates a new instance of OrgsClientsMarvis.
// It takes a baseController as a parameter and returns a pointer to the OrgsClientsMarvis.
func NewOrgsClientsMarvis(baseController baseController) *OrgsClientsMarvis {
	orgsClientsMarvis := OrgsClientsMarvis{baseController: baseController}
	return &orgsClientsMarvis
}

// GetOrgMarvisClientInsights takes context, orgId, marvisclientId, duration, interval, start, end, limit, page as parameters and
// returns an models.ApiResponse with models.MarvisClientInsights data and
// an error if there was an issue with the request or response.
// Return time-series metrics for a specific Marvis Client device. For the full list of supported metric field names and example values, refer to [List Insight Metrics]($e/Constants%20Definitions/listInsightMetrics) under `/api/v1/const/insight_metrics`.
func (o *OrgsClientsMarvis) GetOrgMarvisClientInsights(
	ctx context.Context,
	orgId uuid.UUID,
	marvisclientId uuid.UUID,
	duration *string,
	interval *string,
	start *string,
	end *string,
	limit *int,
	page *int) (
	models.ApiResponse[models.MarvisClientInsights],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		"/api/v1/orgs/%v/insights/marvisclient/%v/marvisclient-metrics",
	)
	req.AppendTemplateParams(orgId, marvisclientId)
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
	if duration != nil {
		req.QueryParam("duration", *duration)
	}
	if interval != nil {
		req.QueryParam("interval", *interval)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result models.MarvisClientInsights
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MarvisClientInsights](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountOrgMarvisClientEvents takes context, orgId, distinct, mType, deviceId, wifiMac, wifiIp, hostname, ssid, bssid, channel, preBssid, preChannel, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count Marvis Client events by a distinct field.
func (o *OrgsClientsMarvis) CountOrgMarvisClientEvents(
	ctx context.Context,
	orgId uuid.UUID,
	distinct *string,
	mType *string,
	deviceId *string,
	wifiMac *string,
	wifiIp *string,
	hostname *string,
	ssid *string,
	bssid *string,
	channel *string,
	preBssid *string,
	preChannel *string,
	limit *int,
	start *string,
	end *string,
	duration *string) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/marvisclients/events/count")
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
	if mType != nil {
		req.QueryParam("type", *mType)
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
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if bssid != nil {
		req.QueryParam("bssid", *bssid)
	}
	if channel != nil {
		req.QueryParam("channel", *channel)
	}
	if preBssid != nil {
		req.QueryParam("pre_bssid", *preBssid)
	}
	if preChannel != nil {
		req.QueryParam("pre_channel", *preChannel)
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

// SearchOrgMarvisClientEvents takes context, orgId, mType, deviceId, wifiMac, wifiIp, hostname, ssid, bssid, channel, preBssid, preChannel, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.MarvisClientEventsSearch data and
// an error if there was an issue with the request or response.
// Search Marvis Client events across the organization.
func (o *OrgsClientsMarvis) SearchOrgMarvisClientEvents(
	ctx context.Context,
	orgId uuid.UUID,
	mType *string,
	deviceId *string,
	wifiMac *string,
	wifiIp *string,
	hostname *string,
	ssid *string,
	bssid *string,
	channel *string,
	preBssid *string,
	preChannel *string,
	limit *int,
	start *string,
	end *string,
	duration *string) (
	models.ApiResponse[models.MarvisClientEventsSearch],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		"/api/v1/orgs/%v/marvisclients/events/search",
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
	if mType != nil {
		req.QueryParam("type", *mType)
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
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if bssid != nil {
		req.QueryParam("bssid", *bssid)
	}
	if channel != nil {
		req.QueryParam("channel", *channel)
	}
	if preBssid != nil {
		req.QueryParam("pre_bssid", *preBssid)
	}
	if preChannel != nil {
		req.QueryParam("pre_channel", *preChannel)
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

	var result models.MarvisClientEventsSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MarvisClientEventsSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteOrgMarvisClient takes context, orgId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Marvis Client
func (o *OrgsClientsMarvis) DeleteOrgMarvisClient(
	ctx context.Context,
	orgId uuid.UUID) (
	*http.Response,
	error) {
	req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/stats/marvisclients")
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
