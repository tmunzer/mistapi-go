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

// SitesStatsPorts represents a controller struct.
type SitesStatsPorts struct {
	baseController
}

// NewSitesStatsPorts creates a new instance of SitesStatsPorts.
// It takes a baseController as a parameter and returns a pointer to the SitesStatsPorts.
func NewSitesStatsPorts(baseController baseController) *SitesStatsPorts {
	sitesStatsPorts := SitesStatsPorts{baseController: baseController}
	return &sitesStatsPorts
}

// CountSiteSwOrGwPorts takes context, siteId, distinct, fullDuplex, mac, neighborMac, neighborPortDesc, neighborSystemName, poeDisabled, poeMode, poeOn, portId, portMac, powerDraw, txPkts, rxPkts, rxBytes, txBps, rxBps, txMcastPkts, txBcastPkts, rxMcastPkts, rxBcastPkts, speed, stpState, stpRole, authState, up, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count switch and gateway port statistics for a site, optionally grouped by the `distinct` field and filtered by port, neighbor, PoE, STP, traffic, and time attributes. Use [Count Org Switch/Gateway Ports]($e/Orgs%20Stats%20-%20Ports/countOrgSwOrGwPorts) to count switch and gateway port statistics across the organization.
func (s *SitesStatsPorts) CountSiteSwOrGwPorts(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SitePortsCountDistinctEnum,
	fullDuplex *bool,
	mac *string,
	neighborMac *string,
	neighborPortDesc *string,
	neighborSystemName *string,
	poeDisabled *bool,
	poeMode *string,
	poeOn *bool,
	portId *string,
	portMac *string,
	powerDraw *float64,
	txPkts *int,
	rxPkts *int,
	rxBytes *int,
	txBps *int,
	rxBps *int,
	txMcastPkts *int,
	txBcastPkts *int,
	rxMcastPkts *int,
	rxBcastPkts *int,
	speed *int,
	stpState *models.PortStpStateEnum,
	stpRole *models.PortStpRoleEnum,
	authState *models.PortAuthStateEnum,
	up *bool,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/ports/count")
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
	if distinct != nil {
		req.QueryParam("distinct", *distinct)
	}
	if fullDuplex != nil {
		req.QueryParam("full_duplex", *fullDuplex)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if neighborMac != nil {
		req.QueryParam("neighbor_mac", *neighborMac)
	}
	if neighborPortDesc != nil {
		req.QueryParam("neighbor_port_desc", *neighborPortDesc)
	}
	if neighborSystemName != nil {
		req.QueryParam("neighbor_system_name", *neighborSystemName)
	}
	if poeDisabled != nil {
		req.QueryParam("poe_disabled", *poeDisabled)
	}
	if poeMode != nil {
		req.QueryParam("poe_mode", *poeMode)
	}
	if poeOn != nil {
		req.QueryParam("poe_on", *poeOn)
	}
	if portId != nil {
		req.QueryParam("port_id", *portId)
	}
	if portMac != nil {
		req.QueryParam("port_mac", *portMac)
	}
	if powerDraw != nil {
		req.QueryParam("power_draw", *powerDraw)
	}
	if txPkts != nil {
		req.QueryParam("tx_pkts", *txPkts)
	}
	if rxPkts != nil {
		req.QueryParam("rx_pkts", *rxPkts)
	}
	if rxBytes != nil {
		req.QueryParam("rx_bytes", *rxBytes)
	}
	if txBps != nil {
		req.QueryParam("tx_bps", *txBps)
	}
	if rxBps != nil {
		req.QueryParam("rx_bps", *rxBps)
	}
	if txMcastPkts != nil {
		req.QueryParam("tx_mcast_pkts", *txMcastPkts)
	}
	if txBcastPkts != nil {
		req.QueryParam("tx_bcast_pkts", *txBcastPkts)
	}
	if rxMcastPkts != nil {
		req.QueryParam("rx_mcast_pkts", *rxMcastPkts)
	}
	if rxBcastPkts != nil {
		req.QueryParam("rx_bcast_pkts", *rxBcastPkts)
	}
	if speed != nil {
		req.QueryParam("speed", *speed)
	}
	if stpState != nil {
		req.QueryParam("stp_state", *stpState)
	}
	if stpRole != nil {
		req.QueryParam("stp_role", *stpRole)
	}
	if authState != nil {
		req.QueryParam("auth_state", *authState)
	}
	if up != nil {
		req.QueryParam("up", *up)
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

// SearchSiteSwOrGwPorts takes context, siteId, deviceType, authState, fullDuplex, lteImsi, lteIccid, lteImei, mac, neighborMac, neighborPortDesc, neighborSystemName, poeDisabled, poeMode, poeOn, poePriority, portId, portMac, speed, stpState, stpRole, up, xcvrPartNumber, limit, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseSwitchPortSearch data and
// an error if there was an issue with the request or response.
// Search switch and gateway port statistics for a site.
// Returns ports that match the search criteria, including current or most recent port status and statistics within the hour.
// Traffic information (Tx/Rx) is reported as cumulative counters since the last device reboot. Use [Search Org Switch/Gateway Ports]($e/Orgs%20Stats%20-%20Ports/searchOrgSwOrGwPorts) to search switch and gateway port statistics across the organization.
func (s *SitesStatsPorts) SearchSiteSwOrGwPorts(
	ctx context.Context,
	siteId uuid.UUID,
	deviceType *models.SearchOrgSwOrGwPortsTypeEnum,
	authState *models.PortAuthStateEnum,
	fullDuplex *bool,
	lteImsi *string,
	lteIccid *string,
	lteImei *string,
	mac *string,
	neighborMac *string,
	neighborPortDesc *string,
	neighborSystemName *string,
	poeDisabled *bool,
	poeMode *string,
	poeOn *bool,
	poePriority *models.PoePriorityEnum,
	portId *string,
	portMac *string,
	speed *int,
	stpState *models.PortStpStateEnum,
	stpRole *models.PortStpRoleEnum,
	up *bool,
	xcvrPartNumber *string,
	limit *int,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseSwitchPortSearch],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/ports/search")
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
	if deviceType != nil {
		req.QueryParam("device_type", *deviceType)
	}
	if authState != nil {
		req.QueryParam("auth_state", *authState)
	}
	if fullDuplex != nil {
		req.QueryParam("full_duplex", *fullDuplex)
	}
	if lteImsi != nil {
		req.QueryParam("lte_imsi", *lteImsi)
	}
	if lteIccid != nil {
		req.QueryParam("lte_iccid", *lteIccid)
	}
	if lteImei != nil {
		req.QueryParam("lte_imei", *lteImei)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if neighborMac != nil {
		req.QueryParam("neighbor_mac", *neighborMac)
	}
	if neighborPortDesc != nil {
		req.QueryParam("neighbor_port_desc", *neighborPortDesc)
	}
	if neighborSystemName != nil {
		req.QueryParam("neighbor_system_name", *neighborSystemName)
	}
	if poeDisabled != nil {
		req.QueryParam("poe_disabled", *poeDisabled)
	}
	if poeMode != nil {
		req.QueryParam("poe_mode", *poeMode)
	}
	if poeOn != nil {
		req.QueryParam("poe_on", *poeOn)
	}
	if poePriority != nil {
		req.QueryParam("poe_priority", *poePriority)
	}
	if portId != nil {
		req.QueryParam("port_id", *portId)
	}
	if portMac != nil {
		req.QueryParam("port_mac", *portMac)
	}
	if speed != nil {
		req.QueryParam("speed", *speed)
	}
	if stpState != nil {
		req.QueryParam("stp_state", *stpState)
	}
	if stpRole != nil {
		req.QueryParam("stp_role", *stpRole)
	}
	if up != nil {
		req.QueryParam("up", *up)
	}
	if xcvrPartNumber != nil {
		req.QueryParam("xcvr_part_number", *xcvrPartNumber)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.ResponseSwitchPortSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseSwitchPortSearch](decoder)
	return models.NewApiResponse(result, resp), err
}
