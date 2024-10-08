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

// UtilitiesWAN represents a controller struct.
type UtilitiesWAN struct {
    baseController
}

// NewUtilitiesWAN creates a new instance of UtilitiesWAN.
// It takes a baseController as a parameter and returns a pointer to the UtilitiesWAN.
func NewUtilitiesWAN(baseController baseController) *UtilitiesWAN {
    utilitiesWAN := UtilitiesWAN{baseController: baseController}
    return &utilitiesWAN
}

// ClearSiteSsrArpCache takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Clear ARP cache for SSR, SRX and Switch
// Clear the entire ARP cache or a subset if arguments are provided.
// *Note*: port_id is optional if neither vlan nor ip is specified
func (u *UtilitiesWAN) ClearSiteSsrArpCache(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsClearArp) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/clear_arp", siteId, deviceId),
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

// ClearSiteSsrBgpRoutes takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Clear routes associated with one or all BGP neighbors
func (u *UtilitiesWAN) ClearSiteSsrBgpRoutes(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsClearBgp) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/clear_bgp", siteId, deviceId),
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
        "400": {Message: "parameter neighbor absent"},
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

// ClearSiteDeviceSession takes context, siteId, deviceId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Clear session
func (u *UtilitiesWAN) ClearSiteDeviceSession(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsClearSession) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/clear_session", siteId, deviceId),
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

// ReleaseSiteSsrDhcpLease takes context, siteId, deviceId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Releases an active DHCP lease.
func (u *UtilitiesWAN) ReleaseSiteSsrDhcpLease(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsReleaseDhcp) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/release_dhcp", siteId, deviceId),
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

// TestSiteSsrDnsResolution takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// DNS resolutions are performed on the Device.
// The output will be available through websocket. As there can be multiple command issued against the same SSR at the same time and the output all goes through the same websocket stream, `session` is used for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
// ##### Example output from ws stream
// ```
// Router      | Hostname               | Resolved | Last Resolved        | Expiration
// -------------|------------------------|----------|----------------------|---------------------
// test-device | xxx.yyy.net            | Y        | 2022-03-28T03:56:49Z | 2022-03-28T03:57:49Z
// ```
func (u *UtilitiesWAN) TestSiteSsrDnsResolution(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/resolve_dns", siteId, deviceId),
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
    
    var result models.WebsocketSession
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSession](decoder)
    return models.NewApiResponse(result, resp), err
}

// RunSiteSrxTopCommand takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with models.WebsocketSessionWithUrl data and
// an error if there was an issue with the request or response.
// Run top command on switches and SRX. The output will be available through websocket. 
// As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
func (u *UtilitiesWAN) RunSiteSrxTopCommand(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.WebsocketSessionWithUrl],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/run_top", siteId, deviceId),
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

// ServicePingFromSsr takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Ping from SSR
// Service Ping can be performed from the Device. The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, session is introduced for demux.
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
func (u *UtilitiesWAN) ServicePingFromSsr(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsServicePing) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/service_ping", siteId, deviceId),
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

// GetSiteSsrOspfDatabase takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get OSPF Database from the Device. The output will be available through websocket. 
// As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
// #### Example output from ws stream
// ```
// ===== ==================== ========== ======= ======== ================ =================== =================
// Vrf   Neighbor Router ID   Priority   State   Uptime   Dead Timer Due   Interface Address   Interface State
// ===== ==================== ========== ======= ======== ================ =================== =================
// 1.0.0.3                     1   Full       852               38   172.16.3.2          Backup
// 1.0.0.4                     1   Full       811               33   172.16.3.2          DROther
// 1.0.0.3                     1   Full       852               38   172.16.4.2          Backup
// 1.0.0.4                     1   Full       811               34   172.16.4.2          DROther
// ```
func (u *UtilitiesWAN) GetSiteSsrOspfDatabase(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowOspfDatabase) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/show_ospf_database", siteId, deviceId),
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

// GetSiteSsrOspfInterface takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get OSPF interfaces from the Device. The output will be available through websocket. 
// As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
// #### Example output from ws stream
// ```
// ===== ================== =================== ============== =============== =========== ========= ===========
// Vrf   Device Interface   Network Interface   Interface Up   IP Address      OSPF Type   Area ID   Area Type
// ===== ================== =================== ============== =============== =========== ========= ===========
// net1               g1                          True   172.16.1.2/24   Broadcast   0.0.0.0   default
// net3               g3                          True   172.16.3.2/24   Broadcast   0.0.0.0   default
// net4               g4                          True   172.16.4.2/24   Broadcast   0.0.0.4   default
// ```
func (u *UtilitiesWAN) GetSiteSsrOspfInterface(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowOspfInterfaces) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/show_ospf_interfaces", siteId, deviceId),
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

// GetSiteSsrOspfNeighbors takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get OSPF Neighbors from the Device. The output will be available through websocket. 
// As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
// #### Example output from ws stream
// ```
// ===== ==================== ========== ======= ======== ================ =================== =================
// Vrf   Neighbor Router ID   Priority   State   Uptime   Dead Timer Due   Interface Address   Interface State
// ===== ==================== ========== ======= ======== ================ =================== =================
// 1.0.0.3                     1   Full       852               38   172.16.3.2          Backup
// 1.0.0.4                     1   Full       811               33   172.16.3.2          DROther
// 1.0.0.3                     1   Full       852               38   172.16.4.2          Backup
// 1.0.0.4                     1   Full       811               34   172.16.4.2          DROther
// ```
func (u *UtilitiesWAN) GetSiteSsrOspfNeighbors(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowOspfNeighbors) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/show_ospf_neighbors", siteId, deviceId),
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

// GetSiteSsrOspfSummary takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get OSPF summary from the Device. The output will be available through websocket. 
// As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
// #### Example output from ws stream
// ```
// ===== =========== ========== ============= ==================== ========= =========== =============
// Vrf   Router ID   ABR Type   ASBR Router   External LSA Count   Area ID   Area Type   Area Border
// Router
// ===== =========== ========== ============= ==================== ========= =========== =============
// 1.0.0.2     cisco            False                    0   0.0.0.0
// 1.0.0.2     cisco            False                    0   0.0.0.4   default
// ```
func (u *UtilitiesWAN) GetSiteSsrOspfSummary(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowOspfSummary) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/show_ospf_summary", siteId, deviceId),
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

// GetSiteSsrAndSrxRoutes takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get routes from SSR, SRX and Switch. 
// The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
// ##### Example output from ws stream
// ```
// admin@labsystem1.fiedler# show bgp neighbors
// BGP neighbor is 192.168.4.1, remote AS 4200000001, local AS 4200000128, external
// link
// BGP version 4, remote router ID 1.1.1.1
// BGP state = Established, up for 00:27:25
// Last read 00:00:25, hold time is 90, keepalive interval is 30 seconds
// Configured hold time is 90, keepalive interval is 30 seconds
// Neighbor capabilities:
// 4 Byte AS: advertised and received
// Route refresh: advertised and received(old &amp; new)
// Address family IPv4 Unicast: advertised and received
// Graceful Restart Capabilty: advertised and received
// Remote Restart timer is 120 seconds
// Address families by peer:
// none
// ...
// ```
func (u *UtilitiesWAN) GetSiteSsrAndSrxRoutes(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowRoute) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/show_route", siteId, deviceId),
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

// GetSiteSsrServicePath takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get service path information of the Device.
// The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, session is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
// ##### Example output from ws stream
// ```
// show service_path
// Service    Service-route     Type              Destination  Next-Hop  Interface  Vector  Cost  Rate  Capacity        State
// Web        web-route1        service_agent     4.4.4.4      1.1.1.2     lan        red     10    1    200/3000       Up*
// Web        web-route1        service_agent     4.4.4.4      1.1.1.3     lan        red     10    1    200/3000       Up
// Web        web-route2        service_agent     5.5.5.5      2.2.2.2     lan       blue     20    2    50/unlimited   Down
// Login      <None>            BgpOverSVR        10.1.1.1     1.2.3.4     wan        red     10    3        -          Up
// Login      <None>            BgpOverSVR        11.1.1.1     1.2.3.4     wan        red     10    1        -          Up
// App1       <None>            Routed                -           -         -          -      -     -        -          -
// App1       learned-routed    Routed                -           -         -          -      -     -        -          -
// ```
func (u *UtilitiesWAN) GetSiteSsrServicePath(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowServicePath) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/show_service_path", siteId, deviceId),
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

// GetSiteSsrAndSrxSessions takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get active sessions passing through the Device.
// The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, session is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }
// ```
// ##### Example output from ws stream
// ```console
// admin@ssr.node# show sessions
// Fri 2020-04-17 16:55:34 UTC
// Node: node1
// ====================================== ===== ============= =========== ========== ====== ======= ================= ========== ================= =========== ================= ========== =================== ========= =================
// Session Id                             Dir   Service       Tenant      Dev Name   VLAN   Proto   Src IP            Src Port   Dest IP           Dest Port   NAT IP            NAT Port   Payload Encrypted   Timeout   Uptime
// ====================================== ===== ============= =========== ========== ====== ======= ================= ========== ================= =========== ================= ========== =================== ========= =================
// 01187fb8-765a-45e5-ae90-37d77f15e292   fwd   Internet      lanSubnet   lan           0   udp     192.168.0.28         44674   35.166.173.18          9930   96.230.191.130       19569   false                   154   0 days  0:00:28
// 01187fb8-765a-45e5-ae90-37d77f15e292   rev   Internet      lanSubnet   wan           0   udp     35.166.173.18         9930   96.230.191.130        19569   0.0.0.0                  0   false                   154   0 days  0:00:28
// 0859a4ae-bcff-4aa6-b812-79a5236a6c13   fwd   Internet      lanSubnet   lan           0   tcp     192.168.0.41         60843   17.249.171.246          443   96.230.191.130       51941   false                     2   0 days  0:00:10
// admin@node0.90ec7732e327# show sessions by-id 262900cd-bca8-443a-8aab-e5c93d147ab5
// Wed 2024-06-26 20:37:48 UTC
// Retrieving session information...
// =======================================================================================================================================================================================
// 90ec7732e327.node0    Session ID: 262900cd-bca8-443a-8aab-e5c93d147ab5
// =======================================================================================================================================================================================
// Service Name:                      Internet
// Service Route Name:                Internet-sr-local-breakout-1-node0
// Session Source:                    SourceType: PUBLIC
// Session Type:                      HTTPS
// Service Class:                     service-class-0-low
// Source Tenant:                     LAN
// Peer Name:                         N/A
// Inter Node:                        N/A
// Inter Router:                      N/A
// Ingress Source Nat:                N/A
// Payload Security Policy:           internal
// Payload Encrypted:                 True
// Common Name Info:                  N/A
// Tcp Time To Establish:             76
// Tls Time To Establish:             76
// Domain Name:                       N/A
// Uri:                               N/A
// Category:                          N/A
// ```"
func (u *UtilitiesWAN) GetSiteSsrAndSrxSessions(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowSession) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/show_session", siteId, deviceId),
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
