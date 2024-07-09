# Sites Devices Stats

```go
sitesDevicesStats := client.SitesDevicesStats()
```

## Class Name

`SitesDevicesStats`

## Methods

* [Count Site Bgp Stats](../../doc/controllers/sites-devices-stats.md#count-site-bgp-stats)
* [Count Site Sw or Gw Ports](../../doc/controllers/sites-devices-stats.md#count-site-sw-or-gw-ports)
* [Count Site Switch Ports](../../doc/controllers/sites-devices-stats.md#count-site-switch-ports)
* [Get Site All Clients Stats by Device](../../doc/controllers/sites-devices-stats.md#get-site-all-clients-stats-by-device)
* [Get Site Device Stats](../../doc/controllers/sites-devices-stats.md#get-site-device-stats)
* [Get Site Gateway Metrics](../../doc/controllers/sites-devices-stats.md#get-site-gateway-metrics)
* [Get Site Mx Edge Stats](../../doc/controllers/sites-devices-stats.md#get-site-mx-edge-stats)
* [List Site Devices Stats](../../doc/controllers/sites-devices-stats.md#list-site-devices-stats)
* [List Site Mx Edges Stats](../../doc/controllers/sites-devices-stats.md#list-site-mx-edges-stats)
* [Search Site Bgp Stats](../../doc/controllers/sites-devices-stats.md#search-site-bgp-stats)
* [Search Site Sw or Gw Ports](../../doc/controllers/sites-devices-stats.md#search-site-sw-or-gw-ports)
* [Search Site Switch Ports](../../doc/controllers/sites-devices-stats.md#search-site-switch-ports)


# Count Site Bgp Stats

Count BGP Stats

```go
CountSiteBgpStats(
    ctx context.Context,
    siteId uuid.UUID,
    state *string,
    distinct *string) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `state` | `*string` | Query, Optional | - |
| `distinct` | `*string` | Query, Optional | - |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





apiResponse, err := sitesDevicesStats.CountSiteBgpStats(ctx, siteId, nil, nil)
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
  "distinct": "string",
  "end": 0,
  "limit": 0,
  "results": [
    {
      "count": 0,
      "property": "string"
    }
  ],
  "start": 0,
  "total": 0
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Count Site Sw or Gw Ports

Count by Distinct Attributes of Switch/Gateway Ports

```go
CountSiteSwOrGwPorts(
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
    stpState *models.CountSiteSwOrGwPortsStpStateEnum,
    stpRole *models.CountSiteSwOrGwPortsStpRoleEnum,
    authState *models.CountSiteSwOrGwPortsAuthStateEnum,
    up *bool,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SitePortsCountDistinctEnum`](../../doc/models/site-ports-count-distinct-enum.md) | Query, Optional | - |
| `fullDuplex` | `*bool` | Query, Optional | indicates full or half duplex |
| `mac` | `*string` | Query, Optional | device identifier |
| `neighborMac` | `*string` | Query, Optional | Chassis identifier of the chassis type listed |
| `neighborPortDesc` | `*string` | Query, Optional | Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” |
| `neighborSystemName` | `*string` | Query, Optional | Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” |
| `poeDisabled` | `*bool` | Query, Optional | is the POE configured not be disabled. |
| `poeMode` | `*string` | Query, Optional | poe mode depending on class E.g. “802.3at” |
| `poeOn` | `*bool` | Query, Optional | is the device attached to POE |
| `portId` | `*string` | Query, Optional | interface name |
| `portMac` | `*string` | Query, Optional | interface mac address |
| `powerDraw` | `*float64` | Query, Optional | Amount of power being used by the interface at the time the command is executed. Unit in watts. |
| `txPkts` | `*int` | Query, Optional | Output packets |
| `rxPkts` | `*int` | Query, Optional | Input packets |
| `rxBytes` | `*int` | Query, Optional | Input bytes |
| `txBps` | `*int` | Query, Optional | Output rate |
| `rxBps` | `*int` | Query, Optional | Input rate |
| `txMcastPkts` | `*int` | Query, Optional | Multicast output packets |
| `txBcastPkts` | `*int` | Query, Optional | Broadcast output packets |
| `rxMcastPkts` | `*int` | Query, Optional | Multicast input packets |
| `rxBcastPkts` | `*int` | Query, Optional | Broadcast input packets |
| `speed` | `*int` | Query, Optional | port speed |
| `stpState` | [`*models.CountSiteSwOrGwPortsStpStateEnum`](../../doc/models/count-site-sw-or-gw-ports-stp-state-enum.md) | Query, Optional | if `up`==`true` |
| `stpRole` | [`*models.CountSiteSwOrGwPortsStpRoleEnum`](../../doc/models/count-site-sw-or-gw-ports-stp-role-enum.md) | Query, Optional | if `up`==`true` |
| `authState` | [`*models.CountSiteSwOrGwPortsAuthStateEnum`](../../doc/models/count-site-sw-or-gw-ports-auth-state-enum.md) | Query, Optional | if `up`==`true` && has Authenticator role |
| `up` | `*bool` | Query, Optional | indicates if interface is up |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SitePortsCountDistinctEnum("mac")



















































page := 1

limit := 100





duration := "10m"

apiResponse, err := sitesDevicesStats.CountSiteSwOrGwPorts(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &page, &limit, nil, nil, &duration)
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
  "distinct": "string",
  "end": 0,
  "limit": 0,
  "results": [
    {
      "count": 0,
      "property": "string"
    }
  ],
  "start": 0,
  "total": 0
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Count Site Switch Ports

Count by Distinct Attributes of Switch/Gateway Ports

```go
CountSiteSwitchPorts(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteSwitchPortsCountDistinctEnum,
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
    stpState *models.CountSiteSwitchPortsStpStateEnum,
    stpRole *models.CountSiteSwitchPortsStpRoleEnum,
    authState *models.CountSiteSwitchPortsAuthStateEnum,
    up *bool,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteSwitchPortsCountDistinctEnum`](../../doc/models/site-switch-ports-count-distinct-enum.md) | Query, Optional | - |
| `fullDuplex` | `*bool` | Query, Optional | indicates full or half duplex |
| `mac` | `*string` | Query, Optional | device identifier |
| `neighborMac` | `*string` | Query, Optional | Chassis identifier of the chassis type listed |
| `neighborPortDesc` | `*string` | Query, Optional | Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” |
| `neighborSystemName` | `*string` | Query, Optional | Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” |
| `poeDisabled` | `*bool` | Query, Optional | is the POE configured not be disabled. |
| `poeMode` | `*string` | Query, Optional | poe mode depending on class E.g. “802.3at” |
| `poeOn` | `*bool` | Query, Optional | is the device attached to POE |
| `portId` | `*string` | Query, Optional | interface name |
| `portMac` | `*string` | Query, Optional | interface mac address |
| `powerDraw` | `*float64` | Query, Optional | Amount of power being used by the interface at the time the command is executed. Unit in watts. |
| `txPkts` | `*int` | Query, Optional | Output packets |
| `rxPkts` | `*int` | Query, Optional | Input packets |
| `rxBytes` | `*int` | Query, Optional | Input bytes |
| `txBps` | `*int` | Query, Optional | Output rate |
| `rxBps` | `*int` | Query, Optional | Input rate |
| `txMcastPkts` | `*int` | Query, Optional | Multicast output packets |
| `txBcastPkts` | `*int` | Query, Optional | Broadcast output packets |
| `rxMcastPkts` | `*int` | Query, Optional | Multicast input packets |
| `rxBcastPkts` | `*int` | Query, Optional | Broadcast input packets |
| `speed` | `*int` | Query, Optional | port speed |
| `stpState` | [`*models.CountSiteSwitchPortsStpStateEnum`](../../doc/models/count-site-switch-ports-stp-state-enum.md) | Query, Optional | if `up`==`true` |
| `stpRole` | [`*models.CountSiteSwitchPortsStpRoleEnum`](../../doc/models/count-site-switch-ports-stp-role-enum.md) | Query, Optional | if `up`==`true` |
| `authState` | [`*models.CountSiteSwitchPortsAuthStateEnum`](../../doc/models/count-site-switch-ports-auth-state-enum.md) | Query, Optional | if `up`==`true` |
| `up` | `*bool` | Query, Optional | indicates if interface is up |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteSwitchPortsCountDistinctEnum("mac")



















































page := 1

limit := 100





duration := "10m"

apiResponse, err := sitesDevicesStats.CountSiteSwitchPorts(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &page, &limit, nil, nil, &duration)
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
  "distinct": "string",
  "end": 0,
  "limit": 0,
  "results": [
    {
      "count": 0,
      "property": "string"
    }
  ],
  "start": 0,
  "total": 0
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site All Clients Stats by Device

Get wireless client stat by Device

```go
GetSiteAllClientsStatsByDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[[][]models.ClientWirelessStats],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[][]models.ClientWirelessStats`](../../doc/models/client-wireless-stats.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesDevicesStats.GetSiteAllClientsStatsByDevice(ctx, siteId, deviceId)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Device Stats

Get Site Device Stats Details

```go
GetSiteDeviceStats(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.StatsDevice2],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.StatsDevice2`](../../doc/models/containers/stats-device-2.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesDevicesStats.GetSiteDeviceStats(ctx, siteId, deviceId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    if r, ok := responseBody.AsApStats(); ok {
        fmt.Println("Value narrowed down to models.ApStats: ", *r)
    } else if r, ok := responseBody.AsSwitchStats(); ok {
        fmt.Println("Value narrowed down to models.SwitchStats: ", *r)
    } else if r, ok := responseBody.AsGatewayStats(); ok {
        fmt.Println("Value narrowed down to models.GatewayStats: ", *r)
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Gateway Metrics

Get Site Gateway Metrics

```go
GetSiteGatewayMetrics(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.GatewayMetrics],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.GatewayMetrics`](../../doc/models/gateway-metrics.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesDevicesStats.GetSiteGatewayMetrics(ctx, siteId)
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
  "config_success": 99.9,
  "version_compliance": {
    "major_version": {
      "SRX320": {
        "major_count": 0,
        "major_version": "19.4R2-S1.2"
      }
    },
    "score": 99.9,
    "type": "gateway"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Mx Edge Stats

Get One Site MxEdge Stats

```go
GetSiteMxEdgeStats(
    ctx context.Context,
    siteId uuid.UUID,
    mxedgeId uuid.UUID,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.MxedgeStats],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mxedgeId` | `uuid.UUID` | Template, Required | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.MxedgeStats`](../../doc/models/mxedge-stats.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mxedgeId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





duration := "10m"

apiResponse, err := sitesDevicesStats.GetSiteMxEdgeStats(ctx, siteId, mxedgeId, nil, nil, &duration)
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
  "cpu_stat": {
    "cpus": {
      "cpu0": {
        "idle": 89,
        "interrupt": 0,
        "system": 8,
        "usage": 10,
        "user": 1
      },
      "cpu1": {
        "idle": 81,
        "interrupt": 0,
        "system": 4,
        "usage": 18,
        "user": 13
      },
      "cpu2": {
        "idle": 81,
        "interrupt": 0,
        "system": 4,
        "usage": 18,
        "user": 13
      },
      "cpu3": {
        "idle": 2,
        "interrupt": 0,
        "system": 50,
        "usage": 97,
        "user": 46
      }
    },
    "idle": 62,
    "interrupt": 0,
    "system": 17,
    "usage": 37,
    "user": 19
  },
  "created_time": 1632684398,
  "for_site": false,
  "id": "00000000-0000-0000-1000-020000a80cb4",
  "ip_stat": {
    "ip": "192.168.1.244",
    "ips": {
      "ens18": "192.168.1.244/24,fe80::104c:ffff:fee0:caf8/64"
    },
    "macs": {
      "ens18": "e4434b217044"
    }
  },
  "lag_stat": {
    "lacp0": {
      "active_ports": [
        "port0",
        "port1"
      ]
    }
  },
  "last_seen": 1633721215,
  "mac": "020000a80cb4",
  "memory_stat": {
    "active": 394936320,
    "available": 4699291648,
    "buffers": 107646976,
    "cached": 478060544,
    "free": 4330659840,
    "inactive": 211980288,
    "swap_cached": 0,
    "swap_free": 1022357504,
    "swap_total": 1022357504,
    "total": 8365957120,
    "usage": 48
  },
  "model": "ME-VM",
  "modified_time": 1633643629,
  "mxagent_registered": true,
  "mxcluster_id": "678bc339-7635-4556-bbc0-e77ad493ef8b",
  "name": "me-vm-1",
  "num_tunnels": 0,
  "oob_ip_config": {
    "dns": [
      "8.8.8.8",
      "1.1.1.1"
    ],
    "gateway": "10.0.0.1",
    "ip": "10.0.0.10",
    "netmask": "255.255.255.0",
    "type": "static"
  },
  "org_id": "11b08247-b1ee-4152-9b25-312b323ce480",
  "port_stat": {
    "port0": {
      "full_duplex": true,
      "mac": "9e294e49091d",
      "rx_bytes": 646898375700,
      "rx_errors": 0,
      "rx_pkts": 8784449574,
      "speed": 10000,
      "state": "forwarding",
      "tx_bytes": 647200748038,
      "tx_errors": 0,
      "tx_pkts": 8788647466,
      "up": true
    },
    "port1": {
      "full_duplex": true,
      "mac": "a270fe53437e",
      "rx_bytes": 647200437652,
      "rx_errors": 0,
      "rx_pkts": 8788644886,
      "speed": 10000,
      "state": "forwarding",
      "tx_bytes": 646898681650,
      "tx_errors": 0,
      "tx_pkts": 8784452092,
      "up": true
    }
  },
  "sensor_stat": {},
  "serial": "string",
  "service_stat": {
    "mxagent": {
      "ext_ip": "99.0.86.164",
      "last_seen": 1633721215,
      "package_state": "Installed",
      "package_version": "3.1.1037-1",
      "running_state": "Running",
      "uptime": 21240
    },
    "tunterm": {
      "ext_ip": "99.0.86.164",
      "last_seen": 1633721203,
      "package_state": "Installed",
      "package_version": "0.1.2449+deb10",
      "running_state": "Running",
      "uptime": 76261
    }
  },
  "services": [
    "tunterm"
  ],
  "site_id": "00000000-0000-0000-0000-000000000000",
  "status": "connected",
  "tunterm_ip_config": {
    "gateway": "192.168.11.1",
    "ip": "192.168.11.91",
    "netmask": "255.255.255.0"
  },
  "tunterm_port_config": {
    "downstream_ports": [
      "0",
      "1"
    ],
    "separate_upstream_downstream": false,
    "upstream_ports": [
      "0",
      "1"
    ]
  },
  "tunterm_registered": true,
  "tunterm_stat": {
    "monitoring_failed": false
  },
  "uptime": 76281,
  "virtualization_type": "KVM"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Site Devices Stats

Get List of Site Devices Stats

```go
ListSiteDevicesStats(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.DeviceTypeWithAllEnum,
    status *models.StatDeviceStatusFilterEnum,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.ListSiteDevicesStatsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeWithAllEnum`](../../doc/models/device-type-with-all-enum.md) | Query, Optional | - |
| `status` | [`*models.StatDeviceStatusFilterEnum`](../../doc/models/stat-device-status-filter-enum.md) | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.ListSiteDevicesStatsResponse`](../../doc/models/containers/list-site-devices-stats-response.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeWithAllEnum("ap")

status := models.StatDeviceStatusFilterEnum("all")

page := 1

limit := 100

apiResponse, err := sitesDevicesStats.ListSiteDevicesStats(ctx, siteId, &mType, &status, &page, &limit)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    for _, item := range responseBody {
        if i, ok := item.AsApStats(); ok {
            fmt.Println("Value narrowed down to models.ApStats: ", *i)
        } else if i, ok := item.AsSwitchStats(); ok {
            fmt.Println("Value narrowed down to models.SwitchStats: ", *i)
        } else if i, ok := item.AsGatewayStats(); ok {
            fmt.Println("Value narrowed down to models.GatewayStats: ", *i)
        }
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
[
  {
    "ble_config": {
      "beacon_rate": 3,
      "beacon_rate_model": "custom",
      "beam_disabled": [
        1,
        3,
        6
      ],
      "power": 10,
      "power_mode": "custom"
    },
    "ble_stat": {
      "beacon_rate": 3,
      "eddystone_uid_enabled": false,
      "eddystone_uid_freq_msec": 200,
      "eddystone_uid_instance": "5c5b35000001",
      "eddystone_uid_namespace": "2818e3868dec25629ede",
      "eddystone_url_enabled": true,
      "eddystone_url_freq_msec": 100,
      "eddystone_url_url": "https://www.abc.com",
      "ibeacon_enabled": true,
      "ibeacon_major": 13,
      "ibeacon_minor": 138,
      "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
      "major": 12345,
      "minors": [
        201,
        202,
        203,
        204,
        205,
        206,
        207,
        208
      ],
      "power": 10,
      "rx_bytes": 135,
      "rx_pkts": 135,
      "tx_bytes": 5231513353,
      "tx_pkts": 135135135,
      "tx_resets": 0,
      "uuid": "ada72f8f-1643-e5c6-94db-f2a5636f1a64"
    },
    "cert_expiry": 1534534392,
    "ext_ip": "73.92.124.103",
    "fwupdate": {
      "progress": 10,
      "status": "inprogress",
      "status_id": 5,
      "timestamp": 1428949501
    },
    "iot_stat": {
      "DI2": {
        "value": 0
      }
    },
    "ip": "10.2.9.159",
    "ip_config": {
      "dns": [
        "8.8.8.8",
        "4.4.4.4"
      ],
      "dns_suffix": [
        ".mist.local",
        ".mist.com"
      ],
      "gateway": "10.2.1.254",
      "ip": "10.2.1.1",
      "netmask": "255.255.255.0",
      "type": "static"
    },
    "ip_stat": {
      "dns": [
        "8.8.8.8",
        "4.4.4.4"
      ],
      "dns_suffix": [
        ".mist.local",
        ".mist.com"
      ],
      "gateway": "10.2.1.254",
      "gateway6": "2607:f8b0:4005:808::1",
      "ip": "10.2.1.1",
      "ip6": "2607:f8b0:4005:808::2004",
      "ips": {
        "vlan1": "10.2.1.1/24,2607:f8b0:4005:808::1/32",
        "vlan193": "10.73.1.31/16",
        "vlan3157": "10.72.11.14/24"
      },
      "netmask": "255.255.255.0",
      "netmask6": "/32"
    },
    "l2tp_stat": {
      "7dae216d-7c98-a51b-e068-dd7d477b7216": {
        "sessions": [
          {
            "local_sid": 31,
            "remote_id": "vpn1",
            "remote_sid": 13,
            "state": "established"
          }
        ],
        "state": "established_with_sessions",
        "uptime": 135,
        "wxtunnel_id": "7dae216d-7c98-a51b-e068-dd7d477b7216"
      }
    },
    "last_seen": 1470417522,
    "last_trouble": {
      "code": "03",
      "timestamp": 1428949501
    },
    "led": {
      "brightness": 255,
      "enabled": true
    },
    "lldp_stat": {
      "chassis_id": "63:68:61:73:73:69",
      "lldp_med_supported": false,
      "mgmt_addr": "10.1.5.2",
      "port_desc": "2/26",
      "power_allocated": 15500,
      "power_draw": 15000,
      "power_request_count": 3,
      "power_requested": 25500,
      "system_desc": "HP J9729A 2920-48G-POE+ Switch",
      "system_name": "TC2-OWL-Stack-01"
    },
    "locating": false,
    "mac": "5c5b35000010",
    "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
    "mesh_downlinks": {
      "00000000-0000-0000-1000-5c5b356be59f": {
        "band": "24",
        "channel": 7,
        "idle_time": 3,
        "last_seen": 1470417522,
        "proto": "a",
        "rssi": -65,
        "rx_bps": 12,
        "rx_bytes": 217416,
        "rx_packets": 2337,
        "rx_rate": 65,
        "rx_retries": 5,
        "snr": 31,
        "tx_bps": 6,
        "tx_bytes": 175132,
        "tx_packets": 1566,
        "tx_rate": 65,
        "tx_retries": 500
      }
    },
    "mesh_uplink": {
      "band": "24",
      "channel": 7,
      "idle_time": 3,
      "last_seen": 1470417522,
      "proto": "a",
      "rssi": -65,
      "rx_bps": 12,
      "rx_bytes": 217416,
      "rx_packets": 2337,
      "rx_rate": 65,
      "rx_retries": 5,
      "snr": 31,
      "tx_bps": 6,
      "tx_bytes": 175132,
      "tx_packets": 1566,
      "tx_rate": 65,
      "tx_retries": 500,
      "uplink_ap_id": "00000000-0000-0000-1000-5c5b35000010"
    },
    "model": "AP200",
    "name": "conference room",
    "num_clients": 10,
    "port_stat": {
      "eth0": {
        "full_duplex": true,
        "rx_bytes": 2056,
        "rx_errors": 0,
        "rx_pkts": 670,
        "speed": 1000,
        "tx_bytes": 2056,
        "tx_pkts": 670,
        "up": true
      },
      "eth1": {
        "up": false
      },
      "module": {
        "up": false
      }
    },
    "power_budget": -12000,
    "power_src": "PoE 802.3af",
    "radio_config": {
      "band_24": {
        "bandwidth": 20,
        "channel": 0,
        "dynamic_chaining_enabled": false,
        "power": 0,
        "rx_chain": 4,
        "tx_chain": 4
      },
      "band_5": {
        "bandwidth": 40,
        "channel": 0,
        "dynamic_chaining_enabled": false,
        "power": 0,
        "rx_chain": 4,
        "tx_chain": 1
      },
      "scanning_enabled": true
    },
    "radio_stat": {
      "band_24": {
        "bandwidth": 20,
        "channel": 6,
        "mac": "5c5b350004a0",
        "num_clients": 6,
        "power": 19,
        "rx_bytes": 8504737800,
        "rx_pkts": 57731964,
        "tx_bytes": 211166512114,
        "tx_pkts": 812058566
      },
      "band_5": {
        "bandwidth": 80,
        "channel": 44,
        "mac": "5c5b350004b0",
        "num_clients": 4,
        "power": 15,
        "rx_bytes": 10366616,
        "rx_pkts": 38603,
        "tx_bytes": 50877568,
        "tx_pkts": 145496
      }
    },
    "rx_bps": 60003,
    "rx_bytes": 8515104416,
    "rx_pkts": 57770567,
    "serial": "FXLH2015170017",
    "status": "connected",
    "tx_bps": 634301,
    "tx_bytes": 211217389682,
    "tx_pkts": 812204062,
    "type": "ap",
    "uptime": 13500,
    "usb_stat": {
      "channel": 3,
      "connected": true,
      "last_activity": 1586873254,
      "type": "imagotag",
      "up": true
    },
    "version": "1.0.0",
    "x": 53.5,
    "y": 173.1
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Site Mx Edges Stats

Get List of Site MxEdges Stats

```go
ListSiteMxEdgesStats(
    ctx context.Context,
    siteId uuid.UUID,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[[]models.MxedgeStats],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`[]models.MxedgeStats`](../../doc/models/mxedge-stats.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100





duration := "10m"

apiResponse, err := sitesDevicesStats.ListSiteMxEdgesStats(ctx, siteId, &page, &limit, nil, nil, &duration)
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
[
  {
    "cpu_stat": {
      "cpus": {
        "cpu0": {
          "idle": 89,
          "interrupt": 0,
          "system": 8,
          "usage": 10,
          "user": 1
        },
        "cpu1": {
          "idle": 81,
          "interrupt": 0,
          "system": 4,
          "usage": 18,
          "user": 13
        },
        "cpu2": {
          "idle": 81,
          "interrupt": 0,
          "system": 4,
          "usage": 18,
          "user": 13
        },
        "cpu3": {
          "idle": 2,
          "interrupt": 0,
          "system": 50,
          "usage": 97,
          "user": 46
        }
      },
      "idle": 62,
      "interrupt": 0,
      "system": 17,
      "usage": 37,
      "user": 19
    },
    "created_time": 1632684398,
    "for_site": false,
    "id": "00000000-0000-0000-1000-020000a80cb4",
    "ip_stat": {
      "ip": "192.168.1.244",
      "ips": {
        "ens18": "192.168.1.244/24,fe80::104c:ffff:fee0:caf8/64"
      },
      "macs": {
        "ens18": "e4434b217044"
      }
    },
    "lag_stat": {
      "lacp0": {
        "active_ports": [
          "port0",
          "port1"
        ]
      }
    },
    "last_seen": 1633721215,
    "mac": "020000a80cb4",
    "memory_stat": {
      "active": 394936320,
      "available": 4699291648,
      "buffers": 107646976,
      "cached": 478060544,
      "free": 4330659840,
      "inactive": 211980288,
      "swap_cached": 0,
      "swap_free": 1022357504,
      "swap_total": 1022357504,
      "total": 8365957120,
      "usage": 48
    },
    "model": "ME-VM",
    "modified_time": 1633643629,
    "mxagent_registered": true,
    "mxcluster_id": "678bc339-7635-4556-bbc0-e77ad493ef8b",
    "name": "me-vm-1",
    "num_tunnels": 0,
    "oob_ip_config": {
      "dns": [
        "8.8.8.8",
        "1.1.1.1"
      ],
      "gateway": "10.0.0.1",
      "ip": "10.0.0.10",
      "netmask": "255.255.255.0",
      "type": "static"
    },
    "org_id": "11b08247-b1ee-4152-9b25-312b323ce480",
    "port_stat": {
      "port0": {
        "full_duplex": true,
        "mac": "9e294e49091d",
        "rx_bytes": 646898375700,
        "rx_errors": 0,
        "rx_pkts": 8784449574,
        "speed": 10000,
        "state": "forwarding",
        "tx_bytes": 647200748038,
        "tx_errors": 0,
        "tx_pkts": 8788647466,
        "up": true
      },
      "port1": {
        "full_duplex": true,
        "mac": "a270fe53437e",
        "rx_bytes": 647200437652,
        "rx_errors": 0,
        "rx_pkts": 8788644886,
        "speed": 10000,
        "state": "forwarding",
        "tx_bytes": 646898681650,
        "tx_errors": 0,
        "tx_pkts": 8784452092,
        "up": true
      }
    },
    "sensor_stat": {},
    "serial": "string",
    "service_stat": {
      "mxagent": {
        "ext_ip": "99.0.86.164",
        "last_seen": 1633721215,
        "package_state": "Installed",
        "package_version": "3.1.1037-1",
        "running_state": "Running",
        "uptime": 21240
      },
      "tunterm": {
        "ext_ip": "99.0.86.164",
        "last_seen": 1633721203,
        "package_state": "Installed",
        "package_version": "0.1.2449+deb10",
        "running_state": "Running",
        "uptime": 76261
      }
    },
    "services": [
      "tunterm"
    ],
    "site_id": "00000000-0000-0000-0000-000000000000",
    "status": "connected",
    "tunterm_ip_config": {
      "gateway": "192.168.11.1",
      "ip": "192.168.11.91",
      "netmask": "255.255.255.0"
    },
    "tunterm_port_config": {
      "downstream_ports": [
        "0",
        "1"
      ],
      "separate_upstream_downstream": false,
      "upstream_ports": [
        "0",
        "1"
      ]
    },
    "tunterm_registered": true,
    "tunterm_stat": {
      "monitoring_failed": false
    },
    "uptime": 76281,
    "virtualization_type": "KVM"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Search Site Bgp Stats

Search BGP Stats

```go
SearchSiteBgpStats(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.ResponseSearchBgps],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.ResponseSearchBgps`](../../doc/models/response-search-bgps.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesDevicesStats.SearchSiteBgpStats(ctx, siteId)
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
  "end": 0,
  "limit": 0,
  "results": [
    {
      "evpn_overlay": true,
      "for_overlay": true,
      "local_as": 65000,
      "mac": "020001c04668",
      "neighbor": "15.8.3.5",
      "neighbor_as": 65000,
      "neighbor_mac": "020001c04600",
      "node": "node0",
      "org_id": "0c160b7f-1027-4cd1-923b-744534c4b070",
      "rx_pkts": 63366,
      "rx_routes": 60,
      "site_id": "725a8d34-a126-4f2c-b990-d1219421cb75",
      "state": "established",
      "timestamp": 1666251056.07,
      "tx_pkts": 1735,
      "tx_routes": 60,
      "up": true,
      "uptime": 31355,
      "vrf_name": "default"
    }
  ],
  "start": 0,
  "total": 0
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Search Site Sw or Gw Ports

Search Switch / Gateway Ports

```go
SearchSiteSwOrGwPorts(
    ctx context.Context,
    siteId uuid.UUID,
    fullDuplex *bool,
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
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `fullDuplex` | `*bool` | Query, Optional | indicates full or half duplex |
| `mac` | `*string` | Query, Optional | device identifier |
| `deviceType` | [`*models.SearchSiteSwOrGwPortsDeviceTypeEnum`](../../doc/models/search-site-sw-or-gw-ports-device-type-enum.md) | Query, Optional | device type |
| `neighborMac` | `*string` | Query, Optional | Chassis identifier of the chassis type listed |
| `neighborPortDesc` | `*string` | Query, Optional | Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” |
| `neighborSystemName` | `*string` | Query, Optional | Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” |
| `poeDisabled` | `*bool` | Query, Optional | is the POE configured not be disabled. |
| `poeMode` | `*string` | Query, Optional | poe mode depending on class E.g. “802.3at” |
| `poeOn` | `*bool` | Query, Optional | is the device attached to POE |
| `portId` | `*string` | Query, Optional | interface name |
| `portMac` | `*string` | Query, Optional | interface mac address |
| `powerDraw` | `*float64` | Query, Optional | Amount of power being used by the interface at the time the command is executed. Unit in watts. |
| `txPkts` | `*int` | Query, Optional | Output packets |
| `rxPkts` | `*int` | Query, Optional | Input packets |
| `rxBytes` | `*int` | Query, Optional | Input bytes |
| `txBps` | `*int` | Query, Optional | Output rate |
| `rxBps` | `*int` | Query, Optional | Input rate |
| `txErrors` | `*int` | Query, Optional | Output errors |
| `rxErrors` | `*int` | Query, Optional | Input errors |
| `txMcastPkts` | `*int` | Query, Optional | Multicast output packets |
| `txBcastPkts` | `*int` | Query, Optional | Broadcast output packets |
| `rxMcastPkts` | `*int` | Query, Optional | Multicast input packets |
| `rxBcastPkts` | `*int` | Query, Optional | Broadcast input packets |
| `speed` | `*int` | Query, Optional | port speed |
| `macLimit` | `*int` | Query, Optional | Limit on number of dynamically learned macs |
| `macCount` | `*int` | Query, Optional | Number of mac addresses in the forwarding table |
| `up` | `*bool` | Query, Optional | indicates if interface is up |
| `active` | `*bool` | Query, Optional | indicates if interface is active/inactive |
| `jitter` | `*float64` | Query, Optional | Last sampled jitter of the interface |
| `loss` | `*float64` | Query, Optional | Last sampled loss of the interface |
| `latency` | `*float64` | Query, Optional | Last sampled latency of the interface |
| `stpState` | [`*models.SearchSiteSwOrGwPortsStpStateEnum`](../../doc/models/search-site-sw-or-gw-ports-stp-state-enum.md) | Query, Optional | if `up`==`true` |
| `stpRole` | [`*models.SearchSiteSwOrGwPortsStpRoleEnum`](../../doc/models/search-site-sw-or-gw-ports-stp-role-enum.md) | Query, Optional | if `up`==`true` |
| `xcvrPartNumber` | `*string` | Query, Optional | Optic Slot Partnumber, Check for null/empty |
| `authState` | [`*models.SearchSiteSwOrGwPortsAuthStateEnum`](../../doc/models/search-site-sw-or-gw-ports-auth-state-enum.md) | Query, Optional | if `up`==`true` && has Authenticator role |
| `lteImsi` | `*string` | Query, Optional | LTE IMSI value, Check for null/empty |
| `lteIccid` | `*string` | Query, Optional | LTE ICCID value, Check for null/empty |
| `lteImei` | `*string` | Query, Optional | LTE IMEI value, Check for null/empty |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseSwitchPortSearch`](../../doc/models/response-switch-port-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")













































































limit := 100





duration := "10m"

apiResponse, err := sitesDevicesStats.SearchSiteSwOrGwPorts(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
  "end": 1513177200,
  "limit": 10,
  "results": [
    {
      "active": true,
      "auth_state": "init",
      "for_site": true,
      "full_duplex": true,
      "jitter": 0,
      "latency": 0,
      "loss": 0,
      "lte_iccid": "string",
      "lte_imei": "string",
      "lte_imsi": "string",
      "mac": "5c4527a96580",
      "mac_count": 0,
      "mac_limit": 0,
      "neighbor_mac": "64d814353400",
      "neighbor_port_desc": "GigabitEthernet1/0/21",
      "neighbor_system_name": "CORP-D-SW-2",
      "org_id": "c168ddee-c14c-11e5-8e81-1258369c38a9",
      "poe_disabled": true,
      "poe_mode": "802.3af",
      "poe_on": true,
      "port_id": "ge-0/0/0",
      "port_mac": "5c4527a96580",
      "port_usage": "lan",
      "power_draw": 0,
      "rx_bcast_pkts": 0,
      "rx_bps": 0,
      "rx_bytes": 4563443626,
      "rx_errors": 0,
      "rx_mcast_pkts": 0,
      "rx_pkts": 0,
      "site_id": "c1698122-c14c-11e5-8e81-1258369c38a9",
      "speed": 1000,
      "stp_role": "designated",
      "stp_state": "forwarding",
      "tx_bcast_pkts": 0,
      "tx_bps": 0,
      "tx_bytes": 11299516780,
      "tx_errors": 0,
      "tx_mcast_pkts": 0,
      "tx_pkts": 492176,
      "type": "gateway",
      "up": true,
      "xcvr_part_number": "string"
    }
  ],
  "start": 1511967600,
  "total": 100
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Search Site Switch Ports

Search Switch / Gateway Ports

```go
SearchSiteSwitchPorts(
    ctx context.Context,
    siteId uuid.UUID,
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
    stpState *models.SearchSiteSwitchPortsStpStateEnum,
    stpRole *models.SearchSiteSwitchPortsStpRoleEnum,
    authState *models.SearchSiteSwitchPortsAuthStateEnum,
    up *bool,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseSwitchPortSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `fullDuplex` | `*bool` | Query, Optional | indicates full or half duplex |
| `mac` | `*string` | Query, Optional | device identifier |
| `neighborMac` | `*string` | Query, Optional | Chassis identifier of the chassis type listed |
| `neighborPortDesc` | `*string` | Query, Optional | Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” |
| `neighborSystemName` | `*string` | Query, Optional | Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” |
| `poeDisabled` | `*bool` | Query, Optional | is the POE configured not be disabled. |
| `poeMode` | `*string` | Query, Optional | poe mode depending on class E.g. “802.3at” |
| `poeOn` | `*bool` | Query, Optional | is the device attached to POE |
| `portId` | `*string` | Query, Optional | interface name |
| `portMac` | `*string` | Query, Optional | interface mac address |
| `powerDraw` | `*float64` | Query, Optional | Amount of power being used by the interface at the time the command is executed. Unit in watts. |
| `txPkts` | `*int` | Query, Optional | Output packets |
| `rxPkts` | `*int` | Query, Optional | Input packets |
| `rxBytes` | `*int` | Query, Optional | Input bytes |
| `txBps` | `*int` | Query, Optional | Output rate |
| `rxBps` | `*int` | Query, Optional | Input rate |
| `txMcastPkts` | `*int` | Query, Optional | Multicast output packets |
| `txBcastPkts` | `*int` | Query, Optional | Broadcast output packets |
| `rxMcastPkts` | `*int` | Query, Optional | Multicast input packets |
| `rxBcastPkts` | `*int` | Query, Optional | Broadcast input packets |
| `speed` | `*int` | Query, Optional | port speed |
| `stpState` | [`*models.SearchSiteSwitchPortsStpStateEnum`](../../doc/models/search-site-switch-ports-stp-state-enum.md) | Query, Optional | if `up`==`true` |
| `stpRole` | [`*models.SearchSiteSwitchPortsStpRoleEnum`](../../doc/models/search-site-switch-ports-stp-role-enum.md) | Query, Optional | if `up`==`true` |
| `authState` | [`*models.SearchSiteSwitchPortsAuthStateEnum`](../../doc/models/search-site-switch-ports-auth-state-enum.md) | Query, Optional | if `up`==`true` && has Authenticator role |
| `up` | `*bool` | Query, Optional | indicates if interface is up |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseSwitchPortSearch`](../../doc/models/response-switch-port-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



















































limit := 100





duration := "10m"

apiResponse, err := sitesDevicesStats.SearchSiteSwitchPorts(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
  "end": 1513177200,
  "limit": 10,
  "results": [
    {
      "active": true,
      "auth_state": "init",
      "for_site": true,
      "full_duplex": true,
      "jitter": 0,
      "latency": 0,
      "loss": 0,
      "lte_iccid": "string",
      "lte_imei": "string",
      "lte_imsi": "string",
      "mac": "5c4527a96580",
      "mac_count": 0,
      "mac_limit": 0,
      "neighbor_mac": "64d814353400",
      "neighbor_port_desc": "GigabitEthernet1/0/21",
      "neighbor_system_name": "CORP-D-SW-2",
      "org_id": "c168ddee-c14c-11e5-8e81-1258369c38a9",
      "poe_disabled": true,
      "poe_mode": "802.3af",
      "poe_on": true,
      "port_id": "ge-0/0/0",
      "port_mac": "5c4527a96580",
      "port_usage": "lan",
      "power_draw": 0,
      "rx_bcast_pkts": 0,
      "rx_bps": 0,
      "rx_bytes": 4563443626,
      "rx_errors": 0,
      "rx_mcast_pkts": 0,
      "rx_pkts": 0,
      "site_id": "c1698122-c14c-11e5-8e81-1258369c38a9",
      "speed": 1000,
      "stp_role": "designated",
      "stp_state": "forwarding",
      "tx_bcast_pkts": 0,
      "tx_bps": 0,
      "tx_bytes": 11299516780,
      "tx_errors": 0,
      "tx_mcast_pkts": 0,
      "tx_pkts": 492176,
      "type": "gateway",
      "up": true,
      "xcvr_part_number": "string"
    }
  ],
  "start": 1511967600,
  "total": 100
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

