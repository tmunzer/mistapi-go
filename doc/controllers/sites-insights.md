# Sites Insights

```go
sitesInsights := client.SitesInsights()
```

## Class Name

`SitesInsights`

## Methods

* [Get Site Insight Metrics](../../doc/controllers/sites-insights.md#get-site-insight-metrics)
* [Get Site Insight Metrics for Client](../../doc/controllers/sites-insights.md#get-site-insight-metrics-for-client)
* [Get Site Insight Metrics for Device](../../doc/controllers/sites-insights.md#get-site-insight-metrics-for-device)


# Get Site Insight Metrics

Get Site Insight Metrics
See metrics possibilities at /api/v1/const/insight_metrics

```go
GetSiteInsightMetrics(
    ctx context.Context,
    siteId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    interval *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.InsightMetrics],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | See /api/v1/const/insight_metrics for available metrics |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.InsightMetrics`](../../doc/models/insight-metrics.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

metric := "metric8"





duration := "10m"

interval := "10m"

limit := 100

page := 1

apiResponse, err := sitesInsights.GetSiteInsightMetrics(ctx, siteId, metric, nil, nil, &duration, &interval, &limit, &page)
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
  "end": 0,
  "interval": 0,
  "results": [
    {}
  ],
  "start": 0
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


# Get Site Insight Metrics for Client

Get Client Insight Metrics
See metrics possibilities at /api/v1/const/insight_metrics

```go
GetSiteInsightMetricsForClient(
    ctx context.Context,
    siteId uuid.UUID,
    clientMac string,
    metric string,
    start *int,
    end *int,
    duration *string,
    interval *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.InsightMetrics],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `clientMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `metric` | `string` | Template, Required | See /api/v1/const/insight_metrics for available metrics |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.InsightMetrics`](../../doc/models/insight-metrics.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

clientMac := "0000000000ab"

metric := "metric8"





duration := "10m"

interval := "10m"

limit := 100

page := 1

apiResponse, err := sitesInsights.GetSiteInsightMetricsForClient(ctx, siteId, clientMac, metric, nil, nil, &duration, &interval, &limit, &page)
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
  "end": 0,
  "interval": 0,
  "results": [
    {}
  ],
  "start": 0
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


# Get Site Insight Metrics for Device

Get AP Insight Metrics
See metrics possibilities at /api/v1/const/insight_metrics

```go
GetSiteInsightMetricsForDevice(
    ctx context.Context,
    siteId uuid.UUID,
    metric string,
    deviceMac string,
    start *int,
    end *int,
    duration *string,
    interval *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseDeviceMetrics],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | See /api/v1/const/insight_metrics for available metrics |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.ResponseDeviceMetrics`](../../doc/models/response-device-metrics.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

metric := "metric8"

deviceMac := "0000000000ab"





duration := "10m"

interval := "10m"

limit := 100

page := 1

apiResponse, err := sitesInsights.GetSiteInsightMetricsForDevice(ctx, siteId, metric, deviceMac, nil, nil, &duration, &interval, &limit, &page)
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
  "end": 1604347200,
  "interval": 3600,
  "limit": 168,
  "page": 1,
  "results": [
    10,
    11,
    12,
    12,
    10,
    9,
    9,
    9,
    10,
    10,
    11,
    11,
    11,
    11,
    11,
    11,
    11,
    10,
    11,
    11,
    10,
    11,
    11,
    10
  ],
  "rt": [
    "2020-11-01 20:00:00+00:00",
    "2020-11-01 21:00:00+00:00",
    "2020-11-01 22:00:00+00:00",
    "2020-11-01 23:00:00+00:00",
    "2020-11-02 00:00:00+00:00",
    "2020-11-02 01:00:00+00:00",
    "2020-11-02 02:00:00+00:00",
    "2020-11-02 03:00:00+00:00",
    "2020-11-02 04:00:00+00:00",
    "2020-11-02 05:00:00+00:00",
    "2020-11-02 06:00:00+00:00",
    "2020-11-02 07:00:00+00:00",
    "2020-11-02 08:00:00+00:00",
    "2020-11-02 09:00:00+00:00",
    "2020-11-02 10:00:00+00:00",
    "2020-11-02 11:00:00+00:00",
    "2020-11-02 12:00:00+00:00",
    "2020-11-02 13:00:00+00:00",
    "2020-11-02 14:00:00+00:00",
    "2020-11-02 15:00:00+00:00",
    "2020-11-02 16:00:00+00:00",
    "2020-11-02 17:00:00+00:00",
    "2020-11-02 18:00:00+00:00",
    "2020-11-02 19:00:00+00:00"
  ],
  "start": 1604260800
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

