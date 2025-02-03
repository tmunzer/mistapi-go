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
* [Release Site Ssr Dhcp Lease](../../doc/controllers/utilities-wan.md#release-site-ssr-dhcp-lease)
* [Run Site Srx Top Command](../../doc/controllers/utilities-wan.md#run-site-srx-top-command)
* [Service Ping From Ssr](../../doc/controllers/utilities-wan.md#service-ping-from-ssr)
* [Show Site Ssr and Srx Routes](../../doc/controllers/utilities-wan.md#show-site-ssr-and-srx-routes)
* [Show Site Ssr and Srx Sessions](../../doc/controllers/utilities-wan.md#show-site-ssr-and-srx-sessions)
* [Show Site Ssr Ospf Database](../../doc/controllers/utilities-wan.md#show-site-ssr-ospf-database)
* [Show Site Ssr Ospf Interfaces](../../doc/controllers/utilities-wan.md#show-site-ssr-ospf-interfaces)
* [Show Site Ssr Ospf Neighbors](../../doc/controllers/utilities-wan.md#show-site-ssr-ospf-neighbors)
* [Show Site Ssr Ospf Summary](../../doc/controllers/utilities-wan.md#show-site-ssr-ospf-summary)
* [Show Site Ssr Service Path](../../doc/controllers/utilities-wan.md#show-site-ssr-service-path)
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
    AdditionalProperties: map[string]interface{}{
        "port_id": interface{}("ge-0/0/1.10"),
    },
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
    Ip:                   models.ToPointer("10.1.1.1"),
    PortId:               models.ToPointer("wan"),
    Vlan:                 models.ToPointer(1000),
    Vrf:                  models.ToPointer("guest"),
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
    Neighbor:             "all",
    Type:                 models.UtilsClearBgpTypeEnum_IN,
    Vrf:                  models.ToPointer("TestVrf"),
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
| 400 | Parameter neighbor absent | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Release Site Ssr Dhcp Lease

Releases an active DHCP lease.

The output will be available through websocket. As there can be multiple command issued  against the same Device at the same time and the output all goes through the same websocket stream, session is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
    "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
}```



#### Example output from ws stream

```json
{  
  "channel": "/sites/d6fb4f96-xxxx-xxxx-xxxx-xxxxxxxxxxxx/devices/00000000-0000-0000-1000-xxxxxxxxxxxx/cmd",
  "event": "data",
  "data": {
    "session": "9106e908-74dc-4a4f-9050-9c2adcaf44a5",
    "raw": "Running traceroute...\ntraceroute to 8.8.8.8, 64 hops max\n 0  192.168.1.1 1 ms  192.168.1.1 1 ms  192.168.1.1 1 ms\n 1  80.10.236.81 2 ms  80.10.236.81 4 ms  80.10.236.81 2 ms\n 2  193.253.80.250 3 ms  193.253.80.250 2 ms  193.253.80.250 2 ms\n 3  193.252.159.41 2 ms  193.252.159.41 1 ms  193.252.159.41 3 ms\n"
  }
}
```

"

```go
ReleaseSiteSsrDhcpLease(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsReleaseDhcp) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsReleaseDhcp`](../../doc/models/utils-release-dhcp.md) | Body, Optional | - |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsReleaseDhcp{
    PortId:               "ge-0/0/1.10",
}

apiResponse, err := utilitiesWAN.ReleaseSiteSsrDhcpLease(ctx, siteId, deviceId, &body)
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
    Count:                models.ToPointer(10),
    Host:                 "1.1.1.1",
    Service:              "web-session",
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


# Show Site Ssr and Srx Routes

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
ShowSiteSsrAndSrxRoutes(
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
| `body` | [`*models.UtilsShowRoute`](../../doc/models/utils-show-route.md) | Body, Optional | All attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowRoute{
    Duration:             models.ToPointer(0),
    Interval:             models.ToPointer(0),
    Neighbor:             models.ToPointer("192.168.4.1"),
    Prefix:               models.ToPointer("192.168.0.5/30"),
    Protocol:             models.ToPointer(models.UtilsShowRouteProtocolEnum_BGP),
    Route:                models.ToPointer("advertised"),
    Vrf:                  models.ToPointer("default"),
}

apiResponse, err := utilitiesWAN.ShowSiteSsrAndSrxRoutes(ctx, siteId, deviceId, &body)
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


# Show Site Ssr and Srx Sessions

Get active sessions passing through the Device.

The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, session is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

`json { "subscribe": "/sites/{site_id}/devices/{device_id}/cmd" }`

#### Example output from ws stream

`json { "channel": "/sites/d6fb4f96-xxxx-xxxx-xxxx-xxxxxxxxxxxx/devices/00000000-0000-0000-1000-xxxxxxxxxxxx/cmd", "event": "data", "data": { "session": "f517bf29-1141-41ae-a084-17cacb0ccb57", "raw": "{\"status\":\"SUCCESS\",\"finished\":true,\"rows\":[{\"session_id\":\"a04b1cc7-dcc1-40a6-a010-0fe46ca38551\",\"direction\":\"forward\",\"service\":\"internet\",\"tenant\":\"SRV.PRD-Core\",\"device_interface\":\"ge-0/0/3\",\"network_interface\":\"ge-0/0/3.100\",\"protocol\":\"TCP\",\"source_ip\":\"10.3.20.101\",\"source_port\":45733,\"destination_ip\":\"13.38.46.35\",\"destination_port\":443,\"nat_ip\":\"192.168.1.115\",\"nat_port\":45256,\"payload_encrypted\":false,\"timeout\":1581,\"uptime\":319},{\"session_id\":\"a04b1cc7-dcc1-40a6-a010-0fe46ca38551\",\"direction\":\"reverse\",\"service\":\"internet\",\"tenant\":\"SRV.PRD-Core\",\"device_interface\":\"ge-0/0/0\",\"network_interface\":\"ge-0/0/0\",\"protocol\":\"TCP\",\"source_ip\":\"13.38.46.35\",\"source_port\":443,\"destination_ip\":\"192.168.1.115\",\"destination_port\":45256,\"nat_ip\":\"0.0.0.0\",\"nat_port\":0,\"payload_encrypted\":false,\"timeout\":1581,\"uptime\":319}]}\n" } }`

```go
ShowSiteSsrAndSrxSessions(
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
    Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
    ServiceName:          models.ToPointer("any"),
}

apiResponse, err := utilitiesWAN.ShowSiteSsrAndSrxSessions(ctx, siteId, deviceId, &body)
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


# Show Site Ssr Ospf Database

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
ShowSiteSsrOspfDatabase(
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
| `body` | [`*models.UtilsShowOspfDatabase`](../../doc/models/utils-show-ospf-database.md) | Body, Optional | All attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowOspfDatabase{
    SelfOriginate:        models.ToPointer(false),
    Vrf:                  models.ToPointer("lan"),
}

apiResponse, err := utilitiesWAN.ShowSiteSsrOspfDatabase(ctx, siteId, deviceId, &body)
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


# Show Site Ssr Ospf Interfaces

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
ShowSiteSsrOspfInterfaces(
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
| `body` | [`*models.UtilsShowOspfInterfaces`](../../doc/models/utils-show-ospf-interfaces.md) | Body, Optional | All attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowOspfInterfaces{
    PortId:               models.ToPointer("ge-0/0/3"),
    Vrf:                  models.ToPointer("lan"),
}

apiResponse, err := utilitiesWAN.ShowSiteSsrOspfInterfaces(ctx, siteId, deviceId, &body)
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


# Show Site Ssr Ospf Neighbors

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
ShowSiteSsrOspfNeighbors(
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
| `body` | [`*models.UtilsShowOspfNeighbors`](../../doc/models/utils-show-ospf-neighbors.md) | Body, Optional | All attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowOspfNeighbors{
    Neighbor:             models.ToPointer("10.1.1.1"),
    PortId:               models.ToPointer("ge-0/0/3"),
    Vrf:                  models.ToPointer("lan"),
}

apiResponse, err := utilitiesWAN.ShowSiteSsrOspfNeighbors(ctx, siteId, deviceId, &body)
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


# Show Site Ssr Ospf Summary

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
ShowSiteSsrOspfSummary(
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
| `body` | [`*models.UtilsShowOspfSummary`](../../doc/models/utils-show-ospf-summary.md) | Body, Optional | All attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowOspfSummary{
    Vrf:                  models.ToPointer("lan"),
}

apiResponse, err := utilitiesWAN.ShowSiteSsrOspfSummary(ctx, siteId, deviceId, &body)
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


# Show Site Ssr Service Path

Get service path information of the Device.

The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, session is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

`json { "subscribe": "/sites/{site_id}/devices/{device_id}/cmd" }`

#### Example output from ws stream

`json { "channel": "/sites/d6fb4f96-xxxx-xxxx-xxxx-xxxxxxxxxxxx/devices/00000000-0000-0000-1000-xxxxxxxxxxxx/cmd", "event": "data", "data": { "session":"5cb8a6db-d11a-42cd-bed7-19e9f29e637", "raw":"{\"status\":\"SUCCESS\",\"finished\":true,\"rows\":[{\"service\":\"management\",\"type\":\"service-agent\",\"network_interface\":\"ge-0/0/0\",\"destination\":\"\",\"gateway_ip\":\"192.168.1.1\",\"vector\":\"\",\"cost\":0,\"rate\":0,\"state\":\"Up\",\"capacity\":\"0/unlimited\",\"meetsSLA\":\"Yes\"},{\"service\":\"management\",\"type\":\"service-agent\",\"network_interface\":\"ge-0/0/1\",\"destination\":\"\",\"gateway_ip\":\"192.168.0.1\",\"vector\":\"\",\"cost\":0,\"rate\":0,\"state\":\"Up\",\"capacity\":\"0/unlimited\",\"meetsSLA\":\"Yes\"}]}" } }`

```go
ShowSiteSsrServicePath(
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
    Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
    ServiceName:          models.ToPointer("any"),
}

apiResponse, err := utilitiesWAN.ShowSiteSsrServicePath(ctx, siteId, deviceId, &body)
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

