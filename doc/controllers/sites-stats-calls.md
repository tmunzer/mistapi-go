# Sites Stats-Calls

```go
sitesStatsCalls := client.SitesStatsCalls()
```

## Class Name

`SitesStatsCalls`

## Methods

* [Count Site Calls](../../doc/controllers/sites-stats-calls.md#count-site-calls)
* [Get Site Calls Summary](../../doc/controllers/sites-stats-calls.md#get-site-calls-summary)
* [List Site Troubleshoot Calls](../../doc/controllers/sites-stats-calls.md#list-site-troubleshoot-calls)
* [Search Site Calls](../../doc/controllers/sites-stats-calls.md#search-site-calls)
* [Troubleshoot Site Call](../../doc/controllers/sites-stats-calls.md#troubleshoot-site-call)


# Count Site Calls

Count by Distinct Attributes of Calls

```go
CountSiteCalls(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.CountSiteCallsDistinctEnum,
    rating *int,
    app *string,
    start *int,
    end *int,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.CountSiteCallsDistinctEnum`](../../doc/models/count-site-calls-distinct-enum.md) | Query, Optional | **Default**: `"mac"` |
| `rating` | `*int` | Query, Optional | Feedback rating (e.g. "rating=1" or "rating=1,2")<br><br>**Constraints**: `>= 1`, `<= 5` |
| `app` | `*string` | Query, Optional | - |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.CountSiteCallsDistinctEnum_MAC

rating := 5

app := "zoom"





limit := 100

apiResponse, err := sitesStatsCalls.CountSiteCalls(ctx, siteId, &distinct, &rating, &app, nil, nil, &limit)
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


# Get Site Calls Summary

Summarized, aggregated stats for the site calls

```go
GetSiteCallsSummary(
    ctx context.Context,
    siteId uuid.UUID,
    apMac *string,
    app *string,
    start *int,
    end *int) (
    models.ApiResponse[models.ResponseStatsCallsSummary],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `apMac` | `*string` | Query, Optional | AP MAC, optional |
| `app` | `*string` | Query, Optional | APp name (`zoom` or `teams`). default is both. Optional |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseStatsCallsSummary](../../doc/models/response-stats-calls-summary.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apMac := "001122334455"

app := "zoom"





apiResponse, err := sitesStatsCalls.GetSiteCallsSummary(ctx, siteId, &apMac, &app, nil, nil)
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
  "bad_minutes": 5566,
  "bad_minutes_client": 526,
  "bad_minutes_site_wan": 3612,
  "bad_minutes_wireless": 1428,
  "num_aps": 1,
  "num_users": 3,
  "total_minutes": 575217
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


# List Site Troubleshoot Calls

Summary of calls troubleshoot by site

```go
ListSiteTroubleshootCalls(
    ctx context.Context,
    siteId uuid.UUID,
    ap *string,
    meetingId *string,
    mac *string,
    app *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseCallTroubleshootSummary],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `ap` | `*string` | Query, Optional | AP MAC |
| `meetingId` | `*string` | Query, Optional | meeting_id |
| `mac` | `*string` | Query, Optional | Device identifier |
| `app` | `*string` | Query, Optional | Third party app name |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCallTroubleshootSummary](../../doc/models/response-call-troubleshoot-summary.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ap := "001122334455"

meetingId := "1234567890abcdef"

mac := "001122334455"

app := "zoom"





duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesStatsCalls.ListSiteTroubleshootCalls(ctx, siteId, &ap, &meetingId, &mac, &app, nil, nil, &duration, &limit, &page)
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
  "mac": "983a78ea4a44",
  "meeting_id": "b784d744-9a7c-4fad-9af0-f78858a319b1",
  "results": [
    {
      "audio_in": {
        "ap_num_clients": 45.48306793636746,
        "ap_rtt": 6.352042701509264,
        "client_cpu": 9.323452578650581,
        "client_radio_band": 1.5325644097982958E-06,
        "client_rssi": 17.251008563571506,
        "client_tx_bytes": 1.8379177401463191,
        "client_tx_rates": 10.668423069847954,
        "client_tx_retries": 43.323209603627525,
        "client_vpn_distance": 112.4420166015625,
        "expected": 29.74261474609375,
        "radio_bandwidth": -0.1533682727151447,
        "radio_channel": 0.662909648484654,
        "radio_util": 27.891777674357098,
        "radio_util_interference": 4.38913492154744,
        "site_num_clients": -0.2855822932389047,
        "site_wan_avg_upload_mpbs": -0.988989942603641,
        "site_wan_jitter": 0.7875519659784105,
        "site_wan_rtt": 15.094849904378256,
        "site_wan_upload_mpbs": -0.8131117953194512
      },
      "audio_out": {
        "ap_num_clients": 45.48306793636746,
        "ap_rtt": 6.352042701509264,
        "client_cpu": 9.323452578650581,
        "client_radio_band": 1.5325644097982958E-06,
        "client_rssi": 17.251008563571506,
        "client_tx_bytes": 1.8379177401463191,
        "client_tx_rates": 10.668423069847954,
        "client_tx_retries": 43.323209603627525,
        "client_vpn_distance": 112.4420166015625,
        "expected": 29.74261474609375,
        "radio_bandwidth": -0.1533682727151447,
        "radio_channel": 0.662909648484654,
        "radio_util": 27.891777674357098,
        "radio_util_interference": 4.38913492154744,
        "site_num_clients": -0.2855822932389047,
        "site_wan_avg_upload_mpbs": -0.988989942603641,
        "site_wan_jitter": 0.7875519659784105,
        "site_wan_rtt": 15.094849904378256,
        "site_wan_upload_mpbs": -0.8131117953194512
      },
      "timestamp": 1695425115,
      "video_in": {
        "ap_num_clients": 45.48306793636746,
        "ap_rtt": 6.352042701509264,
        "client_cpu": 9.323452578650581,
        "client_radio_band": 1.5325644097982958E-06,
        "client_rssi": 17.251008563571506,
        "client_tx_bytes": 1.8379177401463191,
        "client_tx_rates": 10.668423069847954,
        "client_tx_retries": 43.323209603627525,
        "client_vpn_distance": 112.4420166015625,
        "expected": 29.74261474609375,
        "radio_bandwidth": -0.1533682727151447,
        "radio_channel": 0.662909648484654,
        "radio_util": 27.891777674357098,
        "radio_util_interference": 4.38913492154744,
        "site_num_clients": -0.2855822932389047,
        "site_wan_avg_upload_mpbs": -0.988989942603641,
        "site_wan_jitter": 0.7875519659784105,
        "site_wan_rtt": 15.094849904378256,
        "site_wan_upload_mpbs": -0.8131117953194512
      },
      "video_out": {
        "ap_num_clients": 45.48306793636746,
        "ap_rtt": 6.352042701509264,
        "client_cpu": 9.323452578650581,
        "client_radio_band": 1.5325644097982958E-06,
        "client_rssi": 17.251008563571506,
        "client_tx_bytes": 1.8379177401463191,
        "client_tx_rates": 10.668423069847954,
        "client_tx_retries": 43.323209603627525,
        "client_vpn_distance": 112.4420166015625,
        "expected": 29.74261474609375,
        "radio_bandwidth": -0.1533682727151447,
        "radio_channel": 0.662909648484654,
        "radio_util": 27.891777674357098,
        "radio_util_interference": 4.38913492154744,
        "site_num_clients": -0.2855822932389047,
        "site_wan_avg_upload_mpbs": -0.988989942603641,
        "site_wan_jitter": 0.7875519659784105,
        "site_wan_rtt": 15.094849904378256,
        "site_wan_upload_mpbs": -0.8131117953194512
      }
    }
  ]
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


# Search Site Calls

Search Calls

```go
SearchSiteCalls(
    ctx context.Context,
    siteId uuid.UUID,
    mac *string,
    app *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseStatsCalls],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | Device identifier |
| `app` | `*string` | Query, Optional | Third party app name |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseStatsCalls](../../doc/models/response-stats-calls.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mac := "001122334455"

app := "zoom"

limit := 100





duration := "10m"

apiResponse, err := sitesStatsCalls.SearchSiteCalls(ctx, siteId, &mac, &app, &limit, nil, nil, &duration)
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


# Troubleshoot Site Call

Troubleshoot a call

```go
TroubleshootSiteCall(
    ctx context.Context,
    siteId uuid.UUID,
    clientMac string,
    meetingId string,
    mac *string,
    app *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.CallTroubleshoot],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `clientMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |
| `meetingId` | `string` | Query, Required | meeting_id |
| `mac` | `*string` | Query, Optional | Device identifier |
| `app` | `*string` | Query, Optional | Third party app name |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.CallTroubleshoot](../../doc/models/call-troubleshoot.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

clientMac := "0000000000ab"

meetingId := "1234567890abcdef"

mac := "001122334455"

app := "zoom"





duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesStatsCalls.TroubleshootSiteCall(ctx, siteId, clientMac, meetingId, &mac, &app, nil, nil, &duration, &limit, &page)
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
  "mac": "983a78ea4a44",
  "meeting_id": "b784d744-9a7c-4fad-9af0-f78858a319b1",
  "results": [
    {
      "audio_in": {
        "ap_num_clients": -0.6565111,
        "ap_rtt": 0.16559607,
        "client_cpu": 3.7028809,
        "client_n_streams": 0.15803306,
        "client_radio_band": 0.5576923,
        "client_rssi": -1.0839354,
        "client_rx_bytes": 2.2622051,
        "client_rx_rates": 0.62357205,
        "client_rx_retries": 0.26726437,
        "client_tx_bytes": 0.15803306,
        "client_tx_rates": 0.62357205,
        "client_tx_retries": 0.77553505,
        "client_vpn_distance": 1.6474955,
        "client_wifi_version": 0.18267937,
        "expected": 30.941595,
        "radio_bandwidth": -0.06538621,
        "radio_channel": -0.73391086,
        "radio_tx_power": 0.10027129,
        "radio_util": 12.770318,
        "radio_util_interference": -3.079999,
        "site_num_clients": 0.017364305,
        "wan_avg_download_mbps": 1.4803165,
        "wan_avg_upload_mbps": -0.038184267,
        "wan_jitter": 5.9680853,
        "wan_max_download_mbps": 1.4803165,
        "wan_max_upload_mbps": -0.038184267,
        "wan_rtt": 46.77899
      },
      "audio_out": {
        "ap_num_clients": -0.6565111,
        "ap_rtt": 0.16559607,
        "client_cpu": 3.7028809,
        "client_n_streams": 0.15803306,
        "client_radio_band": 0.5576923,
        "client_rssi": -1.0839354,
        "client_rx_bytes": 2.2622051,
        "client_rx_rates": 0.62357205,
        "client_rx_retries": 0.26726437,
        "client_tx_bytes": 0.15803306,
        "client_tx_rates": 0.62357205,
        "client_tx_retries": 0.77553505,
        "client_vpn_distance": 1.6474955,
        "client_wifi_version": 0.18267937,
        "expected": 30.941595,
        "radio_bandwidth": -0.06538621,
        "radio_channel": -0.73391086,
        "radio_tx_power": 0.10027129,
        "radio_util": 12.770318,
        "radio_util_interference": -3.079999,
        "site_num_clients": 0.017364305,
        "wan_avg_download_mbps": 1.4803165,
        "wan_avg_upload_mbps": -0.038184267,
        "wan_jitter": 5.9680853,
        "wan_max_download_mbps": 1.4803165,
        "wan_max_upload_mbps": -0.038184267,
        "wan_rtt": 46.77899
      },
      "timestamp": 1695425115,
      "video_in": {
        "ap_num_clients": -0.6565111,
        "ap_rtt": 0.16559607,
        "client_cpu": 3.7028809,
        "client_n_streams": 0.15803306,
        "client_radio_band": 0.5576923,
        "client_rssi": -1.0839354,
        "client_rx_bytes": 2.2622051,
        "client_rx_rates": 0.62357205,
        "client_rx_retries": 0.26726437,
        "client_tx_bytes": 0.15803306,
        "client_tx_rates": 0.62357205,
        "client_tx_retries": 0.77553505,
        "client_vpn_distance": 1.6474955,
        "client_wifi_version": 0.18267937,
        "expected": 30.941595,
        "radio_bandwidth": -0.06538621,
        "radio_channel": -0.73391086,
        "radio_tx_power": 0.10027129,
        "radio_util": 12.770318,
        "radio_util_interference": -3.079999,
        "site_num_clients": 0.017364305,
        "wan_avg_download_mbps": 1.4803165,
        "wan_avg_upload_mbps": -0.038184267,
        "wan_jitter": 5.9680853,
        "wan_max_download_mbps": 1.4803165,
        "wan_max_upload_mbps": -0.038184267,
        "wan_rtt": 46.77899
      },
      "video_out": {
        "ap_num_clients": -0.6565111,
        "ap_rtt": 0.16559607,
        "client_cpu": 3.7028809,
        "client_n_streams": 0.15803306,
        "client_radio_band": 0.5576923,
        "client_rssi": -1.0839354,
        "client_rx_bytes": 2.2622051,
        "client_rx_rates": 0.62357205,
        "client_rx_retries": 0.26726437,
        "client_tx_bytes": 0.15803306,
        "client_tx_rates": 0.62357205,
        "client_tx_retries": 0.77553505,
        "client_vpn_distance": 1.6474955,
        "client_wifi_version": 0.18267937,
        "expected": 30.941595,
        "radio_bandwidth": -0.06538621,
        "radio_channel": -0.73391086,
        "radio_tx_power": 0.10027129,
        "radio_util": 12.770318,
        "radio_util_interference": -3.079999,
        "site_num_clients": 0.017364305,
        "wan_avg_download_mbps": 1.4803165,
        "wan_avg_upload_mbps": -0.038184267,
        "wan_jitter": 5.9680853,
        "wan_max_download_mbps": 1.4803165,
        "wan_max_upload_mbps": -0.038184267,
        "wan_rtt": 46.77899
      }
    }
  ]
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

