# Orgs Stats-Ports

```go
orgsStatsPorts := client.OrgsStatsPorts()
```

## Class Name

`OrgsStatsPorts`

## Methods

* [Count Org Sw or Gw Ports](../../doc/controllers/orgs-stats-ports.md#count-org-sw-or-gw-ports)
* [Search Org Sw or Gw Ports](../../doc/controllers/orgs-stats-ports.md#search-org-sw-or-gw-ports)


# Count Org Sw or Gw Ports

Count by Distinct Attributes of Switch/Gateway Ports at the Org level

```go
CountOrgSwOrGwPorts(
    ctx context.Context,
    orgId uuid.UUID,
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
    siteId *uuid.UUID,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SitePortsCountDistinctEnum`](../../doc/models/site-ports-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `full_duplex`, `mac`, `neighbor_mac`, `neighbor_port_desc`, `neighbor_system_name`, `poe_disabled`, `poe_mode`, `poe_on`, `port_id`, `port_mac`, `speed`, `up`<br><br>**Default**: `"mac"` |
| `fullDuplex` | `*bool` | Query, Optional | Indicates full or half duplex |
| `mac` | `*string` | Query, Optional | Filter results by MAC address. Accepts multiple comma-separated values. |
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
| `siteId` | `*uuid.UUID` | Query, Optional | Filter results by site identifier |
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

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SitePortsCountDistinctEnum_MAC

fullDuplex := true

mac := "5c5b53010101,5c5b53020202"

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

siteId := uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811")

duration := "10m"

limit := 100

apiResponse, err := orgsStatsPorts.CountOrgSwOrGwPorts(ctx, orgId, &distinct, &fullDuplex, &mac, &neighborMac, &neighborPortDesc, &neighborSystemName, &poeDisabled, &poeMode, &poeOn, &portId, &portMac, &powerDraw, &txPkts, &rxPkts, &rxBytes, &txBps, &rxBps, &txMcastPkts, &txBcastPkts, &rxMcastPkts, &rxBcastPkts, &speed, nil, nil, nil, &up, &siteId, nil, nil, &duration, &limit)
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


# Search Org Sw or Gw Ports

Search Switch / Gateway Ports Stats.
Returns a list of switch/gateway ports stats that match the search criteria.

The response provide current/last port status and statistics within the hour.
Traffic information (Tx/Rx) are cumulative counters since the last device reboot.

```go
SearchOrgSwOrGwPorts(
    ctx context.Context,
    orgId uuid.UUID,
    deviceType *string,
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
    models.ApiResponse[models.ResponsePortStatsSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceType` | `*string` | Query, Optional | Type of device. enum: `switch`, `gateway`, `all`. Accepts multiple comma-separated values. |
| `authState` | [`*models.PortAuthStateEnum`](../../doc/models/port-auth-state-enum.md) | Query, Optional | Authentication state used to filter port results when `up`==`true` and the port has an authenticator role. enum: `""`, `authenticated`, `authenticating`, `held`, `init` |
| `fullDuplex` | `*bool` | Query, Optional | Indicates full or half duplex |
| `lteImsi` | `*string` | Query, Optional | LTE IMSI value, Check for null/empty |
| `lteIccid` | `*string` | Query, Optional | LTE ICCID value, Check for null/empty |
| `lteImei` | `*string` | Query, Optional | LTE IMEI value, Check for null/empty |
| `mac` | `*string` | Query, Optional | Filter results by MAC address. Accepts multiple comma-separated values. |
| `neighborMac` | `*string` | Query, Optional | Chassis identifier of the chassis type listed. Accepts multiple comma-separated values. |
| `neighborPortDesc` | `*string` | Query, Optional | Description supplied by the system on the interface E.g. "GigabitEthernet2/0/39". Accepts multiple comma-separated values. |
| `neighborSystemName` | `*string` | Query, Optional | Name supplied by the system on the interface E.g. neighbor system name E.g. "Kumar-Acc-SW.mist.local". Accepts multiple comma-separated values. |
| `poeDisabled` | `*bool` | Query, Optional | Is the POE configured not be disabled. |
| `poeMode` | `*string` | Query, Optional | POE mode depending on class E.g. "802.3at" |
| `poeOn` | `*bool` | Query, Optional | Is the device attached to POE |
| `poePriority` | [`*models.PoePriorityEnum`](../../doc/models/poe-priority-enum.md) | Query, Optional | PoE priority used to filter switch port results. enum: `low`, `high` |
| `portId` | `*string` | Query, Optional | Filter results by port identifier. Accepts multiple comma-separated values. |
| `portMac` | `*string` | Query, Optional | Filter results by port MAC address. Accepts multiple comma-separated values. |
| `speed` | `*int` | Query, Optional | Filter results by port speed |
| `stpState` | [`*models.PortStpStateEnum`](../../doc/models/port-stp-state-enum.md) | Query, Optional | STP state used to filter port results when `up`==`true`. enum: `""`, `blocking`, `disabled`, `forwarding`, `learning`, `listening` |
| `stpRole` | [`*models.PortStpRoleEnum`](../../doc/models/port-stp-role-enum.md) | Query, Optional | STP role used to filter port results when `up`==`true`. enum: `""`, `alternate`, `backup`, `designated`, `disabled`, `root`, `root-prevented` |
| `up` | `*bool` | Query, Optional | Indicates if interface is up |
| `xcvrPartNumber` | `*string` | Query, Optional | Optic Slot Partnumber, Check for null/empty |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: Example response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponsePortStatsSearch](../../doc/models/response-port-stats-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceType := "gateway,switch"

lteImsi := "310260000000001"

lteIccid := "89014103211118510720"

lteImei := "123456789012345"

mac := "5c5b53010101,5c5b53020202"

neighborMac := "5c5b53010101,5c5b53020202"

neighborPortDesc := "ge-0-2,ge-0-1"

neighborSystemName := "sdwan-newyork,toronto-srx-2_node1"

portId := "ge-0/0/1,ge-0/0/2"

portMac := "5c5b53010101,5c5b53020202"

xcvrPartNumber := "SFP-10G-SR"

limit := 100

sort := "-site_id"

apiResponse, err := orgsStatsPorts.SearchOrgSwOrGwPorts(ctx, orgId, &deviceType, nil, nil, &lteImsi, &lteIccid, &lteImei, &mac, &neighborMac, &neighborPortDesc, &neighborSystemName, nil, nil, nil, nil, &portId, &portMac, nil, nil, nil, nil, &xcvrPartNumber, &limit, &sort, nil)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

