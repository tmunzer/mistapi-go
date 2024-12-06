# Sites Stats-Ports

```go
sitesStatsPorts := client.SitesStatsPorts()
```

## Class Name

`SitesStatsPorts`

## Methods

* [Count Site Sw or Gw Ports](../../doc/controllers/sites-stats-ports.md#count-site-sw-or-gw-ports)
* [Search Site Sw or Gw Ports](../../doc/controllers/sites-stats-ports.md#search-site-sw-or-gw-ports)


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
    stpState *models.CountPortsStpStateEnum,
    stpRole *models.CountPortsStpRoleEnum,
    authState *models.CountPortsAuthStateEnum,
    up *bool,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SitePortsCountDistinctEnum`](../../doc/models/site-ports-count-distinct-enum.md) | Query, Optional | **Default**: `"mac"` |
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
| `stpState` | [`*models.CountPortsStpStateEnum`](../../doc/models/count-ports-stp-state-enum.md) | Query, Optional | if `up`==`true` |
| `stpRole` | [`*models.CountPortsStpRoleEnum`](../../doc/models/count-ports-stp-role-enum.md) | Query, Optional | if `up`==`true` |
| `authState` | [`*models.CountPortsAuthStateEnum`](../../doc/models/count-ports-auth-state-enum.md) | Query, Optional | if `up`==`true` && has Authenticator role |
| `up` | `*bool` | Query, Optional | indicates if interface is up |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SitePortsCountDistinctEnum("mac")























































duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesStatsPorts.CountSiteSwOrGwPorts(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |

## Response Type

[`models.ResponseSwitchPortSearch`](../../doc/models/response-switch-port-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")













































































limit := 100





duration := "10m"

apiResponse, err := sitesStatsPorts.SearchSiteSwOrGwPorts(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

