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
    start *string,
    end *string,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SitePortsCountDistinctEnum`](../../doc/models/site-ports-count-distinct-enum.md) | Query, Optional | **Default**: `"mac"` |
| `fullDuplex` | `*bool` | Query, Optional | Indicates full or half duplex |
| `mac` | `*string` | Query, Optional | Device identifier |
| `neighborMac` | `*string` | Query, Optional | Chassis identifier of the chassis type listed |
| `neighborPortDesc` | `*string` | Query, Optional | Description supplied by the system on the interface E.g. "GigabitEthernet2/0/39" |
| `neighborSystemName` | `*string` | Query, Optional | Name supplied by the system on the interface E.g. neighbor system name E.g. "Kumar-Acc-SW.mist.local" |
| `poeDisabled` | `*bool` | Query, Optional | Is the POE configured not be disabled. |
| `poeMode` | `*string` | Query, Optional | POE mode depending on class E.g. "802.3at" |
| `poeOn` | `*bool` | Query, Optional | Is the device attached to POE |
| `portId` | `*string` | Query, Optional | Interface name |
| `portMac` | `*string` | Query, Optional | Interface mac address |
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
| `speed` | `*int` | Query, Optional | Port speed |
| `stpState` | [`*models.CountPortsStpStateEnum`](../../doc/models/count-ports-stp-state-enum.md) | Query, Optional | If `up`==`true` |
| `stpRole` | [`*models.CountPortsStpRoleEnum`](../../doc/models/count-ports-stp-role-enum.md) | Query, Optional | If `up`==`true` |
| `authState` | [`*models.CountPortsAuthStateEnum`](../../doc/models/count-ports-auth-state-enum.md) | Query, Optional | If `up`==`true` && has Authenticator role |
| `up` | `*bool` | Query, Optional | Indicates if interface is up |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SitePortsCountDistinctEnum_MAC

fullDuplex := true

mac := "5c5b350e0410"

neighborMac := "5c5b350e0410"

neighborPortDesc := "ge-2/0/39"

neighborSystemName := "Kumar-Acc-SW.mist.local"

poeDisabled := false

poeMode := "802.3at"

poeOn := true

portId := "ge-2/0/39"

portMac := "5c5b350e0410"

powerDraw := float64(15.4)

txPkts := 1000000

rxPkts := 1000000

rxBytes := 1000000

txBps := 1000000

rxBps := 1000000

txMcastPkts := 1000000

txBcastPkts := 1000000

rxMcastPkts := 1000000

rxBcastPkts := 1000000

speed := 1000000000

up := true

duration := "10m"

limit := 100

apiResponse, err := sitesStatsPorts.CountSiteSwOrGwPorts(ctx, siteId, &distinct, &fullDuplex, &mac, &neighborMac, &neighborPortDesc, &neighborSystemName, &poeDisabled, &poeMode, &poeOn, &portId, &portMac, &powerDraw, &txPkts, &rxPkts, &rxBytes, &txBps, &rxBps, &txMcastPkts, &txBcastPkts, &rxMcastPkts, &rxBcastPkts, &speed, nil, nil, nil, &up, nil, nil, &duration, &limit)
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
    disabled *bool,
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
    opticsBiasCurrent *float64,
    opticsTxPower *float64,
    opticsRxPower *float64,
    opticsModuleTemperature *float64,
    opticsModuleVoltage *float64,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string) (
    models.ApiResponse[models.ResponseSwitchPortSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `fullDuplex` | `*bool` | Query, Optional | Indicates full or half duplex |
| `disabled` | `*bool` | Query, Optional | Indicates if interface is disabled |
| `mac` | `*string` | Query, Optional | Device identifier |
| `deviceType` | [`*models.SearchSiteSwOrGwPortsDeviceTypeEnum`](../../doc/models/search-site-sw-or-gw-ports-device-type-enum.md) | Query, Optional | Device type |
| `neighborMac` | `*string` | Query, Optional | Chassis identifier of the chassis type listed |
| `neighborPortDesc` | `*string` | Query, Optional | Description supplied by the system on the interface E.g. "GigabitEthernet2/0/39" |
| `neighborSystemName` | `*string` | Query, Optional | Name supplied by the system on the interface E.g. neighbor system name E.g. "Kumar-Acc-SW.mist.local" |
| `poeDisabled` | `*bool` | Query, Optional | Is the POE configured not be disabled. |
| `poeMode` | `*string` | Query, Optional | POE mode depending on class E.g. "802.3at" |
| `poeOn` | `*bool` | Query, Optional | Is the device attached to POE |
| `portId` | `*string` | Query, Optional | Interface name |
| `portMac` | `*string` | Query, Optional | Interface mac address |
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
| `speed` | `*int` | Query, Optional | Port speed |
| `macLimit` | `*int` | Query, Optional | Limit on number of dynamically learned macs |
| `macCount` | `*int` | Query, Optional | Number of mac addresses in the forwarding table |
| `up` | `*bool` | Query, Optional | Indicates if interface is up |
| `active` | `*bool` | Query, Optional | Indicates if interface is active/inactive |
| `jitter` | `*float64` | Query, Optional | Last sampled jitter of the interface |
| `loss` | `*float64` | Query, Optional | Last sampled loss of the interface |
| `latency` | `*float64` | Query, Optional | Last sampled latency of the interface |
| `stpState` | [`*models.SearchSiteSwOrGwPortsStpStateEnum`](../../doc/models/search-site-sw-or-gw-ports-stp-state-enum.md) | Query, Optional | If `up`==`true` |
| `stpRole` | [`*models.SearchSiteSwOrGwPortsStpRoleEnum`](../../doc/models/search-site-sw-or-gw-ports-stp-role-enum.md) | Query, Optional | If `up`==`true` |
| `xcvrPartNumber` | `*string` | Query, Optional | Optic Slot Partnumber, Check for null/empty |
| `authState` | [`*models.SearchSiteSwOrGwPortsAuthStateEnum`](../../doc/models/search-site-sw-or-gw-ports-auth-state-enum.md) | Query, Optional | If `up`==`true` && has Authenticator role |
| `lteImsi` | `*string` | Query, Optional | LTE IMSI value, Check for null/empty |
| `lteIccid` | `*string` | Query, Optional | LTE ICCID value, Check for null/empty |
| `lteImei` | `*string` | Query, Optional | LTE IMEI value, Check for null/empty |
| `opticsBiasCurrent` | `*float64` | Query, Optional | Bias current of the optics in mA |
| `opticsTxPower` | `*float64` | Query, Optional | Transmit power of the optics in dBm |
| `opticsRxPower` | `*float64` | Query, Optional | Receive power of the optics in dBm |
| `opticsModuleTemperature` | `*float64` | Query, Optional | Temperature of the optics module in Celsius |
| `opticsModuleVoltage` | `*float64` | Query, Optional | Voltage of the optics module in mV |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSwitchPortSearch](../../doc/models/response-switch-port-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

fullDuplex := true

disabled := false

mac := "5c5b350e0410"

neighborMac := "5c5b350e0410"

neighborPortDesc := "ge-2/0/39"

neighborSystemName := "Kumar-Acc-SW.mist.local"

poeDisabled := false

poeMode := "802.3at"

poeOn := true

portId := "ge-2/0/39"

portMac := "5c5b350e0410"

powerDraw := float64(15.4)

txPkts := 1000000

rxPkts := 1000000

rxBytes := 1000000

txBps := 1000000

rxBps := 1000000

txErrors := 0

rxErrors := 0

txMcastPkts := 100

txBcastPkts := 100

rxMcastPkts := 100

rxBcastPkts := 100

speed := 1000

macLimit := 1000

macCount := 10

up := true

active := true

jitter := float64(0.456)

loss := float64(0.01)

latency := float64(0.123)

xcvrPartNumber := "SFP-10G-SR"

lteImsi := "310260000000001"

lteIccid := "89014103211118510720"

lteImei := "123456789012345"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := sitesStatsPorts.SearchSiteSwOrGwPorts(ctx, siteId, &fullDuplex, &disabled, &mac, nil, &neighborMac, &neighborPortDesc, &neighborSystemName, &poeDisabled, &poeMode, &poeOn, &portId, &portMac, &powerDraw, &txPkts, &rxPkts, &rxBytes, &txBps, &rxBps, &txErrors, &rxErrors, &txMcastPkts, &txBcastPkts, &rxMcastPkts, &rxBcastPkts, &speed, &macLimit, &macCount, &up, &active, &jitter, &loss, &latency, nil, nil, &xcvrPartNumber, nil, &lteImsi, &lteIccid, &lteImei, nil, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort)
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

