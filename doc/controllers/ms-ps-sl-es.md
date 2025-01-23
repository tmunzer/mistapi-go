# MS Ps SL Es

```go
mSPsSLEs := client.MSPsSLEs()
```

## Class Name

`MSPsSLEs`


# Get Msp Sle

Get MSP SLEs (all/worst Orgs ...)

```go
GetMspSle(
    ctx context.Context,
    mspId uuid.UUID,
    metric string,
    sle *string,
    duration *string,
    interval *string,
    start *int,
    end *int) (
    models.ApiResponse[models.InsightMetrics],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | see /api/v1/const/insight_metrics for available metrics |
| `sle` | `*string` | Query, Optional | see /api/v1/const/insight_metrics for more details |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |

## Response Type

[`models.InsightMetrics`](../../doc/models/insight-metrics.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

metric := "metric8"



duration := "10m"

interval := "10m"





apiResponse, err := mSPsSLEs.GetMspSle(ctx, mspId, metric, nil, &duration, &interval, nil, nil)
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
  "end": 1643706000,
  "interval": 3600,
  "limit": 100,
  "results": [
    {
      "ap-availability": 0.9919400860511628,
      "ap-health": 0.967607512909879,
      "capacity": 0.7484652273070254,
      "coverage": 0.91217567374857,
      "num_aps": 13,
      "num_clients": 12,
      "org_id": "ab0aca7a-d45c-469b-b3bb-4fe240642d0b",
      "roaming": 0.991735537682683,
      "roaming-exp": 0.991735537682683,
      "successful-connect": 0.46052632135780236,
      "throughput": 0.6775702123846302,
      "time-to-connect": 0.9349112447196916
    },
    {
      "ap-availability": 0.9990384613092129,
      "ap-health": 0.48201754375507955,
      "capacity": 0.9702673450306101,
      "coverage": 0.8335392334930375,
      "num_aps": 1,
      "num_clients": 6,
      "org_id": "49ff76e0-a283-4e7d-b38d-041f1e9aff3c",
      "roaming": 1,
      "roaming-exp": 1,
      "successful-connect": 1,
      "throughput": 0,
      "time-to-connect": 1
    },
    {
      "ap-availability": 1,
      "ap-health": 0.982456140612301,
      "capacity": 1,
      "coverage": 0.9276041182442488,
      "num_aps": 2,
      "num_clients": 3,
      "org_id": "9b9b48f1-15a4-459e-86cc-9cbec9005983",
      "roaming": 1,
      "roaming-exp": 1,
      "successful-connect": 1,
      "throughput": 1,
      "time-to-connect": 0.8125
    },
    {
      "ap-availability": 0.9981132070973234,
      "ap-health": 0.9991228068084047,
      "capacity": 1,
      "coverage": 1,
      "num_aps": 1,
      "num_clients": 0,
      "org_id": "eb0e1671-7a6b-472b-94c3-c187dafe5274",
      "roaming": 1,
      "roaming-exp": 1,
      "successful-connect": 1,
      "throughput": 0,
      "time-to-connect": 0.5
    }
  ],
  "start": 1643670000
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

