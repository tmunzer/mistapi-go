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

// OrgsClientsWireless represents a controller struct.
type OrgsClientsWireless struct {
	baseController
}

// NewOrgsClientsWireless creates a new instance of OrgsClientsWireless.
// It takes a baseController as a parameter and returns a pointer to the OrgsClientsWireless.
func NewOrgsClientsWireless(baseController baseController) *OrgsClientsWireless {
	orgsClientsWireless := OrgsClientsWireless{baseController: baseController}
	return &orgsClientsWireless
}

// CountOrgWirelessClients takes context, orgId, distinct, mac, hostname, device, os, model, ap, vlan, ssid, ip, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Org Wireless Clients
func (o *OrgsClientsWireless) CountOrgWirelessClients(
	ctx context.Context,
	orgId uuid.UUID,
	distinct *models.OrgClientsCountDistinctEnum,
	mac *string,
	hostname *string,
	device *string,
	os *string,
	model *string,
	ap *string,
	vlan *string,
	ssid *string,
	ip *string,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/clients/count")
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
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if hostname != nil {
		req.QueryParam("hostname", *hostname)
	}
	if device != nil {
		req.QueryParam("device", *device)
	}
	if os != nil {
		req.QueryParam("os", *os)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if ap != nil {
		req.QueryParam("ap", *ap)
	}
	if vlan != nil {
		req.QueryParam("vlan", *vlan)
	}
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if ip != nil {
		req.QueryParam("ip", *ip)
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

// CountOrgWirelessClientEvents takes context, orgId, distinct, mType, reasonCode, ssid, ap, proto, band, wlanId, siteId, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Client-Events
func (o *OrgsClientsWireless) CountOrgWirelessClientEvents(
	ctx context.Context,
	orgId uuid.UUID,
	distinct *models.SiteClientEventsCountDistinctEnum,
	mType *string,
	reasonCode *int,
	ssid *string,
	ap *string,
	proto *models.Dot11ProtoEnum,
	band *models.Dot11BandEnum,
	wlanId *string,
	siteId *uuid.UUID,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/clients/events/count")
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
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if reasonCode != nil {
		req.QueryParam("reason_code", *reasonCode)
	}
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if ap != nil {
		req.QueryParam("ap", *ap)
	}
	if proto != nil {
		req.QueryParam("proto", *proto)
	}
	if band != nil {
		req.QueryParam("band", *band)
	}
	if wlanId != nil {
		req.QueryParam("wlan_id", *wlanId)
	}
	if siteId != nil {
		req.QueryParam("site_id", *siteId)
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

// SearchOrgWirelessClientEvents takes context, orgId, mType, reasonCode, ssid, ap, keyMgmt, proto, band, wlanId, nacruleId, start, end, duration, sort, limit, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseEventsSearch data and
// an error if there was an issue with the request or response.
// Get Org Clients Events
func (o *OrgsClientsWireless) SearchOrgWirelessClientEvents(
	ctx context.Context,
	orgId uuid.UUID,
	mType *string,
	reasonCode *int,
	ssid *string,
	ap *string,
	keyMgmt *models.ClientKeyMgmtEnum,
	proto *models.Dot11ProtoEnum,
	band *models.Dot11BandEnum,
	wlanId *uuid.UUID,
	nacruleId *uuid.UUID,
	start *string,
	end *string,
	duration *string,
	sort *string,
	limit *int,
	searchAfter *string) (
	models.ApiResponse[models.ResponseEventsSearch],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/clients/events/search")
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
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if reasonCode != nil {
		req.QueryParam("reason_code", *reasonCode)
	}
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if ap != nil {
		req.QueryParam("ap", *ap)
	}
	if keyMgmt != nil {
		req.QueryParam("key_mgmt", *keyMgmt)
	}
	if proto != nil {
		req.QueryParam("proto", *proto)
	}
	if band != nil {
		req.QueryParam("band", *band)
	}
	if wlanId != nil {
		req.QueryParam("wlan_id", *wlanId)
	}
	if nacruleId != nil {
		req.QueryParam("nacrule_id", *nacruleId)
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
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.ResponseEventsSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseEventsSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchOrgWirelessClients takes context, orgId, siteId, mac, ip, hostname, band, device, os, model, ap, pskId, pskName, username, vlan, ssid, text, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseClientSearch data and
// an error if there was an issue with the request or response.
// Search Org Wireless Clients
func (o *OrgsClientsWireless) SearchOrgWirelessClients(
	ctx context.Context,
	orgId uuid.UUID,
	siteId *uuid.UUID,
	mac *string,
	ip *string,
	hostname *string,
	band *string,
	device *string,
	os *string,
	model *string,
	ap *string,
	pskId *string,
	pskName *string,
	username *string,
	vlan *string,
	ssid *string,
	text *string,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseClientSearch],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/clients/search")
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
	if siteId != nil {
		req.QueryParam("site_id", *siteId)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if ip != nil {
		req.QueryParam("ip", *ip)
	}
	if hostname != nil {
		req.QueryParam("hostname", *hostname)
	}
	if band != nil {
		req.QueryParam("band", *band)
	}
	if device != nil {
		req.QueryParam("device", *device)
	}
	if os != nil {
		req.QueryParam("os", *os)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if ap != nil {
		req.QueryParam("ap", *ap)
	}
	if pskId != nil {
		req.QueryParam("psk_id", *pskId)
	}
	if pskName != nil {
		req.QueryParam("psk_name", *pskName)
	}
	if username != nil {
		req.QueryParam("username", *username)
	}
	if vlan != nil {
		req.QueryParam("vlan", *vlan)
	}
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if text != nil {
		req.QueryParam("text", *text)
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
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.ResponseClientSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseClientSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountOrgWirelessClientsSessions takes context, orgId, distinct, ap, band, clientFamily, clientManufacture, clientModel, clientOs, ssid, wlanId, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Org Wireless Clients Sessions
func (o *OrgsClientsWireless) CountOrgWirelessClientsSessions(
	ctx context.Context,
	orgId uuid.UUID,
	distinct *models.OrgClientSessionsCountDistinctEnum,
	ap *string,
	band *models.Dot11BandEnum,
	clientFamily *string,
	clientManufacture *string,
	clientModel *string,
	clientOs *string,
	ssid *string,
	wlanId *uuid.UUID,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/clients/sessions/count")
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
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}
	if ap != nil {
		req.QueryParam("ap", *ap)
	}
	if band != nil {
		req.QueryParam("band", *band)
	}
	if clientFamily != nil {
		req.QueryParam("client_family", *clientFamily)
	}
	if clientManufacture != nil {
		req.QueryParam("client_manufacture", *clientManufacture)
	}
	if clientModel != nil {
		req.QueryParam("client_model", *clientModel)
	}
	if clientOs != nil {
		req.QueryParam("client_os", *clientOs)
	}
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if wlanId != nil {
		req.QueryParam("wlan_id", *wlanId)
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

// SearchOrgWirelessClientSessions takes context, orgId, ap, band, clientFamily, clientManufacture, clientModel, clientUsername, clientOs, ssid, wlanId, pskId, pskName, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.SearchWirelessClientSession data and
// an error if there was an issue with the request or response.
// Search Org Wireless Clients Sessions
func (o *OrgsClientsWireless) SearchOrgWirelessClientSessions(
	ctx context.Context,
	orgId uuid.UUID,
	ap *string,
	band *models.Dot11BandEnum,
	clientFamily *string,
	clientManufacture *string,
	clientModel *string,
	clientUsername *string,
	clientOs *string,
	ssid *string,
	wlanId *uuid.UUID,
	pskId *string,
	pskName *string,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.SearchWirelessClientSession],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/clients/sessions/search")
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
	if ap != nil {
		req.QueryParam("ap", *ap)
	}
	if band != nil {
		req.QueryParam("band", *band)
	}
	if clientFamily != nil {
		req.QueryParam("client_family", *clientFamily)
	}
	if clientManufacture != nil {
		req.QueryParam("client_manufacture", *clientManufacture)
	}
	if clientModel != nil {
		req.QueryParam("client_model", *clientModel)
	}
	if clientUsername != nil {
		req.QueryParam("client_username", *clientUsername)
	}
	if clientOs != nil {
		req.QueryParam("client_os", *clientOs)
	}
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if wlanId != nil {
		req.QueryParam("wlan_id", *wlanId)
	}
	if pskId != nil {
		req.QueryParam("psk_id", *pskId)
	}
	if pskName != nil {
		req.QueryParam("psk_name", *pskName)
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
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.SearchWirelessClientSession
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.SearchWirelessClientSession](decoder)
	return models.NewApiResponse(result, resp), err
}
