package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// UtilitiesCommon represents a controller struct.
type UtilitiesCommon struct {
    baseController
}

// NewUtilitiesCommon creates a new instance of UtilitiesCommon.
// It takes a baseController as a parameter and returns a pointer to the UtilitiesCommon.
func NewUtilitiesCommon(baseController baseController) *UtilitiesCommon {
    utilitiesCommon := UtilitiesCommon{baseController: baseController}
    return &utilitiesCommon
}

// ArpFromDevice takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// ARP can be performed on the Device. The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
// ##### Example output from ws stream
// ```json
// { 
// "event": "data", 
// "channel": "/sites/4ac1dcf4-9d8b-7211-65c4-057819f0862b/devices/00000000-0000-0000-1000-5c5b350e0060/cmd", 
// "data": { 
// "session": "session_id", 
// "raw": 
// "Output": "\tMAC\t\tDEV\tVLAN\tRx Packets\t\t Rx Bytes\t\tTx Packets\t\t Tx Bytes\tFlows\tIdle sec\n-----------------------------------------------------------------------------------------------------------------------"
// } 
// }
// ```
func (u *UtilitiesCommon) ArpFromDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.HaClusterNode) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/arp", siteId, deviceId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.WebsocketSession
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSession](decoder)
    return models.NewApiResponse(result, resp), err
}

// ClearSiteDeviceMacTable takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Clear MAC Table from the Device.
// The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
func (u *UtilitiesCommon) ClearSiteDeviceMacTable(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsMacTable) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/clear_mac_table", siteId, deviceId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.WebsocketSession
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSession](decoder)
    return models.NewApiResponse(result, resp), err
}

// ClearSiteDevicePolicyHitCount takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with models.WebsocketSessionWithUrl data and
// an error if there was an issue with the request or response.
// Clear application policy hit counts for all the policies
func (u *UtilitiesCommon) ClearSiteDevicePolicyHitCount(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.WebsocketSessionWithUrl],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/clear_policy_hit_count", siteId, deviceId),
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
    
    var result models.WebsocketSessionWithUrl
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSessionWithUrl](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteDeviceConfigCmd takes context, siteId, deviceId, sort as parameters and
// returns an models.ApiResponse with models.ResponseDeviceConfigCli data and
// an error if there was an issue with the request or response.
// Get Config CLI Commands
// For a brown-field switch deployment where we adopted the switch through Adoption Command, we do not wipe out / overwrite the existing config automatically. Instead, we generate CLI commands that we would have generated. The user can inspect, modify, and incorporate this into their existing config manually.
// Once they feel comfortable about the config we generate, they can enable allow_mist_config where we will take full control of their config like a claimed switch
func (u *UtilitiesCommon) GetSiteDeviceConfigCmd(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    sort *bool) (
    models.ApiResponse[models.ResponseDeviceConfigCli],
    error) {
    req := u.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/config_cmd", siteId, deviceId),
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
    if sort != nil {
        req.QueryParam("sort", *sort)
    }
    
    var result models.ResponseDeviceConfigCli
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseDeviceConfigCli](decoder)
    return models.NewApiResponse(result, resp), err
}

// StartSiteLocateDevice takes context, siteId, deviceId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// ### Access Points
// Locate an Access Point by blinking it's LED.
// It is a persisted state that has to be stopped by calling Stop Locating API
// ### Switches
// Locate a Switch by blinking all port LEDs. 
// By default, request is sent to `master` switch and LEDs will keep flashing for 5 minutes.
// In case of virtual chassis (VC) the desired member mac has to be passed in the request payload. 
// At anypoint, only one VC member can be requested to flash the LED. 
// To stop LED flashing before the duration ends /unlocate API request can be made. 
// If /unlocate API is not called LED will continue to flash on device for the given duration. 
// Default duration is 5 minutes and 120 minutes is the maximum.
func (u *UtilitiesCommon) StartSiteLocateDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.LocateSwitch) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/locate", siteId, deviceId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// MonitorSiteDeviceTraffic takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSessionWithUrl data and
// an error if there was an issue with the request or response.
// Monitor traffic on switches and SRX.
// * JUNOS uses cmd "monitor interface <port>" to monitor traffic on particular <port>
// * JUNOS uses cmd "monitor interface traffic" to monitor traffic on all ports
func (u *UtilitiesCommon) MonitorSiteDeviceTraffic(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsMonitorTraffic) (
    models.ApiResponse[models.WebsocketSessionWithUrl],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/monitor_traffic", siteId, deviceId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.WebsocketSessionWithUrl
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSessionWithUrl](decoder)
    return models.NewApiResponse(result, resp), err
}

// PingFromDevice takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Ping from AP, Switch and SSR
// Ping can be performed from the Device. The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
// ##### Example output from ws stream
// ```json
// {
// "event": "data",
// "channel": "/sites/4ac1dcf4-9d8b-7211-65c4-057819f0862b/devices/00000000-0000-0000-1000-5c5b350e0060/cmd",
// "data": {
// "session": "session_id",
// "raw": "64 bytes from 23.211.0.110: seq=8 ttl=58 time=12.323 ms\n"
// }
// }
// ```
func (u *UtilitiesCommon) PingFromDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsPing) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/ping", siteId, deviceId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.WebsocketSession
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSession](decoder)
    return models.NewApiResponse(result, resp), err
}

// ReadoptSiteOctermDevice takes context, siteId, deviceId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// For the octerm devices, the device ID must come from fpc0. However, for a VC, the users may change the original fpc0 from CLI. To fix the issue, the readopt API could be used to trigger the readopt process so the device would get the corret device ID to connect the cloud.
func (u *UtilitiesCommon) ReadoptSiteOctermDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/readopt", siteId, deviceId),
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// ReleaseSiteDeviceDhcpLease takes context, siteId, deviceId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Releases an active DHCP lease.
func (u *UtilitiesCommon) ReleaseSiteDeviceDhcpLease(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsReleaseDhcpLeases) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/release_dhcp_leases", siteId, deviceId),
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
        "400": {Message: "Parameter `port ` absent"},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// ReprovisionSiteOctermDevice takes context, siteId, deviceId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// To force one device to reprovision itself again.
func (u *UtilitiesCommon) ReprovisionSiteOctermDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/reprovision", siteId, deviceId),
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// GetSiteDeviceZtpPassword takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with models.RootPasswordString data and
// an error if there was an issue with the request or response.
// In the case where soemthing happens during/after ZTP, the root-password is modified (required for ZTP to set up outbound-ssh) but the user-defined password config has not be configured. This API can be used to retrieve the temporary password.
func (u *UtilitiesCommon) GetSiteDeviceZtpPassword(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.RootPasswordString],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/request_ztp_password", siteId, deviceId),
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
    
    var result models.RootPasswordString
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RootPasswordString](decoder)
    return models.NewApiResponse(result, resp), err
}

// RestartSiteDevice takes context, siteId, deviceId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Restart / Reboot a device
func (u *UtilitiesCommon) RestartSiteDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsDevicesRestart) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/restart", siteId, deviceId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// CreateSiteDeviceShellSession takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with models.WebsocketSessionWithUrl data and
// an error if there was an issue with the request or response.
// Create Shell Session
func (u *UtilitiesCommon) CreateSiteDeviceShellSession(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.WebsocketSessionWithUrl],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/shell", siteId, deviceId),
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
    
    var result models.WebsocketSessionWithUrl
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSessionWithUrl](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteDeviceArpTable takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get ARP Table from the Device.
// The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
func (u *UtilitiesCommon) GetSiteDeviceArpTable(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowArp) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/show_arp", siteId, deviceId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.WebsocketSession
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSession](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteDeviceBgpSummary takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get BGP Summary from SSR, SRX and Switch.
// The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\"\
// }
// ```
// ##### Example output from ws stream
// ```
// Tue 2024-04-23 16:36:06 UTC
// Retrieving bgp entries...
// BGP table version is 354, local router ID is 10.224.8.16, vrf id 0
// Default local pref 100, local AS 65000
// Status codes:  s suppressed, d damped, h history, * valid, > best, = multipath,
// i internal, r RIB_failure, S Stale, R Removed
// Nexthop codes: @NNN nexthop's vrf id, < announce-nh-self
// Origin codes:  i - IGP, e - EGP, ? - incomplete
// RPKI validation codes: V valid, I invalid, N Not found
// Network                                      Next Hop                                  Metric LocPrf Weight Path
// *> 161.161.161.0/24
// ```"
func (u *UtilitiesCommon) GetSiteDeviceBgpSummary(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowBgpRummary) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/show_bgp_rummary", siteId, deviceId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.WebsocketSession
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSession](decoder)
    return models.NewApiResponse(result, resp), err
}

// ShowSiteDeviceDhcpLeases takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSessionWithUrl data and
// an error if there was an issue with the request or response.
// Shows DHCP leases
func (u *UtilitiesCommon) ShowSiteDeviceDhcpLeases(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowDhcpLeases) (
    models.ApiResponse[models.WebsocketSessionWithUrl],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/show_dhcp_leases", siteId, deviceId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.WebsocketSessionWithUrl
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSessionWithUrl](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteDeviceEvpnDatabase takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get EVPN Database from the Device. The output will be available through websocket.
func (u *UtilitiesCommon) GetSiteDeviceEvpnDatabase(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowEvpnDatabase) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/show_evpn_database", siteId, deviceId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.WebsocketSession
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSession](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteDeviceForwardingTable takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get forwarding table from the Device. The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
// ##### Example output from ws stream
// ```
// Mon 2024-05-20 16:47:30 UTC Retrieving fib entries… Entry Count: 3268 Capacity:    22668 ==================== ====== ======= ================== ===== ====================== =========== =========== ====== IP Prefix            Port   Proto   Tenant             VRF   Service                Next Hops   Vector      Cost ==================== ====== ======= ================== ===== ====================== =========== =========== ====== 0.0.0.0/0               0   None    Old_Mgmt           -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10 branch1-Kiosk      -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10 branch1-MGT        -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10 3.1.1.0/24              0   None    Old_Mgmt           -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10 branch1-Kiosk      -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10 branch1-MGT        -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10
// ```
func (u *UtilitiesCommon) GetSiteDeviceForwardingTable(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowForwardingTable) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/show_forwarding_table", siteId, deviceId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.WebsocketSession
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSession](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteDeviceMacTable takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get MAC Table from the Device.
// The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
func (u *UtilitiesCommon) GetSiteDeviceMacTable(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsMacTable) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/show_mac_table", siteId, deviceId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.WebsocketSession
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSession](decoder)
    return models.NewApiResponse(result, resp), err
}

// UploadSiteDeviceSupportFile takes context, siteId, deviceId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Support / Upload device support files
// #### Info Param
// **Parameter**|**Description** 
// :-------------: |:-------------: |:-------------: 
// process|Upload 1 file with output of show system processes extensive
// outbound-ssh|Upload 1 file that concatenates all /var/log/outbound-ssh.log* files
// messages|Upload 1 to 10 /var/log/messages* files
// core-dumps|Upload all core dump files, if any
// full|string|Upload 1 file with output of request support information, 1 file that concatenates all /var/log/outbound-ssh.log files, all core dump files, the 3 most recent /var/log/messages files, and Mist agent logs (for Junos devices running the Mist agent)
// var-logs|Upload all non-empty files in the /var/log/ directory
// jma-logs|Upload Mist agent logs (for Junos devices running the Mist agent only)
func (u *UtilitiesCommon) UploadSiteDeviceSupportFile(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsSendSupportLogs) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/support", siteId, deviceId),
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
        "400": {Message: "Device not online"},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// TracerouteFromDevice takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Traceroute can be performed from the Device.
// The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
func (u *UtilitiesCommon) TracerouteFromDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsTraceroute) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/traceroute", siteId, deviceId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.WebsocketSession
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSession](decoder)
    return models.NewApiResponse(result, resp), err
}

// StopSiteLocateDevice takes context, siteId, deviceId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Stop Locate a Device
func (u *UtilitiesCommon) StopSiteLocateDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/unlocate", siteId, deviceId),
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}
