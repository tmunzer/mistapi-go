# Sites Rogues

```go
sitesRogues := client.SitesRogues()
```

## Class Name

`SitesRogues`

## Methods

* [Count Site Rogue Events](../../doc/controllers/sites-rogues.md#count-site-rogue-events)
* [Get Site Rogue AP](../../doc/controllers/sites-rogues.md#get-site-rogue-ap)
* [List Site Rogue a Ps](../../doc/controllers/sites-rogues.md#list-site-rogue-a-ps)
* [List Site Rogue Clients](../../doc/controllers/sites-rogues.md#list-site-rogue-clients)
* [Search Site Rogue Events](../../doc/controllers/sites-rogues.md#search-site-rogue-events)


# Count Site Rogue Events

Count Rogue Events

```go
CountSiteRogueEvents(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteRogueEventsCountDistinctEnum,
    mType *models.RogueTypeEnum,
    ssid *string,
    bssid *string,
    apMac *string,
    channel *string,
    seenOnLan *bool,
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
| `distinct` | [`*models.SiteRogueEventsCountDistinctEnum`](../../doc/models/site-rogue-events-count-distinct-enum.md) | Query, Optional | **Default**: `"bssid"` |
| `mType` | [`*models.RogueTypeEnum`](../../doc/models/rogue-type-enum.md) | Query, Optional | - |
| `ssid` | `*string` | Query, Optional | ssid of the network detected as threat |
| `bssid` | `*string` | Query, Optional | bssid of the network detected as threat |
| `apMac` | `*string` | Query, Optional | mac of the device that had strongest signal strength for ssid/bssid pair |
| `channel` | `*string` | Query, Optional | channel over which ap_mac heard ssid/bssid pair |
| `seenOnLan` | `*bool` | Query, Optional | whether the reporting AP see a wireless client (on LAN) connecting to it |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteRogueEventsCountDistinctEnum("bssid")













limit := 100





duration := "10m"

apiResponse, err := sitesRogues.CountSiteRogueEvents(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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


# Get Site Rogue AP

Get Rogue AP Details

```go
GetSiteRogueAP(
    ctx context.Context,
    siteId uuid.UUID,
    rogueBssid string) (
    models.ApiResponse[models.RogueDetails],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `rogueBssid` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |

## Response Type

[`models.RogueDetails`](../../doc/models/rogue-details.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

rogueBssid := "0000000000ab"

apiResponse, err := sitesRogues.GetSiteRogueAP(ctx, siteId, rogueBssid)
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
  "manufacture": "Intel Corporate",
  "seen_as_client": true
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


# List Site Rogue a Ps

Get List of Site Rogue/Neighbor APs

```go
ListSiteRogueAPs(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.RogueTypeEnum,
    limit *int,
    start *int,
    end *int,
    duration *string,
    interval *string) (
    models.ApiResponse[models.ResponseInsightRogue],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.RogueTypeEnum`](../../doc/models/rogue-type-enum.md) | Query, Optional | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |

## Response Type

[`models.ResponseInsightRogue`](../../doc/models/response-insight-rogue.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



limit := 100





duration := "10m"

interval := "10m"

apiResponse, err := sitesRogues.ListSiteRogueAPs(ctx, siteId, nil, &limit, nil, nil, &duration, &interval)
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
  "end": 1428954000,
  "limit": 100,
  "next": "/api/v1/sites/a3eda150-ab3f-11e4-aa18-13e21dd250cc/rogues?start=1498482000&end=1498485600&limit=10&interval=1h&type=others",
  "results": [
    {
      "ap_mac": "5c5b350e021c",
      "avg_rssi": -72,
      "bssid": "d8-97-ba-76-b5-aa",
      "channel": "11",
      "num_aps": 4,
      "ssid": "xfinitywifi",
      "times_heard": 8
    }
  ],
  "start": 1428939600
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


# List Site Rogue Clients

Get List of Site Rogue Clients

```go
ListSiteRogueClients(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    start *int,
    end *int,
    duration *string,
    interval *string) (
    models.ApiResponse[models.ResponseInsightRogueClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |

## Response Type

[`models.ResponseInsightRogueClient`](../../doc/models/response-insight-rogue-client.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100





duration := "10m"

interval := "10m"

apiResponse, err := sitesRogues.ListSiteRogueClients(ctx, siteId, &limit, nil, nil, &duration, &interval)
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
  "end": 1428954000,
  "limit": 100,
  "next": "/api/v1/sites/a3eda150-ab3f-11e4-aa18-13e21dd250cc/rogues/clients?start=1498482000&end=1498485600&limit=10&interval=1h",
  "results": [
    {
      "annotation": "whitelist",
      "ap_mac": "5c-5b-35-0e-02-1c",
      "avg_rssi": -63.9,
      "band": "5",
      "bssid": "d8-97-ba-76-b5-aa",
      "client_mac": "34-f8-32-13-57-c2",
      "num_aps": 2
    }
  ],
  "start": 1428939600
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


# Search Site Rogue Events

Search Rogue Events

```go
SearchSiteRogueEvents(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.RogueTypeEnum,
    ssid *string,
    bssid *string,
    apMac *string,
    channel *int,
    seenOnLan *bool,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseEventsRogueSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.RogueTypeEnum`](../../doc/models/rogue-type-enum.md) | Query, Optional | - |
| `ssid` | `*string` | Query, Optional | ssid of the network detected as threat |
| `bssid` | `*string` | Query, Optional | bssid of the network detected as threat |
| `apMac` | `*string` | Query, Optional | mac of the device that had strongest signal strength for ssid/bssid pair |
| `channel` | `*int` | Query, Optional | channel over which ap_mac heard ssid/bssid pair |
| `seenOnLan` | `*bool` | Query, Optional | whether the reporting AP see a wireless client (on LAN) connecting to it |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |

## Response Type

[`models.ResponseEventsRogueSearch`](../../doc/models/response-events-rogue-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")













limit := 100





duration := "10m"

apiResponse, err := sitesRogues.SearchSiteRogueEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
  "end": 1538074800,
  "limit": 10,
  "results": [
    {
      "ap": "5c5b350e10030",
      "bssid": "38ff363c8c4c",
      "channel": 136,
      "rssi": -54,
      "ssid": "MyHomeNetwork",
      "timestamp": 1538074612
    },
    {
      "ap": "5c5b350e10030",
      "bssid": "60d02c2394cc",
      "channel": 11,
      "rssi": -59,
      "ssid": "Home-Office",
      "timestamp": 1538074612
    }
  ],
  "start": 1538071200,
  "total": 2
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

