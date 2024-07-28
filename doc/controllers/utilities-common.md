# Utilities Common

```go
utilitiesCommon := client.UtilitiesCommon()
```

## Class Name

`UtilitiesCommon`

## Methods

* [Arp From Device](../../doc/controllers/utilities-common.md#arp-from-device)
* [Clear Site Device Mac Table](../../doc/controllers/utilities-common.md#clear-site-device-mac-table)
* [Create Site Device Shell Session](../../doc/controllers/utilities-common.md#create-site-device-shell-session)
* [Get Site Device Arp Table](../../doc/controllers/utilities-common.md#get-site-device-arp-table)
* [Get Site Device Bgp Summary](../../doc/controllers/utilities-common.md#get-site-device-bgp-summary)
* [Get Site Device Config Cmd](../../doc/controllers/utilities-common.md#get-site-device-config-cmd)
* [Get Site Device Evpn Database](../../doc/controllers/utilities-common.md#get-site-device-evpn-database)
* [Get Site Device Forwarding Table](../../doc/controllers/utilities-common.md#get-site-device-forwarding-table)
* [Get Site Device Mac Table](../../doc/controllers/utilities-common.md#get-site-device-mac-table)
* [Get Site Device Ztp Password](../../doc/controllers/utilities-common.md#get-site-device-ztp-password)
* [Monitor Site Device Traffic](../../doc/controllers/utilities-common.md#monitor-site-device-traffic)
* [Ping From Device](../../doc/controllers/utilities-common.md#ping-from-device)
* [Readopt Site Octerm Device](../../doc/controllers/utilities-common.md#readopt-site-octerm-device)
* [Reprovision Site Octerm Device](../../doc/controllers/utilities-common.md#reprovision-site-octerm-device)
* [Restart Site Device](../../doc/controllers/utilities-common.md#restart-site-device)
* [Start Site Locate Device](../../doc/controllers/utilities-common.md#start-site-locate-device)
* [Stop Site Locate Device](../../doc/controllers/utilities-common.md#stop-site-locate-device)
* [Traceroute From Device](../../doc/controllers/utilities-common.md#traceroute-from-device)
* [Upload Site Device Support File](../../doc/controllers/utilities-common.md#upload-site-device-support-file)


# Arp From Device

ARP can be performed on the Device. The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.

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
   "raw": 
   "Output": "\tMAC\t\tDEV\tVLAN\tRx Packets\t\t Rx Bytes\t\tTx Packets\t\t Tx Bytes\tFlows\tIdle sec\n-----------------------------------------------------------------------------------------------------------------------"
  } 
}
```

```go
ArpFromDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.HaClusterNode) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.HaClusterNode`](../../doc/models/ha-cluster-node.md) | Body, Optional | - |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.HaClusterNode{
    Node: models.ToPointer(models.HaClusterNodeEnum("node0")),
}

apiResponse, err := utilitiesCommon.ArpFromDevice(ctx, siteId, deviceId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Clear Site Device Mac Table

Clear MAC Table from the Device.

The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
    "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
}
```

```go
ClearSiteDeviceMacTable(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsMacTable) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsMacTable`](../../doc/models/utils-mac-table.md) | Body, Optional | all attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsMacTable{
    MacAddress: models.ToPointer("f8c1165c6400"),
    PortId:     models.ToPointer("ge-0/0/0.0"),
    VlanId:     models.ToPointer("ge-0/0/0.0"),
}

apiResponse, err := utilitiesCommon.ClearSiteDeviceMacTable(ctx, siteId, deviceId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Create Site Device Shell Session

Create Shell Session

```go
CreateSiteDeviceShellSession(
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

apiResponse, err := utilitiesCommon.CreateSiteDeviceShellSession(ctx, siteId, deviceId)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Device Arp Table

Get ARP Table from the Device.

The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
    "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
}
```

```go
GetSiteDeviceArpTable(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowArp) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsShowArp`](../../doc/models/utils-show-arp.md) | Body, Optional | all attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowArp{
    Duration: models.ToPointer(0),
    Interval: models.ToPointer(0),
    Ip:       models.ToPointer("192.168.30.7"),
    PortId:   models.ToPointer("ge-0/0/0.0"),
    Vrf:      models.ToPointer("guest"),
}

apiResponse, err := utilitiesCommon.GetSiteDeviceArpTable(ctx, siteId, deviceId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Device Bgp Summary

Get BGP Summary from SSR, SRX and Switch.

The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
    \"subscribe\": \"/sites/{site_id}/devices/{device_id}/cmd\"\

}
```

##### Example output from ws stream

```
Tue 2024-04-23 16:36:06 UTC
Retrieving bgp entries...
BGP table version is 354, local router ID is 10.224.8.16, vrf id 0
Default local pref 100, local AS 65000
Status codes:  s suppressed, d damped, h history, * valid, > best, = multipath,
              i internal, r RIB_failure, S Stale, R Removed
Nexthop codes: @NNN nexthop's vrf id, < announce-nh-self
Origin codes:  i - IGP, e - EGP, ? - incomplete
RPKI validation codes: V valid, I invalid, N Not found

  Network                                      Next Hop                                  Metric LocPrf Weight Path
*> 161.161.161.0/24
```"
```

```go
GetSiteDeviceBgpSummary(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowBgpRummary) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsShowBgpRummary`](../../doc/models/utils-show-bgp-rummary.md) | Body, Optional | all attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



apiResponse, err := utilitiesCommon.GetSiteDeviceBgpSummary(ctx, siteId, deviceId, nil)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Device Config Cmd

Get Config CLI Commands
For a brown-field switch deployment where we adopted the switch through Adoption Command, we do not wipe out / overwrite the existing config automatically. Instead, we generate CLI commands that we would have generated. The user can inspect, modify, and incorporate this into their existing config manually.

Once they feel comfortable about the config we generate, they can enable allow_mist_config where we will take full control of their config like a claimed switch

```go
GetSiteDeviceConfigCmd(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    sort *bool) (
    models.ApiResponse[models.ResponseDeviceConfigCli],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `sort` | `*bool` | Query, Optional | Make output cmds sorted (for better readability) or not. |

## Response Type

[`models.ResponseDeviceConfigCli`](../../doc/models/response-device-config-cli.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

sort := false

apiResponse, err := utilitiesCommon.GetSiteDeviceConfigCmd(ctx, siteId, deviceId, &sort)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "cli": [
    "set system hostname corp-a135"
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Device Evpn Database

Get EVPN Database from the Device. The output will be available through websocket.

```go
GetSiteDeviceEvpnDatabase(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowEvpnDatabase) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsShowEvpnDatabase`](../../doc/models/utils-show-evpn-database.md) | Body, Optional | all attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowEvpnDatabase{
    Mac:    models.ToPointer("f8c1165c6400"),
    PortId: models.ToPointer("ge-0/0/0.0"),
}

apiResponse, err := utilitiesCommon.GetSiteDeviceEvpnDatabase(ctx, siteId, deviceId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Device Forwarding Table

Get ARP Table from the Device.

Get forwarding table from the Device. The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
    "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
}
```

##### Example output from ws stream

```
Mon 2024-05-20 16:47:30 UTC Retrieving fib entries… Entry Count: 3268 Capacity:    22668 ==================== ====== ======= ================== ===== ====================== =========== =========== ====== IP Prefix            Port   Proto   Tenant             VRF   Service                Next Hops   Vector      Cost ==================== ====== ======= ================== ===== ====================== =========== =========== ====== 0.0.0.0/0               0   None    Old_Mgmt           -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10 branch1-Kiosk      -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10 branch1-MGT        -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10 3.1.1.0/24              0   None    Old_Mgmt           -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10 branch1-Kiosk      -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10 branch1-MGT        -     internet-wan_and_lte   1-2.0       broadband      1 1-4.0       lte           10

```

```go
GetSiteDeviceForwardingTable(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowForwardingTable) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsShowForwardingTable`](../../doc/models/utils-show-forwarding-table.md) | Body, Optional | all attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowForwardingTable{
    Prefix:          models.ToPointer("3.1.1.0/24"),
    ServiceIp:       models.ToPointer("3.1.1.10"),
    ServiceName:     models.ToPointer("internet-wan_and_lte"),
    ServicePort:     models.ToPointer(32768),
    ServiceProtocol: models.ToPointer("udp"),
    ServiceTenant:   models.ToPointer("branch1-wifi-mgt"),
    Vrf:             models.ToPointer("guest"),
}

apiResponse, err := utilitiesCommon.GetSiteDeviceForwardingTable(ctx, siteId, deviceId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Device Mac Table

Get MAC Table from the Device.

The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
    "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
}
```

```go
GetSiteDeviceMacTable(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsMacTable) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsMacTable`](../../doc/models/utils-mac-table.md) | Body, Optional | all attributes are optional |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsMacTable{
    MacAddress: models.ToPointer("f8c1165c6400"),
    PortId:     models.ToPointer("ge-0/0/0.0"),
    VlanId:     models.ToPointer("ge-0/0/0.0"),
}

apiResponse, err := utilitiesCommon.GetSiteDeviceMacTable(ctx, siteId, deviceId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Device Ztp Password

In the case where soemthing happens during/after ZTP, the root-password is modified (required for ZTP to set up outbound-ssh) but the user-defined password config has not be configured. This API can be used to retrieve the temporary password.

```go
GetSiteDeviceZtpPassword(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.RootPasswordString],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.RootPasswordString`](../../doc/models/root-password-string.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesCommon.GetSiteDeviceZtpPassword(ctx, siteId, deviceId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "root_password": "ef8070ef8f924edb592e1819ed64b31172ab8de9d5cde75d3f46acd9506202ab9b1cbb97e381c5aa11037f17e5ed7b4b609461cd813d944670549d410ef82f2e"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Monitor Site Device Traffic

Monitor traffic on switches and SRX.

JUNOS uses cmd “monitor interface” to monitor traffic on particular JUNOS uses cmd “monitor interface traffic” to monitor traffic on all ports

```go
MonitorSiteDeviceTraffic(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsMonitorTraffic) (
    models.ApiResponse[models.WebsocketSessionWithUrl],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsMonitorTraffic`](../../doc/models/utils-monitor-traffic.md) | Body, Optional | - |

## Response Type

[`models.WebsocketSessionWithUrl`](../../doc/models/websocket-session-with-url.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsMonitorTraffic{
    Port: models.ToPointer("ge-0/0/1"),
}

apiResponse, err := utilitiesCommon.MonitorSiteDeviceTraffic(ctx, siteId, deviceId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Ping From Device

Ping from AP, Switch and SSR

Ping can be performed from the Device. The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.

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
PingFromDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsPing) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsPing`](../../doc/models/utils-ping.md) | Body, Optional | Request Body |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsPing{
    Count:           models.ToPointer(10),
    Host:            "1.1.1.1",
}

apiResponse, err := utilitiesCommon.PingFromDevice(ctx, siteId, deviceId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Readopt Site Octerm Device

For the octerm devices, the device ID must come from fpc0. However, for a VC, the users may change the original fpc0 from CLI. To fix the issue, the readopt API could be used to trigger the readopt process so the device would get the corret device ID to connect the cloud.

```go
ReadoptSiteOctermDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := utilitiesCommon.ReadoptSiteOctermDevice(ctx, siteId, deviceId)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Reprovision Site Octerm Device

To force one device to reprovision itself again.

```go
ReprovisionSiteOctermDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := utilitiesCommon.ReprovisionSiteOctermDevice(ctx, siteId, deviceId)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Restart Site Device

Restart / Reboot a device

```go
RestartSiteDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsDevicesRestart) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsDevicesRestart`](../../doc/models/utils-devices-restart.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



resp, err := utilitiesCommon.RestartSiteDevice(ctx, siteId, deviceId, nil)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Start Site Locate Device

### Access Points

Locate an Access Point by blinking it's LED.
It is a persisted state that has to be stopped by calling Stop Locating API

### Switches

Locate a Switch by blinking all port LEDs.
By default, request is sent to `master` switch and LEDs will keep flashing for 5 minutes.
In case of virtual chassis (VC) the desired member mac has to be passed in the request payload.
At anypoint, only one VC member can be requested to flash the LED.
To stop LED flashing before the duration ends /unlocate API request can be made.
If /unlocate API is not called LED will continue to flash on device for the given duration.
Default duration is 5 minutes and 120 minutes is the maximum.

```go
StartSiteLocateDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.LocateSwitch) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.LocateSwitch`](../../doc/models/locate-switch.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.LocateSwitch{
    Duration: models.ToPointer(5),
    Mac:      models.ToPointer("f01c2d4ff760"),
}

resp, err := utilitiesCommon.StartSiteLocateDevice(ctx, siteId, deviceId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Stop Site Locate Device

Stop Locate a Device

```go
StopSiteLocateDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := utilitiesCommon.StopSiteLocateDevice(ctx, siteId, deviceId)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Traceroute From Device

Traceroute can be performed from the Device.

The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
    "subscribe": "/sites/{site_id}/devices/{device_id}/cmd"
}
```

```go
TracerouteFromDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsTraceroute) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsTraceroute`](../../doc/models/utils-traceroute.md) | Body, Optional | Request Body |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsTraceroute{
    Host:     models.ToPointer("string"),
    Port:     models.ToPointer(33434),
    Protocol: models.ToPointer(models.UtilsTracerouteProtocolEnum("udp")),
}

apiResponse, err := utilitiesCommon.TracerouteFromDevice(ctx, siteId, deviceId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Upload Site Device Support File

Support / Upload device support files

#### Info Param

**Parameter**|**Description**
:-------------: |:-------------: |:-------------:
process|Upload 1 file with output of show system processes extensive
outbound-ssh|Upload 1 file that concatenates all /var/log/outbound-ssh.log* files
messages|Upload 1 to 10 /var/log/messages* files
core-dumps|Upload all core dump files, if any
full|string|Upload 1 file with output of request support information, 1 file that concatenates all /var/log/outbound-ssh.log files, all core dump files, the 3 most recent /var/log/messages files, and Mist agent logs (for Junos devices running the Mist agent)
var-logs|Upload all non-empty files in the /var/log/ directory
jma-logs|Upload Mist agent logs (for Junos devices running the Mist agent only)

```go
UploadSiteDeviceSupportFile(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsSendSupportLogs) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsSendSupportLogs`](../../doc/models/utils-send-support-logs.md) | Body, Optional | Request Body |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsSendSupportLogs{
    Info:             models.ToPointer(models.UtilsSendSupportLogsInfoEnum("full")),
}

resp, err := utilitiesCommon.UploadSiteDeviceSupportFile(ctx, siteId, deviceId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Device not online | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

