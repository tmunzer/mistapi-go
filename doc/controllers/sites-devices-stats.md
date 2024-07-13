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
    models.ApiResponse[models.GetSiteDeviceStatsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.GetSiteDeviceStatsResponse`](../../doc/models/containers/get-site-device-stats-response.md)

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

## Example Response *(as JSON)*

```json
{
  "auto_upgrade_stat": {
    "lastcheck": 1720594762
  },
  "ble_stat": {
    "beacon_enabled": true,
    "beacon_rate": 4,
    "eddystone_uid_enabled": false,
    "eddystone_uid_freq_msec": 1000,
    "eddystone_uid_instance": "5c5b35d0077b",
    "eddystone_uid_namespace": "9777c1a06ef611e68bbf",
    "eddystone_url_enabled": false,
    "eddystone_url_freq_msec": 1000,
    "eddystone_url_url": "",
    "ibeacon_enabled": false,
    "ibeacon_freq_msec": 1000,
    "ibeacon_major": 894,
    "ibeacon_minor": 9328,
    "ibeacon_uuid": "af010e2b-f829-4975-b49e-2e896ed1d627",
    "major": 894,
    "minors": [
      9328,
      9329,
      9330,
      9331,
      9332,
      9333,
      9334,
      9335,
      -1
    ],
    "power": 8,
    "rx_bytes": 158500843,
    "rx_pkts": 3549163,
    "tx_bytes": 509640,
    "tx_pkts": 85411,
    "tx_resets": 0,
    "uuid": "af010e2b-f829-4975-b49e-2e896ed1d627"
  },
  "config_reverted": false,
  "cpu_system": 21921854,
  "cpu_user": 7496631,
  "cpu_util": 5,
  "created_time": 1718228350,
  "env_stat": {
    "accel_x": -0.092,
    "accel_y": 0.004,
    "accel_z": -1.02,
    "ambient_temp": 43,
    "attitude": 0,
    "cpu_temp": 53,
    "humidity": 9,
    "magne_x": 0,
    "magne_y": 0,
    "magne_z": 0,
    "pressure": 968,
    "vcore_voltage": 0
  },
  "ext_ip": "66.129.234.28",
  "hw_rev": "C02",
  "id": "00000000-0000-0000-1000-5c5b35d0077b",
  "inactive_wired_vlans": [],
  "ip": "192.168.95.3",
  "ip_stat": {
    "dhcp_server": "192.168.95.1",
    "dns": [
      "8.8.8.8"
    ],
    "gateway": "192.168.95.1",
    "ip": "192.168.95.3",
    "ip6": "fe80:0:0:0:5e5b:35ff:fed0:77b",
    "ips": {
      "vlan1": "192.168.95.3/24,fe80:0:0:0:5e5b:35ff:fed0:77b/64"
    },
    "netmask": "255.255.255.0",
    "netmask6": "/64"
  },
  "last_seen": 1720595866,
  "last_trouble": {
    "code": "07",
    "timestamp": 1720039666
  },
  "lldp_stat": {
    "chassis_id": "d0:07:ca:f5:21:00",
    "lldp_med_supported": false,
    "mgmt_addr": "100.123.105.1",
    "mgmt_addrs": [
      "100.123.105.1"
    ],
    "port_desc": "ge-0/0/4",
    "port_id": "ge-0/0/4",
    "power_allocated": 0,
    "power_draw": 0,
    "power_request_count": 0,
    "power_requested": 0,
    "system_desc": "Juniper Networks, Inc. ex4300-48t internet router, kernel JUNOS 20.4R3-S7.2, Build date: 2023-04-21 19:47:18 UTC Copyright (c) 1996-2023 Juniper Networks, Inc.",
    "system_name": "Phoenix-Switch"
  },
  "mac": "5c5b35d0077b",
  "mem_total_kb": 505468,
  "mem_used_kb": 202096,
  "model": "AP43",
  "modified_time": 1718530662,
  "mount": "faceup",
  "name": "Phoenix",
  "notes": "",
  "num_clients": 1,
  "org_id": "af010e2b-f829-4975-b49e-2e896ed1d627",
  "port_stat": {
    "eth0": {
      "full_duplex": true,
      "rx_bytes": 1284143195,
      "rx_errors": 0,
      "rx_peak_bps": 17585,
      "rx_pkts": 5199816,
      "speed": 1000,
      "tx_bytes": 1283744961,
      "tx_peak_bps": 26484,
      "tx_pkts": 3990463,
      "up": true
    },
    "eth1": {
      "full_duplex": false,
      "rx_bytes": 0,
      "rx_errors": 0,
      "rx_peak_bps": 0,
      "rx_pkts": 0,
      "speed": 0,
      "tx_bytes": 0,
      "tx_peak_bps": 0,
      "tx_pkts": 0,
      "up": false
    }
  },
  "power_budget": 8400,
  "power_constrained": false,
  "power_src": "DC Input",
  "radio_config": {},
  "radio_stat": {
    "band_24": {
      "bandwidth": 20,
      "channel": 11,
      "mac": "5c5b35dea810",
      "noise_floor": -80,
      "num_clients": 0,
      "power": 17,
      "rx_bytes": 12948211,
      "rx_pkts": 65292,
      "tx_bytes": 19071943,
      "tx_pkts": 76926,
      "usage": "24",
      "util_all": 24,
      "util_non_wifi": 2,
      "util_rx_in_bss": 0,
      "util_rx_other_bss": 17,
      "util_tx": 4,
      "util_undecodable_wifi": 0,
      "util_unknown_wifi": 1
    },
    "band_5": {
      "bandwidth": 40,
      "channel": 36,
      "mac": "5c5b35dea7f0",
      "noise_floor": -90,
      "num_clients": 1,
      "power": 17,
      "rx_bytes": 578362619,
      "rx_pkts": 2687577,
      "tx_bytes": 1199571353,
      "tx_pkts": 2479302,
      "usage": "5",
      "util_all": 13,
      "util_non_wifi": 0,
      "util_rx_in_bss": 0,
      "util_rx_other_bss": 10,
      "util_tx": 1,
      "util_undecodable_wifi": 0,
      "util_unknown_wifi": 1
    }
  },
  "rx_bps": 9276,
  "rx_bytes": 591310830,
  "rx_pkts": 2752869,
  "serial": "A0703200709E6",
  "site_id": "46fc665e-9706-4296-8fe2-78f42f2e67e4",
  "status": "connected",
  "switch_redundancy": {
    "num_redundant_aps": 1
  },
  "tx_bps": 8067,
  "tx_bytes": 1218643296,
  "tx_pkts": 2556228,
  "type": "ap",
  "uptime": 1593120,
  "version": "0.14.29313"
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
    "auto_upgrade_stat": {
      "lastcheck": 1720594762
    },
    "ble_stat": {
      "beacon_enabled": true,
      "beacon_rate": 4,
      "eddystone_uid_enabled": false,
      "eddystone_uid_freq_msec": 1000,
      "eddystone_uid_instance": "5c5b35d0077b",
      "eddystone_uid_namespace": "9777c1a06ef611e68bbf",
      "eddystone_url_enabled": false,
      "eddystone_url_freq_msec": 1000,
      "eddystone_url_url": "",
      "ibeacon_enabled": false,
      "ibeacon_freq_msec": 1000,
      "ibeacon_major": 894,
      "ibeacon_minor": 9328,
      "ibeacon_uuid": "af010e2b-f829-4975-b49e-2e896ed1d627",
      "major": 894,
      "minors": [
        9328,
        9329,
        9330,
        9331,
        9332,
        9333,
        9334,
        9335,
        -1
      ],
      "power": 8,
      "rx_bytes": 158500843,
      "rx_pkts": 3549163,
      "tx_bytes": 509640,
      "tx_pkts": 85411,
      "tx_resets": 0,
      "uuid": "af010e2b-f829-4975-b49e-2e896ed1d627"
    },
    "config_reverted": false,
    "cpu_system": 21921854,
    "cpu_user": 7496631,
    "cpu_util": 5,
    "created_time": 1718228350,
    "env_stat": {
      "accel_x": -0.092,
      "accel_y": 0.004,
      "accel_z": -1.02,
      "ambient_temp": 43,
      "attitude": 0,
      "cpu_temp": 53,
      "humidity": 9,
      "magne_x": 0,
      "magne_y": 0,
      "magne_z": 0,
      "pressure": 968,
      "vcore_voltage": 0
    },
    "ext_ip": "66.129.234.28",
    "hw_rev": "C02",
    "id": "00000000-0000-0000-1000-5c5b35d0077b",
    "inactive_wired_vlans": [],
    "ip": "192.168.95.3",
    "ip_stat": {
      "dhcp_server": "192.168.95.1",
      "dns": [
        "8.8.8.8"
      ],
      "gateway": "192.168.95.1",
      "ip": "192.168.95.3",
      "ip6": "fe80:0:0:0:5e5b:35ff:fed0:77b",
      "ips": {
        "vlan1": "192.168.95.3/24,fe80:0:0:0:5e5b:35ff:fed0:77b/64"
      },
      "netmask": "255.255.255.0",
      "netmask6": "/64"
    },
    "last_seen": 1720595866,
    "last_trouble": {
      "code": "07",
      "timestamp": 1720039666
    },
    "lldp_stat": {
      "chassis_id": "d0:07:ca:f5:21:00",
      "lldp_med_supported": false,
      "mgmt_addr": "100.123.105.1",
      "mgmt_addrs": [
        "100.123.105.1"
      ],
      "port_desc": "ge-0/0/4",
      "port_id": "ge-0/0/4",
      "power_allocated": 0,
      "power_draw": 0,
      "power_request_count": 0,
      "power_requested": 0,
      "system_desc": "Juniper Networks, Inc. ex4300-48t internet router, kernel JUNOS 20.4R3-S7.2, Build date: 2023-04-21 19:47:18 UTC Copyright (c) 1996-2023 Juniper Networks, Inc.",
      "system_name": "Phoenix-Switch"
    },
    "mac": "5c5b35d0077b",
    "mem_total_kb": 505468,
    "mem_used_kb": 202096,
    "model": "AP43",
    "modified_time": 1718530662,
    "mount": "faceup",
    "name": "Phoenix",
    "notes": "",
    "num_clients": 1,
    "org_id": "af010e2b-f829-4975-b49e-2e896ed1d627",
    "port_stat": {
      "eth0": {
        "full_duplex": true,
        "rx_bytes": 1284143195,
        "rx_errors": 0,
        "rx_peak_bps": 17585,
        "rx_pkts": 5199816,
        "speed": 1000,
        "tx_bytes": 1283744961,
        "tx_peak_bps": 26484,
        "tx_pkts": 3990463,
        "up": true
      },
      "eth1": {
        "full_duplex": false,
        "rx_bytes": 0,
        "rx_errors": 0,
        "rx_peak_bps": 0,
        "rx_pkts": 0,
        "speed": 0,
        "tx_bytes": 0,
        "tx_peak_bps": 0,
        "tx_pkts": 0,
        "up": false
      }
    },
    "power_budget": 8400,
    "power_constrained": false,
    "power_src": "DC Input",
    "radio_config": {},
    "radio_stat": {
      "band_24": {
        "bandwidth": 20,
        "channel": 11,
        "mac": "5c5b35dea810",
        "noise_floor": -80,
        "num_clients": 0,
        "power": 17,
        "rx_bytes": 12948211,
        "rx_pkts": 65292,
        "tx_bytes": 19071943,
        "tx_pkts": 76926,
        "usage": "24",
        "util_all": 24,
        "util_non_wifi": 2,
        "util_rx_in_bss": 0,
        "util_rx_other_bss": 17,
        "util_tx": 4,
        "util_undecodable_wifi": 0,
        "util_unknown_wifi": 1
      },
      "band_5": {
        "bandwidth": 40,
        "channel": 36,
        "mac": "5c5b35dea7f0",
        "noise_floor": -90,
        "num_clients": 1,
        "power": 17,
        "rx_bytes": 578362619,
        "rx_pkts": 2687577,
        "tx_bytes": 1199571353,
        "tx_pkts": 2479302,
        "usage": "5",
        "util_all": 13,
        "util_non_wifi": 0,
        "util_rx_in_bss": 0,
        "util_rx_other_bss": 10,
        "util_tx": 1,
        "util_undecodable_wifi": 0,
        "util_unknown_wifi": 1
      }
    },
    "rx_bps": 9276,
    "rx_bytes": 591310830,
    "rx_pkts": 2752869,
    "serial": "A0703200709E6",
    "site_id": "46fc665e-9706-4296-8fe2-78f42f2e67e4",
    "status": "connected",
    "switch_redundancy": {
      "num_redundant_aps": 1
    },
    "tx_bps": 8067,
    "tx_bytes": 1218643296,
    "tx_pkts": 2556228,
    "type": "ap",
    "uptime": 1593120,
    "version": "0.14.29313"
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

