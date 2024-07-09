# Sites Skyatp

```go
sitesSkyatp := client.SitesSkyatp()
```

## Class Name

`SitesSkyatp`

## Methods

* [Count Site Skyatp Events](../../doc/controllers/sites-skyatp.md#count-site-skyatp-events)
* [Search Site Skyatp Events](../../doc/controllers/sites-skyatp.md#search-site-skyatp-events)


# Count Site Skyatp Events

Count by Distinct Attributes of Skyatp Events (WIP)

```go
CountSiteSkyatpEvents(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteSkyAtpEventsCountDistinctEnum,
    mType *string,
    mac *string,
    deviceMac *string,
    threatLevel *int,
    ipAddress *string,
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
| `distinct` | [`*models.SiteSkyAtpEventsCountDistinctEnum`](../../doc/models/site-sky-atp-events-count-distinct-enum.md) | Query, Optional | - |
| `mType` | `*string` | Query, Optional | event type, e.g. cc, fs, mw |
| `mac` | `*string` | Query, Optional | client MAC |
| `deviceMac` | `*string` | Query, Optional | device MAC |
| `threatLevel` | `*int` | Query, Optional | threat level |
| `ipAddress` | `*string` | Query, Optional | - |
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

distinct := models.SiteSkyAtpEventsCountDistinctEnum("type")









ipAddress := "192.168.1.1"

limit := 100





duration := "10m"

apiResponse, err := sitesSkyatp.CountSiteSkyatpEvents(ctx, siteId, &distinct, nil, nil, nil, nil, &ipAddress, &limit, nil, nil, &duration)
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


# Search Site Skyatp Events

Search Skyatp Events (WIP)

```go
SearchSiteSkyatpEvents(
    ctx context.Context,
    siteId uuid.UUID,
    mType *string,
    mac *string,
    deviceMac *string,
    threatLevel *int,
    ipAddress *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseEventsSkyAtpSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | event type, e.g. cc, fs, mw |
| `mac` | `*string` | Query, Optional | client MAC |
| `deviceMac` | `*string` | Query, Optional | device MAC |
| `threatLevel` | `*int` | Query, Optional | threat level |
| `ipAddress` | `*string` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseEventsSkyAtpSearch`](../../doc/models/response-events-sky-atp-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")









ipAddress := "192.168.1.1"

limit := 100





duration := "10m"

apiResponse, err := sitesSkyatp.SearchSiteSkyatpEvents(ctx, siteId, nil, nil, nil, nil, &ipAddress, &limit, nil, nil, &duration)
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
  "end": 1513176951,
  "limit": 10,
  "results": [
    {
      "device_mac": "658279bb1fa4",
      "ip": "172.16.0.11",
      "mac": "b019c66c8348",
      "org_id": "3139f2c2-fac6-11e5-8156-0242ac110006",
      "site_id": "70e0f468-fc13-11e5-85ad-0242ac110008",
      "threat_level": 7,
      "timestamp": 1592524478,
      "type": "cc"
    }
  ],
  "start": 1512572151,
  "total": 1
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

