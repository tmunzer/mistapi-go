# Sites Stats-Ospf

```go
sitesStatsOspf := client.SitesStatsOspf()
```

## Class Name

`SitesStatsOspf`

## Methods

* [Count Site Ospf Stats](../../doc/controllers/sites-stats-ospf.md#count-site-ospf-stats)
* [Search Site Ospf Stats](../../doc/controllers/sites-stats-ospf.md#search-site-ospf-stats)


# Count Site Ospf Stats

Count by Distinct Attributes of OSPF peers stats

```go
CountSiteOspfStats(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.OspfPeerStatsCountDistinctEnum,
    start *string,
    end *string,
    limit *int,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OspfPeerStatsCountDistinctEnum`](../../doc/models/ospf-peer-stats-count-distinct-enum.md) | Query, Optional | - |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

sort := "-site_id"

apiResponse, err := sitesStatsOspf.CountSiteOspfStats(ctx, siteId, nil, nil, nil, &limit, &sort, nil)
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


# Search Site Ospf Stats

Search OSPF Neighbor Stats

```go
SearchSiteOspfStats(
    ctx context.Context,
    siteIdTemplate uuid.UUID,
    siteId *string,
    mac *string,
    vrfName *string,
    peerIp *string,
    start *string,
    end *string,
    limit *int,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.OspfPeerStatsSearchResult],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteIdTemplate` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*string` | Query, Optional | - |
| `mac` | `*string` | Query, Optional | - |
| `vrfName` | `*string` | Query, Optional | - |
| `peerIp` | `*string` | Query, Optional | - |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.OspfPeerStatsSearchResult](../../doc/models/ospf-peer-stats-search-result.md).

## Example Usage

```go
ctx := context.Background()

siteIdTemplate := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

sort := "-site_id"

apiResponse, err := sitesStatsOspf.SearchSiteOspfStats(ctx, siteIdTemplate, nil, nil, nil, nil, nil, nil, &limit, &sort, nil)
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

