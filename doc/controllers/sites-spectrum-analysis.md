# Sites Spectrum Analysis

```go
sitesSpectrumAnalysis := client.SitesSpectrumAnalysis()
```

## Class Name

`SitesSpectrumAnalysis`

## Methods

* [Get Site Running Spectrum Analysis](../../doc/controllers/sites-spectrum-analysis.md#get-site-running-spectrum-analysis)
* [Initiate Site Analyze Spectrum](../../doc/controllers/sites-spectrum-analysis.md#initiate-site-analyze-spectrum)
* [List Site Spectrum Analysis](../../doc/controllers/sites-spectrum-analysis.md#list-site-spectrum-analysis)


# Get Site Running Spectrum Analysis

Get the running spectrum analysis for a site

```go
GetSiteRunningSpectrumAnalysis(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.ResponseRunningSpectrumAnalysis],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseRunningSpectrumAnalysis](../../doc/models/response-running-spectrum-analysis.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesSpectrumAnalysis.GetSiteRunningSpectrumAnalysis(ctx, siteId)
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
  "band": "5",
  "device_id": "00000000-0000-0000-1000-5c5b35bd76bb",
  "duration": 600,
  "format": "stream",
  "started_time": 1435080709
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


# Initiate Site Analyze Spectrum

Initiate a spectrum analysis for a site

The output will be available through websocket. As there can be multiple command
issued against the same device at the same time and the output all goes through
the same websocket stream, session is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

`json { "subscribe": "/sites/{site_id}/analyze_spectrum" } `

#### Example output from ws stream

```json
{
  "event": "data",
  "channel": "/sites/4ac1dcf4-9d8b-7211-65c4-057819f0862b/analyze_spectrum",
  "data": {
      "session": "session_id",

      "fft_samples": [
          {
              "frequency": 2437.0,
              "rssi / signal ?": -93
          },
          ...
      ],

      "channel_usage": [
          {
              "channel": 36,
              "noise": -78,

              "wifi": 0.13,
              "non_wifi": 0.08
          },
          ...
      ]
  }
}     
```

```go
InitiateSiteAnalyzeSpectrum(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.SpectrumAnalysis) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SpectrumAnalysis`](../../doc/models/spectrum-analysis.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.WebsocketSession](../../doc/models/websocket-session.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SpectrumAnalysis{
    Band:                 models.SpectrumAnalysisBandEnum_ENUM5,
    DeviceId:             models.ToPointer(uuid.MustParse("00000000-0000-0000-1000-5c5b35bd76bb")),
    Duration:             models.ToPointer(600),
    Format:               models.ToPointer(models.SpectrumAnalysisFormatEnum_STREAM),
}

apiResponse, err := sitesSpectrumAnalysis.InitiateSiteAnalyzeSpectrum(ctx, siteId, &body)
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


# List Site Spectrum Analysis

List the past spectrum analysis for a site

```go
ListSiteSpectrumAnalysis(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    start *string,
    end *string,
    duration *string) (
    models.ApiResponse[models.ResponsePastSpectrumAnalysis],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponsePastSpectrumAnalysis](../../doc/models/response-past-spectrum-analysis.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

duration := "10m"

apiResponse, err := sitesSpectrumAnalysis.ListSiteSpectrumAnalysis(ctx, siteId, &limit, nil, nil, &duration)
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
  "end": 1694708579,
  "limit": 10,
  "results": [
    {
      "band": "5",
      "channel_usage": [
        {
          "channel": 36,
          "noise": -78,
          "non_wifi": 0.08,
          "wifi": 0.13
        }
      ],
      "fft_samples": [
        {
          "frequency": 2437,
          "rssi / signal ?": -93
        }
      ],
      "mac": "5c5b35bd76bb",
      "org_id": "f2695c32-0e83-4936-b1b2-96fc88051213",
      "timestamp": 1694098696
    }
  ],
  "start": 1694622179,
  "total": 4
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

