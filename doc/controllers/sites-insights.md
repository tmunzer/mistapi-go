# Sites Insights

```go
sitesInsights := client.SitesInsights()
```

## Class Name

`SitesInsights`

## Methods

* [Get Site Insight Metrics](../../doc/controllers/sites-insights.md#get-site-insight-metrics)
* [Get Site Insight Metrics for AP](../../doc/controllers/sites-insights.md#get-site-insight-metrics-for-ap)
* [Get Site Insight Metrics for Client](../../doc/controllers/sites-insights.md#get-site-insight-metrics-for-client)
* [Get Site Insight Metrics for Device](../../doc/controllers/sites-insights.md#get-site-insight-metrics-for-device)
* [Get Site Insight Metrics for Gateway](../../doc/controllers/sites-insights.md#get-site-insight-metrics-for-gateway)
* [Get Site Insight Metrics for Mx Edge](../../doc/controllers/sites-insights.md#get-site-insight-metrics-for-mx-edge)
* [Get Site Insight Metrics for Switch](../../doc/controllers/sites-insights.md#get-site-insight-metrics-for-switch)


# Get Site Insight Metrics

Get Site Insight Metrics

```go
GetSiteInsightMetrics(
    ctx context.Context,
    siteId uuid.UUID,
    metrics string,
    start *string,
    end *string,
    duration *string,
    interval *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.InsightMetrics],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `metrics` | `string` | Query, Required | Comma separated Metric names, e.g. `num_clients,num_aps`. See possible values at [List Insight Metrics](/#operations/listInsightMetrics) |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.InsightMetrics](../../doc/models/insight-metrics.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

metrics := "num_clients,num_aps"

duration := "10m"

interval := "10m"

limit := 100

page := 1

apiResponse, err := sitesInsights.GetSiteInsightMetrics(ctx, siteId, metrics, nil, nil, &duration, &interval, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Get Site Insight Metrics for AP

Get AP Insight Metrics

```go
GetSiteInsightMetricsForAP(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    metrics string,
    start *string,
    end *string,
    duration *string,
    interval *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseDeviceMetrics],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `metrics` | `string` | Query, Required | Comma separated Metric names, e.g. `num_clients,num_stressed_clients`. See possible values at [List Insight Metrics](/#operations/listInsightMetrics) |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceMetrics](../../doc/models/response-device-metrics.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

metrics := "num_clients,num_stressed_clients"

duration := "10m"

interval := "10m"

limit := 100

page := 1

apiResponse, err := sitesInsights.GetSiteInsightMetricsForAP(ctx, siteId, deviceId, metrics, nil, nil, &duration, &interval, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Get Site Insight Metrics for Client

Get Client Insight Metrics

```go
GetSiteInsightMetricsForClient(
    ctx context.Context,
    siteId uuid.UUID,
    clientMac string,
    metrics string,
    start *string,
    end *string,
    duration *string,
    interval *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.InsightMetrics],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `clientMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `metrics` | `string` | Query, Required | Comma separated Metric names, e.g. `top-app-by-num_client,top-app-by-bytes`. See possible values at [List Insight Metrics](/#operations/listInsightMetrics) |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.InsightMetrics](../../doc/models/insight-metrics.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

clientMac := "0000000000ab"

metrics := "top-app-by-num_client,top-app-by-bytes"

duration := "10m"

interval := "10m"

limit := 100

page := 1

apiResponse, err := sitesInsights.GetSiteInsightMetricsForClient(ctx, siteId, clientMac, metrics, nil, nil, &duration, &interval, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Get Site Insight Metrics for Device

Get AP Insight Metrics
See metrics possibilities at [List Insight Metrics](../../doc/controllers/constants-definitions.md#list-insight-metrics)

```go
GetSiteInsightMetricsForDevice(
    ctx context.Context,
    siteId uuid.UUID,
    metric string,
    deviceMac string,
    portId *string,
    start *string,
    end *string,
    duration *string,
    interval *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseDeviceMetrics],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | See [List Insight Metrics](../../doc/controllers/constants-definitions.md#list-insight-metrics) for available metrics |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `portId` | `*string` | Query, Optional | Port ID of the device, e.g. `ge-0/0/1`. Can be used with metrics related to interfaces, e.g. `rx_bytes`. |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceMetrics](../../doc/models/response-device-metrics.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

metric := "metric8"

deviceMac := "0000000000ab"

portId := "ge-0/0/1"

duration := "10m"

interval := "10m"

limit := 100

page := 1

apiResponse, err := sitesInsights.GetSiteInsightMetricsForDevice(ctx, siteId, metric, deviceMac, &portId, nil, nil, &duration, &interval, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Get Site Insight Metrics for Gateway

Get Gateway Insight Metrics

```go
GetSiteInsightMetricsForGateway(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    metrics string,
    portId *string,
    start *string,
    end *string,
    duration *string,
    interval *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseDeviceMetrics],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `metrics` | `string` | Query, Required | Comma separated Metric names, e.g. `tx_bps,rx_bps`. See possible values at [List Insight Metrics](/#operations/listInsightMetrics) |
| `portId` | `*string` | Query, Optional | Port ID of the gateway device, e.g. `ge-0/0/1` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceMetrics](../../doc/models/response-device-metrics.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

metrics := "tx_bps,rx_bps"

portId := "ge-0/0/1"

duration := "10m"

interval := "10m"

limit := 100

page := 1

apiResponse, err := sitesInsights.GetSiteInsightMetricsForGateway(ctx, siteId, deviceId, metrics, &portId, nil, nil, &duration, &interval, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Get Site Insight Metrics for Mx Edge

Get MxEdge Insight Metrics
See metrics possibilities at [List Insight Metrics](/#operations/listInsightMetrics)

```go
GetSiteInsightMetricsForMxEdge(
    ctx context.Context,
    siteId uuid.UUID,
    metric string,
    deviceMac string,
    portId *string,
    start *string,
    end *string,
    duration *string,
    interval *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseDeviceMetrics],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | See [List Insight Metrics](../../doc/controllers/constants-definitions.md#list-insight-metrics) for available metrics |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `portId` | `*string` | Query, Optional | Port ID of the MxEdge device, e.g. `port0`. Can be used with metrics related to interfaces, e.g. `rx_bytes`. |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceMetrics](../../doc/models/response-device-metrics.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

metric := "metric8"

deviceMac := "0000000000ab"

portId := "port0"

duration := "10m"

interval := "10m"

limit := 100

page := 1

apiResponse, err := sitesInsights.GetSiteInsightMetricsForMxEdge(ctx, siteId, metric, deviceMac, &portId, nil, nil, &duration, &interval, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Get Site Insight Metrics for Switch

Get Switch Insight Metrics
See metrics possibilities at [List Insight Metrics](/#operations/listInsightMetrics)

```go
GetSiteInsightMetricsForSwitch(
    ctx context.Context,
    siteId uuid.UUID,
    metric string,
    deviceMac string,
    portId *string,
    start *string,
    end *string,
    duration *string,
    interval *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseDeviceMetrics],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | See [List Insight Metrics](../../doc/controllers/constants-definitions.md#list-insight-metrics) for available metrics |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `portId` | `*string` | Query, Optional | Port ID of the switch device, e.g. `ge-0/0/1`. Can be used with metrics related to interfaces, e.g. `rx_bytes`. |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceMetrics](../../doc/models/response-device-metrics.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

metric := "metric8"

deviceMac := "0000000000ab"

portId := "ge-0/0/1"

duration := "10m"

interval := "10m"

limit := 100

page := 1

apiResponse, err := sitesInsights.GetSiteInsightMetricsForSwitch(ctx, siteId, metric, deviceMac, &portId, nil, nil, &duration, &interval, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

