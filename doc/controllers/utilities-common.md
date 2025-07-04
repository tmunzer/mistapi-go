# Utilities Common

```go
utilitiesCommon := client.UtilitiesCommon()
```

## Class Name

`UtilitiesCommon`

## Methods

* [Arp From Device](../../doc/controllers/utilities-common.md#arp-from-device)
* [Bounce Device Port](../../doc/controllers/utilities-common.md#bounce-device-port)
* [Clear Site Device Mac Table](../../doc/controllers/utilities-common.md#clear-site-device-mac-table)
* [Clear Site Device Policy Hit Count](../../doc/controllers/utilities-common.md#clear-site-device-policy-hit-count)
* [Create Site Device Shell Session](../../doc/controllers/utilities-common.md#create-site-device-shell-session)
* [Get Site Device Config Cmd](../../doc/controllers/utilities-common.md#get-site-device-config-cmd)
* [Get Site Device Ztp Password](../../doc/controllers/utilities-common.md#get-site-device-ztp-password)
* [Monitor Site Device Traffic](../../doc/controllers/utilities-common.md#monitor-site-device-traffic)
* [Ping From Device](../../doc/controllers/utilities-common.md#ping-from-device)
* [Readopt Site Octerm Device](../../doc/controllers/utilities-common.md#readopt-site-octerm-device)
* [Release Site Device Dhcp Lease](../../doc/controllers/utilities-common.md#release-site-device-dhcp-lease)
* [Reprovision Site Octerm Device](../../doc/controllers/utilities-common.md#reprovision-site-octerm-device)
* [Restart Site Device](../../doc/controllers/utilities-common.md#restart-site-device)
* [Restart Site Multiple Devices](../../doc/controllers/utilities-common.md#restart-site-multiple-devices)
* [Show Site Device Arp Table](../../doc/controllers/utilities-common.md#show-site-device-arp-table)
* [Show Site Device Bgp Summary](../../doc/controllers/utilities-common.md#show-site-device-bgp-summary)
* [Show Site Device Dhcp Leases](../../doc/controllers/utilities-common.md#show-site-device-dhcp-leases)
* [Show Site Device Evpn Database](../../doc/controllers/utilities-common.md#show-site-device-evpn-database)
* [Show Site Device Forwarding Table](../../doc/controllers/utilities-common.md#show-site-device-forwarding-table)
* [Show Site Device Mac Table](../../doc/controllers/utilities-common.md#show-site-device-mac-table)
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSession](../../doc/models/websocket-session.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.HaClusterNode{
    Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Bounce Device Port

Port Bounce can be performed from Switch/Gateway.

**Note:** Ports starting with vme, ae, irb, and HA control ports (for SSR only) are not supported

The output will be available through websocket. As there can be multiple command issued against the same AP at the same time and the output all goes through the same websocket stream, session is introduced for demux.

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
        "raw": "Port bounce complete."
    }
}
```

```go
BounceDevicePort(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsBouncePort) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsBouncePort`](../../doc/models/utils-bounce-port.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsBouncePort{
    Ports:                []string{
        "ge-0/0/0",
        "ge-0/0/1",
    },
}

resp, err := utilitiesCommon.BounceDevicePort(ctx, siteId, deviceId, &body)
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
| `body` | [`*models.UtilsMacTable`](../../doc/models/utils-mac-table.md) | Body, Optional | All attributes are optional |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSession](../../doc/models/websocket-session.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsMacTable{
    MacAddress:           models.ToPointer("f8c1165c6400"),
    PortId:               models.ToPointer("ge-0/0/0.0"),
    VlanId:               models.ToPointer("ge-0/0/0.0"),
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Clear Site Device Policy Hit Count

Clear application policy hit counts for all the policies

```go
ClearSiteDevicePolicyHitCount(
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSessionWithUrl](../../doc/models/websocket-session-with-url.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := utilitiesCommon.ClearSiteDevicePolicyHitCount(ctx, siteId, deviceId)
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


# Create Site Device Shell Session

Create Shell Session

```go
CreateSiteDeviceShellSession(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.ShellNode) (
    models.ApiResponse[models.WebsocketSessionWithUrl],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.ShellNode`](../../doc/models/shell-node.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSessionWithUrl](../../doc/models/websocket-session-with-url.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



apiResponse, err := utilitiesCommon.CreateSiteDeviceShellSession(ctx, siteId, deviceId, nil)
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
| `sort` | `*bool` | Query, Optional | Make output cmds sorted (for better readability) or not.<br><br>**Default**: `false` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceConfigCli](../../doc/models/response-device-config-cli.md).

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Site Device Ztp Password

In the case where something happens during/after ZTP, the root-password is modified (required for ZTP to set up outbound-ssh) but the user-defined password config has not be configured. This API can be used to retrieve the temporary password.

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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.RootPasswordString](../../doc/models/root-password-string.md).

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Monitor Site Device Traffic

Monitor traffic on switches and SRX.

* JUNOS uses cmd "monitor interface <port>" to monitor traffic on particular <port>
* JUNOS uses cmd "monitor interface traffic" to monitor traffic on all ports

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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSessionWithUrl](../../doc/models/websocket-session-with-url.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsMonitorTraffic{
    Port:                 models.ToPointer("ge-0/0/1"),
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSession](../../doc/models/websocket-session.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsPing{
    Count:                models.ToPointer(10),
    Host:                 "1.1.1.1",
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Readopt Site Octerm Device

For the octerm devices, the device ID must come from fpc0. However, for a VC, the users may change the original fpc0 from CLI. To fix the issue, the readopt API could be used to trigger the readopt process so the device would get the correct device ID to connect the cloud.

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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Release Site Device Dhcp Lease

Releases an active DHCP lease.

```go
ReleaseSiteDeviceDhcpLease(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsReleaseDhcpLeases) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsReleaseDhcpLeases`](../../doc/models/utils-release-dhcp-leases.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsReleaseDhcpLeases{
    PortId:               "ge-0/0/1.10",
}

resp, err := utilitiesCommon.ReleaseSiteDeviceDhcpLease(ctx, siteId, deviceId, &body)
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Restart Site Multiple Devices

Note that only the devices that are connected will be restarted.

```go
RestartSiteMultipleDevices(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.UtilsDevicesRestartMulti) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsDevicesRestartMulti`](../../doc/models/utils-devices-restart-multi.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsDevicesRestartMulti{
    DeviceIds:            []uuid.UUID{
        uuid.MustParse("00000000-0000-0000-1000-5c5b35584a6f"),
        uuid.MustParse("00000000-0000-0000-1000-5c5b350ea3b3"),
    },
}

resp, err := utilitiesCommon.RestartSiteMultipleDevices(ctx, siteId, &body)
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


# Show Site Device Arp Table

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
ShowSiteDeviceArpTable(
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
| `body` | [`*models.UtilsShowArp`](../../doc/models/utils-show-arp.md) | Body, Optional | All attributes are optional |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSession](../../doc/models/websocket-session.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowArp{
    Duration:             models.ToPointer(0),
    Interval:             models.ToPointer(0),
    Ip:                   models.ToPointer("192.168.30.7"),
    PortId:               models.ToPointer("ge-0/0/0.0"),
    Vrf:                  models.ToPointer("guest"),
}

apiResponse, err := utilitiesCommon.ShowSiteDeviceArpTable(ctx, siteId, deviceId, &body)
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


# Show Site Device Bgp Summary

Get BGP Summary from SSR, SRX and Switch.

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
ShowSiteDeviceBgpSummary(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowBgpSummary) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsShowBgpSummary`](../../doc/models/utils-show-bgp-summary.md) | Body, Optional | All attributes are optional |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSession](../../doc/models/websocket-session.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



apiResponse, err := utilitiesCommon.ShowSiteDeviceBgpSummary(ctx, siteId, deviceId, nil)
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


# Show Site Device Dhcp Leases

Shows DHCP leases

```go
ShowSiteDeviceDhcpLeases(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.UtilsShowDhcpLeases) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UtilsShowDhcpLeases`](../../doc/models/utils-show-dhcp-leases.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSession](../../doc/models/websocket-session.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowDhcpLeases{
    Network:              "guest",
}

apiResponse, err := utilitiesCommon.ShowSiteDeviceDhcpLeases(ctx, siteId, deviceId, &body)
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


# Show Site Device Evpn Database

Get EVPN Database from the Device. The output will be available through websocket.

```go
ShowSiteDeviceEvpnDatabase(
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
| `body` | [`*models.UtilsShowEvpnDatabase`](../../doc/models/utils-show-evpn-database.md) | Body, Optional | All attributes are optional |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSession](../../doc/models/websocket-session.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowEvpnDatabase{
    Duration:             models.ToPointer(0),
    Interval:             models.ToPointer(0),
    Mac:                  models.ToPointer("f8c1165c6400"),
    PortId:               models.ToPointer("ge-0/0/0.0"),
}

apiResponse, err := utilitiesCommon.ShowSiteDeviceEvpnDatabase(ctx, siteId, deviceId, &body)
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


# Show Site Device Forwarding Table

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
ShowSiteDeviceForwardingTable(
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
| `body` | [`*models.UtilsShowForwardingTable`](../../doc/models/utils-show-forwarding-table.md) | Body, Optional | All attributes are optional |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSession](../../doc/models/websocket-session.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsShowForwardingTable{
    Prefix:               models.ToPointer("3.1.1.0/24"),
    ServiceIp:            models.ToPointer("3.1.1.10"),
    ServiceName:          models.ToPointer("internet-wan_and_lte"),
    ServicePort:          models.ToPointer(32768),
    ServiceProtocol:      models.ToPointer("udp"),
    ServiceTenant:        models.ToPointer("branch1-wifi-mgt"),
    Vrf:                  models.ToPointer("guest"),
}

apiResponse, err := utilitiesCommon.ShowSiteDeviceForwardingTable(ctx, siteId, deviceId, &body)
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


# Show Site Device Mac Table

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
ShowSiteDeviceMacTable(
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
| `body` | [`*models.UtilsMacTable`](../../doc/models/utils-mac-table.md) | Body, Optional | All attributes are optional |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSession](../../doc/models/websocket-session.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsMacTable{
    MacAddress:           models.ToPointer("f8c1165c6400"),
    PortId:               models.ToPointer("ge-0/0/0.0"),
    VlanId:               models.ToPointer("ge-0/0/0.0"),
}

apiResponse, err := utilitiesCommon.ShowSiteDeviceMacTable(ctx, siteId, deviceId, &body)
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.LocateSwitch{
    Duration:             models.ToPointer(5),
    Mac:                  models.ToPointer("f01c2d4ff760"),
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Traceroute From Device

Traceroute can be performed from the Device.

The output will be available through websocket. As there can be multiple command issued against the same Device at the same time and the output all goes through the same websocket stream, session is introduced for demux.

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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSession](../../doc/models/websocket-session.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsTraceroute{
    Host:                 models.ToPointer("string"),
    Port:                 models.ToPointer(33434),
    Protocol:             models.ToPointer(models.UtilsTracerouteProtocolEnum_UDP),
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Upload Site Device Support File

Support / Upload device support files

#### Info Param

| Name | Type | Description |
| --- | --- | --- |
| process | string | Upload 1 file with output of show system processes extensive |
| outbound-ssh | string | Upload 1 file that concatenates all /var/log/outbound-ssh.log* files |
| messages | string | Upload 1 to 10 /var/log/messages* files |
| core-dumps | string | Upload all core dump files, if any |
| full | string | Upload 1 file with output of request support information, 1 file that concatenates all /var/log/outbound-ssh.log files, all core dump files, the 3 most recent /var/log/messages files, and Mist agent logs (for Junos devices running the Mist agent) |
| var-logs | string | Upload all non-empty files in the /var/log/ directory |
| jma-logs | string | Upload Mist agent logs (for Junos devices running the Mist agent only) |
"

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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UtilsSendSupportLogs{
    Info:                 models.ToPointer(models.UtilsSendSupportLogsInfoEnum_FULL),
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

