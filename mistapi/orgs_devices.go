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
// Get List of Org Devices
func (o *OrgsDevices) ListOrgDevices(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponseOrgDevices],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/devices", orgId),
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    
    var result models.ResponseOrgDevices
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseOrgDevices](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountOrgDevices takes context, orgId, distinct, hostname, siteId, model, mac, version, ipAddress, mxtunnelStatus, mxedgeId, lldpSystemName, lldpSystemDesc, lldpPortId, lldpMgmtAddr, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count Org Devices
func (o *OrgsDevices) CountOrgDevices(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgDevicesCountDistinctEnum,
    hostname *string,
    siteId *string,
    model *string,
    mac *string,
    version *string,
    ipAddress *string,
    mxtunnelStatus *models.CountOrgDevicesMxtunnelStatusEnum,
    mxedgeId *string,
    lldpSystemName *string,
    lldpSystemDesc *string,
    lldpPortId *string,
    lldpMgmtAddr *string,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/devices/count", orgId),
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
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
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if version != nil {
        req.QueryParam("version", *version)
    }
    if ipAddress != nil {
        req.QueryParam("ip_address", *ipAddress)
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

// CountOrgDeviceEvents takes context, orgId, distinct, siteId, ap, apfw, model, text, timestamp, mType, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count Org Devices Events
func (o *OrgsDevices) CountOrgDeviceEvents(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgDevicesEventsCountDistinctEnum,
    siteId *string,
    ap *string,
    apfw *string,
    model *string,
    text *string,
    timestamp *string,
    mType *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/devices/events/count", orgId),
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
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
    if timestamp != nil {
        req.QueryParam("timestamp", *timestamp)
    }
    if mType != nil {
        req.QueryParam("type", *mType)
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

// SearchOrgDeviceEvents takes context, orgId, mac, model, deviceType, text, timestamp, mType, lastBy, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseDeviceEventsSearch data and
// an error if there was an issue with the request or response.
// Search Org Devices Events
func (o *OrgsDevices) SearchOrgDeviceEvents(
    ctx context.Context,
    orgId uuid.UUID,
    mac *string,
    model *string,
    deviceType *models.DeviceTypeEnum,
    text *string,
    timestamp *string,
    mType *string,
    lastBy *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseDeviceEventsSearch],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/devices/events/search", orgId),
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
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
    if timestamp != nil {
        req.QueryParam("timestamp", *timestamp)
    }
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if lastBy != nil {
        req.QueryParam("last_by", *lastBy)
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
    
    var result models.ResponseDeviceEventsSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseDeviceEventsSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountOrgDeviceLastConfigs takes context, orgId, mType, distinct, start, end, limit as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Counts the number of entries in device config history for distinct field with given filters
func (o *OrgsDevices) CountOrgDeviceLastConfigs(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeEnum,
    distinct *models.OrgDevicesLastConfigsCountDistinctEnum,
    start *int,
    end *int,
    limit *int) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/devices/last_config/count", orgId),
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
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
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    
    var result models.RepsonseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchOrgDeviceLastConfigs takes context, orgId, mType, mac, name, version, start, end, limit, duration as parameters and
// returns an models.ApiResponse with models.ResponseConfigHistorySearch data and
// an error if there was an issue with the request or response.
// Search Device Last Configs
func (o *OrgsDevices) SearchOrgDeviceLastConfigs(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeEnum,
    mac *string,
    name *string,
    version *string,
    start *int,
    end *int,
    limit *int,
    duration *string) (
    models.ApiResponse[models.ResponseConfigHistorySearch],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/devices/last_config/search", orgId),
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    if mType != nil {
        req.QueryParam("type", *mType)
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
    
    var result models.ResponseConfigHistorySearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseConfigHistorySearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListOrgApsMacs takes context, orgId, page, limit as parameters and
// returns an models.ApiResponse with []models.ApRadioMac data and
// an error if there was an issue with the request or response.
// For some scenarios like E911 or security systems, the BSSIDs are required to identify which AP the client is connecting to. Then the location of the AP can be used as the approximate location of the client.
// Each radio MAC can have 16 BSSIDs (enumerate the last octet from 0-F)
func (o *OrgsDevices) ListOrgApsMacs(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.ApRadioMac],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/devices/radio_macs", orgId),
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
    
    var result []models.ApRadioMac
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ApRadioMac](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchOrgDevices takes context, orgId, hostname, siteId, model, mac, version, powerConstrained, ipAddress, mxtunnelStatus, mxedgeId, lldpSystemName, lldpSystemDesc, lldpPortId, lldpMgmtAddr, band24Bandwith, band5Bandwith, band6Bandwith, band24Channel, band5Channel, band6Channel, eth0PortSpeed, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseDeviceSearch data and
// an error if there was an issue with the request or response.
// Search Org Devices
func (o *OrgsDevices) SearchOrgDevices(
    ctx context.Context,
    orgId uuid.UUID,
    hostname *string,
    siteId *string,
    model *string,
    mac *string,
    version *string,
    powerConstrained *bool,
    ipAddress *string,
    mxtunnelStatus *models.SearchOrgDevicesMxtunnelStatusEnum,
    mxedgeId *string,
    lldpSystemName *string,
    lldpSystemDesc *string,
    lldpPortId *string,
    lldpMgmtAddr *string,
    band24Bandwith *int,
    band5Bandwith *int,
    band6Bandwith *int,
    band24Channel *int,
    band5Channel *int,
    band6Channel *int,
    eth0PortSpeed *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseDeviceSearch],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/devices/search", orgId),
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    if hostname != nil {
        req.QueryParam("hostname", *hostname)
    }
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if model != nil {
        req.QueryParam("model", *model)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if version != nil {
        req.QueryParam("version", *version)
    }
    if powerConstrained != nil {
        req.QueryParam("power_constrained", *powerConstrained)
    }
    if ipAddress != nil {
        req.QueryParam("ip_address", *ipAddress)
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
    if band24Bandwith != nil {
        req.QueryParam("band_24_bandwith", *band24Bandwith)
    }
    if band5Bandwith != nil {
        req.QueryParam("band_5_bandwith", *band5Bandwith)
    }
    if band6Bandwith != nil {
        req.QueryParam("band_6_bandwith", *band6Bandwith)
    }
    if band24Channel != nil {
        req.QueryParam("band_24_channel", *band24Channel)
    }
    if band5Channel != nil {
        req.QueryParam("band_5_channel", *band5Channel)
    }
    if band6Channel != nil {
        req.QueryParam("band_6_channel", *band6Channel)
    }
    if eth0PortSpeed != nil {
        req.QueryParam("eth0_port_speed", *eth0PortSpeed)
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
    
    var result models.ResponseDeviceSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseDeviceSearch](decoder)
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
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/ocdevices/outbound_ssh_cmd", orgId),
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
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
