# Sites Stats-Beacons

```go
sitesStatsBeacons := client.SitesStatsBeacons()
```

## Class Name

`SitesStatsBeacons`


# List Site Beacons Stats

Get List of Site Beacons Stats

```go
ListSiteBeaconsStats(
    ctx context.Context,
    siteId uuid.UUID,
    start *string,
    end *string,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.StatsBeacon],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.StatsBeacon](../../doc/models/stats-beacon.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesStatsBeacons.ListSiteBeaconsStats(ctx, siteId, nil, nil, &duration, &limit, &page)
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

