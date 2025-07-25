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
// Count by Distinct Attributes of Switch/Gateway Ports
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
    stpState *models.CountPortsStpStateEnum,
    stpRole *models.CountPortsStpRoleEnum,
    authState *models.CountPortsAuthStateEnum,
    up *bool,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/ports/count")
    req.AppendTemplateParams(siteId)
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

// SearchSiteSwOrGwPorts takes context, siteId, fullDuplex, disabled, mac, deviceType, neighborMac, neighborPortDesc, neighborSystemName, poeDisabled, poeMode, poeOn, portId, portMac, powerDraw, txPkts, rxPkts, rxBytes, txBps, rxBps, txErrors, rxErrors, txMcastPkts, txBcastPkts, rxMcastPkts, rxBcastPkts, speed, macLimit, macCount, up, active, jitter, loss, latency, stpState, stpRole, xcvrPartNumber, authState, lteImsi, lteIccid, lteImei, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseSwitchPortSearch data and
// an error if there was an issue with the request or response.
// Search Switch / Gateway Ports
func (s *SitesStatsPorts) SearchSiteSwOrGwPorts(
    ctx context.Context,
    siteId uuid.UUID,
    fullDuplex *bool,
    disabled *bool,
    mac *string,
    deviceType *models.SearchSiteSwOrGwPortsDeviceTypeEnum,
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
    txErrors *int,
    rxErrors *int,
    txMcastPkts *int,
    txBcastPkts *int,
    rxMcastPkts *int,
    rxBcastPkts *int,
    speed *int,
    macLimit *int,
    macCount *int,
    up *bool,
    active *bool,
    jitter *float64,
    loss *float64,
    latency *float64,
    stpState *models.SearchSiteSwOrGwPortsStpStateEnum,
    stpRole *models.SearchSiteSwOrGwPortsStpRoleEnum,
    xcvrPartNumber *string,
    authState *models.SearchSiteSwOrGwPortsAuthStateEnum,
    lteImsi *string,
    lteIccid *string,
    lteImei *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseSwitchPortSearch],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/stats/ports/search")
    req.AppendTemplateParams(siteId)
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
    if fullDuplex != nil {
        req.QueryParam("full_duplex", *fullDuplex)
    }
    if disabled != nil {
        req.QueryParam("disabled", *disabled)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if deviceType != nil {
        req.QueryParam("device_type", *deviceType)
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
    if txErrors != nil {
        req.QueryParam("tx_errors", *txErrors)
    }
    if rxErrors != nil {
        req.QueryParam("rx_errors", *rxErrors)
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
    if macLimit != nil {
        req.QueryParam("mac_limit", *macLimit)
    }
    if macCount != nil {
        req.QueryParam("mac_count", *macCount)
    }
    if up != nil {
        req.QueryParam("up", *up)
    }
    if active != nil {
        req.QueryParam("active", *active)
    }
    if jitter != nil {
        req.QueryParam("jitter", *jitter)
    }
    if loss != nil {
        req.QueryParam("loss", *loss)
    }
    if latency != nil {
        req.QueryParam("latency", *latency)
    }
    if stpState != nil {
        req.QueryParam("stp_state", *stpState)
    }
    if stpRole != nil {
        req.QueryParam("stp_role", *stpRole)
    }
    if xcvrPartNumber != nil {
        req.QueryParam("xcvr_part_number", *xcvrPartNumber)
    }
    if authState != nil {
        req.QueryParam("auth_state", *authState)
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
    
    var result models.ResponseSwitchPortSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseSwitchPortSearch](decoder)
    return models.NewApiResponse(result, resp), err
}
