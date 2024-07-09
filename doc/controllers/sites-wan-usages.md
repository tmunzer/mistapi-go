# Sites WAN Usages

```go
sitesWANUsages := client.SitesWANUsages()
```

## Class Name

`SitesWANUsages`

## Methods

* [Count Site Wan Usage](../../doc/controllers/sites-wan-usages.md#count-site-wan-usage)
* [Search Site Wan Usage](../../doc/controllers/sites-wan-usages.md#search-site-wan-usage)


# Count Site Wan Usage

Count Site WAN Uages

```go
CountSiteWanUsage(
    ctx context.Context,
    siteId uuid.UUID,
    mac *string,
    peerMac *string,
    portId *string,
    peerPortId *string,
    policy *string,
    tenant *string,
    pathType *string,
    distinct *models.WanUsagesCountDisctinctEnum,
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
| `mac` | `*string` | Query, Optional | MAC address |
| `peerMac` | `*string` | Query, Optional | Peer MAC address |
| `portId` | `*string` | Query, Optional | Port ID for the device |
| `peerPortId` | `*string` | Query, Optional | Peer Port ID for the device |
| `policy` | `*string` | Query, Optional | policy for the wan path |
| `tenant` | `*string` | Query, Optional | tenant network in which the packet is sent |
| `pathType` | `*string` | Query, Optional | path_type of the port |
| `distinct` | [`*models.WanUsagesCountDisctinctEnum`](../../doc/models/wan-usages-count-disctinct-enum.md) | Query, Optional | - |
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















distinct := models.WanUsagesCountDisctinctEnum("policy")

page := 1

limit := 100





duration := "10m"

apiResponse, err := sitesWANUsages.CountSiteWanUsage(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, &distinct, &page, &limit, nil, nil, &duration)
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


# Search Site Wan Usage

Search Site WAN Uages

```go
SearchSiteWanUsage(
    ctx context.Context,
    siteId uuid.UUID,
    mac *string,
    peerMac *string,
    portId *string,
    peerPortId *string,
    policy *string,
    tenant *string,
    pathType *string,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.SearchWanUsage],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | MAC address |
| `peerMac` | `*string` | Query, Optional | Peer MAC address |
| `portId` | `*string` | Query, Optional | Port ID for the device |
| `peerPortId` | `*string` | Query, Optional | Peer Port ID for the device |
| `policy` | `*string` | Query, Optional | policy for the wan path |
| `tenant` | `*string` | Query, Optional | tenant network in which the packet is sent |
| `pathType` | `*string` | Query, Optional | path_type of the port |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.SearchWanUsage`](../../doc/models/search-wan-usage.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")















page := 1

limit := 100





duration := "10m"

apiResponse, err := sitesWANUsages.SearchSiteWanUsage(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, &page, &limit, nil, nil, &duration)
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

