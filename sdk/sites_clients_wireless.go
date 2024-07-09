package mistapi

import (
	"context"
	"fmt"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi/sdk/errors"
	"github.com/tmunzer/mistapi/sdk/models"
)

// SitesClientsWireless represents a controller struct.
type SitesClientsWireless struct {
	baseController
}

// NewSitesClientsWireless creates a new instance of SitesClientsWireless.
// It takes a baseController as a parameter and returns a pointer to the SitesClientsWireless.
func NewSitesClientsWireless(baseController baseController) *SitesClientsWireless {
	sitesClientsWireless := SitesClientsWireless{baseController: baseController}
	return &sitesClientsWireless
}

// CountSiteWirelessClients takes context, siteId, distinct, ssid, ap, ipAddress, vlan, hostname, os, model, device, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Clients
func (s *SitesClientsWireless) CountSiteWirelessClients(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteClientsCountDistinctEnum,
	ssid *string,
	ap *string,
	ipAddress *string,
	vlan *string,
	hostname *string,
	os *string,
	model *string,
	device *string,
	page *int,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/clients/count", siteId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if ap != nil {
		req.QueryParam("ap", *ap)
	}
	if ipAddress != nil {
		req.QueryParam("ip_address", *ipAddress)
	}
	if vlan != nil {
		req.QueryParam("vlan", *vlan)
	}
	if hostname != nil {
		req.QueryParam("hostname", *hostname)
	}
	if os != nil {
		req.QueryParam("os", *os)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if device != nil {
		req.QueryParam("device", *device)
	}
	if page != nil {
		req.QueryParam("page", *page)
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

	var result models.RepsonseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteWirelessClientEvents takes context, siteId, distinct, mType, reasonCode, ssid, ap, proto, band, wlanId, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Client-Events
func (s *SitesClientsWireless) CountSiteWirelessClientEvents(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteClientEventsCountDistinctEnum,
	mType *string,
	reasonCode *int,
	ssid *string,
	ap *string,
	proto *models.Dot11ProtoEnum,
	band *models.Dot11BandEnum,
	wlanId *string,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/clients/events/count", siteId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
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

	var result models.RepsonseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteWirelessClientEvents takes context, siteId, mType, reasonCode, ssid, ap, proto, band, wlanId, nacruleId, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseEventsSearch data and
// an error if there was an issue with the request or response.
// Get Site Clients Events
func (s *SitesClientsWireless) SearchSiteWirelessClientEvents(
	ctx context.Context,
	siteId uuid.UUID,
	mType *string,
	reasonCode *int,
	ssid *string,
	ap *string,
	proto *models.Dot11ProtoEnum,
	band *models.Dot11BandEnum,
	wlanId *string,
	nacruleId *string,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.ResponseEventsSearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/clients/events/search", siteId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
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

	var result models.ResponseEventsSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseEventsSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteWirelessClients takes context, siteId, mac, ipAddress, hostname, device, os, model, ap, ssid, text, nacruleId, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseClientSearch data and
// an error if there was an issue with the request or response.
// Search Wireless Clients
// **NOTE**: fuzzy logic can be used with ‘*’, supported filters: mac, hostname, device, os, model. E.g. /clients/search?device=Mac*&hostname=jerry
func (s *SitesClientsWireless) SearchSiteWirelessClients(
	ctx context.Context,
	siteId uuid.UUID,
	mac *string,
	ipAddress *string,
	hostname *string,
	device *string,
	os *string,
	model *string,
	ap *string,
	ssid *string,
	text *string,
	nacruleId *string,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.ResponseClientSearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/clients/search", siteId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if ipAddress != nil {
		req.QueryParam("ip_address", *ipAddress)
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
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if text != nil {
		req.QueryParam("text", *text)
	}
	if nacruleId != nil {
		req.QueryParam("nacrule_id", *nacruleId)
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

	var result models.ResponseClientSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseClientSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteWirelessClientSessions takes context, siteId, distinct, ap, band, clientFamily, clientManufacture, clientModel, clientOs, ssid, wlanId, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of Client Sessions
func (s *SitesClientsWireless) CountSiteWirelessClientSessions(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteClientSessionsCountDistinctEnum,
	ap *string,
	band *models.Dot11BandEnum,
	clientFamily *string,
	clientManufacture *string,
	clientModel *string,
	clientOs *string,
	ssid *string,
	wlanId *string,
	page *int,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/clients/sessions/count", siteId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
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
	if page != nil {
		req.QueryParam("page", *page)
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

	var result models.RepsonseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteWirelessClientSessions takes context, siteId, ap, band, clientFamily, clientManufacture, clientModel, clientUsername, clientOs, ssid, wlanId, pskId, pskName, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseClientSessionsSearch data and
// an error if there was an issue with the request or response.
// Search Client Sessions
func (s *SitesClientsWireless) SearchSiteWirelessClientSessions(
	ctx context.Context,
	siteId uuid.UUID,
	ap *string,
	band *models.Dot11BandEnum,
	clientFamily *string,
	clientManufacture *string,
	clientModel *string,
	clientUsername *string,
	clientOs *string,
	ssid *string,
	wlanId *string,
	pskId *uuid.UUID,
	pskName *string,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.ResponseClientSessionsSearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/clients/sessions/search", siteId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
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

	var result models.ResponseClientSessionsSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseClientSessionsSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteEventsForClient takes context, siteId, clientMac, mType, proto, band, channel, wlanId, ssid, start, end, page, limit, duration as parameters and
// returns an models.ApiResponse with models.ResponseClientEventsSearch data and
// an error if there was an issue with the request or response.
// Get the list of events for a specific client
func (s *SitesClientsWireless) GetSiteEventsForClient(
	ctx context.Context,
	siteId uuid.UUID,
	clientMac string,
	mType *string,
	proto *models.Dot11ProtoEnum,
	band *models.Dot11BandEnum,
	channel *string,
	wlanId *string,
	ssid *string,
	start *int,
	end *int,
	page *int,
	limit *int,
	duration *string) (
	models.ApiResponse[models.ResponseClientEventsSearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/clients/%v/events", siteId, clientMac),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if proto != nil {
		req.QueryParam("proto", *proto)
	}
	if band != nil {
		req.QueryParam("band", *band)
	}
	if channel != nil {
		req.QueryParam("channel", *channel)
	}
	if wlanId != nil {
		req.QueryParam("wlan_id", *wlanId)
	}
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if start != nil {
		req.QueryParam("start", *start)
	}
	if end != nil {
		req.QueryParam("end", *end)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if duration != nil {
		req.QueryParam("duration", *duration)
	}

	var result models.ResponseClientEventsSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseClientEventsSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSiteWirelessClientsStats takes context, siteId, wired, limit, start, end, duration as parameters and
// returns an models.ApiResponse with []models.StatsClientAnyOf data and
// an error if there was an issue with the request or response.
// Get List of Site All Clients Stats Details
func (s *SitesClientsWireless) ListSiteWirelessClientsStats(
	ctx context.Context,
	siteId uuid.UUID,
	wired *bool,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[[]models.StatsClientAnyOf],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/clients", siteId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	if wired != nil {
		req.QueryParam("wired", *wired)
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

	var result []models.StatsClientAnyOf
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.StatsClientAnyOf](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteWirelessClientStats takes context, siteId, clientMac, wired as parameters and
// returns an models.ApiResponse with []models.StatsClientAnyOf data and
// an error if there was an issue with the request or response.
// Get Site Client Stats Details
func (s *SitesClientsWireless) GetSiteWirelessClientStats(
	ctx context.Context,
	siteId uuid.UUID,
	clientMac string,
	wired *bool) (
	models.ApiResponse[[]models.StatsClientAnyOf],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/clients/%v", siteId, clientMac),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	if wired != nil {
		req.QueryParam("wired", *wired)
	}

	var result []models.StatsClientAnyOf
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.StatsClientAnyOf](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteWirelessClientsStatsByMap takes context, siteId, mapId, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with [][]models.ClientWirelessStats data and
// an error if there was an issue with the request or response.
// Get Site Clients Stats By Map
func (s *SitesClientsWireless) GetSiteWirelessClientsStatsByMap(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID,
	page *int,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[[][]models.ClientWirelessStats],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/maps/%v/clients", siteId, mapId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})
	if page != nil {
		req.QueryParam("page", *page)
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

	var result [][]models.ClientWirelessStats
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[][]models.ClientWirelessStats](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListSiteUnconnectedClientStats takes context, siteId, mapId as parameters and
// returns an models.ApiResponse with []models.UnconnectedClientStat data and
// an error if there was an issue with the request or response.
// Get List of Site Unconnected Client Location
func (s *SitesClientsWireless) ListSiteUnconnectedClientStats(
	ctx context.Context,
	siteId uuid.UUID,
	mapId uuid.UUID) (
	models.ApiResponse[[]models.UnconnectedClientStat],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/stats/maps/%v/unconnected_clients", siteId, mapId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
	})

	var result []models.UnconnectedClientStat
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.UnconnectedClientStat](decoder)
	return models.NewApiResponse(result, resp), err
}