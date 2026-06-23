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

// OrgsDevices represents a controller struct.
type OrgsDevices struct {
	baseController
}

// NewOrgsDevices creates a new instance of OrgsDevices.
// It takes a baseController as a parameter and returns a pointer to the OrgsDevices.
func NewOrgsDevices(baseController baseController) *OrgsDevices {
	orgsDevices := OrgsDevices{baseController: baseController}
	return &orgsDevices
}

// ListOrgDevices takes context, orgId as parameters and
// returns an models.ApiResponse with models.ResponseOrgDevices data and
// an error if there was an issue with the request or response.
// List devices in the organization, including APs, switches and gateways managed or monitored by Mist.
func (o *OrgsDevices) ListOrgDevices(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[models.ResponseOrgDevices],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/devices")
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

	var result models.ResponseOrgDevices
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseOrgDevices](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountOrgDevices takes context, orgId, distinct, hostname, siteId, model, managed, mac, version, ip, mxtunnelStatus, mxedgeId, lldpSystemName, lldpSystemDesc, lldpPortId, lldpMgmtAddr, mType, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count organization device records, optionally grouped by `distinct` and filtered by device identifiers, model, LLDP attributes, Mist Edge, tunnel status, device type, and time range.
func (o *OrgsDevices) CountOrgDevices(
	ctx context.Context,
	orgId uuid.UUID,
	distinct *models.OrgDevicesCountDistinctEnum,
	hostname *string,
	siteId *uuid.UUID,
	model *string,
	managed *string,
	mac *string,
	version *string,
	ip *string,
	mxtunnelStatus *models.CountOrgDevicesMxtunnelStatusEnum,
	mxedgeId *uuid.UUID,
	lldpSystemName *string,
	lldpSystemDesc *string,
	lldpPortId *string,
	lldpMgmtAddr *string,
	mType *models.DeviceTypeDefaultApEnum,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/devices/count")
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
	if hostname != nil {
		req.QueryParam("hostname", *hostname)
	}
	if siteId != nil {
		req.QueryParam("site_id", *siteId)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if managed != nil {
		req.QueryParam("managed", *managed)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if version != nil {
		req.QueryParam("version", *version)
	}
	if ip != nil {
		req.QueryParam("ip", *ip)
	}
	if mxtunnelStatus != nil {
		req.QueryParam("mxtunnel_status", *mxtunnelStatus)
	}
	if mxedgeId != nil {
		req.QueryParam("mxedge_id", *mxedgeId)
	}
	if lldpSystemName != nil {
		req.QueryParam("lldp_system_name", *lldpSystemName)
	}
	if lldpSystemDesc != nil {
		req.QueryParam("lldp_system_desc", *lldpSystemDesc)
	}
	if lldpPortId != nil {
		req.QueryParam("lldp_port_id", *lldpPortId)
	}
	if lldpMgmtAddr != nil {
		req.QueryParam("lldp_mgmt_addr", *lldpMgmtAddr)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
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

// CountOrgDeviceEvents takes context, orgId, distinct, siteId, ap, apfw, model, text, mType, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count device event records across the organization, optionally grouped by `distinct` and filtered by site, AP, firmware, model, event text, event type, and time range.
func (o *OrgsDevices) CountOrgDeviceEvents(
	ctx context.Context,
	orgId uuid.UUID,
	distinct *models.OrgDevicesEventsCountDistinctEnum,
	siteId *uuid.UUID,
	ap *string,
	apfw *string,
	model *string,
	text *string,
	mType *string,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/devices/events/count")
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
	if siteId != nil {
		req.QueryParam("site_id", *siteId)
	}
	if ap != nil {
		req.QueryParam("ap", *ap)
	}
	if apfw != nil {
		req.QueryParam("apfw", *apfw)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if text != nil {
		req.QueryParam("text", *text)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
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

// SearchOrgDeviceEvents takes context, orgId, mac, model, deviceType, text, mType, lastBy, includes, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseDeviceEventsSearch data and
// an error if there was an issue with the request or response.
// Search device event records across the organization with filters for MAC address, model, device type, event text, event type, additional event indices, and time range.
func (o *OrgsDevices) SearchOrgDeviceEvents(
	ctx context.Context,
	orgId uuid.UUID,
	mac *string,
	model *string,
	deviceType *string,
	text *string,
	mType *string,
	lastBy *string,
	includes *string,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseDeviceEventsSearch],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/devices/events/search")
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
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if deviceType != nil {
		req.QueryParam("device_type", *deviceType)
	}
	if text != nil {
		req.QueryParam("text", *text)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if lastBy != nil {
		req.QueryParam("last_by", *lastBy)
	}
	if includes != nil {
		req.QueryParam("includes", *includes)
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

	var result models.ResponseDeviceEventsSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseDeviceEventsSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountOrgDeviceLastConfigs takes context, orgId, mType, distinct, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count device config history records across the organization, optionally grouped by `distinct` and filtered by device type and time range.
func (o *OrgsDevices) CountOrgDeviceLastConfigs(
	ctx context.Context,
	orgId uuid.UUID,
	mType *models.DeviceTypeDefaultApEnum,
	distinct *models.OrgDevicesLastConfigsCountDistinctEnum,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/devices/last_config/count")
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

// SearchOrgDeviceLastConfigs takes context, orgId, deviceType, mac, name, version, certExpiryDuration, start, end, limit, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseConfigHistorySearch data and
// an error if there was an issue with the request or response.
// Search device config history records across the organization with filters for device type, MAC address, name, software version, certificate-expiry duration, and time range.
func (o *OrgsDevices) SearchOrgDeviceLastConfigs(
	ctx context.Context,
	orgId uuid.UUID,
	deviceType *models.LastConfigDeviceTypeEnum,
	mac *string,
	name *string,
	version *string,
	certExpiryDuration *string,
	start *string,
	end *string,
	limit *int,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseConfigHistorySearch],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/devices/last_config/search")
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
	if deviceType != nil {
		req.QueryParam("device_type", *deviceType)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if name != nil {
		req.QueryParam("name", *name)
	}
	if version != nil {
		req.QueryParam("version", *version)
	}
	if certExpiryDuration != nil {
		req.QueryParam("cert_expiry_duration", *certExpiryDuration)
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
	if duration != nil {
		req.QueryParam("duration", *duration)
	}
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.ResponseConfigHistorySearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseConfigHistorySearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListOrgApsMacs takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.ApRadioMac data and
// an error if there was an issue with the request or response.
// For some scenarios like E911 or security systems, the BSSIDs are required to identify which AP the client is connecting to. Then the location of the AP can be used as the approximate location of the client.
// Each radio MAC can have up to 16 BSSIDs. These are derived by incrementing the least significant hexadecimal digit (last nibble) of the MAC address from 0 to F, while keeping the remaining bits unchanged.
func (o *OrgsDevices) ListOrgApsMacs(
	ctx context.Context,
	orgId uuid.UUID,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.ApRadioMac],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/devices/radio_macs")
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
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.ApRadioMac
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ApRadioMac](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchOrgDevices takes context, orgId, band24Channel, band5Channel, band6Channel, band24Bandwidth, band5Bandwidth, band6Bandwidth, band24Power, band5Power, band6Power, clustered, eth0PortSpeed, evpntopoId, extIp, hostname, ip, lastConfigStatus, lastHostname, lldpMgmtAddr, lldpPortId, lldpSystemDesc, lldpSystemName, mac, model, mxedgeId, mxedgeIds, mxtunnelStatus, node, node0Mac, node1Mac, powerConstrained, radiusStats, siteId, stats, t128agentVersion, mType, version, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseDeviceSearch data and
// an error if there was an issue with the request or response.
// Search organization devices with filters for AP radio attributes, gateway HA attributes, switch EVPN attributes, LLDP data, MAC address, IP address, model, software version, site, Mist Edge, and time range. Set `stats=true` to include device stats in the response.
func (o *OrgsDevices) SearchOrgDevices(
	ctx context.Context,
	orgId uuid.UUID,
	band24Channel *int,
	band5Channel *int,
	band6Channel *int,
	band24Bandwidth *int,
	band5Bandwidth *int,
	band6Bandwidth *int,
	band24Power *int,
	band5Power *int,
	band6Power *int,
	clustered *bool,
	eth0PortSpeed *int,
	evpntopoId *uuid.UUID,
	extIp *string,
	hostname *string,
	ip *string,
	lastConfigStatus *string,
	lastHostname *string,
	lldpMgmtAddr *string,
	lldpPortId *string,
	lldpSystemDesc *string,
	lldpSystemName *string,
	mac *string,
	model *string,
	mxedgeId *uuid.UUID,
	mxedgeIds *string,
	mxtunnelStatus *models.SearchOrgDevicesMxtunnelStatusEnum,
	node *models.HaClusterNodeEnum,
	node0Mac *string,
	node1Mac *string,
	powerConstrained *bool,
	radiusStats *string,
	siteId *string,
	stats *bool,
	t128agentVersion *string,
	mType *models.DeviceTypeDefaultApEnum,
	version *string,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseDeviceSearch],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/devices/search")
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
	if band24Channel != nil {
		req.QueryParam("band_24_channel", *band24Channel)
	}
	if band5Channel != nil {
		req.QueryParam("band_5_channel", *band5Channel)
	}
	if band6Channel != nil {
		req.QueryParam("band_6_channel", *band6Channel)
	}
	if band24Bandwidth != nil {
		req.QueryParam("band_24_bandwidth", *band24Bandwidth)
	}
	if band5Bandwidth != nil {
		req.QueryParam("band_5_bandwidth", *band5Bandwidth)
	}
	if band6Bandwidth != nil {
		req.QueryParam("band_6_bandwidth", *band6Bandwidth)
	}
	if band24Power != nil {
		req.QueryParam("band_24_power", *band24Power)
	}
	if band5Power != nil {
		req.QueryParam("band_5_power", *band5Power)
	}
	if band6Power != nil {
		req.QueryParam("band_6_power", *band6Power)
	}
	if clustered != nil {
		req.QueryParam("clustered", *clustered)
	}
	if eth0PortSpeed != nil {
		req.QueryParam("eth0_port_speed", *eth0PortSpeed)
	}
	if evpntopoId != nil {
		req.QueryParam("evpntopo_id", *evpntopoId)
	}
	if extIp != nil {
		req.QueryParam("ext_ip", *extIp)
	}
	if hostname != nil {
		req.QueryParam("hostname", *hostname)
	}
	if ip != nil {
		req.QueryParam("ip", *ip)
	}
	if lastConfigStatus != nil {
		req.QueryParam("last_config_status", *lastConfigStatus)
	}
	if lastHostname != nil {
		req.QueryParam("last_hostname", *lastHostname)
	}
	if lldpMgmtAddr != nil {
		req.QueryParam("lldp_mgmt_addr", *lldpMgmtAddr)
	}
	if lldpPortId != nil {
		req.QueryParam("lldp_port_id", *lldpPortId)
	}
	if lldpSystemDesc != nil {
		req.QueryParam("lldp_system_desc", *lldpSystemDesc)
	}
	if lldpSystemName != nil {
		req.QueryParam("lldp_system_name", *lldpSystemName)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if mxedgeId != nil {
		req.QueryParam("mxedge_id", *mxedgeId)
	}
	if mxedgeIds != nil {
		req.QueryParam("mxedge_ids", *mxedgeIds)
	}
	if mxtunnelStatus != nil {
		req.QueryParam("mxtunnel_status", *mxtunnelStatus)
	}
	if node != nil {
		req.QueryParam("node", *node)
	}
	if node0Mac != nil {
		req.QueryParam("node0_mac", *node0Mac)
	}
	if node1Mac != nil {
		req.QueryParam("node1_mac", *node1Mac)
	}
	if powerConstrained != nil {
		req.QueryParam("power_constrained", *powerConstrained)
	}
	if radiusStats != nil {
		req.QueryParam("radius_stats", *radiusStats)
	}
	if siteId != nil {
		req.QueryParam("site_id", *siteId)
	}
	if stats != nil {
		req.QueryParam("stats", *stats)
	}
	if t128agentVersion != nil {
		req.QueryParam("t128agent_version", *t128agentVersion)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if version != nil {
		req.QueryParam("version", *version)
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

	var result models.ResponseDeviceSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseDeviceSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// ListOrgDevicesSummary takes context, orgId as parameters and
// returns an models.ApiResponse with models.ResponseOrgDevicesSummary data and
// an error if there was an issue with the request or response.
// Return aggregate organization device counts by device category and assignment state.
func (o *OrgsDevices) ListOrgDevicesSummary(
	ctx context.Context,
	orgId uuid.UUID) (
	models.ApiResponse[models.ResponseOrgDevicesSummary],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/devices/summary")
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

	var result models.ResponseOrgDevicesSummary
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseOrgDevicesSummary](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetOrgJuniperDevicesCommand takes context, orgId, siteId as parameters and
// returns an models.ApiResponse with models.ResponseDeviceConfigCmd data and
// an error if there was an issue with the request or response.
// Get Org Juniper Devices command
// Juniper devices can be managed/adopted by Mist. Currently outbound-ssh + netconf is used.
// A few lines of CLI commands are generated per-Org, allowing the Juniper devices to phone home to Mist.
func (o *OrgsDevices) GetOrgJuniperDevicesCommand(
	ctx context.Context,
	orgId uuid.UUID,
	siteId *string) (
	models.ApiResponse[models.ResponseDeviceConfigCmd],
	error) {
	req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/ocdevices/outbound_ssh_cmd")
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
	if siteId != nil {
		req.QueryParam("site_id", *siteId)
	}

	var result models.ResponseDeviceConfigCmd
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseDeviceConfigCmd](decoder)
	return models.NewApiResponse(result, resp), err
}
