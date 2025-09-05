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

Count Site WAN Usages

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
    distinct *models.WanUsagesCountDistinctEnum,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
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
| `policy` | `*string` | Query, Optional | Policy for the wan path |
| `tenant` | `*string` | Query, Optional | Tenant network in which the packet is sent |
| `pathType` | `*string` | Query, Optional | path_type of the port |
| `distinct` | [`*models.WanUsagesCountDistinctEnum`](../../doc/models/wan-usages-count-distinct-enum.md) | Query, Optional | **Default**: `"policy"` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mac := "001122334455"

peerMac := "001122334455"

portId := "ge-0/0/0"

peerPortId := "ge-0/0/0"

pathType := "primary"

distinct := models.WanUsagesCountDistinctEnum_POLICY

duration := "10m"

limit := 100

apiResponse, err := sitesWANUsages.CountSiteWanUsage(ctx, siteId, &mac, &peerMac, &portId, &peerPortId, nil, nil, &pathType, &distinct, nil, nil, &duration, &limit)
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


# Search Site Wan Usage

Search Site WAN Usages

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
    limit *int,
    page *int,
    start *int,
    end *int,
    duration *string,
    sort *string) (
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
| `policy` | `*string` | Query, Optional | Policy for the wan path |
| `tenant` | `*string` | Query, Optional | Tenant network in which the packet is sent |
| `pathType` | `*string` | Query, Optional | path_type of the port |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SearchWanUsage](../../doc/models/search-wan-usage.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mac := "001122334455"

peerMac := "001122334455"

portId := "ge-0/0/0"

peerPortId := "ge-0/0/0"

policy := "primary"

pathType := "primary"

limit := 100

page := 1

duration := "10m"

sort := "-site_id"

apiResponse, err := sitesWANUsages.SearchSiteWanUsage(ctx, siteId, &mac, &peerMac, &portId, &peerPortId, &policy, nil, &pathType, &limit, &page, nil, nil, &duration, &sort)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

