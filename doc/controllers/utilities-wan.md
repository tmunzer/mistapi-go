# Utilities WAN

```go
utilitiesWAN := client.UtilitiesWAN()
```

## Class Name

`UtilitiesWAN`

## Methods

* [Clear Site Device Session](../../doc/controllers/utilities-wan.md#clear-site-device-session)
* [Clear Site Ssr Arp Cache](../../doc/controllers/utilities-wan.md#clear-site-ssr-arp-cache)
* [Clear Site Ssr Bgp Routes](../../doc/controllers/utilities-wan.md#clear-site-ssr-bgp-routes)
* [Get Site Ssr and Srx Routes](../../doc/controllers/utilities-wan.md#get-site-ssr-and-srx-routes)
* [Get Site Ssr and Srx Sessions](../../doc/controllers/utilities-wan.md#get-site-ssr-and-srx-sessions)
* [Get Site Ssr Ospf Database](../../doc/controllers/utilities-wan.md#get-site-ssr-ospf-database)
* [Get Site Ssr Ospf Interface](../../doc/controllers/utilities-wan.md#get-site-ssr-ospf-interface)
* [Get Site Ssr Ospf Neighbors](../../doc/controllers/utilities-wan.md#get-site-ssr-ospf-neighbors)
* [Get Site Ssr Ospf Summary](../../doc/controllers/utilities-wan.md#get-site-ssr-ospf-summary)
* [Get Site Ssr Service Path](../../doc/controllers/utilities-wan.md#get-site-ssr-service-path)
* [Release Site Ssr Dhcp Lease](../../doc/controllers/utilities-wan.md#release-site-ssr-dhcp-lease)
* [Run Site Srx Top Command](../../doc/controllers/utilities-wan.md#run-site-srx-top-command)
* [Service Ping From Ssr](../../doc/controllers/utilities-wan.md#service-ping-from-ssr)
* [Test Site Ssr Dns Resolution](../../doc/controllers/utilities-wan.md#test-site-ssr-dns-resolution)


# Clear Site Device Session

Clear session

```go
ClearSiteDeviceSession(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsClearSession) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsClearSession`](../../doc/models/utils-clear-session.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsClearSession{
}

resp, err := utilitiesWAN.ClearSiteDeviceSession(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Clear Site Ssr Arp Cache

Clear ARP cache for SSR, SRX and Switch

Clear the entire ARP cache or a subset if arguments are provided.

*Note*: port_id is optional if neither vlan nor ip is specified

```go
ClearSiteSsrArpCache(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsClearArp) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsClearArp`](../../doc/models/utils-clear-arp.md) | Body, Optional | - |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsClearArp{
    Ip:     models.ToPointer("10.1.1.1"),
    PortId: models.ToPointer("wan"),
    Vlan:   models.ToPointer(1000),
    Vrf:    models.ToPointer("guest"),
}

apiResponse, err := utilitiesWAN.ClearSiteSsrArpCache(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Clear Site Ssr Bgp Routes

Clear routes associated with one or all BGP neighbors

```go
ClearSiteSsrBgpRoutes(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsClearBgp) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsClearBgp`](../../doc/models/utils-clear-bgp.md) | Body, Optional | - |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsClearBgp{
    Neighbor: "all",
    Type:     models.UtilsClearBgpTypeEnum("in"),
    Vrf:      models.ToPointer("TestVrf"),
}

apiResponse, err := utilitiesWAN.ClearSiteSsrBgpRoutes(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | parameter neighbor absent | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Site Ssr and Srx Routes

Get routes from SSR, SRX and Switch.

The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
    "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
}
```

##### Example output from ws stream

```
admin@labsystem1.fiedler# show bgp neighbors
BGP neighbor is 192.168.4.1, remote AS 4200000001, local AS 4200000128, external
link
  BGP version 4, remote router ID 1.1.1.1
  BGP state = Established, up for 00:27:25
  Last read 00:00:25, hold time is 90, keepalive interval is 30 seconds
  Configured hold time is 90, keepalive interval is 30 seconds
  Neighbor capabilities:
    4 Byte AS: advertised and received
    Route refresh: advertised and received(old &amp; new)
    Address family IPv4 Unicast: advertised and received
    Graceful Restart Capabilty: advertised and received
      Remote Restart timer is 120 seconds
      Address families by peer:
        none
        ...
```

```go
GetSiteSsrAndSrxRoutes(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowRoute) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsShowRoute`](../../doc/models/utils-show-route.md) | Body, Optional | all attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowRoute{
    Duration: models.ToPointer(0),
    Interval: models.ToPointer(0),
    Neighbor: models.ToPointer("192.168.4.1"),
    Prefix:   models.ToPointer("192.168.0.5/30"),
    Protocol: models.ToPointer(models.UtilsShowRouteProtocolEnum("bgp")),
    Route:    models.ToPointer("advertised"),
    Vrf:      models.ToPointer("default"),
}

apiResponse, err := utilitiesWAN.GetSiteSsrAndSrxRoutes(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Site Ssr and Srx Sessions

Get active sessions passing through the Device.

The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, session is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
    "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
  }
```

##### Example output from ws stream

```console
admin@ssr.node# show sessions
Fri 2020-04-17 16:55:34 UTC

Node: node1

====================================== ===== ============= =========== ========== ====== ======= ================= ========== ================= =========== ================= ========== =================== ========= =================
Session Id                             Dir   Service       Tenant      Dev Name   VLAN   Proto   Src IP            Src Port   Dest IP           Dest Port   NAT IP            NAT Port   Payload Encrypted   Timeout   Uptime
====================================== ===== ============= =========== ========== ====== ======= ================= ========== ================= =========== ================= ========== =================== ========= =================
01187fb8-765a-45e5-ae90-37d77f15e292   fwd   Internet      lanSubnet   lan           0   udp     192.168.0.28         44674   35.166.173.18          9930   96.230.191.130       19569   false                   154   0 days  0:00:28
01187fb8-765a-45e5-ae90-37d77f15e292   rev   Internet      lanSubnet   wan           0   udp     35.166.173.18         9930   96.230.191.130        19569   0.0.0.0                  0   false                   154   0 days  0:00:28
0859a4ae-bcff-4aa6-b812-79a5236a6c13   fwd   Internet      lanSubnet   lan           0   tcp     192.168.0.41         60843   17.249.171.246          443   96.230.191.130       51941   false                     2   0 days  0:00:10


admin@node0.90ec7732e327# show sessions by-id 262900cd-bca8-443a-8aab-e5c93d147ab5
Wed 2024-06-26 20:37:48 UTC
Retrieving session information...

=======================================================================================================================================================================================
90ec7732e327.node0    Session ID: 262900cd-bca8-443a-8aab-e5c93d147ab5
=======================================================================================================================================================================================
Service Name:                      Internet
Service Route Name:                Internet-sr-local-breakout-1-node0
Session Source:                    SourceType: PUBLIC
Session Type:                      HTTPS
Service Class:                     service-class-0-low
Source Tenant:                     LAN
Peer Name:                         N/A
Inter Node:                        N/A
Inter Router:                      N/A
Ingress Source Nat:                N/A
Payload Security Policy:           internal
Payload Encrypted:                 True
Common Name Info:                  N/A
Tcp Time To Establish:             76
Tls Time To Establish:             76
Domain Name:                       N/A
Uri:                               N/A
Category:                          N/A
```"
```

```go
GetSiteSsrAndSrxSessions(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowSession) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsShowSession`](../../doc/models/utils-show-session.md) | Body, Optional | - |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowSession{
    Node:        models.ToPointer(models.HaClusterNodeEnum("node0")),
    ServiceName: models.ToPointer("any"),
}

apiResponse, err := utilitiesWAN.GetSiteSsrAndSrxSessions(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Site Ssr Ospf Database

Get OSPF Database from the Device. The output will be available through websocket.

As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
  "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
}
```

#### Example output from ws stream

```
===== ==================== ========== ======= ======== ================ =================== =================
Vrf   Neighbor Router ID   Priority   State   Uptime   Dead Timer Due   Interface Address   Interface State
===== ==================== ========== ======= ======== ================ =================== =================
      1.0.0.3                     1   Full       852               38   172.16.3.2          Backup
      1.0.0.4                     1   Full       811               33   172.16.3.2          DROther
      1.0.0.3                     1   Full       852               38   172.16.4.2          Backup
      1.0.0.4                     1   Full       811               34   172.16.4.2          DROther
```

```go
GetSiteSsrOspfDatabase(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowOspfDatabase) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsShowOspfDatabase`](../../doc/models/utils-show-ospf-database.md) | Body, Optional | all attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowOspfDatabase{
    SelfOriginate: models.ToPointer(false),
    Vrf:           models.ToPointer("lan"),
}

apiResponse, err := utilitiesWAN.GetSiteSsrOspfDatabase(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Site Ssr Ospf Interface

Get OSPF interfaces from the Device. The output will be available through websocket.

As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
  "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
}
```

#### Example output from ws stream

```
===== ================== =================== ============== =============== =========== ========= ===========
Vrf   Device Interface   Network Interface   Interface Up   IP Address      OSPF Type   Area ID   Area Type
===== ================== =================== ============== =============== =========== ========= ===========
      net1               g1                          True   172.16.1.2/24   Broadcast   0.0.0.0   default
      net3               g3                          True   172.16.3.2/24   Broadcast   0.0.0.0   default
      net4               g4                          True   172.16.4.2/24   Broadcast   0.0.0.4   default
```

```go
GetSiteSsrOspfInterface(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowOspfInterfaces) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsShowOspfInterfaces`](../../doc/models/utils-show-ospf-interfaces.md) | Body, Optional | all attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowOspfInterfaces{
    PortId: models.ToPointer("ge-0/0/3"),
    Vrf:    models.ToPointer("lan"),
}

apiResponse, err := utilitiesWAN.GetSiteSsrOspfInterface(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Site Ssr Ospf Neighbors

Get OSPF Neighbors from the Device. The output will be available through websocket.

As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
  "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
}
```

#### Example output from ws stream

```
===== ==================== ========== ======= ======== ================ =================== =================
Vrf   Neighbor Router ID   Priority   State   Uptime   Dead Timer Due   Interface Address   Interface State
===== ==================== ========== ======= ======== ================ =================== =================
      1.0.0.3                     1   Full       852               38   172.16.3.2          Backup
      1.0.0.4                     1   Full       811               33   172.16.3.2          DROther
      1.0.0.3                     1   Full       852               38   172.16.4.2          Backup
      1.0.0.4                     1   Full       811               34   172.16.4.2          DROther
```

```go
GetSiteSsrOspfNeighbors(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowOspfNeighbors) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsShowOspfNeighbors`](../../doc/models/utils-show-ospf-neighbors.md) | Body, Optional | all attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowOspfNeighbors{
    Neighbor: models.ToPointer("10.1.1.1"),
    PortId:   models.ToPointer("ge-0/0/3"),
    Vrf:      models.ToPointer("lan"),
}

apiResponse, err := utilitiesWAN.GetSiteSsrOspfNeighbors(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Site Ssr Ospf Summary

Get OSPF summary from the Device. The output will be available through websocket.

As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
  "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
}
```

#### Example output from ws stream

```
===== =========== ========== ============= ==================== ========= =========== =============
Vrf   Router ID   ABR Type   ASBR Router   External LSA Count   Area ID   Area Type   Area Border
                                                                                      Router
===== =========== ========== ============= ==================== ========= =========== =============
      1.0.0.2     cisco            False                    0   0.0.0.0
      1.0.0.2     cisco            False                    0   0.0.0.4   default
```

```go
GetSiteSsrOspfSummary(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowOspfSummary) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsShowOspfSummary`](../../doc/models/utils-show-ospf-summary.md) | Body, Optional | all attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowOspfSummary{
    Vrf:  models.ToPointer("lan"),
}

apiResponse, err := utilitiesWAN.GetSiteSsrOspfSummary(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Site Ssr Service Path

Get service path information of the Device.

The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, session is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
    "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
}
```

##### Example output from ws stream

```
show service_path

Service    Service-route     Type              Destination  Next-Hop  Interface  Vector  Cost  Rate  Capacity        State

Web        web-route1        service_agent     4.4.4.4      1.1.1.2     lan        red     10    1    200/3000       Up*
Web        web-route1        service_agent     4.4.4.4      1.1.1.3     lan        red     10    1    200/3000       Up
Web        web-route2        service_agent     5.5.5.5      2.2.2.2     lan       blue     20    2    50/unlimited   Down
Login      <None>            BgpOverSVR        10.1.1.1     1.2.3.4     wan        red     10    3        -          Up
Login      <None>            BgpOverSVR        11.1.1.1     1.2.3.4     wan        red     10    1        -          Up
App1       <None>            Routed                -           -         -          -      -     -        -          -
App1       learned-routed    Routed                -           -         -          -      -     -        -          -
```

```go
GetSiteSsrServicePath(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowServicePath) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsShowServicePath`](../../doc/models/utils-show-service-path.md) | Body, Optional | - |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowServicePath{
    Node:        models.ToPointer(models.HaClusterNodeEnum("node0")),
    ServiceName: models.ToPointer("any"),
}

apiResponse, err := utilitiesWAN.GetSiteSsrServicePath(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Release Site Ssr Dhcp Lease

Releases an active DHCP lease.

```go
ReleaseSiteSsrDhcpLease(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsReleaseDhcp) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsReleaseDhcp`](../../doc/models/utils-release-dhcp.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsReleaseDhcp{
    PortId: "ge-0/0/1.10",
}

resp, err := utilitiesWAN.ReleaseSiteSsrDhcpLease(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Parameter `port` absent | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Run Site Srx Top Command

Run top command on switches and SRX. The output will be available through websocket.

As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
  "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
}
```

```go
RunSiteSrxTopCommand(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.WebsocketSessionWithUrl],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.WebsocketSessionWithUrl`](../../doc/models/websocket-session-with-url.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesWAN.RunSiteSrxTopCommand(ctx, siteId, deviceId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Service Ping From Ssr

Ping from SSR

Service Ping can be performed from the Device. The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, session is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
    "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
}
```

##### Example output from ws stream

```json
{
    "event": "data",
    "channel": "/sites/4ac1dcf4-9d8b-7211-65c4-057819f0862b/devices/00000000-0000-0000-1000-5c5b350e0060/cmd",
    "data": {
        "session": "session_id",
        "raw": "64 bytes from 23.211.0.110: seq=8 ttl=58 time=12.323 ms\n"
    }
}
```

```go
ServicePingFromSsr(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsServicePing) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsServicePing`](../../doc/models/utils-service-ping.md) | Body, Optional | Request Body |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsServicePing{
    Count:   models.ToPointer(10),
    Host:    "1.1.1.1",
    Service: "web-session",
}

apiResponse, err := utilitiesWAN.ServicePingFromSsr(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Test Site Ssr Dns Resolution

DNS resolutions are performed on the Device.

The output will be available through websocket. As there can be multiple command issued against the same SSR at the same time and the output all goes through the same websocket stream, `session` is used for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
    "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
}
```

##### Example output from ws stream

```
 Router      | Hostname               | Resolved | Last Resolved        | Expiration
-------------|------------------------|----------|----------------------|---------------------
 test-device | xxx.yyy.net            | Y        | 2022-03-28T03:56:49Z | 2022-03-28T03:57:49Z
```

```go
TestSiteSsrDnsResolution(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesWAN.TestSiteSsrDnsResolution(ctx, siteId, deviceId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

