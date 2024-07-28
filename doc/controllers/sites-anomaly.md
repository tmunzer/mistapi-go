# Sites Anomaly

```go
sitesAnomaly := client.SitesAnomaly()
```

## Class Name

`SitesAnomaly`

## Methods

* [Get Site Anomaly Events](../../doc/controllers/sites-anomaly.md#get-site-anomaly-events)
* [Get Site Anomaly Events for Client](../../doc/controllers/sites-anomaly.md#get-site-anomaly-events-for-client)
* [Get Site Anomaly Eventsfor Device](../../doc/controllers/sites-anomaly.md#get-site-anomaly-eventsfor-device)


# Get Site Anomaly Events

Get Site Anomaly Events

```go
GetSiteAnomalyEvents(
    ctx context.Context,
    siteId uuid.UUID,
    metric string) (
    models.ApiResponse[models.ResponseAnomalySearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | see /api/v1/const/insight_metrics for available metrics |

## Response Type

[`models.ResponseAnomalySearch`](../../doc/models/response-anomaly-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

metric := "metric8"

apiResponse, err := sitesAnomaly.GetSiteAnomalyEvents(ctx, siteId, metric)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Anomaly Events for Client

Get Client Anomaly Events

```go
GetSiteAnomalyEventsForClient(
    ctx context.Context,
    siteId uuid.UUID,
    clientMac string,
    metric string) (
    models.ApiResponse[models.ResponseAnomalySearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `clientMac` | `string` | Template, Required | - |
| `metric` | `string` | Template, Required | see /api/v1/const/insight_metrics for available metrics |

## Response Type

[`models.ResponseAnomalySearch`](../../doc/models/response-anomaly-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

clientMac := "0000000000ab"

metric := "metric8"

apiResponse, err := sitesAnomaly.GetSiteAnomalyEventsForClient(ctx, siteId, clientMac, metric)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Anomaly Eventsfor Device

Get Device Anomaly Events

```go
GetSiteAnomalyEventsforDevice(
    ctx context.Context,
    siteId uuid.UUID,
    metric string,
    deviceMac string) (
    models.ApiResponse[models.ResponseAnomalySearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | see /api/v1/const/insight_metrics for available metrics |
| `deviceMac` | `string` | Template, Required | - |

## Response Type

[`models.ResponseAnomalySearch`](../../doc/models/response-anomaly-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

metric := "metric8"

deviceMac := "0000000000ab"

apiResponse, err := sitesAnomaly.GetSiteAnomalyEventsforDevice(ctx, siteId, metric, deviceMac)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

