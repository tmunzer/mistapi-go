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

Count switch and gateway port statistics for a site, optionally grouped by the `distinct` field and filtered by port, neighbor, PoE, STP, traffic, and time attributes. Use [Count Org Switch/Gateway Ports](../../doc/controllers/orgs-stats-ports.md#count-org-sw-or-gw-ports) to count switch and gateway port statistics across the organization.

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
    stpState *models.PortStpStateEnum,
    stpRole *models.PortStpRoleEnum,
    authState *models.PortAuthStateEnum,
    up *bool,
    start *string,
    end *string,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SitePortsCountDistinctEnum`](../../doc/models/site-ports-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `full_duplex`, `mac`, `neighbor_mac`, `neighbor_port_desc`, `neighbor_system_name`, `poe_disabled`, `poe_mode`, `poe_on`, `port_id`, `port_mac`, `speed`, `up`<br><br>**Default**: `"mac"` |
| `fullDuplex` | `*bool` | Query, Optional | Indicates full or half duplex |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |
| `neighborMac` | `*string` | Query, Optional | Chassis identifier of the chassis type listed |
| `neighborPortDesc` | `*string` | Query, Optional | Description supplied by the system on the interface E.g. "GigabitEthernet2/0/39" |
| `neighborSystemName` | `*string` | Query, Optional | Name supplied by the system on the interface E.g. neighbor system name E.g. "Kumar-Acc-SW.mist.local" |
| `poeDisabled` | `*bool` | Query, Optional | Is the POE configured not be disabled. |
| `poeMode` | `*string` | Query, Optional | POE mode depending on class E.g. "802.3at" |
| `poeOn` | `*bool` | Query, Optional | Is the device attached to POE |
| `portId` | `*string` | Query, Optional | Filter results by port identifier |
| `portMac` | `*string` | Query, Optional | Filter results by port MAC address |
| `powerDraw` | `*float64` | Query, Optional | Amount of power being used by the interface at the time the command is executed. Unit in watts. |
| `txPkts` | `*int` | Query, Optional | Filter results by transmitted packet count |
| `rxPkts` | `*int` | Query, Optional | Filter results by received packet count |
| `rxBytes` | `*int` | Query, Optional | Filter results by received byte count |
| `txBps` | `*int` | Query, Optional | Filter results by transmit rate |
| `rxBps` | `*int` | Query, Optional | Filter results by receive rate |
| `txMcastPkts` | `*int` | Query, Optional | Filter results by transmitted multicast packet count |
| `txBcastPkts` | `*int` | Query, Optional | Filter results by transmitted broadcast packet count |
| `rxMcastPkts` | `*int` | Query, Optional | Filter results by received multicast packet count |
| `rxBcastPkts` | `*int` | Query, Optional | Filter results by received broadcast packet count |
| `speed` | `*int` | Query, Optional | Filter results by port speed |
| `stpState` | [`*models.PortStpStateEnum`](../../doc/models/port-stp-state-enum.md) | Query, Optional | STP state used to filter port results when `up`==`true`. enum: `""`, `blocking`, `disabled`, `forwarding`, `learning`, `listening` |
| `stpRole` | [`*models.PortStpRoleEnum`](../../doc/models/port-stp-role-enum.md) | Query, Optional | STP role used to filter port results when `up`==`true`. enum: `""`, `alternate`, `backup`, `designated`, `disabled`, `root`, `root-prevented` |
| `authState` | [`*models.PortAuthStateEnum`](../../doc/models/port-auth-state-enum.md) | Query, Optional | Authentication state used to filter port results when `up`==`true` and the port has an authenticator role. enum: `""`, `authenticated`, `authenticating`, `held`, `init` |
| `up` | `*bool` | Query, Optional | Indicates if interface is up |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

**200**: Result of Count

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
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Search Site Sw or Gw Ports

Search switch and gateway port statistics for a site.
Returns ports that match the search criteria, including current or most recent port status and statistics within the hour.

Traffic information (Tx/Rx) is reported as cumulative counters since the last device reboot. Use [Search Org Switch/Gateway Ports](../../doc/controllers/orgs-stats-ports.md#search-org-sw-or-gw-ports) to search switch and gateway port statistics across the organization.

```go
SearchSiteSwOrGwPorts(
    ctx context.Context,
    siteId uuid.UUID,
    deviceType *models.SearchOrgSwOrGwPortsTypeEnum,
    authState *models.PortAuthStateEnum,
    fullDuplex *bool,
    lteImsi *string,
    lteIccid *string,
    lteImei *string,
    mac *string,
    neighborMac *string,
    neighborPortDesc *string,
    neighborSystemName *string,
    poeDisabled *bool,
    poeMode *string,
    poeOn *bool,
    poePriority *models.PoePriorityEnum,
    portId *string,
    portMac *string,
    speed *int,
    stpState *models.PortStpStateEnum,
    stpRole *models.PortStpRoleEnum,
    up *bool,
    xcvrPartNumber *string,
    limit *int,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseSwitchPortSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceType` | [`*models.SearchOrgSwOrGwPortsTypeEnum`](../../doc/models/search-org-sw-or-gw-ports-type-enum.md) | Query, Optional | Type of device. enum: `switch`, `gateway`, `all`<br><br>**Default**: `"all"` |
| `authState` | [`*models.PortAuthStateEnum`](../../doc/models/port-auth-state-enum.md) | Query, Optional | Authentication state used to filter port results when `up`==`true` and the port has an authenticator role. enum: `""`, `authenticated`, `authenticating`, `held`, `init` |
| `fullDuplex` | `*bool` | Query, Optional | Indicates full or half duplex |
| `lteImsi` | `*string` | Query, Optional | LTE IMSI value, Check for null/empty |
| `lteIccid` | `*string` | Query, Optional | LTE ICCID value, Check for null/empty |
| `lteImei` | `*string` | Query, Optional | LTE IMEI value, Check for null/empty |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |
| `neighborMac` | `*string` | Query, Optional | Chassis identifier of the chassis type listed |
| `neighborPortDesc` | `*string` | Query, Optional | Description supplied by the system on the interface E.g. "GigabitEthernet2/0/39" |
| `neighborSystemName` | `*string` | Query, Optional | Name supplied by the system on the interface E.g. neighbor system name E.g. "Kumar-Acc-SW.mist.local" |
| `poeDisabled` | `*bool` | Query, Optional | Is the POE configured not be disabled. |
| `poeMode` | `*string` | Query, Optional | POE mode depending on class E.g. "802.3at" |
| `poeOn` | `*bool` | Query, Optional | Is the device attached to POE |
| `poePriority` | [`*models.PoePriorityEnum`](../../doc/models/poe-priority-enum.md) | Query, Optional | PoE priority used to filter switch port results. enum: `low`, `high` |
| `portId` | `*string` | Query, Optional | Filter results by port identifier |
| `portMac` | `*string` | Query, Optional | Filter results by port MAC address |
| `speed` | `*int` | Query, Optional | Filter results by port speed |
| `stpState` | [`*models.PortStpStateEnum`](../../doc/models/port-stp-state-enum.md) | Query, Optional | STP state used to filter port results when `up`==`true`. enum: `""`, `blocking`, `disabled`, `forwarding`, `learning`, `listening` |
| `stpRole` | [`*models.PortStpRoleEnum`](../../doc/models/port-stp-role-enum.md) | Query, Optional | STP role used to filter port results when `up`==`true`. enum: `""`, `alternate`, `backup`, `designated`, `disabled`, `root`, `root-prevented` |
| `up` | `*bool` | Query, Optional | Indicates if interface is up |
| `xcvrPartNumber` | `*string` | Query, Optional | Optic Slot Partnumber, Check for null/empty |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: List of Switch Ports Stats

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSwitchPortSearch](../../doc/models/response-switch-port-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceType := models.SearchOrgSwOrGwPortsTypeEnum_ALL

lteImsi := "310260000000001"

lteIccid := "89014103211118510720"

lteImei := "123456789012345"

xcvrPartNumber := "SFP-10G-SR"

limit := 100

sort := "-site_id"

apiResponse, err := sitesStatsPorts.SearchSiteSwOrGwPorts(ctx, siteId, &deviceType, nil, nil, &lteImsi, &lteIccid, &lteImei, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &xcvrPartNumber, &limit, &sort, nil)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

