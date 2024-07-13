# Orgs Devices Stats

```go
orgsDevicesStats := client.OrgsDevicesStats()
```

## Class Name

`OrgsDevicesStats`

## Methods

* [Count Org Bgp Stats](../../doc/controllers/orgs-devices-stats.md#count-org-bgp-stats)
* [Count Org Switch Ports](../../doc/controllers/orgs-devices-stats.md#count-org-switch-ports)
* [List Org Devices Stats](../../doc/controllers/orgs-devices-stats.md#list-org-devices-stats)
* [Search Org Bgp Stats](../../doc/controllers/orgs-devices-stats.md#search-org-bgp-stats)
* [Search Org Sw or Gw Ports](../../doc/controllers/orgs-devices-stats.md#search-org-sw-or-gw-ports)


# Count Org Bgp Stats

Count Org BGP Stats

```go
CountOrgBgpStats(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsDevicesStats.CountOrgBgpStats(ctx, orgId)
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


# Count Org Switch Ports

Count by Distinct Attributes of Switch/Gateway Ports

```go
CountOrgSwitchPorts(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgSwitchPortCountDistinctEnum,
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
    stpState *models.CountOrgSwitchPortsStpStateEnum,
    stpRole *models.CountOrgSwitchPortsStpRoleEnum,
    authState *models.CountOrgSwitchPortsAuthStateEnum,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgSwitchPortCountDistinctEnum`](../../doc/models/org-switch-port-count-distinct-enum.md) | Query, Optional | - |
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
| `stpState` | [`*models.CountOrgSwitchPortsStpStateEnum`](../../doc/models/count-org-switch-ports-stp-state-enum.md) | Query, Optional | if `up`==`true` |
| `stpRole` | [`*models.CountOrgSwitchPortsStpRoleEnum`](../../doc/models/count-org-switch-ports-stp-role-enum.md) | Query, Optional | if `up`==`true` |
| `authState` | [`*models.CountOrgSwitchPortsAuthStateEnum`](../../doc/models/count-org-switch-ports-auth-state-enum.md) | Query, Optional | if `up`==`true` |
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

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgSwitchPortCountDistinctEnum("mac")



















































page := 1

limit := 100





duration := "10m"

apiResponse, err := orgsDevicesStats.CountOrgSwitchPorts(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &page, &limit, nil, nil, &duration)
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


# List Org Devices Stats

Get List of Org Devices stats
This API renders some high-level device stats, pagination is assumed and returned in response header (as the response is an array)

```go
ListOrgDevicesStats(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeWithAllEnum,
    status *models.DeviceStatusEnum,
    siteId *string,
    mac *string,
    evpntopoId *string,
    evpnUnused *string,
    fields *string,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[[]models.ListOrgDevicesStatsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeWithAllEnum`](../../doc/models/device-type-with-all-enum.md) | Query, Optional | - |
| `status` | [`*models.DeviceStatusEnum`](../../doc/models/device-status-enum.md) | Query, Optional | - |
| `siteId` | `*string` | Query, Optional | - |
| `mac` | `*string` | Query, Optional | - |
| `evpntopoId` | `*string` | Query, Optional | EVPN Topology ID |
| `evpnUnused` | `*string` | Query, Optional | if `evpn_unused`==`true`, find EVPN eligible switches which don’t belong to any EVPN Topology yet |
| `fields` | `*string` | Query, Optional | list of additional fields requests, comma separeted, or `fields=*` for all of them |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`[]models.ListOrgDevicesStatsResponse`](../../doc/models/containers/list-org-devices-stats-response.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeWithAllEnum("ap")

status := models.DeviceStatusEnum("all")









fields := "field1,field2"

page := 1

limit := 100





duration := "10m"

apiResponse, err := orgsDevicesStats.ListOrgDevicesStats(ctx, orgId, &mType, &status, nil, nil, nil, nil, &fields, &page, &limit, nil, nil, &duration)
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


# Search Org Bgp Stats

Search Org BGP Stats

```go
SearchOrgBgpStats(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponseSearchBgps],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.ResponseSearchBgps`](../../doc/models/response-search-bgps.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsDevicesStats.SearchOrgBgpStats(ctx, orgId)
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


# Search Org Sw or Gw Ports

Search Switch / Gateway Ports

```go
SearchOrgSwOrGwPorts(
    ctx context.Context,
    orgId uuid.UUID,
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
    stpState *models.SearchOrgSwOrGwPortsStpStateEnum,
    stpRole *models.SearchOrgSwOrGwPortsStpRoleEnum,
    authState *models.SearchOrgSwOrGwPortsAuthStateEnum,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponsePortStatsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
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
| `stpState` | [`*models.SearchOrgSwOrGwPortsStpStateEnum`](../../doc/models/search-org-sw-or-gw-ports-stp-state-enum.md) | Query, Optional | if `up`==`true` |
| `stpRole` | [`*models.SearchOrgSwOrGwPortsStpRoleEnum`](../../doc/models/search-org-sw-or-gw-ports-stp-role-enum.md) | Query, Optional | if `up`==`true` |
| `authState` | [`*models.SearchOrgSwOrGwPortsAuthStateEnum`](../../doc/models/search-org-sw-or-gw-ports-auth-state-enum.md) | Query, Optional | if `up`==`true` && has Authenticator role |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponsePortStatsSearch`](../../doc/models/response-port-stats-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



























































limit := 100





duration := "10m"

apiResponse, err := orgsDevicesStats.SearchOrgSwOrGwPorts(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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

