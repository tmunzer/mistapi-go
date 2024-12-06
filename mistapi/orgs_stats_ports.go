package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// OrgsStatsPorts represents a controller struct.
type OrgsStatsPorts struct {
    baseController
}

// NewOrgsStatsPorts creates a new instance of OrgsStatsPorts.
// It takes a baseController as a parameter and returns a pointer to the OrgsStatsPorts.
func NewOrgsStatsPorts(baseController baseController) *OrgsStatsPorts {
    orgsStatsPorts := OrgsStatsPorts{baseController: baseController}
    return &orgsStatsPorts
}

// SearchOrgSwOrGwPorts takes context, orgId, fullDuplex, mac, neighborMac, neighborPortDesc, neighborSystemName, poeDisabled, poeMode, poeOn, portId, portMac, powerDraw, txPkts, rxPkts, rxBytes, txBps, rxBps, txErrors, rxErrors, txMcastPkts, txBcastPkts, rxMcastPkts, rxBcastPkts, speed, macLimit, macCount, up, stpState, stpRole, authState, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponsePortStatsSearch data and
// an error if there was an issue with the request or response.
// Search Switch / Gateway Ports
func (o *OrgsStatsPorts) SearchOrgSwOrGwPorts(
    ctx context.Context,
    orgId uuid.UUID,
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
    stpState *models.SearchOrgSwOrGwPortsStpStateEnum,
    stpRole *models.SearchOrgSwOrGwPortsStpRoleEnum,
    authState *models.SearchOrgSwOrGwPortsAuthStateEnum,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponsePortStatsSearch],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/stats/ports/search", orgId),
    )
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
    if stpState != nil {
        req.QueryParam("stp_state", *stpState)
    }
    if stpRole != nil {
        req.QueryParam("stp_role", *stpRole)
    }
    if authState != nil {
        req.QueryParam("auth_state", *authState)
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
    
    var result models.ResponsePortStatsSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponsePortStatsSearch](decoder)
    return models.NewApiResponse(result, resp), err
}
