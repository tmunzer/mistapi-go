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
    req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/%v/clear_arp")
    req.AppendTemplateParams(siteId, deviceId)
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
    req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/%v/clear_bgp")
    req.AppendTemplateParams(siteId, deviceId)
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
        "400": {Message: "Parameter neighbor absent"},
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
    req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/%v/clear_session")
    req.AppendTemplateParams(siteId, deviceId)
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
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Releases an active DHCP lease.
// The output will be available through websocket. As there can be multiple command issued  against the same Device at the same time and the output all goes through the same websocket stream, session is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json
// {
// "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
// }```
// #### Example output from ws stream
// ```json
// {  
// "channel": "/sites/d6fb4f96-xxxx-xxxx-xxxx-xxxxxxxxxxxx/devices/00000000-0000-0000-1000-xxxxxxxxxxxx/cmd",
// "event": "data",
// "data": {
// "session": "9106e908-74dc-4a4f-9050-9c2adcaf44a5",
// "raw": "Running traceroute...\ntraceroute to 8.8.8.8, 64 hops max\n 0  192.168.1.1 1 ms  192.168.1.1 1 ms  192.168.1.1 1 ms\n 1  80.10.236.81 2 ms  80.10.236.81 4 ms  80.10.236.81 2 ms\n 2  193.253.80.250 3 ms  193.253.80.250 2 ms  193.253.80.250 2 ms\n 3  193.252.159.41 2 ms  193.252.159.41 1 ms  193.252.159.41 3 ms\n"
// }
// }
// ```
// "
func (u *UtilitiesWAN) ReleaseSiteSsrDhcpLease(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsReleaseDhcp) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/%v/release_dhcp")
    req.AppendTemplateParams(siteId, deviceId)
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
    
    var result models.WebsocketSession
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSession](decoder)
    return models.NewApiResponse(result, resp), err
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
    req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/%v/resolve_dns")
    req.AppendTemplateParams(siteId, deviceId)
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
    req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/%v/run_top")
    req.AppendTemplateParams(siteId, deviceId)
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
    req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/%v/service_ping")
    req.AppendTemplateParams(siteId, deviceId)
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

// ShowSiteGatewayOspfDatabase takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get OSPF Database from SSR and SRX. The output will be available through websocket. 
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
func (u *UtilitiesWAN) ShowSiteGatewayOspfDatabase(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowOspfDatabase) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      "/api/v1/sites/%v/devices/%v/show_ospf_database",
    )
    req.AppendTemplateParams(siteId, deviceId)
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

// ShowSiteGatewayOspfInterfaces takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get OSPF interfaces from SSR and SRX. The output will be available through websocket. 
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
func (u *UtilitiesWAN) ShowSiteGatewayOspfInterfaces(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowOspfInterfaces) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      "/api/v1/sites/%v/devices/%v/show_ospf_interfaces",
    )
    req.AppendTemplateParams(siteId, deviceId)
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

// ShowSiteGatewayOspfNeighbors takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get OSPF Neighbors from SSR and SRX. The output will be available through websocket. 
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
func (u *UtilitiesWAN) ShowSiteGatewayOspfNeighbors(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowOspfNeighbors) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      "/api/v1/sites/%v/devices/%v/show_ospf_neighbors",
    )
    req.AppendTemplateParams(siteId, deviceId)
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

// ShowSiteGatewayOspfSummary takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get OSPF summary from SSR and SRX. The output will be available through websocket. 
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
func (u *UtilitiesWAN) ShowSiteGatewayOspfSummary(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowOspfSummary) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      "/api/v1/sites/%v/devices/%v/show_ospf_summary",
    )
    req.AppendTemplateParams(siteId, deviceId)
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

// ShowSiteSsrAndSrxRoutes takes context, siteId, deviceId, body as parameters and
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
// Graceful Restart Capability: advertised and received
// Remote Restart timer is 120 seconds
// Address families by peer:
// none
// ...
// ```
func (u *UtilitiesWAN) ShowSiteSsrAndSrxRoutes(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowRoute) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/%v/show_route")
    req.AppendTemplateParams(siteId, deviceId)
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

// ShowSiteSsrServicePath takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get service path information of the Device.
// The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, session is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json { "subscribe": "/sites/{site_id}/devices/{device_id}/cmd" } ```
// #### Example output from ws stream
// ```json { "channel": "/sites/d6fb4f96-xxxx-xxxx-xxxx-xxxxxxxxxxxx/devices/00000000-0000-0000-1000-xxxxxxxxxxxx/cmd", "event": "data", "data": { "session":"5cb8a6db-d11a-42cd-bed7-19e9f29e637", "raw":"{\"status\":\"SUCCESS\",\"finished\":true,\"rows\":[{\"service\":\"management\",\"type\":\"service-agent\",\"network_interface\":\"ge-0/0/0\",\"destination\":\"\",\"gateway_ip\":\"192.168.1.1\",\"vector\":\"\",\"cost\":0,\"rate\":0,\"state\":\"Up\",\"capacity\":\"0/unlimited\",\"meetsSLA\":\"Yes\"},{\"service\":\"management\",\"type\":\"service-agent\",\"network_interface\":\"ge-0/0/1\",\"destination\":\"\",\"gateway_ip\":\"192.168.0.1\",\"vector\":\"\",\"cost\":0,\"rate\":0,\"state\":\"Up\",\"capacity\":\"0/unlimited\",\"meetsSLA\":\"Yes\"}]}" } } ```
func (u *UtilitiesWAN) ShowSiteSsrServicePath(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowServicePath) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      "/api/v1/sites/%v/devices/%v/show_service_path",
    )
    req.AppendTemplateParams(siteId, deviceId)
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

// ShowSiteSsrAndSrxSessions takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.WebsocketSession data and
// an error if there was an issue with the request or response.
// Get active sessions passing through the Device.
// The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, session is introduced for demux.
// #### Subscribe to Device Command outputs
// `WS /api-ws/v1/stream`
// ```json { "subscribe": "/sites/{site_id}/devices/{device_id}/cmd" }```
// #### Example output from ws stream
// ```json { "channel": "/sites/d6fb4f96-xxxx-xxxx-xxxx-xxxxxxxxxxxx/devices/00000000-0000-0000-1000-xxxxxxxxxxxx/cmd", "event": "data", "data": { "session": "f517bf29-1141-41ae-a084-17cacb0ccb57", "raw": "{\"status\":\"SUCCESS\",\"finished\":true,\"rows\":[{\"session_id\":\"a04b1cc7-dcc1-40a6-a010-0fe46ca38551\",\"direction\":\"forward\",\"service\":\"internet\",\"tenant\":\"SRV.PRD-Core\",\"device_interface\":\"ge-0/0/3\",\"network_interface\":\"ge-0/0/3.100\",\"protocol\":\"TCP\",\"source_ip\":\"10.3.20.101\",\"source_port\":45733,\"destination_ip\":\"13.38.46.35\",\"destination_port\":443,\"nat_ip\":\"192.168.1.115\",\"nat_port\":45256,\"payload_encrypted\":false,\"timeout\":1581,\"uptime\":319},{\"session_id\":\"a04b1cc7-dcc1-40a6-a010-0fe46ca38551\",\"direction\":\"reverse\",\"service\":\"internet\",\"tenant\":\"SRV.PRD-Core\",\"device_interface\":\"ge-0/0/0\",\"network_interface\":\"ge-0/0/0\",\"protocol\":\"TCP\",\"source_ip\":\"13.38.46.35\",\"source_port\":443,\"destination_ip\":\"192.168.1.115\",\"destination_port\":45256,\"nat_ip\":\"0.0.0.0\",\"nat_port\":0,\"payload_encrypted\":false,\"timeout\":1581,\"uptime\":319}]}\n" } } ```
func (u *UtilitiesWAN) ShowSiteSsrAndSrxSessions(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowSession) (
    models.ApiResponse[models.WebsocketSession],
    error) {
    req := u.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/%v/show_session")
    req.AppendTemplateParams(siteId, deviceId)
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
